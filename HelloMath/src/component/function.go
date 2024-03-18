package component

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//1. Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B.
//solution:
//
//f1: 1→x, 2→x
//f2: 1→x, 2→y
//f3: 1→y, 2→x
//f4: 1→y, 2→y

type Element string

type function struct {
	app.Compo
	a          string
	b          string
	c          string
	Answer1    string
	Answer2    string
	Answer3    string
	theInput   string
	list       []string
	solutionQ1 bool
}

// A := []Element{"a", "b"}
// B := []Element{"1", "2", "3"}
// theMapping := setMapping(A, B, []map[Element]Element{}, 0, make(map[Element]Element))
// printMapping(theMapping)
func setMapping(A, B []Element, theMapping []map[Element]Element, index int, currentMapping map[Element]Element) []map[Element]Element {
	if index == len(A) {
		allMapping := make(map[Element]Element)
		for k, v := range currentMapping {
			allMapping[k] = v
		}
		theMapping = append(theMapping, allMapping)
		return theMapping
	}

	for _, b := range B {
		currentMapping[A[index]] = b
		theMapping = setMapping(A, B, theMapping, index+1, currentMapping)
		delete(currentMapping, A[index])
	}
	return theMapping
}

func (p *function) printMapping(mapping []map[Element]Element) {
	for i, theMapping := range mapping {
		print("f", i+1, ":")
		for a, b := range theMapping {
			fmt.Printf("%s -> %s ", a, b)
		}
		fmt.Println()
	}
}

func (p *function) handleMapping(ctx app.Context, e app.Event) {
	A := []Element{"a", "b"}
	B := []Element{"1", "2", "3"}
	theMapping := setMapping(A, B, []map[Element]Element{}, 0, make(map[Element]Element))
	p.printMapping(theMapping)
	p.Update()
}

func injective(f func(int) int, domain []int) bool {
	theMap := make(map[int]bool)
	for _, value := range domain {
		result := f(value)
		if theMap[result] {
			return false
		}
		theMap[result] = true
	}
	return true
}

func surjective(f func(int) int, domain []int, codomain []int) bool {
	theMap := make(map[int]bool)
	for _, value := range codomain {
		theMap[value] = false
	}
	for _, value := range domain {
		result := f(value)
		theMap[result] = true
	}
	for _, value := range theMap {
		if !value {
			return false
		}
	}
	return true
}

// y=kx+b
func inverseLinear(k, b, y int) int {
	return (y - b) / k
}

//

func preimage() {

}

func composition() {

}

func (p *function) Render() app.UI {
	return app.Div().Class("container-fluid").Body(
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.Div().Class("row").Body(
					app.H5().Class("col-auto").Text("Question"),
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
				app.P().Text("Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B"),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					Value(p.Answer1).
					AutoFocus(true).
					OnChange(p.InputChangeQ1),
				app.Button().
					Class("btn btn-outline-primary btn-sm").
					Text("Submit Answer").
					OnClick(p.SubmitQ1),
			),
		),
		app.Div().Class("row").Body(
			app.Div().Class("col-12").Style("padding", "0").Style("margin", "0").Body(
				app.Button().Class("btn btn-link").Text("Show Answer").OnClick(p.SolutionQ1),
			),
			p.renderSolutionQ1(),
		),

		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.H5().Class("mb-3").Text("Question"),
				app.P().Text("Consider the function 𝑓:{0, 1, … ,9} → {} f(x) = 3x+1"),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					Value(p.Answer2).
					AutoFocus(true).
					OnChange(p.InputChangeQ2),
				app.Button().
					Class("btn btn-outline-primary btn-sm").
					Text("Submit Answer").
					OnClick(p.SubmitQ2),
			),
		),
		app.Div().Class("row").Body(
			app.Div().Class().Style("padding", "0").Style("margin", "0").Body(
				app.Button().Class("btn btn-link").Text("Show Answer").OnClick(p.SolutionQ2),
			),
		),

		app.Div().Class("row").Body(
			app.Div().Class("col-12").Body(
				app.H5().Class("mb-3").Text("Question"),
				app.P().Text("What is the?"),
				app.Input().
					Type("text").
					Style("width", "200px").
					Style("height", "30px").
					Class("form-control mb-3").
					Placeholder("Enter your answer").
					Value(p.Answer3).
					AutoFocus(true).
					OnChange(p.InputChangeQ3),
				app.Button().
					Class("btn btn-outline-primary btn-sm").
					Text("Submit Answer").
					OnClick(p.SubmitQ3),
			),
		),
		app.Div().Class("row").Style("padding", "0").Style("margin", "0").Body(
			app.Div().Class("col-12").Body(
				app.Button().Class("btn btn-link").Text("Show Answer").OnClick(p.SolutionQ3),
			),
		),

		//app.Div().Class("card").Body(
		//	app.Div().Class("card-body").Body(
		//		app.H5().Class("card-title").Text("Injective"),
		//		app.P().Class("card-text").Text("is injective or not"),
		//		app.A().Class("btn btn-primary").Text("Solution"),
		//	),
		//),
		//app.Div().Class("card").Body(
		//	app.Div().Class("card-body").Body(
		//		app.H5().Class("card-title").Text("Surjective"),
		//		app.P().Class("card-text").Text(""),
		//		app.A().Class("btn btn-primary").Text("Solution"),
		//	),
		//),
	)
}

func (p *function) InputChangeQ1(ctx app.Context, e app.Event) {
	p.Answer1 = ctx.JSSrc().Get("value").String()
}
func (p *function) SubmitQ1(ctx app.Context, e app.Event) {
}
func (p *function) SolutionQ1(ctx app.Context, e app.Event) {
	p.solutionQ1 = !p.solutionQ1
	p.Update()
}

func (p *function) renderSolutionQ1() app.UI {
	if p.solutionQ1 {
		return app.Div().Text("Solution")
	}
	return nil
}

func (p *function) InputChangeQ2(ctx app.Context, e app.Event) {
	p.Answer2 = ctx.JSSrc().Get("value").String()
}

func (p *function) SubmitQ2(ctx app.Context, e app.Event) {
}
func (p *function) SolutionQ2(ctx app.Context, e app.Event) {
}

func (p *function) InputChangeQ3(ctx app.Context, e app.Event) {
	p.Answer3 = ctx.JSSrc().Get("value").String()
}
func (p *function) SubmitQ3(ctx app.Context, e app.Event) {
}
func (p *function) SolutionQ3(ctx app.Context, e app.Event) {
}
