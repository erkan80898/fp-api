package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func AwaitResponse(method string, path string, token string) *http.Response {

	client := &http.Client{}

	req, err := http.NewRequest(method, path, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// GET //
func GetDataList(path string, token string) []interface{} {

	resp := AwaitResponse("GET", path, token)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objBody []interface{}

	json.Unmarshal(body, &objBody)
	return objBody

}

func GetDataJson(path string, token string) map[string]interface{} {

	resp := AwaitResponse("GET", path, token)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objBody map[string]interface{}

	json.Unmarshal(body, &objBody)
	return objBody

}

// POST //
func PostDataList(path string, model interface{}, token string) []interface{} {

	resp := AwaitResponse("POST", path, token)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objBody []interface{}

	json.Unmarshal(body, &objBody)
	return objBody
}

func PostDataJson(path string, model interface{}, token string) map[string]interface{} {

	resp := AwaitResponse("POST", path, token)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objBody map[string]interface{}

	json.Unmarshal(body, &objBody)
	return objBody
}
