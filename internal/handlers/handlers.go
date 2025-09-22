package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleMain(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, fmt.Sprint("The server does not support the method", req.Method), http.StatusInternalServerError)
		return
	}

	abs, err := filepath.Abs("./index.html")
	if err != nil {
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	html, err := os.ReadFile(abs)
	if err != nil {
		http.Error(res, "Failed to load file", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(html)
}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, fmt.Sprint("The server does not support the method", req.Method), http.StatusInternalServerError)
		return
	}

	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "Failed to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "Failed to read file", http.StatusInternalServerError)
		return
	}

	text := service.ConvertMorse(string(data))
	name := fmt.Sprintf("../file_%s%s", time.Now().UTC().Format("20060102_150405"), filepath.Ext("*.txt"))

	newFile, err := os.Create(name)
	if err != nil {
		http.Error(res, name, http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(text)
	if err != nil {
		http.Error(res, "Failed to write result to file", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, text)
}
