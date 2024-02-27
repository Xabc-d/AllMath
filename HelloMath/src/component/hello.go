package component

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hello struct {
	app.Compo
	Name string
}

func (h *Hello) Render() app.UI {
	return app.Main().Body(
		&Header{},
		&Sidebar{},
	)
}
