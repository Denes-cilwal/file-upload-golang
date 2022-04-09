package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path"
// 	"text/template"
// )

// var (
// 	templateFile = template.Must(template.ParseFiles("templates/index.html"))
// )

// func main() {
// 	fmt.Println("file upload learning")
// 	log.Println("starting server on port 8080")

// 	http.HandleFunc("/", Home)
// 	http.ListenAndServe(":8080", nil)
// }

// func Home(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		handleUpload(w, r)
// 		return
// 	}
// 	templateFile.ExecuteTemplate(w, "index.html", nil)
// 	w.Write([]byte("Welcome home"))
// }

// func handleUpload(w http.ResponseWriter, r *http.Request) {
// 	// step : specify no of bytes want to read from the body
// 	// r.ParseMultipartForm(10 << 20) // 10 MB

// 	const MAX_UPLOAD_SIZE = 1024 * 1024 //  1 MB

// 	file, fileHeader, err := r.FormFile("image")
// 	fmt.Println(file, "file is")
// 	fmt.Println(fileHeader, "fileHeader is")
// 	fmt.Println(err, "err is")
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}
// 	defer file.Close()

// 	// path base will return file without any slash
// 	fileName := path.Base(fileHeader.Filename)
// 	fmt.Println(fileName, "fileNmae is --->")
// 	dest, err := os.Create(fileName)
// 	if err != nil {
// 		http.Error(w, "InternalServerError", http.StatusInternalServerError)
// 		return
// 	}
// 	defer dest.Close()

// 	if _, err = io.Copy(dest, file); err != nil {
// 		http.Error(w, "InternalServerError", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/success=true", http.StatusSeeOther)
// }
