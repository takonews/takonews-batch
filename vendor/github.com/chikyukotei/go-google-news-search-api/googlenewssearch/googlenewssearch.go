package googlenewssearch

import "net/url"

// QueryParam is a set of query parameters
type QueryParam struct {
	// Language
	Hl string

	// Country
	Ned string

	// Encodding used in articles
	Ie string

	// Encoding used in response
	Oe string

	// Response format
	// rss/atom
	Output string

	// Keyword
	Q string

	// topic scopes results to a particular topic.
	// You cannot use topic in conjunction with either the query argument. If you specify query with topic, the searcher ignores it.
	/* PARAMETERS
	ir Spotlight
	y Society	w World
	b Business	p Politics
	e Entertainment		s Sports
	t Sci/Tech	po Most Popular
	*/
	Topic string
}

const baseURL = "http://news.google.com"

// RequestURL builds request URL from Query parameters
func RequestURL(p *QueryParam) string {
	var URL *url.URL
	URL, err := url.Parse(baseURL)
	if err != nil {
		panic("boom")
	}
	URL.Path += "/news"
	params := url.Values{}
	params.Add("hl", p.Hl)
	params.Add("ned", p.Ned)
	params.Add("ie", p.Ie)
	params.Add("oe", p.Oe)
	params.Add("output", p.Output)
	params.Add("q", p.Q)
	params.Add("topic", p.Topic)
	URL.RawQuery = params.Encode()
	return URL.String()
}
