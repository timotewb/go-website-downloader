package models

type DBType struct {
	Activity ActivityType `json:"activity"`
}

type ActivityType struct {
	ActivityCount int                `json:"count"`
	ActivityData  []ActivityDataType `json:"data"`
}

type ActivityDataType struct {
	Url        string `json:"url"`
	FaviconURL string `json:"favicon_url"`
}