package balance

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

const pathParamNoPresent = "URl UserId is missing"
const internalServerError = "Internal server error"

type Controller interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	serv Service
}

func (c *controller) GetBalance(w http.ResponseWriter, r *http.Request) {

	usrId := strings.TrimPrefix(r.URL.Path, "/balances/")

	if len(usrId) < 1 {
		log.Println(pathParamNoPresent)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(pathParamNoPresent))
		return
	}

	balance, err := c.serv.GetBalance(usrId)

	if err != nil {
		log.Println(internalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalServerError))
		return
	}

	bytes, err := json.Marshal(balance)

	if err != nil {
		log.Println(internalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(internalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewController(serv Service) (Controller, error) {
	if serv == nil {
		return nil, errors.New("service can't be nil")
	}

	return &controller{serv}, nil
}
