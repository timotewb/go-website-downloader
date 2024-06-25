package models

type DBType struct {
	Activity ActivityType `json:"activity"`
	Settings SettingsType `json:"settings"`
}

type ActivityType struct {
	ActivityCount int                `json:"count"`
	ActivityData  []ActivityDataType `json:"data"`
}

type ActivityDataType struct {
	Url        string `json:"url"`
	FaviconURL string `json:"favicon_url"`
	SessionID  string `json:"session_id"`
	StaleFlag  bool   `json:"stale_flag"`
}

type SettingsType struct {
	ContentDir       string `json:"content_dir"`
	ContentDirExists bool   `json:"content_dir_exists"`
	ContentDirWSPort int16  `json:"content_dir_wsport"`
}
