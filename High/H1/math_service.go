package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, _ := strconv.Atoi(nStr)
	result := fib(n)
	json.NewEncoder(w).Encode(map[string]int{"result": result})
}

func main() {
	http.HandleFunc("/fib", fibHandler)
	http.ListenAndServe(":8081", nil)
}
