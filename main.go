package main

import (
	"bytes"
	"crypto/tls"
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ClientCredentials struct {
	Secret string `json:"client_secret"`
	Id     string `json:"client_id"`
}

//
// client_secret, clent_id를 생성하여 base64 타입으로 반환한다.

func GetClientCredentials(url, id, pw string) (string, error) {
	// curl -k -X POST https://localhost:8243/api-application-registration/register -H 'authorization: Basic YWRtaW46YWRtaW4=' -H 'content-type: application/json' -d '{ "applicationName":"appName", "tags":["device_management"]}'
	//
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	body := bytes.NewBufferString(`{ "applicationName":"appName", "tags":["device_management"]}`)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(id, pw)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// Response 체크.
	respBody, err := ioutil.ReadAll(response.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

	var cc ClientCredentials
	json.Unmarshal(respBody, &cc)
	fmt.Printf("%+v\n", cc)
}

func main() {
	encCC := EncodeClientCredentials("amdin", "B1ueb!rd")
	fmt.Println(encCC)

	// curl -k -X POST https://localhost:8243/api-application-registration/register -H 'authorization: Basic YWRtaW46YWRtaW4=' -H 'content-type: application/json' -d '{ "applicationName":"appName", "tags":["device_management"]}'
	//
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: tr,
	}

	body := bytes.NewBufferString(`{ "applicationName":"appName", "tags":["device_management"]}`)
	req, err := http.NewRequest("POST", "https://52.41.242.69:8243/api-application-registration/register", body)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth("admin", "B1ueb!rd")

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// Response 체크.
	respBody, err := ioutil.ReadAll(response.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}

	var cc ClientCredentials
	json.Unmarshal(respBody, &cc)
	fmt.Printf("%+v\n", cc)
}
