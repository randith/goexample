package pwhash

import (
	"testing"
)

/**
 * simplistic minimal testing, given time should extend testing to the interesting bit which is locking access
 */
func Test_inMemoryStore(testing *testing.T)  {
	store := NewInMemoryStore()

	value1 := "value 1"
	key1, err := store.Set(value1)
	if err != nil {
		testing.Error("Unable to set value into store")
	}

	actual1, err := store.Get(key1)
	if err != nil {
		testing.Error("Unable to get value from store")
	}

	if actual1 != value1 {
		testing.Errorf("Value get is incorrect, actual='%s' expected='%s'", actual1, value1)
	}

}
