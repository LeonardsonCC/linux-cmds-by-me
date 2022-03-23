package main

import "log"

func main() {
	dirs, files, err := ListFiles(".")
	if err != nil {
		log.Fatal(err)
	}

	output := NewOutput(dirs, files)
	output.ShowHidden = true
	output.ShowUserOwner = false
	output.ShowGroupOwner = false
	output.Print()
}
