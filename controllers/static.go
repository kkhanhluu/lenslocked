package controllers

import (
	"lenslocked/views"
	"net/http"
)

func StaticHandler(t views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t.Execute(w, nil)
	}
}
