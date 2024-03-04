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
	return app.Div().Class("d-flex").Body(
		app.Div().Class("col-md-3 ").Body(
			app.Div().Class("row").Body(
				app.Div().Class("bg-light sidebar").Style("width", "100%").Style("flex-shrink", "0").Body(
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
			),
		),
		//),
		app.Div().Body(
			app.Text("content"),
		),
	)

}
