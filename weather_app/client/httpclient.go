package client

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func Get(uri string, qparams map[string]string, pparams map[string]string) ([]byte, error) {

	client := &http.Client{}

	for k, v := range pparams {
		uri = strings.Replace(uri, k, v, 1)
	}

	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	if qparams != nil {
		query := req.URL.Query()
		for k, v := range qparams {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}

	log.Print(req.URL.String())

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
