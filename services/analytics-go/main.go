package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{
        "status": "running",
        "service": "Go Analytics Engine",
        "processing_time": "12ms",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    })
    http.ListenAndServe(":8080", nil)
}