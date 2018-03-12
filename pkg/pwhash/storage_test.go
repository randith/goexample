package pwhash

import (
	"testing"
	"time"
)

/**
 * not the best testing as it relies on hard-coded timing
 */
func Test_inMemoryStore(testing *testing.T) {
	store := NewInMemoryStore()

	value1 := "value 1"
	key1, err := store.Set(value1)
	if err != nil {
		testing.Error("Unable to set value into store")
	}

	// it is expected that this is before the actual write occurs
	actual1, err := store.Get(key1)
	if actual1 != "" {
		testing.Error("Expected not to be able to get value immedietely")
	}

	// ensure this is more than the slowSet waits
	time.Sleep(6 * time.Second)

	actual2, err := store.Get(key1)
	if err != nil {
		testing.Error("Unable to get value from the store")
	}

	if actual2 != value1 {
		testing.Errorf("Value get is incorrect, actual='%s' expected='%s'", actual2, value1)
	}

}
