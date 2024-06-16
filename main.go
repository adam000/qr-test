package main

import (
	"log"
	"log/slog"
	"net/http"

	qr "github.com/skip2/go-qrcode"
)

const host = "qr.example.com"

type myHandler struct {
}

// Write the path of the request, plus an arbitrary host, to a QR code.
func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := host + r.URL.String()
	png, err := qr.Encode(path, qr.Medium, 200)
	if err != nil {
		slog.Info("Failed to generate png: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(png)
}

func main() {
	var h myHandler
	log.Fatal(http.ListenAndServe(":8080", h))
}
