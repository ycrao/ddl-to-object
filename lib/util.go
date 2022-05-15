package lib

import (
	"bufio"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"text/template"
	"time"
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

// ReadTemplate read template by target language
func ReadTemplate(language string) (*template.Template, error) {
	templatePath := "./template/" + language + ".template"
	existed, _ := PathExists(templatePath)
	if !existed {
		templatePath = "~/.dto/template/" + language + ".template"
	}
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		errors.New("fail to get default template, please check your template files")
	}
	return tpl, nil
}

// RandomStr return random string
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomInt64Str return random string
func RandomInt64Str(length int) string {
	if length < 18 {
		length = 18
	}
	str := "0123456789135792468"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	// void first letter is 0
	return "8" + string(result)
}
