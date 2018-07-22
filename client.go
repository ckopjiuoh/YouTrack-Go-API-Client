package ytapi

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"fmt"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type YouTrackAPI struct {
	Token  string
	Addr   string
	Debug  bool
	Client *http.Client
}

type Request struct {
	QueryParams *url.Values
	Endpoint    string
	Method      string
	Headers     map[string]string
	Body        io.Reader
}

type Response struct {
}

func NewYouTrackAPIWithClient(token, addr string, client *http.Client, debug bool) (*YouTrackAPI, error) {
	ytclient := &YouTrackAPI{Token: token, Addr: addr, Client: client, Debug: debug}
	if err := ytclient.HealthCheck(); err != nil {
		return nil, err
	}

	return ytclient, nil
}

func NewYouTrackAPI(token, addr string, debug bool) (*YouTrackAPI, error) {
	return NewYouTrackAPIWithClient(token, addr, &http.Client{}, debug)
}

func (c *YouTrackAPI) debugLog(context string, req *http.Request, res *http.Response) {
	if c.Debug {
		log.Println(context)
		if req != nil {
			log.Printf("%v Request: %s %s", time.Now().Format(time.RFC822), req.Method, req.URL)
		}
		if res != nil {
			log.Printf("%v Response: %d", time.Now().Format(time.RFC822), res.StatusCode)
		}
	}
}

func (api *YouTrackAPI) MakeRequest(request *Request, expectedStatusCode int) (*http.Response, error) {
	req, err := http.NewRequest(
		request.Method,
		api.Addr+request.Endpoint+request.QueryParams.Encode(),
		request.Body,
	)

	if err != nil {
		api.debugLog(err.Error(), nil, nil)
		return nil, err
	}

	// Add auth header
	req.Header.Add("Authorization", "Bearer "+api.Token)
	req.Header.Add("Accept", "application/json")

	for k, v := range request.Headers {
		req.Header.Add(k, v)
	}

	res, err := api.Client.Do(req)

	if err != nil {
		api.debugLog(err.Error(), req, res)
		return nil, err
	}

	if res.StatusCode != expectedStatusCode {
		api.debugLog(fmt.Sprintf("Expcted code is %d, but actual: %d", expectedStatusCode, res.StatusCode), req, res)
		return res, fmt.Errorf("Expcted code is %d, but actual: %d", expectedStatusCode, res.StatusCode)
	}

	api.debugLog("Success", req, res)
	return res, err
}
