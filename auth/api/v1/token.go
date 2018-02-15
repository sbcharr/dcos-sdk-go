package v1

import (
	"log"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"

	"github.com/PTC-GLOBAL/dcos-sdk-go/client"
)

type Credential struct {
	Uid string `json:"uid"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

//GetDCOSAccessToken() gets the session token associated with the DCOS login
//Subsequent calls to APIs can be performed with the generated token.
//Token is valid for the session

func GetDCOSAccessToken(userId, password, dcosUrl string) (string, error) {
	httpCli := client.HTTPClient()
	user := Credential{
		Uid: userId, //ask user to provide value
		Password: password, //ask user to provide value
	}
	url := fmt.Sprintf("%s/acs/api/v1/auth/login", dcosUrl) //ask user to pass dcos IP/Domain name
	body, err := json.Marshal(user)
	if err != nil {
		log.Print("error marshaling request body: ", err)
		return "", err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Print("error creating a new request", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpCli.Do(req)
	if err != nil {
		log.Print("error sending request: ", err)
		return "", err
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("error in reading response body: ", err)
		return "", err
	}

	t := Token{}
	if err = json.Unmarshal(respByte, &t); err != nil {
		log.Print("error unmarshaling response: ", err)
		return "", err
	}

	return t.Token, nil
}
