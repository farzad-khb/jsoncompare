package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"time"
)

// defining filemap type
type filemap []map[string]interface{}

func main() {
	// args error handling
	if len(os.Args) != 3 {
		fmt.Println("please provide json file paths: ./main <json file 1 path> <json file 2 path>")
		return
	}

	// measuring execution time
	start := time.Now()
	defer log.Println("- execution time: ", time.Since(start))

	// defining two files for storing each json in filemap data structure
	var file1 filemap
	var file2 filemap
	if err := file1.jsontoFilemap(os.Args[1]); err != nil {
		return
	}

	if err := file2.jsontoFilemap(os.Args[2]); err != nil {
		return
	}

	// turns to false if can't find an object
	equaljsons := true

	// check for number of objects in each json
	if len(file1) != len(file2) {
		equaljsons = false
	}

	// Loops through each object in file1 and searching for it in file2 - O(n^2)
	if equaljsons {

		for _, file1map := range file1 {

			matched := false
			for i := 0; i < len(file2); i++ {

				if reflect.DeepEqual(file1map, file2[i]) {

					matched = true

					// removing the item from the second file (to prevent searching it again)
					file2 = append(file2[:i], file2[i+1:]...)
					break

				}
			}

			// if there was no match for the object in file1, then these two aren't equal
			if !matched {
				fmt.Println("No match found for this object in second file:", file1map)
				equaljsons = false
				break
			}

		}

	}
	// printing the result
	if equaljsons {
		fmt.Println("These two jsons are equal")
	} else {
		fmt.Println("These two jsons are different")
	}

}

//reads json files into filemap
func (fm *filemap) jsontoFilemap(fname string) error {
	// reading first file
	file, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Println("cannot read the file: ", err)
		return err
	}
	json.Unmarshal(file, fm)
	return nil
}
