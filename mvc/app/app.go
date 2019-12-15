package app

import (
	"net/http"

	"github.com/ngavrish/go_microservices_guide/mvc/controllers"
)

// StartApp method
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
