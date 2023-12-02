package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GET //

func GetDataList(path string, token string) []interface{} {
	client := &http.Client{}

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

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
	client := &http.Client{}

	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

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

	client := &http.Client{}

	jsonParam, err := json.Marshal(model)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(jsonParam))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", token)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objBody []interface{}

	json.Unmarshal(body, &objBody)
	return objBody
}

func PostDataJson(path string, model interface{}) map[string]interface{} {

	client := &http.Client{}

	jsonParam, err := json.Marshal(model)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(jsonParam))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-TOKEN", os.Getenv("FLX_API_TOKEN"))

	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var objBody map[string]interface{}

	json.Unmarshal(body, &objBody)
	return objBody
}
