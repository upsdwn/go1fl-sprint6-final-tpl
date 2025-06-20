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

	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(res, "Failed to load file", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(html)
}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, fmt.Sprint("The server does not support the method", req.Method), http.StatusInternalServerError)
		return
	}

	file, handler, err := req.FormFile("myFile")
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

	ext := filepath.Ext(handler.Filename)
	datetime := time.Now().UTC().Format("20060102_150405")
	fileName := fmt.Sprintf("file_%s%s", datetime, ext)

	err = os.WriteFile(fileName, []byte(text), 0o755)
	if err != nil {
		http.Error(res, "Failed to save file", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(text))
}
