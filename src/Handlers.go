package main

import (
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/redis.v3"
)

func IndexStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ms-sequence-generator is runnning! Awesome!!")
}

func SequenceIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sequenceKey := vars["sequenceKey"]
	client := getRedisClient()
	value, err := client.Get(sequenceKey).Int64()
	processSequence(w, r, value, err)
}

func SequenceCreation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sequenceKey := vars["sequenceKey"]
	initialValue := r.FormValue("initialValue")
	if (initialValue == "") {
		initialValue = "0"
	}
	client := getRedisClient()
	err := client.Get(sequenceKey).Err()
	if err == redis.Nil {
		value, errParse := strconv.ParseInt(initialValue, 10, 64)
		if (errParse != nil) {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Initial value must be an integer")
		} else {
			err := client.Set(sequenceKey, initialValue, 0).Err()
			if (err != nil) {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error with Redis")
				panic(err)
			} else {
				encodeSequence(w, r, Sequence{Number: value})
			}
		}
	} else if err != nil {
		 w.WriteHeader(http.StatusInternalServerError)
		 fmt.Fprintf(w, "Error with Redis")
		 panic(err)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "This sequence already exists")
		return
	}
}

func SequenceNext(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sequenceKey := vars["sequenceKey"]
	client := getRedisClient()
	value, err := client.Incr(sequenceKey).Result()
	processSequence(w, r, value, err)
}

func processSequence(w http.ResponseWriter, r *http.Request, value int64, err error) {
	if err != nil {
		if err == redis.Nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Resource Not Found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error with Redis")
		panic(err)
	} else {
		encodeSequence(w, r, Sequence{Number: value})
	}
}

func encodeSequence(w http.ResponseWriter, r *http.Request, s Sequence) {
	if err := json.NewEncoder(w).Encode(s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error encoding sequence")
		panic(err)
		return
	}
}

