package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/*
// if r.ContentLength > MAX_UPLOAD_SIZE {
// 	http.Error(w, "The uploaded image is too big. Please use an image less than 1MB in size", http.StatusBadRequest)
// 	return
// }
because the Content-Length header
can be modified on the client
to be any value regardless
 of the actual file size

*/

var (
	templateFile = template.Must(template.ParseFiles("templates/index.html"))
)

const MAX_UPLOAD_SIZE = 1024 * 1024 //  1MB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		UploadHandler(w, r)
		return
	}
	templateFile.ExecuteTemplate(w, "index.html", nil)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	/*

		The http.MaxBytesReader() method is used to limit the size of
		incoming request bodies. For single file uploads, l
		imiting the size of the request body
		provides a good approximation of limiting the file size.
		The ParseMultipartForm() method subsequently parses the
		request body as multipart/form-data up to the max memory argument.
		If the uploaded file is larger than the argument to ParseMultipartForm(),
		an error will occur.


	*/

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	//  The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	// buff := make([]byte, 512)
	// _, err = file.Read(buff)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// filetype := http.DetectContentType(buff)
	// if filetype != "image/jpeg" && filetype != "image/png" {
	// 	{
	// 		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
	// 		return
	// 	}

	// 	/*

	// 					The DetectContentType() method is provided by the http package for the
	// 					purpose of detecting the content type of the given data. It considers (at most) the first 512 bytes of data to determine the MIME type. This is why we read the first 512 bytes of the file to an empty
	// 					buffer before passing it to the DetectContentType() method.
	// 					 If the resulting filetype is neither a JPEG or PNG, an error is returned.
	// 		 		When we read the first 512 bytes of the uploaded file in order to determine
	// 				 the content type, the underlying file stream pointer moves forward by
	// 				 512 bytes. When io.Copy() is called later, it continues reading from
	// 				  that position resulting in a corrupted image file.
	// 				  The file.Seek() method is used to return the pointer back to
	// 				  the start of the file so that io.Copy() starts from the beginning.

	// 	*/
	// 	_, err := file.Seek(0, io.SeekStart)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// create upload folder if it doesnot exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	fmt.Println(err, "erris")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create a new files in the uploads directory

	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/upload", UploadHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
