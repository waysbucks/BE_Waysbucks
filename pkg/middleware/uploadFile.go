package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	dto "waysbucks/dto/result"
)

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		file, handler, err := r.FormFile("image")

		if err != nil && r.Method == "PATCH" {
			ctx := context.WithValue(r.Context(), "dataFile", "false")
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		if err != nil {
			fmt.Println(handler)
			json.NewEncoder(w).Encode("Error Retriving the File")
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: "Format Image not a JPEG or PNG "}
			json.NewEncoder(w).Encode(response)
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		// fmt.Printf("Uploaded: %+v\n", handler.Filename)
		const MAX_UPLOAD_SIZE = 10 << 20
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Max size is 10mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		tempFile, err := ioutil.TempFile("uploads", "image-*-"+handler.Filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}

		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)

		data := tempFile.Name()
		filename := data[8:]

		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
