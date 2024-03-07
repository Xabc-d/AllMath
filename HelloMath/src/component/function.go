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
	a      string
	b      string
	c      string
	result string
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
func printMapping(mapping []map[Element]Element) {
	for i, theMapping := range mapping {
		print("f", i+1, ":")
		for a, b := range theMapping {
			fmt.Printf("%s -> %s ", a, b)
		}
		fmt.Println()
	}
}

func (p *function) Render() app.UI {
	return app.Div().Class("container").Body(
		app.H1().Class("text-center").Text("function"),
		app.Div().Class("card").Body(
			app.Div().Class("card-body").Body(
				app.H5().Class("card-title").Text("1."),
				app.P().Class("card-text").Text(""),
				app.A().Class("btn btn-primary").Text("Solution"),
			),
		),
		app.Div().Class("card").Body(
			app.Div().Class("card-body").Body(
				app.H5().Class("card-title").Text("Injective"),
				app.P().Class("card-text").Text("is injective or not"),
				app.A().Class("btn btn-primary").Text("Solution"),
			),
		),
		app.Div().Class("card").Body(
			app.Div().Class("card-body").Body(
				app.H5().Class("card-title").Text("Surjective"),
				app.P().Class("card-text").Text(""),
				app.A().Class("btn btn-primary").Text("Solution"),
			),
		),
	)
}