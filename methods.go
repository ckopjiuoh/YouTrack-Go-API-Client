package ytclient

import (
	"fmt"
	"time"
)

func (c *YouTrackClient) HealthCheck() error {
	r, err := c.SendRequest(&Request{
		"",
		"/rest/admin/project",
		"get",
	})

	if err != nil {
		return err
	}

	if r.StatusCode() != 200 {
		return fmt.Errorf("%v, Request: %s, %s, Response status: %d, %v", time.Now().Format(time.RFC822), r.Request.Method, r.Request.URL, r.StatusCode(), string(r.Body()))
	}
	return nil
}
