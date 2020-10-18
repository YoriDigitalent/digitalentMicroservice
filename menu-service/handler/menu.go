package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AddMenu handle add menu
func AddMenu(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(map[string]interface{}{
		"success": true,
	})

	if err != nil {
		fmt.Printf("Failed generate response %s", err)
		return
	}

	w.Write(response)
}