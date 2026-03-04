// Package engfile provides utilities for file operations such as writing data to files.
package engfile

import (
	"fmt"
	"log"
	"os"
)

// ToFile takes any interface, and converts it to bytes, and prints to file.
func ToFile(fName string, input interface{}) bool {
	bInput := []byte(fmt.Sprintf("%+v\n", input))

	// create and write
	f, err := os.Create(fmt.Sprintf("/tmp/%s", fName))
	if err != nil {
		log.Printf("creating file failed: %v", err)
		return false
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			log.Printf("closing file failed: %v", closeErr)
		}
	}()

	_, err = f.Write(bInput)
	if err != nil {
		log.Printf("writing to file failed: %v", err)
		return false
	}

	return true
}
