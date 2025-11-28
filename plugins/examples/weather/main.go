package main

import "C"
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Constants
const (
	ApiUrlFmt = "https://weather.yahoo.com/_atmos/api/locations?woeid=%s"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"
)

func main() {}

//export PluginRun
func PluginRun(cInput *C.char) *C.char {
	// Input is expected to be a WOEID (e.g., "44418" for London)
	woeid := C.GoString(cInput)

	fmt.Printf("[Plugin Weather] Fetching data for WOEID: %s\n", woeid)

	report, err := getFormattedWeather(woeid)
	if err != nil {
		errorMsg := fmt.Sprintf("Error fetching weather: %v", err)
		return C.CString(errorMsg)
	}

	return C.CString(report)
}

// --- Internal Logic ---

type YahooResponse struct {
	Locations []struct {
		Town       struct{ Name string } `json:"town"`
		Region     struct{ Name string } `json:"region"`
		Country    struct{ Name string } `json:"country"`
		Conditions struct {
			Temperature int    `json:"temperature"`
			Unit        string `json:"unit"`
			Text        string `json:"text"`
		} `json:"conditions"`
	} `json:"locations"`
}

func getFormattedWeather(woeid string) (string, error) {
	url := fmt.Sprintf(ApiUrlFmt, woeid)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", UserAgent)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API Status %s", resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	var data YahooResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if len(data.Locations) == 0 {
		return "", fmt.Errorf("Unknown WOEID")
	}

	loc := data.Locations[0]

	// Format string similar to the C example output
	return fmt.Sprintf(
		"Location: %s, %s | Temp: %d%s | Conditions: %s",
		loc.Town.Name,
		loc.Country.Name,
		loc.Conditions.Temperature,
		loc.Conditions.Unit,
		loc.Conditions.Text,
	), nil
}
