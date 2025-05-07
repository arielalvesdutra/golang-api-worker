package main

import (
	"api/usecase"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

)


func main() {
	http.HandleFunc("/contas", getContas)
	err := http.ListenAndServe(":8085", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} 
}


func getContas(w http.ResponseWriter, r *http.Request) {

	contas := usecase.ObterContas()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contas)
}

