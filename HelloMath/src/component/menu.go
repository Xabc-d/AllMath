package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sidebar struct {
	app.Compo
}

func (s *Sidebar) Render() app.UI {
	return app.Div().Class("d-flex flex-column").Style("width", "20%").Body(

		app.Div().Class("row").Body(
			app.Div().Class("bg-light sidebar").Style("width", "100%").Style("flex-shrink", "0").Body(
				app.Ul().Class("nav flex-column").Body(
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().Class("nav-link active").Href("#").Body(
							app.Text("week01"),
						),
					),
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().Class("nav-link active").Href("#").Body(
							app.Text("week02"),
						),
					),
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().Class("nav-link active").Href("#").Body(
							app.Text("week03"),
						),
					),
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().Class("nav-link active").Href("#").Body(
							app.Text("week04"),
						),
					),
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().Class("nav-link active").Href("#").Body(
							app.Text("week05"),
						),
					),
				),
				//app.Div().Style("width", "30%").Body(),
			),
		),
	)
}
