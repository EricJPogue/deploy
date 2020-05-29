package main
import (
	"fmt"
	"io"
	"os"

	"path/filepath"
)

func main() {
	// At least two arguments are needed for any command.
	if (len(os.Args) < 2) {
		help()
		os.Exit(0)
	}

	var secondArg string = os.Args[1]
	if (secondArg == "help") || (secondArg == "-help") || (secondArg == "-h") {
		help()
	} else if (secondArg == "version") || (secondArg == "-version") || (secondArg == "-v") {
		version()
	} else {
		deploy(secondArg)
	} 
}

func check(err error) {
	if err != nil {
		fmt.Printf("Fatal Error: %s\nExiting\n", err)
		failureExitCode := 1
		os.Exit(failureExitCode)
	}
}

func help() {
	fmt.Println("Example Usage:")
	fmt.Println("    deploy version")
	fmt.Println("    deploy [[file-name]]")
	fmt.Println("")
}

func version() {
	fmt.Println("Version: 1.0.1")
}

func deploy(source string) {
	fmt.Printf("Copying \"%s\"", source)

	fileName := filepath.Base(source)
	destination := "/usr/local/bin/" + fileName
	fmt.Printf(" to \"%s\"\n", destination)
	copyFile(source, destination) 

	// Make the destination file executable. 
	err := os.Chmod(destination, 0755)
	check(err)
	fmt.Printf("successfully deployed\n")
}

func copyFile(sourceFileName string, destinationFileName string) {
	sourceFile, err := os.Open(sourceFileName)
	check(err)
	defer sourceFile.Close()

	newFile, err := os.Create(destinationFileName)
	check(err)
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)
	check(err)
}
