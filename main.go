// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START translate_quickstart]

// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func main() {
	ctx := context.Background()

	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	target, err := language.Parse("ru")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	b, err := ioutil.ReadFile("output.txt")

	if err != nil {
		log.Panic(err)
	}

	str := string(b)

	stringss := strings.Split(str, " ")
	file, err := os.OpenFile("translation.txt", os.O_RDWR, 0644)

	translated := []string{}

	for _, s := range stringss {
		translations, err := client.Translate(ctx, []string{s}, target, nil)
		if err != nil {
			log.Fatalf("Failed to translate text: %v", err)
		}
		translated = append(translated, translations[0].Text)
	}
	// file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	file.WriteString(strings.Join(translated, "\n "))
	// fmt.Fprintf(file, string(newStr))
	fmt.Println(str)
}

// func maain() {
// 	ctx := context.Background()

// 	// Creates a client.
// 	client, err := translate.NewClient(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	// Sets the text to translate.
// 	text := "Hello"
// 	// Sets the target language.
// 	target, err := language.Parse("ru")
// 	if err != nil {
// 		log.Fatalf("Failed to parse target language: %v", err)
// 	}

// 	// Translates the text into Russian.
// 	translations, err := client.Translate(ctx, []string{text}, target, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to translate text: %v", err)
// 	}

// 	fmt.Printf("Text: %v\n", text)
// 	fmt.Printf("Translation: %v\n", translations[0].Text)
// 	fmt.Println(translations)
// }

// [END translate_quickstart]
