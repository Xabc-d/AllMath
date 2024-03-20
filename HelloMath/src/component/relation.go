package component

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type relation struct {
	app.Compo
}

func (r *relation) Render() app.UI {
	return app.Div().Class("container-fluid").Body(
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.Div().Class("row").Body(
					app.H5().Class("col-auto").Text("inverse of relation"),
					app.Button().
						Class("col-auto btn btn-outline-primary").
						Style("width", "60px").
						Style("height", "20px").
						Style("padding", "0").
						Style("display", "flex").
						Style("justify-content", "center").
						Style("margin-top", "3px").
						Style("font-size", "0.7rem").
						Text("variation"),
				),
				app.P().Text("Suppose we have two sets A = {1,2,3} and B = {x,y}, and a relation R from A to B defined as follows: R= {(1,x),(2,y),(3,x)}. what is the inverse of this relation."),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					AutoFocus(true),
			),
		),
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
				app.Button().Class("btn btn-link").Text("Show Answer"),
			),
		),

		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.Div().Class("row").Body(
					app.H5().Class("col-auto").Text("composition of relations"),
					app.Button().
						Class("col-auto btn btn-outline-primary").
						Style("width", "60px").
						Style("height", "20px").
						Style("padding", "0").
						Style("display", "flex").
						Style("justify-content", "center").
						Style("margin-top", "3px").
						Style("font-size", "0.7rem").
						Text("variation"),
				),
				app.P().Text("Suppose we have three sets A = {1,2} and B = {x,y} and C = {α,β} and R ⊆ A×B as R = {(1,x),(2,y)} S ⊆ B×C as S = {(x,α),(y,β)}, what is composition R◦S "),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					AutoFocus(true),
			),
		),
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
				app.Button().Class("btn btn-link").Text("Show Answer"),
			),
		),

		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.Div().Class("row").Body(
					app.H5().Class("col-auto").Text("partial order"),
					app.Button().
						Class("col-auto btn btn-outline-primary").
						Style("width", "60px").
						Style("height", "20px").
						Style("padding", "0").
						Style("display", "flex").
						Style("justify-content", "center").
						Style("margin-top", "3px").
						Style("font-size", "0.7rem").
						Text("variation"),
				),
				app.P().Text("Suppose we have a set A= {1, 2, 3, 4} and R={(1,1),(1,2),(1,3)(1,4)(2,2)(2,3)(2,4)(3,3),(3,4),(4,4)}. Does it satisfy partial ordering"),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					AutoFocus(true),
				app.P().Text("Suppose we have a set A= {1, 2, 3, 4}, construct a relation that satisfies the partial order"),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					//Value(p.Answer1).
					AutoFocus(true),
			),
		),
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
				app.Button().Class("btn btn-link").Text("Show Answer"),
			),
		),

		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.Div().Class("row").Body(
					app.H5().Class("col-auto").Text("total order"),
					app.Button().
						Class("col-auto btn btn-outline-primary").
						Style("width", "60px").
						Style("height", "20px").
						Style("padding", "0").
						Style("display", "flex").
						Style("justify-content", "center").
						Style("margin-top", "3px").
						Style("font-size", "0.7rem").
						Text("variation"),
				),
				app.P().Text("Consider the set B = {x, y, z}, construct a total order relation R fot B. Does it satisfy partial ordering"),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					AutoFocus(true),
			),
			app.P().Text("Identify which are total order relationships"),
			app.Input().
				Type("text").
				Style("width", "200px").
				Style("height", "30px").
				Class("form-control mb-3").
				Placeholder("select your answer").
				AutoFocus(true),
		),
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
				app.Button().Class("btn btn-link").Text("Show Answer"),
			),
		),
	)
}

type theRelation struct {
	relation1, relation2 string
}

func Reflexive(set []string, relation map[theRelation]bool) bool {
	for _, element := range set {
		if !relation[theRelation{element, element}] {
			return false
		}
	}
	return true
}

func Antisymmetric(relation map[theRelation]bool) bool {
	for rela := range relation {
		if rela.relation1 != rela.relation2 && relation[theRelation{rela.relation2, rela.relation1}] {
			return false
		}
	}
	return true
}

func Transitive(relation map[theRelation]bool) bool {
	for rela1 := range relation {
		for rela2 := range relation {
			if rela1.relation2 == rela2.relation1 {
				if !relation[theRelation{rela1.relation1, rela2.relation2}] {
					return false
				}
			}
		}
	}
	return true
}

func TotallyOrdered(set []string, relation map[theRelation]bool) bool {
	for _, a := range set {
		for _, b := range set {
			if a != b {
				if !(relation[theRelation{a, b}] || relation[theRelation{b, a}]) {
					return false
				}
			}
		}
	}
	return true
}

func InverseOfRelation(relation map[theRelation]bool) map[theRelation]bool {
	inverse := make(map[theRelation]bool)
	for oriRelation, _ := range relation {
		inverse[theRelation{oriRelation.relation1, oriRelation.relation2}] = true
	}
	return inverse
}
