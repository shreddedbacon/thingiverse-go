package api

import (
	"log"
)

func (thingiAPI *ThingiAPI) GetNewest() (string, error) {
	responseData, err := MakeHTTPRequest("GET", baseURL, thingiAPI.appToken, "newest/", "")
	if err != nil {
		log.Println(err)
	}
	return responseData, err
}
