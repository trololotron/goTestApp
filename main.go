package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/hello/{fullName}", api.Hello).Methods("GET")
	router.HandleFunc("/api/add/{a}/{b}", api.add).Methods("GET")
	router.HandleFunc("/api/sub/{a}/{b}", api.sub).Methods("GET")
	router.HandleFunc("/api/mul/{a}/{b}", api.mul).Methods("GET")
	router.HandleFunc("/api/div/{a}/{b}", api.div).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println(err)
	}
}
