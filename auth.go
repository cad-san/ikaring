package ikaring

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func getOauthQuery(resp *http.Response, id string, password string) (url.Values, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		return nil, err
	}

	query := url.Values{}
	doc.Find("input").Each(func(_ int, s *goquery.Selection) {
		name, ok := s.Attr("name")
		if !ok {
			return
		}

		// parse from docment
		switch name {
		case "client_id":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		case "state":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		case "redirect_uri":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		case "response_type":
			if v, ok := s.Attr("value"); ok {
				query.Add(name, v)
			}
		}
	})

	// fixed value
	query.Add("nintendo_authenticate", "")
	query.Add("nintendo_authorize", "")
	query.Add("scope", "")
	query.Add("lang", "ja-JP")

	// user info
	query.Add("username", id)
	query.Add("password", password)

	return query, nil
}

func getSessionFromCookie(name string, cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
