package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var (
	ErrIndexEmpty = errors.New("no index loaded")
)

const saveFile = "data.json"

// Word -> URLs -> Occurence count
type InvertedIndex = map[string]map[string]uint

type State = string

const (
	EmptyState  State = "empty"
	LoadedState State = "loaded"
)

type data struct {
	Index       InvertedIndex `json:"index"`
	UrlsVisited []string      `json:"urlsVisited"`
}

type DataStore struct {
	state State
	data  data
}

func (ds *DataStore) Load(index InvertedIndex, urlsVisited []string) {
	ds.state = LoadedState
	ds.data = data{index, urlsVisited}
}

func (ds *DataStore) Clear() {
	*ds = *NewDataStore()
}

func (ds *DataStore) State() State {
	return ds.state
}

func (ds *DataStore) Index() (InvertedIndex, error) {
	if ds.state == EmptyState {
		return nil, ErrIndexEmpty
	}
	return ds.data.Index, nil
}

func (ds *DataStore) UrlsVisited() ([]string, error) {
	if ds.state == EmptyState {
		return nil, ErrIndexEmpty
	}
	return ds.data.UrlsVisited, nil
}

func (ds *DataStore) SaveJSON() error {
	if ds.state == EmptyState {
		return ErrIndexEmpty
	}

	jsonData, err := json.MarshalIndent(ds.data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(saveFile, jsonData, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("%q saved to disk\n", saveFile)
	return nil
}

func (ds *DataStore) LoadJSON() error {
	jsonData, err := os.ReadFile(saveFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &ds.data); err != nil {
		return err
	}

	ds.state = LoadedState

	fmt.Printf("%q loaded from disk\n", saveFile)
	return nil
}

func NewDataStore() *DataStore {
	return &DataStore{
		state: EmptyState,
	}
}
