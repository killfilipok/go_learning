package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"unicode"
// )

// func main() {
// 	b, err := ioutil.ReadFile("input.txt")

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	str := string(b)
// 	var newStr []rune
// 	for _, s := range str {
// 		if unicode.IsNumber(s) || s == '.' {
// 			continue
// 		}
// 		newStr = append(newStr, s)

// 	}
// 	file, err := os.OpenFile("output.txt", os.O_RDWR, 0644)

// 	// file, err := os.Create("result.txt")
// 	if err != nil {
// 		log.Fatal("Cannot create file", err)
// 	}
// 	defer file.Close()
// 	file.WriteString(string(newStr))
// 	// fmt.Fprintf(file, string(newStr))
// 	fmt.Println(str)
// }
