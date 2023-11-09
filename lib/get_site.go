package lib

import (
	"encoding/json"
	"fmt"
	m "gwd/models"
	"os"
)

func StartGetSiteJob(r ResponseType) {
	var d m.SessionType
	d.JobCount = 1
	var d1 m.JobDataType
	d1.FaviconURL = "www.favicon.com"
	d1.Url = "www.url.com"
	d.JobData = append(d.JobData, d1)

	dJSON, _ := json.Marshal(d)

	if _, err := os.Stat("session.json"); err == nil {
		fmt.Println("File exists!" + r.Url)
	} else{
		fmt.Println("File NOT found.")
		err = os.WriteFile("output.json", dJSON, 0644)
		if err == nil {
			fmt.Println("File written.")
		} else {
			fmt.Println("File NOT written.")
		}
	}
}