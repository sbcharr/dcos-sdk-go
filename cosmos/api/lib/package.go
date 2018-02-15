package lib

import (

	"net/http"
	"log"
	"encoding/json"
	"github.com/PTC-GLOBAL/dcos-sdk-go/client"
	"fmt"
	"bytes"
	"io/ioutil"
)

type InstallPackageInput struct {
	PackageName string `json:"packageName" required:"true"`
	PackageVersion string `json:"packageVersion,omitempty"`
	AppId string `json:"appId,omitempty"`
	Options interface{} `json:"options,omitempty"`
}

type InstallPackageOutput struct {
	PackageName string `json:"packageName,omitempty"`
	PackageVersion string `json:"packageVersion,omitempty"`
	AppId string `json:"appId" required:"true"`
	PostInstallNotes string `json:"postInstallNotes,omitempty"`
	Cli interface{} `json:"cli,omitempty"`
	ResponseStatus string `json:"responseStatus"`
	ResponseMessage []byte `json:"ResponseMessage,omitempty"`
}


type UninstallPackageInput struct {
	PackageName string `json:"packageName" required:"true"`
	AppId string `json:"appId,omitempty"`
}

type UninstallPackageOutput struct {
	PackageName string `json:"packageName"`
	AppId string `json:"appId"`
	PackageVersion string `json:"packageVersion,omitempty"`
	PostUninstallNotes string `json:"postInstallNotes,omitempty"`
	ResponseStatus string `json:"responseStatus"`
	ResponseMessage []byte `json:"ResponseMessage,omitempty"`
}


/*
 * Runs a service from a Universe package.
 */

func InstallPackage(accessToken, dcosUrl string, packageInput InstallPackageInput) (*InstallPackageOutput, error) {
	installMediaTypeRequest = make(map[string]string)
	installMediaTypeRequest["application"] = "vnd.dcos.package.install-request+json"
	installMediaTypeRequest["charset"] = "utf-8"
	installMediaTypeRequest["version"] = "v1"

	installMediaTypeResponse = make(map[string]string)
	installMediaTypeResponse["application"] = "vnd.dcos.package.install-response+json"
	installMediaTypeResponse["charset"] = "utf-8"
	installMediaTypeResponse["version"] = "v2"

	httpCli := client.HTTPClient()

	url := fmt.Sprintf("%s/package/install", dcosUrl) //ask user to pass dcos IP/Domain name
	body, err := json.Marshal(packageInput)
	if err != nil {
		log.Print("error marshaling request body: ", err)
		return nil, err
	}
	//log.Print(string(body))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Print("error creating a new request", err)
		return nil, err
	}

	req.Header.Set("Authorization", "token=" + accessToken)
	req.Header.Set("Content-Type", "application/vnd.dcos.package.install-request+json;charset=utf-8;version=v1")
	req.Header.Set("Accept", "application/vnd.dcos.package.install-response+json;charset=utf-8;version=v2")
	//req.Header.Set("Content-Type", fmt.Sprintf(`"application/%s;charset=%s;version=%s"`, installMediaTypeRequest["application"], installMediaTypeRequest["charset"], installMediaTypeRequest["version"]))
	//req.Header.Set("Accept", fmt.Sprintf(`"application/%s;charset=%s;version=%s"`, installMediaTypeResponse["application"], installMediaTypeResponse["charset"], installMediaTypeResponse["version"]))

	resp, err := httpCli.Do(req)
	if err != nil {
		log.Print("error sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	//TODO what happens in case of network disconnection, how does client know about the status of the task submitted

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("error in reading response body: ", err)
		return nil, err
	}
	packageOutput := InstallPackageOutput{}

	packageOutput.ResponseStatus = resp.Status
	if resp.StatusCode != 200 {
		if len(respByte) < 1 {
			log.Print("response status other than 200 with nil body")
			//packageOutput.ResponseStatus = resp.Status
			return &packageOutput, nil
		} else {
			log.Print("response status other than 200")
			//packageOutput.ResponseStatus = resp.Status
			packageOutput.ResponseMessage = respByte
			return &packageOutput, nil
		}
	}
	//log.Print(string(respByte))
	//log.Print(resp.Status)

	if err = json.Unmarshal(respByte, &packageOutput); err != nil {
		log.Print("error in unmarshaling response body: ", err)
		return nil, err
	}
	//packageOutput.ResponseStatus = resp.Status
	return &packageOutput, nil
}

func UninstallPackage(accessToken, dcosUrl string, packageInput UninstallPackageInput) (*UninstallPackageOutput, error) {
	httpCli := client.HTTPClient()

	url := fmt.Sprintf("%s/package/uninstall", dcosUrl) //ask user to pass dcos IP/Domain name
	body, err := json.Marshal(packageInput)
	if err != nil {
		log.Print("error marshaling uninstall request body: ", err)
		return nil, err
	}
	//log.Print(string(body))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Print("error creating a new request", err)
		return nil, err
	}

	req.Header.Set("Authorization", "token=" + accessToken)
	req.Header.Set("Content-Type", "application/vnd.dcos.package.uninstall-request+json;charset=utf-8;version=v1")
	req.Header.Set("Accept", "application/vnd.dcos.package.uninstall-response+json;charset=utf-8;version=v1")
	//req.Header.Set("Content-Type", fmt.Sprintf(`"application/%s;charset=%s;version=%s"`, installMediaTypeRequest["application"], installMediaTypeRequest["charset"], installMediaTypeRequest["version"]))
	//req.Header.Set("Accept", fmt.Sprintf(`"application/%s;charset=%s;version=%s"`, installMediaTypeResponse["application"], installMediaTypeResponse["charset"], installMediaTypeResponse["version"]))

	resp, err := httpCli.Do(req)
	if err != nil {
		log.Print("error sending uninstall request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	//TODO what happens in case of network disconnection, how does client know about the status of the task submitted

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("error in reading uninstall response body: ", err)
		return nil, err
	}
	packageOutput := UninstallPackageOutput{}

	packageOutput.ResponseStatus = resp.Status
	if resp.StatusCode != 200 {
		if len(respByte) < 1 {
			log.Print("uninstall response status other than 200 with nil body")
			//packageOutput.ResponseStatus = resp.Status
			return &packageOutput, nil
		} else {
			log.Print("unistall response status other than 200")
			//packageOutput.ResponseStatus = resp.Status
			packageOutput.ResponseMessage = respByte
			return &packageOutput, nil
		}
	}
	//log.Print(string(respByte))
	//log.Print(resp.Status)

	if err = json.Unmarshal(respByte, &packageOutput); err != nil {
		log.Print("error in unmarshaling uninstall response body: ", err)
		return nil, err
	}
	//packageOutput.ResponseStatus = resp.Status
	return &packageOutput, nil
}
