package api

import (
	"time"
)

type Things []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	PublicURL string `json:"public_url"`
	Thumbnail string `json:"thumbnail"`
	Creator   struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		URL       string `json:"url"`
		PublicURL string `json:"public_url"`
		Thumbnail string `json:"thumbnail"`
	} `json:"creator"`
	IsPrivate   bool `json:"is_private"`
	IsPurchased bool `json:"is_purchased"`
	IsPublished bool `json:"is_published"`
}

type ThingFiles []struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Size         int    `json:"size"`
	URL          string `json:"url"`
	PublicURL    string `json:"public_url"`
	DownloadURL  string `json:"download_url"`
	ThreejsURL   string `json:"threejs_url"`
	Thumbnail    string `json:"thumbnail"`
	DefaultImage struct {
		ID    int    `json:"id"`
		URL   string `json:"url"`
		Name  string `json:"name"`
		Sizes []struct {
			Type string `json:"type"`
			Size string `json:"size"`
			URL  string `json:"url"`
		} `json:"sizes"`
		Added time.Time `json:"added"`
	} `json:"default_image"`
	Date          string        `json:"date"`
	FormattedSize string        `json:"formatted_size"`
	MetaData      []interface{} `json:"meta_data"`
	DownloadCount int           `json:"download_count"`
}
