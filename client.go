package ytclient

import (
	"github.com/go-resty/resty"
	"strings"
	"log"
	"time"
	"errors"
)

type YouTrackClient struct {
	Token string
	Addr  string
	Debug bool
}

type Request struct {
	QueryParams string
	Endpoint    string
	Method      string
}

func NewYouTrackClient(token, addr string, debug bool) (*YouTrackClient, error) {
	client := &YouTrackClient{Token: token, Addr: addr, Debug: debug}
	if err := client.HealthCheck(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *YouTrackClient) debugLog(context string, req *resty.Request, res *resty.Response) {
	if c.Debug {
		log.Println(context)
		if req != nil {
			log.Printf("%v Request: %s %s", time.Now().Format(time.RFC822), req.Method, req.URL)
		}
		if res != nil {
			log.Printf("%v Response: %s %s", time.Now().Format(time.RFC822), req.Method, req.URL)
		}
	}
}

func (c *YouTrackClient) SendRequest(request *Request) (*resty.Response, error) {
	r := resty.R().
		SetAuthToken(c.Token).
		SetQueryString(request.QueryParams).
		SetHeader("Accept", "application/json")

	resp := &resty.Response{}
	var err error
	switch strings.ToLower(request.Method) {
	case "get":
		resp, err = r.Get(c.Addr + request.Endpoint)
	case "post":
		resp, err = r.Post(c.Addr + request.Endpoint)
	case "put":
		resp, err = r.Post(c.Addr + request.Endpoint)
	case "delete":
		resp, err = r.Post(c.Addr + request.Endpoint)
	default:
		err = errors.New("method of request is wrong")
		c.debugLog(err.Error(), r, nil)
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
