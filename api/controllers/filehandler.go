package controllers

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/app/api/auth"
	"github.com/app/api/responses"
)

func (server *Server) CreateImage(w http.ResponseWriter, r *http.Request) {

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	println(uid)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = io.WriteString(w, "File "+fileName+" Uploaded successfully")
	_, _ = io.Copy(f, file)
}
