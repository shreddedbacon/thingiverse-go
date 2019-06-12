package api

import (
	"log"
)

func (thingiAPI *ThingiAPI) GetUser(userName string) (string, error) {
	responseData, err := MakeHTTPRequest("GET", baseURL, thingiAPI.appToken, "users/"+userName, "")
	if err != nil {
		log.Println(err)
	}
	return responseData, err
}
