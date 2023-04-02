package service

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type CountBody struct {
	Body string `json:"body"`
}

type CountResponse struct {
	Body int `json:"data"`
}

type CountService struct {
}

func NewCountService() CountService {
	return CountService{}
}

// Save -> calls quote repository save method
func (CountService) Count(text string) (int, error) {

	posturl := "http://counter:8000/counter/"

	body := &CountBody{
		Body: text,
	}

	b, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(b))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	println(res.Status)

	countResponse := &CountResponse{}
	derr := json.NewDecoder(res.Body).Decode(countResponse)
	if derr != nil {
		panic(derr)
	}
	return countResponse.Body, err
}
