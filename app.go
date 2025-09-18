package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) []CityAutocomplete {
	//fmt.Println("ВЫЗОВ ИЗ GO", name)
	token := tokens()
	params := url.Values{}
	params.Add("name", name)
	params.Add("country_code", "RU")
	fullURL := "http://api.edu.cdek.ru/v2/location/suggest/cities" + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
	var cit []CityAutocomplete
	err = json.Unmarshal([]byte(body), &cit)
	if err != nil {
		panic(err)
	}
	//return []string{"Москва", "Санкт-Петербург"}
	return cit
}

func (a *App) OfficesList(name string) []Deliverypoints {
	//fmt.Println("ВЫЗОВ ИЗ GO", name)
	token := tokens()
	params := url.Values{}
	params.Add("city_code", name)
	params.Add("country_code", "RU")
	fullURL := "http://api.edu.cdek.ru/v2/deliverypoints" + "?" + params.Encode()
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	var offices []Deliverypoints
	err = json.Unmarshal([]byte(body), &offices)
	if err != nil {
		panic(err)
	}
	//return []string{"Москва", "Санкт-Петербург"}
	return offices
}

func (a *App) Prints(x string) {
	fmt.Println(x)
}

func (a *App) Printm(x Deliverypoints) {
	fmt.Println(x)
}
