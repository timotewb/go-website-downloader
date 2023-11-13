package lib

// import (
// 	"encoding/json"
// 	"fmt"
// 	m "gwd/models"
// 	"os"
// )

// func StartGetSiteJob(r ResponseType) {
// 	var d m.ActivityType
// 	d.ActivityCount = 1
// 	var d1 m.ActivityDataType
// 	d1.FaviconURL = r.FaviconURL
// 	d1.Url = r.Url
// 	d.ActivityData = append(d.ActivityData, d1)

// 	dJSON, _ := json.Marshal(d)

// 	if _, err := os.Stat("db.json"); err == nil {
// 		fmt.Println("File exists!" + r.Url)
// 	} else{
// 		fmt.Println("File NOT found.")
// 		err = os.WriteFile("output.json", dJSON, 0644)
// 		if err == nil {
// 			fmt.Println("File written.")
// 		} else {
// 			fmt.Println("File NOT written.")
// 		}
// 	}
// }