package datastore

import "fmt"

var _ Store = &StoreImpl{}

type Store interface {
	Get(key string, ds *DataStoreImpl) (string, error)
	GetAll(ds *DataStoreImpl) ([][]string, error)
	Set(key string, val string, ds *DataStoreImpl) error
	Delete(key string, ds *DataStoreImpl) error
}

type StoreImpl struct {
}

var GlobalStore = make(map[string]string)

func (st StoreImpl) Get(key string, ds *DataStoreImpl) (string, error) {
	activeTransaction := ds.Tr.Peek()
	if activeTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
			return val, nil
		}
		return "", fmt.Errorf("key %s not set in global store", key)
	} else {
		if val, ok := activeTransaction.store[key]; ok {
			return val, nil
		}
		return "", fmt.Errorf("key %s not set in active transaction store", key)
	}
}

func (st StoreImpl) GetAll(ds *DataStoreImpl) ([][]string, error) {
	activeTransaction := ds.Tr.Peek()
	if activeTransaction == nil {
		var vals [][]string
		for key, value := range GlobalStore {
			vals = append(vals, []string{key, value})
		}
		return vals, nil
	} else {
		var vals [][]string
		for key, value := range activeTransaction.store {
			vals = append(vals, []string{key, value})
		}
		return vals, nil
	}
}

func (st *StoreImpl) Set(key string, val string, ds *DataStoreImpl) error {
	activeTransaction := ds.Tr.Peek()
	if activeTransaction == nil {
		GlobalStore[key] = val
	} else {
		activeTransaction.store[key] = val
	}

	return nil
}

func (st *StoreImpl) Delete(key string, ds *DataStoreImpl) error {
	activeTransaction := ds.Tr.Peek()
	if activeTransaction == nil {
		if _, ok := GlobalStore[key]; ok {
			delete(GlobalStore, key)
			return nil
		} else {
			return fmt.Errorf("unable to delete key %s as not currently in global store", key)
		}
	} else {
		if _, ok := activeTransaction.store[key]; ok {
			delete(activeTransaction.store, key)
			return nil
		} else {
			return fmt.Errorf("unable to delete key %s as not currently in transaction store", key)
		}
	}
}

func NewStore() Store {
	return newStore()
}

func newStore() Store {
	return &StoreImpl{}
}
