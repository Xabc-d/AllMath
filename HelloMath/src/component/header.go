package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Header struct {
	app.Compo
}

func (h *Header) Render() app.UI {
	return app.Div().Class("container-fluid").Body(
		app.Div().Class("row bg-primary").Style("min-height", "50px").Body(),
		app.Div().Class("row").Style("background-color", "#add8e6").Body(
			app.H1().Style("width", "90%").Style("margin", "0 auto").Body(
				app.Text("HelloMath"),
			),
		),
	)
}
