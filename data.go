package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

func loadData() bool {
	filePath := path.Join(dataPath, dataFile)
	f, _ := exists(filePath)
	if !f {
		file, err := ioutil.ReadFile(filePath)

		if err != nil {
			return false
		}
		err = json.Unmarshal(file, &appData)

		if err != nil {
			return false
		}
	}

	return true
}

func saveData() bool {
	var err error
	var file []byte

	os.MkdirAll(dataPath, os.ModePerm)

	file, err = json.MarshalIndent(appData, "", " ")

	if err != nil {
		return false
	}
	err = ioutil.WriteFile(path.Join(dataPath, dataFile), file, 0644)

	if err != nil {
		return false
	}

	return true
}

func exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, nil
	}
	return err != nil, err
}
