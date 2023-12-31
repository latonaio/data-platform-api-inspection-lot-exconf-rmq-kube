package dpfm_api_input_reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type FileReader struct{}

func NewFileReader() *FileReader {
	return &FileReader{}
}

func (*FileReader) ReadECMC(path string) EC_MC {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("input data read error :%#v", err.Error())
		os.Exit(1)
	}
	ec := EC_MC{}
	err = json.Unmarshal(raw, &ec)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return ec
}

func (*FileReader) ReadSDC(path string) HeaderSDC {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("input data read error :%#v", err.Error())
		os.Exit(1)
	}
	sdc := HeaderSDC{}
	err = json.Unmarshal(raw, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}

func ConvertToSDC(data map[string]interface{}) HeaderSDC {
	raw, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("data marshal error :%#v", err.Error())
		return HeaderSDC{}
	}
	sdc := HeaderSDC{}
	err = json.Unmarshal(raw, &sdc)
	if err != nil {
		fmt.Printf("input data marshal error :%#v", err.Error())
		os.Exit(1)
	}

	return sdc
}
