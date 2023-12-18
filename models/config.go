package models

type ConfigType struct {
	AppDB         string `json:"app_db"`
	AppName       string `json:"app_name"`
	AppDir        string `json:"app_dir"`
	AppDBFileName string `json:"app_db_filename"`
	SessionID     string `json:"session_id"`
}
