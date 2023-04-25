package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/felipecavassana/new_store/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {
	allItems := models.SearchAllItems()
	temp.ExecuteTemplate(w, "Index", allItems)
}

func New(w http.ResponseWriter, _ *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		amount := r.FormValue("amount")
		quantity := r.FormValue("quantity")

		amountFloatConversion, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			log.Println("ERROR - amount conversion:", err)
		}

		quantityIntConversion, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("ERROR - quantity conversion:", err)
		}

		models.NewItem(name, description, amountFloatConversion, quantityIntConversion)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("id")
	models.DeleteItem(itemId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	itemId := r.URL.Query().Get("id")
	item := models.EditItem(itemId)
	temp.ExecuteTemplate(w, "Edit", item)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		amount := r.FormValue("amount")
		quantity := r.FormValue("quantity")

		idIntConversion, err := strconv.Atoi(id)
		if err != nil {
			log.Println("ERROR - ID conversion:", err)
		}

		amountFloatConversion, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			log.Println("ERROR - amount conversion:", err)
		}

		quantityIntConversion, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("ERROR - quantity conversion:", err)
		}

		models.UpdateItem(idIntConversion, name, description, amountFloatConversion, quantityIntConversion)
	}
	http.Redirect(w, r, "/", 301)
}
