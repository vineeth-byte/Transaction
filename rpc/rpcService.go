package rpc

import (
	"log"
	"net/http"
	dto "task/dto"
	"task/service"

	"github.com/gorilla/mux"
)

func Connect() {
	service.Sample = map[string][]*dto.Transaction{}
	r := mux.NewRouter()
	r.HandleFunc("/postTransaction", service.PostTransaction).Methods("POST")
	r.HandleFunc("/getTransactions", service.GetTransactions).Methods("GET")
	r.HandleFunc("/deleteTransaction", service.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/setLocation", service.SetLocation).Methods("POST")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Println(err)
	}
}
