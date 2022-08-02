package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func printResult(w http.ResponseWriter, result any) {
	_, err := fmt.Fprintln(w, result)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func validatePathParam(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(400)
		printResult(w, "Bad Request : "+err.Error())
		return
	}
}

func CalcHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	x, err := strconv.ParseFloat(ps.ByName("x"), 64)
	validatePathParam(w, err)
	y, err := strconv.ParseFloat(ps.ByName("y"), 64)
	validatePathParam(w, err)
	op := ps.ByName("op")

	if err != nil {
		return
	}

	if y == 0 && op == "div" {
		w.WriteHeader(400)
		printResult(w, "Illegal Operation: Divider cannot be 0")
		return
	}

	var result float64

	switch op {
	case "sum":
		result = Sum(x, y)
	case "sub":
		result = Sub(x, y)
	case "mul":
		result = Mul(x, y)
	case "div":
		result, _ = Div(x, y)
	default:
		w.WriteHeader(400)
		printResult(w, "Bad Request : Unsupported Operation : "+op)
		return
	}
	printResult(w, result)
}

func HistoryHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	printResult(w, history)
}

func main() {
	router := httprouter.New()
	router.GET("/calculator/history", HistoryHandler)
	router.GET("/calc/:op/:x/:y", CalcHandler)
	log.Println("Running at 'http://localhost:8080'")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", router))
}
