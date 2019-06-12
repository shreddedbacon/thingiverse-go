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

func (thingiAPI *ThingiAPI) GetPopular() (string, error) {
	responseData, err := MakeHTTPRequest("GET", baseURL, thingiAPI.appToken, "popular/", "")
	if err != nil {
		log.Println(err)
	}
	return responseData, err
}

func (thingiAPI *ThingiAPI) GetFeatured(returnComplete bool) (string, error) {
	query := ""
	if returnComplete {
		query = "{\"return\":\"complete\"}"
	}
	responseData, err := MakeHTTPRequest("GET", baseURL, thingiAPI.appToken, "featured/", query)
	if err != nil {
		log.Println(err)
	}
	return responseData, err
}
