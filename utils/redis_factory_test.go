package utils

import (
	"fmt"
	"testing"
)

var (
	key     = "owie91q23"
	list    = "testList"
	factory = NewFactory(":42024")
)

// Add a new cached item
// Pre-condition.: The specific cached item does not exist in the cache
// Post-condition: the cached item is stored in the cache
func TestAddItem(t *testing.T) {
	t.Log("### TestAddItem ###")

	v := &ShortenedURL{
		Key:     key,
		Expires: 9001,
		Url:     "http://google.com/",
	}
	factory.Add(v)
	if s, err := factory.Exists(key); err != nil || s {
		t.Error("The item did exist in the cache", err)
	} else {
		t.Logf("Successfully added the item with key: %v \n", key)
	}
}

// Get a cached item
// Pre-condition.: a cached item exists in the cache
// Post-condition: the cached item is retrieved from the cache
func TestGetItem(t *testing.T) {
	t.Log("### TestGetItem ###")

	e := &ShortenedURL{}
	if v, err := factory.Get(key, e); err != nil || v == nil {
		t.Error("Could not get item: ", err)
	} else {
		t.Logf("Item: %v\n", v)
	}
	t.Logf("Successfully retrieved item with key: %v \n", key)
}

// Add item to list
// Pre-condition: none
// Post-condition: the item is added to the list with the given key
func TestAddItemList(t *testing.T) {
	t.Log("### TestAddItemList ###")

	values := make([]string, 10)
	for i := 1; i <= 10; i++ {
		values[i-1] = fmt.Sprintf("test%d", i)
	}

	for _, v := range values {
		if _, err := factory.RPush(list, v); err != nil {
			t.Error("Could not add item to list: ", err)
		} else {
			t.Logf("Item added, status: %q\n", v)
		}
	}
}
