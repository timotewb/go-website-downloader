package models

type SessionType struct {
	JobCount int           `json:"job_count"`
	JobData  []JobDataType `json:"job_data"`
}

type JobDataType struct {
	Url        string `json:"url"`
	FaviconURL string `json:"favicon_url"`
}