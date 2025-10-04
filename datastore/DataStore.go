package datastore

import (
	"errors"
	"fmt"
)

// Word -> URLs -> Occurence count
type InvertedIndex = map[string]map[string]uint

type State = string

const (
	EmptyState  State = "empty"
	LoadedState State = "loaded"
)

type DataStore struct {
	state       State
	index       InvertedIndex
	urlsVisited []string
}

func (ds *DataStore) Load(index InvertedIndex, urlsVisited []string) {
	ds.state = LoadedState
	ds.index = index
	ds.urlsVisited = urlsVisited
	fmt.Println("Index loaded")
}

func (ds *DataStore) Clear() {
	*ds = *NewDataStore()
}

func (ds *DataStore) State() State {
	return ds.state
}

func (ds *DataStore) Index() (InvertedIndex, error) {
	if ds.state == EmptyState {
		return nil, errors.New("no index loaded")
	}
	return ds.index, nil
}

func (ds *DataStore) UrlsVisited() ([]string, error) {
	if ds.state == EmptyState {
		return nil, errors.New("no index loaded")
	}
	return ds.urlsVisited, nil
}

func NewDataStore() *DataStore {
	return &DataStore{
		state: EmptyState,
	}
}
