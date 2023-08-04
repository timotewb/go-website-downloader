package lib

import (
	"fmt"
	"runtime"
	"time"
)

func Logger(s string) {
	// _, file, no, ok := runtime.Caller(1)
    // if ok {
    //     fmt.Printf("called from %s#%d - %s\n", file, no, s)
    // }
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		currentTime := time.Now()
		fmt.Printf("%s | %s: %s\n",currentTime.Format("2006-01-02 15:04:05"), details.Name(), s)
	}
}