package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

const BASE_URI = "https://www.googleapis.com/language/translate/v2"

var (
	apiKey = kingpin.Flag("apiKey", "API Key.").OverrideDefaultFromEnvar("TRANSLATE_API_KEY").Short('k').String()
	target = kingpin.Flag("target", "Target language.").Default("en").OverrideDefaultFromEnvar("DEFAULT_LANGUAGE").Short('t').String()
	text   = kingpin.Arg("text", "Text to translate.").Required().Strings()
)

type A struct {
	Data struct {
		Translations []struct {
			DetectedSourceLanguage string `json:"detectedSourceLanguage"`
			TranslatedText         string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func init() {
	kingpin.Parse()
}

func main() {
	res, err := http.Get(fmt.Sprintf("%s?q=%s&target=%s&key=%s", BASE_URI, strings.Join(*text, "+"), *target, *apiKey))
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
