package models

type GalleryType struct {
	SiteName     string `json:"site_name"`
	SiteLocation string `json:"site_location"`
	Favicon      string `json:"favicon"`
}
type ListGalleryType []GalleryType
