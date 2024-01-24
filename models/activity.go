package models

type CheckActivityType struct {
	JobCount     int                `json:"job_count"`
	ActivityData []ActivityDataType `json:"data"`
}
