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
	f, err := os.Create(fmt.Sprintf("/temp/%s", fName))
	if err != nil {
		log.Printf("creating file failed: %v", err)
	}
	defer f.Close()

	_, err = f.Write(bInput)
	if err != nil {
		log.Printf("writing to file failed: %v", err)
	}

	return true
}
