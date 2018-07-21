package helpers

import (
	"io/ioutil"
	"net/http"
)

func ReadBody(r *http.Response) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	r.Body.Close() //  must close

	return bodyBytes, err
}
