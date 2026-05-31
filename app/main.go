package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// MÉTRICA: contador de requisições
var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requisições HTTP",
	},
	[]string{"endpoint"},
)

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func init() {
	prometheus.MustRegister(httpRequests)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.WithLabelValues("/projeto-korp").Inc()

	response := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DISPONIBILIDADE
func healthHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.WithLabelValues("/health").Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.HandleFunc("/health", healthHandler)

	// endpoint de métricas (Prometheus)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor iniciado na porta 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
