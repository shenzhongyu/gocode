package main

import (
	"net/http"
	"html/template"
	"os"
	"io"
	"io/ioutil"
)

const UPLOAD_DIR = "./uploads";

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("./tpl/upload.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerErrors)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
			http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id=" + filename,
		http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(),
		http.StatusInternalServerError)
		return
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name)
	}
	locals["images"] = images ; t, err := template.ParseFiles("list.html")
	if err != nil {
		http.Error(w, err.Error(),
		http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}