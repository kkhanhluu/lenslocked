package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lenslocked/controllers"
	"lenslocked/views"
	"net/http"
	"path/filepath"
)

func parseTemplate(filepath string) (views.Template, error) {
	return views.Parse(
		filepath,
		"templates/shared/header.gohtml",
		"templates/shared/footer.gohtml",
		"templates/shared/nav.gohtml",
	)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.Get("/", controllers.StaticHandler(views.Must(parseTemplate(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(views.Must(parseTemplate(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(views.Must(parseTemplate(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server....")
	http.ListenAndServe(":3000", r)
}
