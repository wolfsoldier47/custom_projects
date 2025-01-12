package controllers

import (
	"encoding/json"
	"fmt"
	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func headerContent(w http.ResponseWriter, res []byte) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	headerContent(w, res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bootDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bootDetails)
	headerContent(w, res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	headerContent(w, res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("couldnt Parse id:", err)
	}
	b := models.DeleteBook(ID)
	res, _ := json.Marshal(b)

	headerContent(w, res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing", err)
	}
	newBook, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		newBook.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		newBook.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		newBook.Publication = updateBook.Publication
	}
	db.Save(&newBook)
	res, _ := json.Marshal(newBook)

	headerContent(w, res)
}
