package main
import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("\nDeploy")

	// At least two arguments are needed for any command.
	if (len(os.Args) < 2) {
		help()
		os.Exit(0)
	}

	var fileName string = os.Args[1]
	fmt.Printf("\nDeploying%s\n", fileName)
}


func help() {
	fmt.Println("Example Usage:")
	fmt.Println("    deploy [[file-name]]")
	fmt.Println("")
}

func copyFile(sourceFileName string, destinationFileName string) bool {
	var returnValue bool = true
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		returnValue = false
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(destinationFileName)
	if err != nil {
		returnValue = false
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)
	if err != nil {
		returnValue = false
	}

	return returnValue
}
