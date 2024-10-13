package main

import (
	"net"
	"net/http"
	"text/template"
	"time"
)

type PageData struct {
	ConnectedToAS44354 bool
	Year               int
}

func main() {
	http.HandleFunc("/", handler("templates/index.html"))
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/locations", handler("templates/locations.html"))
	http.HandleFunc("/peering", handler("templates/peering.html"))
	http.HandleFunc("/contact", handler("templates/contact.html"))
	http.HandleFunc("/static/", staticHandler)
	http.ListenAndServe(":8080", nil)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func generatePageData(r *http.Request) PageData {
	serverIP := r.Header.Get("X-Server-IP")
	return PageData{
		ConnectedToAS44354: isIPv6InRange(serverIP, "2a14:7c0:4b00::/40"),
		Year:               time.Now().Year(),
	}
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

func isIPv6InRange(ipStr string, cidrStr string) bool {
	ip := net.ParseIP(ipStr)
	_, ipNet, _ := net.ParseCIDR(cidrStr)
	return ip != nil && ipNet.Contains(ip)
}
