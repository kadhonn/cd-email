package server

import (
	"bytes"
	"errors"
	"github.com/kadhonn/cd-email/email"
	"log"
	"mime/multipart"
	"net/http"
)

var sender *email.Sender
var password string

func Run(givenSender *email.Sender, givenPassword string) {
	sender = givenSender
	password = givenPassword
	http.HandleFunc("/", renderSiteHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/submit/", submitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderSiteHandler(w http.ResponseWriter, r *http.Request) {
	renderSite(w)
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cache-Control", "max-age=31536000")
	renderFavico(w)
}
func submitHandler(w http.ResponseWriter, r *http.Request) {

	sentPassword := r.FormValue("password")
	if sentPassword != password {
		err := errors.New("passwords didn't match")
		log.Print(err)
		renderError(w, err)
		return
	}

	emailStruct, err := getEmail(r)

	if err != nil {
		log.Print(err)
		renderError(w, err)
		return
	}

	err = sender.Send(emailStruct)

	if err != nil {
		log.Print(err)
		renderError(w, err)
		return
	}

	renderOk(w)
}

func getEmail(request *http.Request) (email.Email, error) {
	err := request.ParseMultipartForm(0)

	if err != nil {
		return email.Email{}, err
	}

	fileHeader := request.MultipartForm.File["file"][0]
	content, err := getFile(fileHeader)

	if err != nil {
		return email.Email{}, err
	}

	return email.Email{Address: request.FormValue("email"), FileName: fileHeader.Filename, Content: content}, nil
}

func getFile(fileHeader *multipart.FileHeader) (content []byte, err error) {
	file, _ := fileHeader.Open()
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
