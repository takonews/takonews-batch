package cron

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chikyukotei/go-google-news-search-api/googlenewssearch"
	"github.com/mauidude/go-readability"
	"github.com/robfig/cron"
	"github.com/takonews/takonews-api/app/models"
	"github.com/takonews/takonews-api/db"
	"github.com/takonews/takonews-batch/cron/rss"
)

var Cron = cron.New()

func init() {
	// sec min hour day month
	Cron.AddFunc("0 */1 * * * ?", storeArticles)
}

func storeArticles() {
	fmt.Println("*******begin storeArticles*******")
	for _, q := range config.QList {
		qp := &googlenewssearch.QueryParam{
			Hl:     "ja",
			Ned:    "us",
			Ie:     "UTF-8",
			Oe:     "UTF-8",
			Output: "rss",
			Q:      q,
		}
		gurl := googlenewssearch.RequestURL(qp)
		itemList, _ := rss.Parse(gurl)
		for _, v := range itemList {
			datetime, _ := time.Parse(time.RFC1123, v.PubDate)
			u, _ := url.Parse(v.Link)
			m, _ := url.ParseQuery(u.RawQuery)
			articleURL, _ := url.Parse(m["url"][0])
			vt := strings.Split(v.Title, " - ")
			fmt.Println(vt)
			if len(vt) != 2 { // title must be [news_title] - [news_site_name]
				continue
			}
			// insert phase
			sql := db.DB
			// insert news_site
			newsSiteName := vt[1]
			var newsSite models.NewsSite
			sql.Model(models.NewsSite{}).FirstOrCreate(&newsSite, models.NewsSite{Name: newsSiteName, URL: articleURL.Host})
			// insert article
			fullText, _ := getFullText(m["url"][0])
			article := models.Article{Title: vt[0], PublishedAt: datetime, URL: m["url"][0], NewsSiteID: newsSite.ID, FullText: fullText}
			var count int
			sql.Model(models.Article{}).Where("url = ?", article.URL).Count(&count)
			if count == 0 {
				sql.Create(&article)
			}
		}
	}
}

func getFullText(url string) (content string, err error) {
	response, _ := http.Get(url)
	defer response.Body.Close()
	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	doc, err := readability.NewDocument(string(html))
	if err != nil {
		return
	}
	content = doc.Content()
	return
}
