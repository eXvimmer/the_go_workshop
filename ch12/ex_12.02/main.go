package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	ErrWorkingFileNotFound = errors.New("the working file is not found")
)

func createBackup(working, backup string) error {
	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		}
	}

	workFile, err := os.Open(working)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(workFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(backup, content, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func addNotes(workingFile, notes string) error {
	notes += "\n"
	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write([]byte(notes)); err != nil {
		return err
	}

	return nil
}

func main() {
	backupFile := "backup.txt"
	workingFile := "notes.txt"
	data := "Contribute to Open Source projects"

	err := createBackup(workingFile, backupFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		note := fmt.Sprintf("%s %d", data, i)

		err := addNotes(workingFile, note)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
