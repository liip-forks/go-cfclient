package cfclient

import "net/url"

// Pagination is used by the V3 apis
type Pagination struct {
	TotalResults int  `json:"total_results"`
	TotalPages   int  `json:"total_pages"`
	First        Link `json:"first"`
	Last         Link `json:"last"`
	Next         *Link `json:"next"`
	Previous     *Link `json:"previous"`
}

// Link is a HATEOAS-style link for v3 apis
type Link struct {
	Href   string `json:"href"`
	Method string `json:"method,omitempty"`
}

func (l *Link) getQuery() (url.Values, error) {
	u, err := url.Parse(l.Href)
	if err != nil {
		return nil, err
	}

	return url.ParseQuery(u.RawQuery)
}
