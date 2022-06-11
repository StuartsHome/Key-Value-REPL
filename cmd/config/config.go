package config

import (
	store "github.com/StuartsHome/key-value-REPL/cmd/datastore"
)

type Base struct {
	Globals *Data
}

type Data struct {
	AppName   string
	DataStore *store.DataStoreImpl
}

func NewData(appName string, dataStore *store.DataStoreImpl) *Data {
	return &Data{
		AppName:   appName,
		DataStore: dataStore,
	}
}
