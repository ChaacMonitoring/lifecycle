package lifecycle_helpers

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type Config map[string]*string

var LifeCycleConfig Config

/* just a banner print */
func banner() {
	fmt.Println("**************************************************")
	fmt.Println("       _._  .__  .__       _  .   .   _		   .__ ")
	fmt.Println("  |     |   |_   |        / `  \\ /   / `  |   |   ")
	fmt.Println("  |     |   |    |--  >< |      |   |     |   |-- ")
	fmt.Println("  |__  _|_  |    |__      \\_,   |    \\_,  |__ |__ ")
	fmt.Println("")
	fmt.Println("**************************************************")
}

/* checking if you still wanna keep the lifecycle up */
func do_you_wanna_continue() {
	var input string
	for {
		fmt.Println("Do you wanna exit. (yes|no):\n\n")

		fmt.Scanf("%s", &input)

		if input == "yes" || input == "y" {
			break
		}
	}
}

/* config from flags */
func ConfigFromFlags() Config {
	var config Config
	config = make(Config)
	config["httpuri"] = flag.String("uri", "0.0.0.0", "what IP to Run HTTP Server at")
	config["httpport"] = flag.String("port", "8080", "what Port to Run HTTP Server at")
	config["db_uri"] = flag.String("db-uri", "0.0.0.0", "what IP to communicate for DB")
	config["db_req_port"] = flag.String("db-req-port", "9797", "what PORT to run ZMQ REQ at")
	config["db_rep_port"] = flag.String("db-rep-port", "9898", "what PORT to run ZMQ REP at")
	config["cpuprofile"] = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	return config
}

/*
putting together base engine for LifeCycle
*/
func LifeCycleEngine() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// remember it will be same DB instance shared across lifecycle package
	if *LifeCycleConfig["cpuprofile"] != "" {
		f, err := os.Create(*LifeCycleConfig["cpuprofile"])
		if err != nil {
			fmt.Println("Error: CPU Profile")
		}
		pprof.StartCPUProfile(f)
		go func() {
			time.Sleep(100 * time.Second)
			pprof.StopCPUProfile()
		}()
	}

	go LifeCycleHTTP()
}

/* LifeCycle Showcase */
func LifeCycle() {
	banner()
	LifeCycleConfig = ConfigFromFlags()
	LifeCycleEngine()
	do_you_wanna_continue()
}
