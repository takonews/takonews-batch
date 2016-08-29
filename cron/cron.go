package cron

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/chikyukotei/go-google-news-search-api/googlenewssearch"
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
			vt := strings.Split(v.Title, " - ")
			if len(vt) != 2 {
				continue
			}
			article := models.Article{Title: vt[0], PublishedAt: datetime, URL: m["url"][0]}
			if db.DB.NewRecord(article) {
				db.DB.Create(&article)
			}
		}
	}
}
