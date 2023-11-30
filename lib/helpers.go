package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetData(path string, model interface{}) interface{} {

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

	var objBody interface{}

	json.Unmarshal(body, &objBody)
	return objBody
}
