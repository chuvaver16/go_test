package client

import (
	"io"
	"log"
	"net/http"
)

func Get(uri string, qparams map[string]string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range qparams {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	//log.Print(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	//log.Print(string(body))

	return body, nil
}
