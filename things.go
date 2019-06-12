package api

import (
	"log"
)

func (thingiAPI *ThingiAPI) GetThing(thingId string) (string, error) {
	responseData, err := MakeHTTPRequest("GET", baseURL, thingiAPI.appToken, "things/"+thingId, "")
	if err != nil {
		log.Println(err)
	}
	return responseData, err
}

func (thingiAPI *ThingiAPI) GetThingFiles(thingId string) (string, error) {
	responseData, err := MakeHTTPRequest("GET", baseURL, thingiAPI.appToken, "things/"+thingId+"/files/", "")
	if err != nil {
		log.Println(err)
	}
	return responseData, err
}
