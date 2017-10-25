package client

import (
	"net/http"
	//"time"
)

func HTTPClient() *http.Client {
	//tr := &http.Transport{IdleConnTimeout: 75 * time.Second}
	//tr.IdleConnTimeout = 75 * time.Second
	//cli := &http.Client{ Transport: tr }
	cli := &http.Client{}
	return cli
}
