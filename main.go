package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Главная"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Заметка"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Неверный запрос", 405)
		return
	}
	w.Write([]byte("Создание новой заметки"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	log.Println("Сервер запущен")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
