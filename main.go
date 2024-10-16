package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		Temperature float64 `json:"temp_c"`
		FeelsLike   float64 `json:"feelslike_c"`
		Humidity    float64 `json:"humidity"`
		Condition   struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

var weather Weather

// https://api.weatherapi.com/v1/current.json?key=be915cf5f218493dbd852745242008&q=Ashgabat
func main() {
	apiKey := "be915cf5f218493dbd852745242008"
	apiUrl := "http://localhost:8000/weather"

	if len(os.Args) < 2 {
		log.Println("Please indicate the name of the city")
		return
	}

	// принимать название города в качестве аргумента командной строки
	city := os.Args[1]

	// Создание параметров запроса
	parameter := url.Values{}
	parameter.Add("key", apiKey)
	parameter.Add("q", city)

	requestUrl := fmt.Sprintf("%s?%s", apiUrl, parameter.Encode())

	fmt.Println(requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Println("Error executing request. ", resp.StatusCode, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("Error! City not found or Incorrect request. Status: ", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ggg", err)
		return
	}

	er := json.Unmarshal(body, &weather)

	if er != nil {
		log.Println("Error parsing JSON!", er)
		return
	}
	fmt.Printf("Current weather in %s:\n", weather.Location.Name)
	fmt.Printf("Temperature :%.1fC\n", weather.Current.Temperature)
	fmt.Printf("Feels like : %.1fC\n", weather.Current.FeelsLike)
	fmt.Printf("Hudimity: %.1f%%\n", weather.Current.Humidity)
	fmt.Printf("Description: %v\n", weather.Current.Condition.Text)

}

//  weatherapi API key: be915cf5f218493dbd852745242008
// openweathermap API key: 6c322c3f91c70af716f54a56ee501533
