package main

import (
	"HelloMath/src/component"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

func main() {
	// Components routing:
	app.Route("/", &component.Hello{Name: "123"})
	app.Route("/hello", &component.Hello{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Icon: app.Icon{
			Default:    "/web/logo.png",       // Specify default favicon.
			AppleTouch: "/web/logo-apple.png", // Specify icon on IOS devices.
		},
		Styles: []string{
			"./web/bootstrap/dist/css/bootstrap.css", // Loads hello.css file.
		},
		Scripts: []string{
			"/web/bootstrap/dist/bootstrap.bundle.js", // Loads hello.js file.
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

}
