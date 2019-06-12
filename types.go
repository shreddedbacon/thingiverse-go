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

type Thing struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	URL       string `json:"url"`
	PublicURL string `json:"public_url"`
	Creator   struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		URL       string `json:"url"`
		PublicURL string `json:"public_url"`
		Thumbnail string `json:"thumbnail"`
	} `json:"creator"`
	Added        time.Time `json:"added"`
	Modified     time.Time `json:"modified"`
	IsPublished  bool      `json:"is_published"`
	IsWip        bool      `json:"is_wip"`
	IsFeatured   bool      `json:"is_featured"`
	LikeCount    int       `json:"like_count"`
	IsLiked      bool      `json:"is_liked"`
	CollectCount int       `json:"collect_count"`
	IsCollected  bool      `json:"is_collected"`
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
	Description      string `json:"description"`
	Instructions     string `json:"instructions"`
	DescriptionHTML  string `json:"description_html"`
	InstructionsHTML string `json:"instructions_html"`
	Details          string `json:"details"`
	DetailsParts     []struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Required string `json:"required,omitempty"`
		Data     []struct {
			Content string `json:"content"`
		} `json:"data,omitempty"`
	} `json:"details_parts"`
	License           string      `json:"license"`
	FilesURL          string      `json:"files_url"`
	ImagesURL         string      `json:"images_url"`
	LikesURL          string      `json:"likes_url"`
	AncestorsURL      string      `json:"ancestors_url"`
	DerivativesURL    string      `json:"derivatives_url"`
	TagsURL           string      `json:"tags_url"`
	CategoriesURL     string      `json:"categories_url"`
	FileCount         int         `json:"file_count"`
	LayoutCount       int         `json:"layout_count"`
	LayoutsURL        string      `json:"layouts_url"`
	IsPrivate         bool        `json:"is_private"`
	IsPurchased       bool        `json:"is_purchased"`
	InLibrary         bool        `json:"in_library"`
	PrintHistoryCount int         `json:"print_history_count"`
	AppID             interface{} `json:"app_id"`
	DownloadCount     int         `json:"download_count"`
	ViewCount         int         `json:"view_count"`
}
