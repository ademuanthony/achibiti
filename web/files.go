package web

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gofrs/uuid"
)

// UploadFile uploads a file to the server
func (s Server) uploadFile(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("file")
	if err != nil {
		s.renderErrorJSON(err.Error(), w)
		return
	}
	defer file.Close()

	id, err := uuid.NewV4()
	if err != nil {
		s.renderErrorJSON("error in generating uuid " + err.Error(), w)
		return
	}
	var name string
	mimeType := handle.Header.Get("Content-Type")
	switch mimeType {
	case "image/jpeg":
		name = fmt.Sprintf("%s.jpeg", id.String())
	case "image/png":

		name = fmt.Sprintf("%s.jpeg", id.String())
	default:
		s.renderErrorJSON("Invalid file format, please upload jpeg or png image", w)
	}

	if err = s.saveFile(w, file, name); err != nil {
		s.renderErrorJSON(err.Error(), w)
		return
	}

	s.renderJSON(map[string]interface{}{"name": name}, w)
}

func (s Server) saveFile(w http.ResponseWriter, file multipart.File, name string) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./web/public/images/"+name, data, 0666)
	if err != nil {
		return err
	}
	return nil
}
