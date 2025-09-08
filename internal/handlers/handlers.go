package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(32 << 20)

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error with upload", http.StatusBadRequest)
		return
	}

	defer file.Close()

	forConv, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error with reading", http.StatusInternalServerError)
		return
	}

	converted, err := service.Convert(string(forConv))
	if err != nil {
		http.Error(w, "error with converting", http.StatusInternalServerError)
		return
	}

	nameForFile := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)

	temporaryFile, err := os.Create(nameForFile)
	if err != nil {
		http.Error(w, "error with creating", http.StatusInternalServerError)
		return
	}

	defer temporaryFile.Close()

	_, err = temporaryFile.WriteString(converted)
	if err != nil {
		http.Error(w, "error with writing", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
