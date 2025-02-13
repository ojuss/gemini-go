package handlers

import (
	"encoding/json"
	"net/http"
	"bytes"
	"io"
	"fmt"
	"gemini-go/ai"
	
)

var reqPayload struct {
	Prompt string `json:"prompt"`
}

var resPayload struct {
	Response string `json:"response"`
}

func AiPromptHandler(w http.ResponseWriter, r *http.Request) {

	// method check doesn't work correctly
	// shows json formatting issue
	// if r.Method != "POST" {

	// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	// }

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

	var buffer bytes.Buffer	
	resPayload.Response = resp

	err = json.NewEncoder(&buffer).Encode(resPayload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// response := resPayload{Response: resp}

	w.Header().Set("Content-Type", "application/json")
	_, err = io.Copy(w, &buffer)

	if err != nil {
		fmt.Println(err)
	}

}