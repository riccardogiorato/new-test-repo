package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Weather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := "https://api.open-meteo.com/v1/forecast?latitude=40.71&longitude=-74.01&current_weather=true"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch weather data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	var weatherData map[string]interface{}
	if err := json.Unmarshal(body, &weatherData); err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		w.Write(body)
		return
	}

	jsonResp, err := json.Marshal(weatherData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		w.Write(body)
		return
	}

	w.Write(jsonResp)
}