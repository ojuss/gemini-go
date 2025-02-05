package handlers

import (
	"encoding/json"
	"net/http"

	"gemini-go/ai"
	
)

var reqPayload struct {
	Prompt string `json:"prompt"`
}

type resPayload struct {
	Response string`json:"response"`
}

func AiPromptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	

	err := json.NewDecoder(r.Body).Decode(&reqPayload)

	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return 
	}

	resp, err := ai.Ai(reqPayload.Prompt)

	if err != nil {
		http.Error(w, "Failed to generate response", http.StatusInternalServerError)
		return
	}

	response := resPayload{Response: resp}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}