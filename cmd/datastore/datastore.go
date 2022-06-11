package datastore

var _ DataStore = &DataStoreImpl{}

type DataStore interface {
}

type DataStoreImpl struct {
	St Store
	Tr TransactionStacker
}

func NewDataStore() *DataStoreImpl {
	return &DataStoreImpl{
		St: NewStore(),
		Tr: NewTransaction(),
	}
}
