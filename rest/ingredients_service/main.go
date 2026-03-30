package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

var client *genai.Client
var modelName string

func initClient() error {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		return fmt.Errorf("GOOGLE_CLOUD_PROJECT environment variable must be set")
	}
	location := os.Getenv("GOOGLE_CLOUD_LOCATION")
	if location == "" {
		location = "europe-west1"
	}

	modelName = os.Getenv("MODEL")
	if modelName == "" {
		modelName = "gemini-2.5-flash"
	}

	var err error
	// Create a new Vertex AI client
	client, err = genai.NewClient(ctx, projectID, location)
	return err
}

func ingredientsHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	// Get the "food" query parameter
	food := r.URL.Query().Get("food")
	if food == "" {
		http.Error(w, `{"error": "Query parameter 'food' is required"}`, http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	// Use Gemini 1.5 Flash
	model := client.GenerativeModel(modelName)
	model.SetTemperature(0.2)
	// Enforce JSON output format
	model.ResponseMIMEType = "application/json"

	// Prompt engineering to ask for exactly a JSON array of strings
	prompt := fmt.Sprintf("List the ingredients for '%s'. Return ONLY a flat JSON array of strings representing the ingredients. Do not wrap it in a markdown block, just output the raw JSON.", food)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Error generating content from Vertex AI: %v", err)
		http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		return
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		log.Printf("No content generated in the response")
		http.Error(w, `{"error": "Failed to generate ingredients"}`, http.StatusInternalServerError)
		return
	}

	// Extract the text part
	part := resp.Candidates[0].Content.Parts[0]
	textPart, ok := part.(genai.Text)
	if !ok {
		log.Printf("Invalid response format from model, expected Text")
		http.Error(w, `{"error": "Invalid response format from model"}`, http.StatusInternalServerError)
		return
	}

	// Return the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(textPart))
}

func main() {
	if err := initClient(); err != nil {
		log.Fatalf("Failed to initialize Vertex AI client: %v\nHave you set GOOGLE_CLOUD_PROJECT?", err)
	}
	defer client.Close()

	http.HandleFunc("/ingredients", ingredientsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
