package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const BASE_URI = "https://www.googleapis.com/language/translate/v2"

var apiKey string

type A struct {
	Data struct {
		Translations []struct {
			DetectedSourceLanguage string `json:"detectedSourceLanguage"`
			TranslatedText         string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func init() {
	apiKey = os.Getenv("TRANSLATE_API_KEY")
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("translate [target-language] [text]")
	}

	targetLanguage := os.Args[1]
	text := strings.Join(os.Args[2:], "+")

	res, err := http.Get(fmt.Sprintf("%s?q=%s&target=%s&key=%s", BASE_URI, text, targetLanguage, apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var a A

	if err := json.NewDecoder(res.Body).Decode(&a); err != nil {
		log.Fatal(err)
	}

	fmt.Println(a.Data.Translations[0].TranslatedText)
}
