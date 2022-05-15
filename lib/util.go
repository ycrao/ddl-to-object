package lib

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

// PathExists path exists?
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


// ReadFile read file as []byte
func ReadFile(path string) ([]byte, error) {
	existed, ex := PathExists(path)
	if existed {
		file, err := os.OpenFile(path, os.O_RDONLY, 0644)
		if err != nil {
			return []byte{}, err
		}
		defer file.Close()
		var content []byte
		content, err = ioutil.ReadAll(file)
		return content, err
	}
	return []byte{}, ex
}

// WriteBytesToFile write bytes content to file
func WriteBytesToFile(filePath string, content []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.Write(content)
	writer.Flush()
	return nil
}

// WriteStringToFile write string content to file
func WriteStringToFile(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
	return nil
}

// VisitLocationInWriteMode when directory not existed, will create it
func VisitLocationInWriteMode(location string) error {
	existed, _ := PathExists(location)
	if !existed {
		path := filepath.Join(".", location)
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return errors.New("create directory fail: " + err.Error())
		}
	}
	return nil
}
