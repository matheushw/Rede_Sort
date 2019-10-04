package utils

import (
    "fmt"
	"os"
)

func Create_config() { //This functions creates the "config" in the script folder
	_, err := os.Create("config")
    if err != nil {
		fmt.Printf("error creating config file: %v", err)
        return
	}
}

func Open_config() (*os.File){ //This function open the "config" file and returns its file descriptor
	f, err := os.OpenFile("config", os.O_RDWR, 0);

	if err != nil {
		fmt.Printf("error opening config file: %v", err)
        return nil
	}
	return f
}

func Write_local_ip() { //This function updates the master's local ip
	f, err := os.OpenFile("config", os.O_RDWR, 0);

	if err != nil {
		fmt.Printf("error opening config file: %v", err)
        return
	}

	_, err = f.WriteString("IP") // TODO

	if err != nil {
		fmt.Printf("error while writing the local ip: %v", err)
        return
	}

	f.Sync()
}