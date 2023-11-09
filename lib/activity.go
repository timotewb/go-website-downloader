package lib

import (
	"encoding/json"
	m "gwd/models"
	"os"
)

func CheckIfJobRunning() (m.CheckJobIsRunningType, error) {

	var st m.SessionType
	var cjir m.CheckJobIsRunningType
	cjir.JobCount = 0

	// check if file exists
	if _, err := os.Stat("output.json"); err == nil {

		// read file in
		data, err := os.ReadFile("output.json")
		if err != nil {
			return cjir, err
		} else {

			// parse to struct
			err = json.Unmarshal(data, &st)
			if err != nil {
				return cjir, err
			} else {
				cjir.JobCount = st.JobCount
			}
		}
	} else {
		return cjir, nil
	}
	return cjir, nil

}