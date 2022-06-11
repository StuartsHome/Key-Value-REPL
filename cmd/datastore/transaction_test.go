package datastore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration Test:
// set value
// start transaction
// set value
// commit
// set value
// set value
// delete value
// readall
// abort transaction
// readall
// abort transaction
// readall
// abort transaction - error

func TestTransactionIntegration_Success(t *testing.T) {
	// Given
	key1, val1 := "1", "1"
	key2, val2 := "2", "2"
	key3, val3 := "3", "3"

	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When - set value
	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	// When - start transaction
	ds.Tr.PushTransaction()

	// When - set value
	err = ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	// When - commit
	err = ds.Tr.Commit()
	require.NoError(t, err)

	// When - set value
	err = ds.St.Set(key2, val2, ds)
	require.NoError(t, err)

	// When - set value
	err = ds.St.Set(key3, val3, ds)
	require.NoError(t, err)

	// When - delete value
	err = ds.St.Delete(key2, ds)
	require.NoError(t, err)

	// When - readall
	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then - readall
	assert.Equal(t, 1, len(got))
	assert.NotContains(t, got, []string{key2, val2})
	assert.Contains(t, got, []string{key3, val3})

	// When - abort transaction
	err = ds.Tr.PopTransaction()
	require.NoError(t, err)

	// When - readall
	got, err = ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then - readall
	assert.Equal(t, 1, len(got))
	assert.NotContains(t, got, []string{key3, val3})
	assert.Contains(t, got, []string{key1, val1})

	// Then - abort transaction
	err = ds.Tr.PopTransaction()
	require.NoError(t, err)

	// When - readall
	got, err = ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then - readall
	assert.Equal(t, 1, len(got))
	assert.Contains(t, got, []string{key1, val1})

	// Then - abort transaction - error
	err = ds.Tr.PopTransaction()
	assert.EqualError(t, err, fmt.Errorf("error: no active transaction.").Error())
}
