package balance

//go:generate mockgen -destination=../balance/mock_client.go -package=balance -source=client.go Client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client interface {
	GetAllMovements(userId string) ([]*Movement, error)
}

type restClient struct {
	url string
}

func (rc *restClient) GetAllMovements(userId string) ([]*Movement, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", rc.url, userId))

	el := make([]*Movement, 0)

	err = rc.parseResponse(resp, err, &el)

	return el, err
}

func (rc *restClient) parseResponse(resp *http.Response, err error, obj interface{}) error {

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New("response status was not OK")
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(obj)

	if err != nil {
		fmt.Println(err)
	}
//	err = json.Unmarshal(bytes, obj)

	return err
}

func NewRestClient(url string) (Client, error) {

	if url == "" {
		return nil, errors.New("url can't be empty")
	}

	return &restClient{url: url}, nil

}
