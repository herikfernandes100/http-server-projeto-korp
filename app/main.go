package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Erro ao gerar resposta", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	log.Println("Servidor iniciado na porta 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
