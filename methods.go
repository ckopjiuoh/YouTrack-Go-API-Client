package ytapi

import (
	"fmt"
	"net/url"
	"time"
)

func (c *YouTrackAPI) HealthCheck() error {
	r, err := c.MakeRequest(&Request{
		&url.Values{},
		"/rest/admin/project",
		GET,
		nil,
		nil,
	})

	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return fmt.Errorf("%v, Request: %s, %s, Response status: %d", time.Now().Format(time.RFC822), r.Request.Method, r.Request.URL, r.StatusCode)
	}

	return nil
}
