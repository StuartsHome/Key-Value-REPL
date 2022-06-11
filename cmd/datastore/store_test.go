package datastore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet_Success_NoTransaction(t *testing.T) {
	// Given
	key := "1"
	val := "2"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	err := ds.St.Set(key, val, ds)
	require.NoError(t, err)

	// Then
	got, err := ds.St.Get(key, ds)
	require.NoError(t, err)
	assert.Equal(t, got, val)
}

func TestGet_Success_Transaction(t *testing.T) {
	// Given
	key1, val1 := "2", "2"
	key2, val2 := "3", "3"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	ds.Tr.PushTransaction()

	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.Tr.Commit()
	require.NoError(t, err)

	err = ds.St.Set(key2, val2, ds)
	require.NoError(t, err)

	// Then
	got, err := ds.St.Get(key2, ds)
	require.NoError(t, err)
	assert.Equal(t, got, val2)

	got, err = ds.St.Get(key1, ds)
	require.Empty(t, got)
	assert.EqualError(t, err, fmt.Errorf("key %s not set in active transaction store", key1).Error())
}

func TestGet_Fail_KeyNotFound(t *testing.T) {
	// Given
	key := "2"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	got, err := ds.St.Get(key, ds)

	// Then - return value empty.
	assert.Empty(t, got)

	expectedErr := fmt.Errorf("key %s not set in global store", key).Error()
	assert.EqualError(t, err, expectedErr)
}

func TestGetAll_Success_NoTransaction(t *testing.T) {
	// Given
	key1, val1 := "3", "3"
	key2, val2 := "4", "4"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.St.Set(key2, val2, ds)
	require.NoError(t, err)

	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then
	assert.Equal(t, 2, len(got))
	assert.Contains(t, got, []string{key1, val1})
	assert.Contains(t, got, []string{key2, val2})
}

func TestGetAll_Success_Transaction(t *testing.T) {
	// Given
	key1, val1 := "5", "5"
	key2, val2 := "6", "6"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	ds.Tr.PushTransaction()

	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.Tr.Commit()
	require.NoError(t, err)

	err = ds.St.Set(key2, val2, ds)
	require.NoError(t, err)

	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then
	assert.Equal(t, 1, len(got))
	assert.NotContains(t, got, []string{key1, val1})
	assert.Contains(t, got, []string{key2, val2})
}

func TestSet_Success_NoTransaction(t *testing.T) {
	// Given
	key1, val1, val2 := "7", "7", "8"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.St.Set(key1, val2, ds)
	require.NoError(t, err)

	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then
	assert.Equal(t, 1, len(got))
	assert.Contains(t, got, []string{key1, val2})
}

func TestSet_Success_Transaction(t *testing.T) {
	// Given
	key1, val1, val2 := "8", "8", "9"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	ds.Tr.PushTransaction()

	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.Tr.Commit()
	require.NoError(t, err)

	err = ds.St.Set(key1, val2, ds)
	require.NoError(t, err)

	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then
	assert.Equal(t, 1, len(got))
	assert.Contains(t, got, []string{key1, val2})
}

func TestDelete_Success_NoTransaction(t *testing.T) {
	// Given
	key1, val1 := "9", "9"
	key2, val2 := "10", "10"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.St.Set(key2, val2, ds)
	require.NoError(t, err)

	err = ds.St.Delete(key1, ds)
	require.NoError(t, err)

	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then
	assert.Equal(t, 1, len(got))
	assert.NotContains(t, got, []string{key1, val1})
	assert.Contains(t, got, []string{key2, val2})
}

func TestDelete_Success_Transaction(t *testing.T) {
	// Given
	key1, val1 := "11", "11"
	key2, val2 := "12", "12"
	key3, val3 := "13", "13"
	ds := NewDataStore()
	GlobalStore = make(map[string]string)

	// When
	ds.Tr.PushTransaction()

	err := ds.St.Set(key1, val1, ds)
	require.NoError(t, err)

	err = ds.Tr.Commit()
	require.NoError(t, err)

	err = ds.St.Set(key2, val2, ds)
	require.NoError(t, err)

	err = ds.St.Set(key3, val3, ds)
	require.NoError(t, err)

	// When - deleting key from previous transaction produces an error.
	err = ds.St.Delete(key1, ds)
	assert.EqualError(t, err, fmt.Errorf("unable to delete key %s as not currently in transaction store", key1).Error())

	err = ds.St.Delete(key3, ds)
	require.NoError(t, err)

	got, err := ds.St.GetAll(ds)
	require.NoError(t, err)

	// Then
	assert.Equal(t, 1, len(got))
	assert.NotContains(t, got, []string{key3, val3})
	assert.Contains(t, got, []string{key2, val2})
}
