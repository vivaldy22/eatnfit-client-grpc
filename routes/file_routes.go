package routes

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/vivaldy22/eatnfit-client/tools/respJson"
	"github.com/vivaldy22/eatnfit-client/tools/vError"
	"github.com/vivaldy22/eatnfit-client/tools/varMux"

	"github.com/gorilla/mux"
)

type fileResponse struct {
	Message  string `json:"message"`
	FileName string `json:"file_name"`
}

type FileService struct{}

func NewFileRoute(r *mux.Router) {
	handler := FileService{}

	r.HandleFunc("/upload/{id}", handler.UploadFile).Methods(http.MethodPost)
	r.HandleFunc("/images/{id}", handler.GetFileByName).Methods(http.MethodGet)
}

func (f *FileService) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("upload-file")

	if err != nil {
		vError.WriteError("Error Retrieving file from form-data", http.StatusBadRequest, err, w)
	} else {
		defer file.Close()
		//fmt.Printf("Upload file: %+v\n", handler.Filename)
		//fmt.Printf("File size: %+v\n", handler.Size)
		//fmt.Printf("MIME Header: %+v\n", handler.Header)

		tempFile, err := ioutil.TempFile("uploaded-images", "temp-*.jpg")

		if err != nil {
			vError.WriteError("Error Creating Temp File", http.StatusBadRequest, err, w)
		} else {
			defer tempFile.Close()
			defer os.Remove(tempFile.Name())

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				vError.WriteError("Error reading file", http.StatusBadRequest, err, w)
			} else {
				tempFile.Write(fileBytes)
				dir, err := os.Getwd()
				if err != nil {
					vError.WriteError("Error getting current directory", http.StatusBadRequest, err, w)
				} else {
					imageName := varMux.GetVarsMux("id", r) + ".jpg"

					fileLocation := filepath.Join(dir, "uploaded-images")
					oldFileLocation := filepath.Join(dir, tempFile.Name())
					newFileLocation := filepath.Join(fileLocation, imageName)
					err = os.Rename(oldFileLocation, newFileLocation)

					if err != nil {
						vError.WriteError("Error renaming file", http.StatusBadRequest, err, w)
					} else {
						respJson.WriteJSON(fileResponse{
							Message:  "Success Upload File",
							FileName: imageName,
						}, w)
					}
				}
			}
		}
	}
}

func (f *FileService) GetFileByName(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		vError.WriteError("Error getting current directory", http.StatusBadRequest, err, w)
	} else {
		vars := mux.Vars(r)
		fileId := vars["id"]
		fileLocation := filepath.Join(dir, "uploaded-images", fileId)

		w.Header().Set("Content-Type", "image/jpeg")
		http.ServeFile(w, r, fileLocation)
	}
}
