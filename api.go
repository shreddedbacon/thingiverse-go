package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
  "fmt"
)

// globals
var (
	baseURL           = "https://api.thingiverse.com/"
	availableLicenses = [9]string{"cc", "cc-sa", "cc-nd", "cc-nc-sa", "cc-nc-nd", "pd0", "gpl", "lgpl", "bsd"}
	netClient         = &http.Client{
		Timeout: time.Second * 10,
	}
)

type thingiAPI interface {
	GetUser(string) (string, error)
	GetThing(string) (string, error)
	GetThingFiles(string) (string, error)
	GetNewest() (string, error)
}

type ThingiAPI struct {
	clientId     string
	clientSecret string
	appToken     string
}

func NewAPI(clientId string, clientSecret string, appToken string) thingiAPI {
	return &ThingiAPI{
		clientId:     clientId,
		clientSecret: clientSecret,
		appToken:     appToken,
	}
}

func main() {
	// example usage
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	appToken := os.Getenv("APPTOKEN")

	api := NewAPI(clientId, clientSecret, appToken)
	userInfo, err := api.GetUser("shreddedbacon")
	if err != nil {
		log.Println(err)
	}
	log.Println(userInfo)
	codeInfo, err := api.GetNewest()
	if err != nil {
		log.Println(err)
	}
	thingiThings := Things{}
	json.Unmarshal([]byte(codeInfo), &thingiThings)
	for _, v := range thingiThings {
		thingInfo, err := api.GetThingFiles(strconv.Itoa(v.ID))
		if err != nil {
			log.Println(err)
		}
		thingiFiles := ThingFiles{}
		json.Unmarshal([]byte(thingInfo), &thingiFiles)
		for _, f := range thingiFiles {
			log.Println("ThingID:", v.ID, "FileName:", f.Name, "DownloadURL:", f.DownloadURL, "CDN:", f.DefaultImage.URL)
		}
	}
}

func MakeHTTPRequest(httpMethod string, apiURL string, accessToken string, apiPath string, queryParam string) (string, error) {
	apiRequest, apiRequestErr := http.NewRequest(httpMethod, apiURL+apiPath, bytes.NewBuffer([]byte(queryParam)))
	if apiRequestErr != nil {
		return "", apiRequestErr
	}
	apiRequest.Header.Add("Authorization", "Bearer "+accessToken)
	apiRequest.Header.Set("Accept", "application/json")
	apiResponse, apiResponseErr := netClient.Do(apiRequest)
	if apiResponseErr != nil {
		return "", apiResponseErr
	}
	defer apiResponse.Body.Close()
	apiResponseBody, _ := ioutil.ReadAll(apiResponse.Body)
	apiReturnBody := string(apiResponseBody)
	if apiResponse.StatusCode != 200 {
		errorMsg := fmt.Sprintf("%v: %v", apiResponse.StatusCode, "Error returned from API")
		return apiReturnBody, errors.New(errorMsg)
	}
	return apiReturnBody, nil
}
