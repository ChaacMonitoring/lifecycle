package lifecycle_helpers

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"time"

	golerror "github.com/abhishekkr/gol/golerror"
)

/*
Lifecycle Handler for HTTP Requests
*/
func LifeCycleHTTP() {
	httpport_num, err_httpport := strconv.Atoi(*LifeCycleConfig["httpport"])
	if err_httpport != nil {
		golerror.Boohoo("Port parameters to bind, error-ed while conversion to number.", true)
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", Index)
	http.HandleFunc("/data", Data)
	http.HandleFunc("/help", F1)
	http.HandleFunc("/status", Status)

	srv := &http.Server{
		Addr:        fmt.Sprintf("%s:%d", *LifeCycleConfig["httpuri"], httpport_num),
		Handler:     http.DefaultServeMux,
		ReadTimeout: time.Duration(5) * time.Second,
	}

	fmt.Printf("access your lifecycle at http://%s:%d\n", *LifeCycleConfig["httpuri"], httpport_num)
	err := srv.ListenAndServe()
	fmt.Println("Game Over:", err)
}
