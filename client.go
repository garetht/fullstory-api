package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/franela/goreq"
	"github.com/garetht/fullstory-api/query"
	"github.com/garetht/fullstory-api/user"
)

type API struct {
	key string
}

const (
	fsBase      = "https://www.fullstory.com"
	fsLogin     = "/loginrequest"
	fsCookieKey = "fs_session"
	fsCsrfKey   = "csrftoken"
)

func (a *API) Init(username string, password string) {
	params := map[string]interface{}{
		"Username": username,
		"Password": password,
	}

	resp := a.post(fsLogin, params)
	cookies := resp.Cookies()

	for _, c := range cookies {
		if c.Name == fsCookieKey {
			a.key = c.Value
			break
		}
	}
}

func (a *API) UserQuery(q *query.FsQuery) (users user.FsUsers) {
	resp := a.post(user.Endpoint, q)

	users = user.FsUsers{}
	resp.Body.FromJsonTo(&users)

	return
}

func (a *API) post(endpoint string, data interface{}) (resp *goreq.Response) {
	url := strings.Join([]string{fsBase, endpoint}, "")
	req := goreq.Request{
		Uri:         url,
		Method:      "POST",
		Body:        data,
		Accept:      "application/json",
		ContentType: "application/json",
		ShowDebug:   true,
	}

	// If the key exists, send it.
	if a.key != "" {
		req = req.WithCookie(&http.Cookie{
			Name:  fsCookieKey,
			Value: a.key,
		}).WithHeader(fsCsrfKey, a.key)
	}

	resp, err := req.Do()

	if err != nil {
		fmt.Printf("%v : error encountered during Fullstory API client POST")
	}

	return
}
