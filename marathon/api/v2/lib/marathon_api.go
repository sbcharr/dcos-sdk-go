package lib

import (
	"fmt"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strings"

	"github.com/PTC-GLOBAL/dcos-sdk-go/client"
)


const (
	TASK_STAGING int = iota
	TASK_RUNNING
	TASK_COMPLETE
	TASK_FAILED
	TASK_OTHER
)

type Deployments struct {
	ID string `json:"id"`
}

type Tasks struct {
	ID string `json:"id"`
	State string `json:"state"`
}

type TaskStatus struct {
	Dpl Deployments `json:"deployments"`
	TasksStaged int `json:"tasksStaged"`
	TasksRunning int `json:"tasksRunning"`
	TasksHealthy int `json:"tasksHealthy"`
	TasksUnhealthy int `json:"tasksUnhealthy"`
	Task Tasks `json:"tasks"`
	Message string `json:"message"`
}

func InstallPackageStatus(appId, accessToken, dcosUrl string) (int, error) {
	httpCli := client.HTTPClient()

	url := fmt.Sprintf("%s/marathon/v2/apps/%s", dcosUrl, appId) //ask user to pass dcos IP/Domain name
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		log.Print("error creating a new request", err)
		return TASK_OTHER, err
	}

	req.Header.Set("Authorization", "token=" + accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpCli.Do(req)
	if err != nil {
		log.Print("error sending request: ", err)
		return TASK_OTHER, err
	}
	defer resp.Body.Close()

	//TODO what happens in case of network disconnection, how does client know about the status of the task submitted

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("error in reading response body: ", err)
		return TASK_OTHER, err
	}

	t := TaskStatus{}
	if err = json.Unmarshal(respByte, &t); err != nil {
		log.Print("error unmarshaling response body; PackageUninstallStatus(): ", err)
		return TASK_OTHER, err
	}
	if t.Task.State == "TASK_STAGING" {
		return TASK_RUNNING, nil
	} else if t.Task.State == "TASK_RUNNING" {
		return TASK_COMPLETE, nil
	}
	return TASK_OTHER, err
}


func UninstallPackageStatus(appId, accessToken, dcosUrl string) (int, error) {
	httpCli := client.HTTPClient()

	url := fmt.Sprintf("%s/marathon/v2/apps/%s", dcosUrl, appId) //ask user to pass dcos IP/Domain name
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		log.Print("error creating a new request", err)
		return TASK_OTHER, err
	}

	req.Header.Set("Authorization", "token=" + accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpCli.Do(req)
	if err != nil {
		log.Print("error sending request: ", err)
		return TASK_OTHER, err
	}
	defer resp.Body.Close()

	//TODO what happens in case of network disconnection, how does client know about the status of the task submitted

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("error in reading response body: ", err)
		return TASK_OTHER, err
	}

	t := TaskStatus{}
	if err = json.Unmarshal(respByte, &t); err != nil {
		log.Print("error unmarshaling response body; PackageUninstallStatus(): ", err)
		return TASK_OTHER, err
	}
	if t.Task.State == "TASK_RUNNING" {
		return TASK_RUNNING, nil
	} else if strings.Contains(t.Message, "does not exist") {
		return TASK_COMPLETE, nil
	}
	return TASK_OTHER, err
}
