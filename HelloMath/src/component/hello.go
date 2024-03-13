package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hello struct {
	app.Compo
	Sidebar Sidebar
	Name    string
}

func (h *Hello) Render() app.UI {
	return app.Main().Body(
		&Header{},
		app.Div().Class("row").Body(
			app.Div().Class("col-2").Body(
				&Sidebar{}),
			app.Div().Class("col-10").Body(
				&function{}),
		),
		//&Header{},
		//&h.Sidebar,
		//&Sidebar{},
		//&Calculator{},
		//&Week{Name: h.Sidebar.SelectedItem},
	)
}
