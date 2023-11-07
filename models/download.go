package models

type ResponseType struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Url        string `json:"url"`
	FaviconURL string `json:"favicon_url"`
}