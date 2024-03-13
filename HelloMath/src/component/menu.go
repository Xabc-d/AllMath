package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Sidebar struct {
	app.Compo
	SelectedItem string
}

func (s *Sidebar) onClick(ctx app.Context, e app.Event) {
	e.PreventDefault()
	v := ctx.JSSrc().Get("text").String()
	//select {}
	fmt.Println(v)
	s.SelectedItem = v
}

func (s *Sidebar) Render() app.UI {
	return app.Div().Class("container-fluid").Body(
		app.Div().Class("row ").Body(
			app.Div().Class("text-bg-light sidebar").Style("height", "100vh").Style("border-right", "1px solid rgba(0, 0, 0, .1)").Body(
				app.Ul().Class("nav flex-column").Body(
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("/week01").Body(
							app.Text("week01"),
						),
					),

					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("#").Body(),
					),

					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("#").Body(
							app.Text("week03"),
						),
					),
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("#").Body(
							app.Text("week04"),
						),
					),
					app.Li().Class("nav-item").Style("width", "50%").Style("margin", "0 auto").Body(
						app.A().OnClick(s.onClick).Class("nav-link active").Href("#").Body(
							app.Text("week05"),
						),
					),
				),
			),
			//app.Div().Class("col-9").Body(
			//	app.Div().Text("content0"),
			//),
		),
		//app.Div().Body(
		//	app.Text("content"),
		//),
	)
}
