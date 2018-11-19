package main

import(
	"net/http"
	"github.com/vttrawick/asrc-site/router"
)

func main() {
	port := ":8080"

	http.HandleFunc("/cluster", router.Cluster)
	http.ListenAndServe(port, nil)
}
