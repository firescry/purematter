package gui

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Pure Matter!"))
}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
