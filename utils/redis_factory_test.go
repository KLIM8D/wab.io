package utils

import (
	"fmt"
	"github.com/ugorji/go/codec"
	"log"
	"testing"
)

var (
	key     = "owie91q23"
	list    = "testList"
	factory = NewFactory(":42024")
)

func TestEncoding(t *testing.T) {
	e := &ShortenedURL{
		Key:     key,
		Expires: 9001,
		Url:     "http://google.com/",
	}
	var b []byte
	var mh codec.MsgpackHandle
	enc := codec.NewEncoderBytes(&b, &mh)
	if err := enc.Encode(e); err != nil {
		t.Error("Could not encode item into msgpack format", err)
	} else {
		log.Printf("Item: (%d) %v\n", len(b), b)
	}

	d := &ShortenedURL{}
	dec := codec.NewDecoderBytes(b, &mh)
	if err := dec.Decode(d); err != nil {
		t.Error("Could not decode item back from msgpack format", err)
	} else {
		log.Printf("Item: %v\n", d)
	}

}

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
		log.Printf("Item: %v\n", v)
		log.Printf("Successfully retrieved item with key: %v \n", key)
	}
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
