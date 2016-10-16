package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//AppList struct
type AppList struct {
	Applications []struct {
		Application  string `json:"application"`
		Integrations []struct {
			Target        string `json:"target"`
			IntegrationID string `json:"integration_id"`
		} `json:"integrations"`
	} `json:"applications"`
}

//LoadMain func
func LoadMain(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "index.html")

	//file, err := ioutil.ReadFile("index.html")

	//b := bytes.NewBuffer(file)
	//if err != nil {
	//	log.Println(err)
	//}
	//b.WriteTo(w)
	//fmt.Println(data)

}

//HandleRequests func
func HandleRequests(w http.ResponseWriter, r *http.Request) {

	appList := AppList{}
	fmt.Println(appList)
	file, err := ioutil.ReadFile("static/node.json")

	if err != nil {
		log.Println(err.Error())
	}
	json.Unmarshal(file, &appList)

	m, _ := json.Marshal(appList)
	w.Write(m)
}

func main() {
	fs := http.FileServer(http.Dir(""))
	r := http.NewServeMux()
	http.Handle("/", fs)
	http.Handle("/getAppList", r)
	r.HandleFunc("/getAppList", HandleRequests)

	//http.ListenAndServe(":3001", nil)
	//mux := http.NewServeMux()

	//r := mux.NewRouter()
	//r.HandleFunc("/", LoadMain).Methods("GET")
	//r.HandleFunc("/getAppList", HandleRequests).Methods("GET")
	//r.Handle("/getAppList", handleRequest)

	http.ListenAndServe(":3001", nil)

}
