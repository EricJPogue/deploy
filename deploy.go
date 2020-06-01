package main
import (
	"fmt"
	"os"
	"time"

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
	} else if (secondArg == "lewis") || (secondArg == "-lewis") || (secondArg == "-l") {
		deployLewis()
	}  else if (secondArg == "app") || (secondArg == "-app") || (secondArg == "-a") {
		if (len(os.Args) < 3) {
			help()
		} else {
			var thirdArg string = os.Args[2]
			deploy(thirdArg)
		}
	} else {
		help()
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
	fmt.Println("    deploy app [[file-name]]")
	fmt.Println("    deploy lewis")
	fmt.Println("")
}

func version() {
	fmt.Println("Version: 1.0.6")
}

func deploy(source string) {
	fmt.Printf("Copying \"%s\"", source)

	fileName := filepath.Base(source)
	destination := "/usr/local/bin/" + fileName
	fmt.Printf(" to \"%s\"\n", destination)
	CopyFile(source, destination) 

	// Make the destination file executable. 
	err := os.Chmod(destination, 0755)
	check(err)
	fmt.Printf("successfully deployed\n")
}

func deployLewis() {
	destDirName := "/Users/eric/Repositories/lewis-education/build"

	t := time.Now()
	backupDirName := destDirName + "-backup-" + t.Format("02-Jan-2006-15-04-05")

	
	// If destination folder exists, rename it. 
	if _, err := os.Stat(destDirName); !os.IsNotExist(err) {
		fmt.Printf("\nRenaming \"%s\" \n    to \"%s\"\n\n", destDirName, backupDirName)
		err := os.Rename(destDirName, backupDirName)
		check(err)
	}

	sourceDirName := "/Users/eric/Repositories/lewis/build"
	fmt.Printf("Copying \"%s\" \n    to \"%s\"\n\n", sourceDirName , destDirName)
	err := CopyDir(sourceDirName, destDirName)
	check(err)
}
