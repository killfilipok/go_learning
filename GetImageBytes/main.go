package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open(`.\image.jpg`)
	defer file.Close()

	// fileInfo, _ := file.Stat()
	// var size int64 = fileInfo.Size()
	// buffer := make([]byte, size)
	// file.Read(buffer)

	// _ = ioutil.WriteFile("file.txt", buffer, os.FileMode(0644))

	// original_image, _, _ := image.Decode(file)

	// buf := new(bytes.Buffer)
	// _ = jpeg.Encode(buf, original_image, nil)
	// send_s3 := buf.Bytes()

	// _ = ioutil.WriteFile("2file.txt", send_s3, os.FileMode(0644))
	// f, _ := os.Open("./coin.jpg")

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	fileDir := `.\new`

	fileBytes, _ := base64.StdEncoding.DecodeString(encoded)

	filetype := http.DetectContentType(fileBytes)

	switch filetype {
	case "image/jpeg":
		fileDir += ".jpeg"
	case "image/jpg":
		fileDir += ".jpg"
	case "image/png":
		fileDir += ".png"
	default:
	}

	file, _ = os.Create(fileDir)

	permissions := 0644
	_ = ioutil.WriteFile(fileDir, []byte(encoded), os.FileMode(permissions))
	tfile, _ := os.Create(`.\text.txt`)
	io.Copy(tfile, strings.NewReader(encoded))
	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	fmt.Println("ENCODED: " + encoded)
}
