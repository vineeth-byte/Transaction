package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	dto "task/dto"
	"time"
)

var loc string

var Sample map[string][]*dto.Transaction

const layout = "2006-01-02T15:04:05Z"

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqbody, _ := ioutil.ReadAll(r.Body)
	var txn *dto.Transaction
	err := json.Unmarshal(reqbody, &txn)
	if err != nil {
		fmt.Println("Error while parsing body :", err)
		w.WriteHeader(400) //if the JSON is invalid
		return
	}
	timeformat, err := time.Parse(layout, txn.Timestamp)
	if err != nil {
		fmt.Println("Error while parsing date :", err)
		w.WriteHeader(422) // invalid date format or not to parsable
		return
	}
	t := time.Now().Format("2006-01-02T15:04:05Z")
	currentime, _ := time.Parse(layout, t)
	diff := currentime.Sub(timeformat).Seconds()
	if diff > 60 {
		fmt.Println("Transaction is older than 60 sec")
		w.WriteHeader(204) // the transaction is older than 60 seconds
		return
	}
	s := Sample[loc]
	s = append(s, txn)
	Sample[loc] = s
	w.WriteHeader(201) //in case of success
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	details := &dto.Details{}
	for _, val := range Sample[loc] {
		fmt.Println(val)
		timeformat, _ := time.Parse(layout, val.Timestamp)
		t := time.Now().Format("2006-01-02T15:04:05Z")
		currentime, _ := time.Parse(layout, t)
		diff := currentime.Sub(timeformat).Seconds()
		fmt.Println(diff)
		if diff <= 60 {
			fmt.Println("entered")
			amount, err := strconv.ParseFloat(val.Amount, 64)
			fmt.Println(err)
			if err == nil {
				details.Sum += float64(amount)
				if details.Max < float64(amount) || details.Max == 0 {
					details.Max = float64(amount)
				}
				if details.Min > float64(amount) || details.Min == 0 {
					details.Min = float64(amount)
				}
				details.Count++
			}
			details.Avg = details.Sum / details.Count
		}
	}
	fmt.Println(details)
	json.NewEncoder(w).Encode(details)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Sample = nil
	json.NewEncoder(w).Encode("Transaction Deleted successfully")
}

func SetLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var city dto.Location
	err = json.Unmarshal(body, &city)
	if err != nil {
		log.Println("error on unmarshalling json", err)
	}
	loc = city.City
	json.NewEncoder(w).Encode("Set Location Successfully")
}
