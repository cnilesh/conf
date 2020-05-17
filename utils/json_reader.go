package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//Read any configuration
func Read(file string, i interface{}) interface{} {
	absPath, absErr := filepath.Abs("/Users/nileshchaudhary/videoconf/src/db/conf/" + file)
	fmt.Println("Abs:", absPath)
	if absErr != nil {
		fmt.Println("error:", absErr)
	}
	// file, _ := os.Open(absPath)
	configFile, _ := ioutil.ReadFile(absPath)
	// defer configFile.Close()
	// decoder := json.NewDecoder(configFile)
	configuration := i
	// err := decoder.Decode(&configuration)
	_ = json.Unmarshal([]byte(configFile), &configuration)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	return configuration
}
