package Ejson

import (
	"io/ioutil"
	"sync"

	"github.com/buger/jsonparser"
)

type Ejson struct {
	JsonFilePath string
	RW           *sync.RWMutex
}

// NewEjson Create a new Ejson
func NewEJsonInit(jsonFilePath string) *Ejson {
	EjsonInstance := &Ejson{
		JsonFilePath: jsonFilePath,
		RW:           &sync.RWMutex{},
	}
	return EjsonInstance
}

// GetStringBytes Get string bytes
func (e *Ejson) Get(key ...string) ([]byte, error) {
	e.RW.RLock()
	data, _ := ioutil.ReadFile(e.JsonFilePath)
	e.RW.RUnlock()
	vb, _, _, err := jsonparser.Get(data, key...)
	if err != nil {
		return nil, err
	}
	//return []byte(str)
	return vb, nil
}

//update creat, Set([]byte(`{"key":"value"}`, "key", "[0]")))
func (e *Ejson) Set(jsonByte []byte, key ...string) error {
	e.RW.Lock()
	defer e.RW.Unlock()
	data, _ := ioutil.ReadFile(e.JsonFilePath)
	res, err := jsonparser.Set(data, []byte(jsonByte), key...)
	if err != nil {
		return err
	}
	ioutil.WriteFile(e.JsonFilePath, res, 0666)
	return nil
}

//Delete
func (e *Ejson) Delete(key ...string) error {
	e.RW.Lock()
	defer e.RW.Unlock()
	data, _ := ioutil.ReadFile(e.JsonFilePath)
	res := jsonparser.Delete(data, key...)
	ioutil.WriteFile(e.JsonFilePath, res, 0666)
	return nil
}
