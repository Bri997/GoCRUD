package main

import (
	"encoding/json"
	"fmt"
	"log"

	"math/rand"
	"net/http"

	// "strconv"
	"github.com/gorilla/mux"
)

type Produce struct {
	ID string `json: "id"`
	Name string `json: "name"`
	Price string `json: "price"`
}

var allProduce []Produce

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq() string {
    b := make([]rune, 16)

    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func getAllProduce(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allProduce)
}

func getProduce(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range allProduce {
		if item.ID == params["ID"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createProduce(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var newProduce Produce
	_ = json.NewDecoder(r.Body).Decode(&newProduce)
	newProduce.ID = randSeq()
	allProduce = append(allProduce, newProduce)
}

func updateProduce(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range allProduce {
		if item.ID == params["ID"]{
			allProduce = append(allProduce[:index], allProduce[index+1:]...)
			var produce Produce
			_ = json.NewDecoder(r.Body).Decode(&produce)
			produce.ID = params["ID"]
			allProduce = append(allProduce, produce)
			json.NewEncoder(w).Encode(produce)
			return
		}
	}
	
}

func deleteProduce(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	 
	for index, item := range allProduce{
		if item.ID == params["ID"]{
			allProduce = append(allProduce[:index], allProduce[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(allProduce)

}

func main() {
	r := mux.NewRouter()

	allProduce = append(allProduce, Produce{
		ID: "1234-5678-9012",
		Name: "Green Bean",
		Price: "1.57",
	})
		allProduce = append(allProduce, Produce{
		ID: "ABCD-1234-5678",
		Name: "Lettuce",
		Price: "1.00",
	})

	r.HandleFunc("/allproduce", getAllProduce).Methods("GET")
	r.HandleFunc("/allproduce/{ID}", getProduce).Methods("GET")
	r.HandleFunc("/allproduce", createProduce).Methods("POST")
	r.HandleFunc("/allproduce/{ID}", updateProduce).Methods("PUT")
	r.HandleFunc("/allproduce/{ID}", deleteProduce).Methods(("DELETE"))

	fmt.Printf("Starting server on port 8080...\n", )
	log.Fatal(http.ListenAndServe(":8080", r))


}