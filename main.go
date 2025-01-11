package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sakshamagrawal07/cache-proxy-server/proxy"
)

func main() {

	PORT := flag.Int("port", 0, "Define desired port for the proxy server to run.")
	ORIGIN := flag.String("origin", "", "Define the URL of the server to which the requests will be forwarded.")

	flag.Parse()

	fmt.Println("PORT : ", *PORT)
	fmt.Println("PORT : ", *ORIGIN)

	if *ORIGIN != "" && *PORT != 0 {
		proxy := proxy.NewProxy(*ORIGIN)

		http.Handle("/", proxy)
		http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {}) //to avoid the browser making 2 requests

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *PORT), nil))
	} else {
		fmt.Println("No parameter passed. Use --port and --origin to start the proxy server.")
		flag.Usage()
	}

	fmt.Println("tst")
}
