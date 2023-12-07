package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AwaitResponse(method string, path string, token string, data interface{}) *http.Response {

	client := &http.Client{}
	var dataByte []byte = nil

	if data != nil {
		dataByte, _ = json.Marshal(data)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(dataByte))

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

func HandleRateLimiting(resp *http.Header) error {
	pool, err := strconv.Atoi(resp.Get("X-Auth-Pool-Size"))
	if err != nil {
		return err
	}
	used, err := strconv.Atoi(resp.Get("X-Auth-Pool-Used"))
	if err != nil {
		return err
	}
	replenishRate, err := strconv.ParseFloat(resp.Get("X-Auth-Replenished-Per-Second"), 10)
	if err != nil {
		return err
	}

	if pool <= used {
		println("RATE LIMIT REACHED : ON HOLD")
		time.Sleep(time.Duration(time.Duration(pool / int(replenishRate)).Seconds()))
	}

	return nil
}

// GET //
func GetDataList(path string, token string) []interface{} {

	resp := AwaitResponse("GET", path, token, nil)

	if err := HandleRateLimiting(&resp.Header); err != nil {
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

	resp := AwaitResponse("GET", path, token, nil)

	if err := HandleRateLimiting(&resp.Header); err != nil {
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

	resp := AwaitResponse("POST", path, token, model)

	if err := HandleRateLimiting(&resp.Header); err != nil {
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

func PostDataJson(path string, model interface{}, token string) map[string]interface{} {

	resp := AwaitResponse("POST", path, token, model)

	if err := HandleRateLimiting(&resp.Header); err != nil {
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
