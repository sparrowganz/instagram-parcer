package instagram

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	DontReadBody = "Body don`t read"
	InvalidStatusCode = "StatusCode != 200"
)

//Request structure
type request struct {
	client http.Client
	url    *string
}

//Create new request
func newRequest(url string) *request {
	req := new(request)
	req.client = http.Client{Timeout: 30 * time.Second}
	req.url = &url
	return req
}

//Send request and read body
func (request *request) send() ([]byte, error) {
	res, err := request.client.Get(*request.url)
	if err != nil {
		return []byte{}, err
	} else if res.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf(InvalidStatusCode)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf(DontReadBody)
	}

	return body, nil
}
