package main

import (
	"net/http"
	"text/template"
)

type PageData struct{}

func main() {
	http.HandleFunc("/", handler("templates/index.html"))
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/geofeed.csv", geofeedHandler)
	http.HandleFunc("/peering", handler("templates/peering.html"))
	http.HandleFunc("/communities", handler("templates/communities.html"))
	http.HandleFunc("/contact", handler("templates/contact.html"))
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":8080", nil)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func geofeedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	http.ServeFile(w, r, "geofeed.csv")
}

func generatePageData(r *http.Request) PageData {
	return PageData{}
}

func handler(templatePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := generatePageData(r)
		renderTemplateWithData(w, templatePath, data)
	}
}

func renderTemplateWithData(w http.ResponseWriter, path string, data interface{}) {
	tmpl, err := template.ParseFiles(path, "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/"+r.URL.Path[len("/static/"):])
}
