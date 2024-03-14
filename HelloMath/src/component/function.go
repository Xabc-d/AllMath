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
	a        string
	b        string
	c        string
	result   string
	theInput string
	list     []string
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
			app.P().Text("Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B."),
		),
		app.Button().
			Text("solution").
			OnClick(p.handleMapping),
		app.H5().Text(p.handleMapping),

		app.Div().Class("row").Body(
			app.P().Text("Given two sets A={1,2} B={x,y}, list all possible functions from set A to set B."),
		),
		app.Input().
			Type("p").
			Value(p.theInput).
			OnChange(p.ValueTo(p.theInput)),

		app.A().Class("btn btn-outline-primary btn-sm").Text("Solution"),

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
