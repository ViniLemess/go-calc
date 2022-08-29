package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

var ErrorLogger = log.New(log.Writer(), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	router := httprouter.New()
	router.GET("/health", healthCheckHandler)
	router.GET("/calc/history", historyHandler)
	router.POST("/calc/:op/:x/:y", calcHandler)
	log.Println("Running at 'http://localhost:8090'")
	if err := http.ListenAndServe("0.0.0.0:8090", router); err != nil {
		log.Fatalln(err)
	}
}

func printResult(w http.ResponseWriter, result string, statusCode int) {
	w.WriteHeader(statusCode)
	_, err := fmt.Fprintln(w, result)
	if err != nil {
		ErrorLogger.Println(err)
		w.WriteHeader(500)
		return
	}
}

func calcHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	x, err := strconv.ParseFloat(ps.ByName("x"), 64)
	if err != nil {
		ErrorLogger.Println(err)
		printResult(w, "Bad Request : Path parameter should be a number but was: "+ps.ByName("x"), 400)
		return
	}
	y, err := strconv.ParseFloat(ps.ByName("y"), 64)
	if err != nil {
		ErrorLogger.Println(err)
		printResult(w, "Bad Request : Path parameter should be a number but was: "+ps.ByName("y"), 400)
		return
	}
	op := ps.ByName("op")

	result, err := calculate(op, x, y)
	if err != nil {
		ErrorLogger.Println(err)
		printResult(w, fmt.Sprintf("Bad Request : %s", err), 400)
		return
	}
	printResult(w, fmt.Sprint(result), 200)
}

func historyHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	printResult(w, fmt.Sprint(history), 200)
}

func healthCheckHandler(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	printResult(w, "Application is running fine!", 200)
}
