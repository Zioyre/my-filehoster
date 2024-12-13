package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFile(wr http.ResponseWriter, rq *http.Request) {
	err := rq.ParseMultipartForm(128 << 20)
	if err != nil {
		http.Error(wr, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, er := rq.FormFile("file")
	if er != nil {
		http.Error(wr, "Error getting the file", http.StatusBadRequest)
		return
	}
	if handler.Size > 128<<20 {
		http.Error(wr, "Max file size is 128mb.", http.StatusRequestEntityTooLarge)
		return
	}
	defer file.Close()

	fname := fmt.Sprintf("%d%s", rand.Intn(1000000), filepath.Ext(handler.Filename))
	destPath := filepath.Join("./files", fname)
	destFile, er := os.Create(destPath)
	if er != nil {
		http.Error(wr, "OOPSIE WOOPSIE!! Uwu We made a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this!", http.StatusInternalServerError)
		return
	}
	defer destFile.Close()
	_, er = io.Copy(destFile, file)
	if er != nil {
		http.Error(wr, "OOPSIE WOOPSIE!! Uwu We made a fucky wucky!! A wittle fucko boingo! The code monkeys at our headquarters are working VEWY HAWD to fix this!", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(wr, "Success! Your file should now be accessible here: https://(your ip/domain)/%s\n", fname)
	fmt.Printf("File successfully saved to: %s\n", destPath)
}

func main() {
	fmt.Println("Server now listening on port 48723")
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":48723", nil)
}
