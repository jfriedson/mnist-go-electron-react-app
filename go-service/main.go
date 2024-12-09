package main

import (
	"fmt"
	"io"
	"net/http"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, POST")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func infer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("error reading body data: %v", err)
			return
		}
		err = r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("error closing body data: %v", err)
			return
		}

		if len(bodyBytes) != 28*28 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid image size %d", len(bodyBytes))
			return
		}

		input := make([]float32, 28*28, 28*28)
		for idx, bodyByte := range bodyBytes {
			input[idx] = float32(bodyByte) / 254.
		}
		fmt.Println(input)

		// TODO: run inference

		fmt.Fprintf(w, "in progress")
	default:
		fmt.Print("method not supported", r.Method)
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
	}
}

func main() {
	// modelArch := loadMnistModelArch()
	// model := loadMnistModel()

	http.HandleFunc("/", CORS(infer))

	fmt.Println("server listening on http://localhost:5122/")
	http.ListenAndServe(":5122", nil)
}
