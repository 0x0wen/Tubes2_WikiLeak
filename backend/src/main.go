package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// IDS("/wiki/Talang_(Swedish_TV_series)", "/wiki/Sweden")
	// BFS("/wiki/Mike_Tyson", "Sweden")
	// Create an HTTP handler
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Allow requests from any origin
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		if req.Method == http.MethodOptions {
			// Set CORS headers for preflight requests
			rw.Header().Set("Access-Control-Allow-Methods", "POST")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			rw.WriteHeader(http.StatusOK)
			return
		}

		if req.Method != http.MethodPost {
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Read the request body
		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(rw, "Failed to read request body", http.StatusBadRequest)
			return
		}

		// Define a struct to unmarshal the JSON data into
		type RequestData struct {
			// Define your data structure here based on the expected JSON format
			Start     string `json:"start"`
			Goal      string `json:"goal"`
			Algorithm string `json:"algorithm"`
			// Add more fields as needed
		}

		// Unmarshal the JSON data into the struct
		var requestData RequestData
		if err := json.Unmarshal(body, &requestData); err != nil {
			http.Error(rw, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		// Now you can use requestData.Start and requestData.Goal in your backend logic
		log.Printf("Start: %s, Goal: %s, Algorithm: %s\n", requestData.Start, requestData.Goal, requestData.Algorithm)

		type ResponseData struct {
			Status string `json:"status"`
			Result Result `json:"result"`
		}

		// Inside the handler function:
		responseData := ResponseData{
			Status: "ok",
		}

		if requestData.Algorithm == "BFS" {
			responseData.Result = BFS(requestData.Start, requestData.Goal)
		} else if requestData.Algorithm == "IDS" {
			responseData.Result = IDS(requestData.Start, requestData.Goal)
		}

		// Marshal the response data to JSON
		respJSON, err := json.Marshal(responseData)
		if err != nil {
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set response headers
		rw.Header().Set("Content-Type", "application/json")
		rw.Header().Set("Content-Length", fmt.Sprint(len(respJSON)))

		// Write the response
		rw.Write(respJSON)
	})

	// Start the server
	log.Println("Server is available at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
