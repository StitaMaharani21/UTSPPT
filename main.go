package main

import (
	"datasimpan"
	"encoding/json"
	"log"
	"net/http"
)

func GetDataStruct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		Data := datasimpan.InputDataStruct()
		datastruct, err := json.Marshal(Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(datastruct)
		return
	}
}

func PostDataStruct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data datasimpan.DataStruct
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&data); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(data)
			return
		} else {
			data := datasimpan.InputDataStruct()
			var response datasimpan.DataStruct
			name := r.FormValue("Nama")
			for _, val := range data {
				if val.Nama == name {
					response = datasimpan.DataStruct{
						Nama:   val.Nama,
						NIM:    val.NIM,
						Alamat: val.Alamat,
					}
					json.NewEncoder(w).Encode(response)
					return
				}
			}
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
	}
}

func main() {
	http.HandleFunc("/getdatastruct", GetDataStruct)
	http.HandleFunc("/postdatastruct", PostDataStruct)
	log.Fatal(http.ListenAndServe(":5050", nil))
}
