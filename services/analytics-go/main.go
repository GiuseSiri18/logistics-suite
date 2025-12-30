package main

import (
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/api/analytics", func(w http.ResponseWriter, r *http.Request) {
        data := map[string]string{"status": "Go service is analyzing logs...", "load": "normal"}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(data)
    })
    http.ListenAndServe(":8080", nil)
}