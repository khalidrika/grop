package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func HOmehandler(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer data.Body.Close()

	var slice []Data

	if err := json.NewDecoder(data.Body).Decode(&slice); err != nil {
		fmt.Println("err")
		return
	}

	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tmp.Execute(w, slice)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	data, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer data.Body.Close()

	var slice []Data

	if err := json.NewDecoder(data.Body).Decode(&slice); err != nil {
		fmt.Println("err")
		return
	}
	temp

	// fmt.Println(r.FormValue("id"))
	// fmt.Println(r.URL.Query().Get("id"))
	id := r.URL.Path[len("/artist/"):]
	fmt.Println(id)
}
