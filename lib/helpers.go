package lib

import (
	Mod "flx/model"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetListingVariantBySku(sku []string) interface{} {

	client := &http.Client{}

	param := Mod.GetListingVariant{
		Skus: sku,
	}

	jsonParam, err := json.Marshal(param)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("GET", Mod.FLX_URL+Mod.LISTING_URL_EXT+Mod.PLURAL_VARIANT_URL_EXT, bytes.NewBuffer(jsonParam))

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
