package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := []byte(`
			<html>
				<head>
					<title>A Simple HTTP Server</title>
				</head>
				<body>
					Check it out!
				</body>
			</html>
		`)

		w.Write(html)
	})

	// Start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ran into an error:", err)
	} else {
		log.Println("Serving on http://localhost:8080")
	}
}
