package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID      int    `json:"id"`
	Payload string `json:"payload"`
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	task := Task{ID: 101, Payload: "секретные_данные_лабораторной_9"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func main() {
	http.HandleFunc("/next-task", getTaskHandler)
	println("Оркестратор запущен на :8082")
	http.ListenAndServe(":8082", nil)
}
