package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var filename string
	fmt.Scanf("%s", &filename)

	inputByteArr, err := ioutil.ReadFile(filename + ".json")

	checkErr := func() {
		if err != nil {
			log.Panic(err)
		}
	}

	checkErr()

	str := string(inputByteArr)

	var m []map[string]*json.RawMessage
	err = json.Unmarshal(inputByteArr, &m)

	checkErr()

	for _, micro := range m {
		for testName, nano := range micro {
			fmt.Println("--- " + testName)

			nanoByteArr, _ := json.Marshal(&nano)

			var nanoNext map[string]*json.RawMessage
			err = json.Unmarshal(nanoByteArr, &nanoNext)

			checkErr()
			for atomID, atom := range nanoNext {
				atomByteArr, _ := json.Marshal(&atom)

				var atomicNext map[string]*json.RawMessage
				err = json.Unmarshal(atomByteArr, &atomicNext)

				checkErr()
				fmt.Println(atomID)

				for cvarkID, cvark := range atomicNext {
					cvarkByteArr, _ := json.Marshal(&cvark)

					fmt.Println(cvarkID + "->" + string(cvarkByteArr))
				}

			}
		}
	}
}
