package linkedlist

import (
	"testing"
)

func TestAppend(t *testing.T) {
	linkedList := &LinkedList[string]{}
	linkedList.Append("5")
	v, err := linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if *v != "5" {
		t.Fail()
	}
	linkedList.Append("10")
	v, err = linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if *v != "5" {
		t.Fail()
	}
	if linkedList.Length != 2 {
		t.Fatal("Length should be 2")
	}
}

func TestPrepend(t *testing.T) {
	linkedList := &LinkedList[string]{}
	linkedList.Prepend("5")
	v, err := linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if *v != "5" {
		t.Fatal()
	}
	linkedList.Prepend("10")
	v, err = linkedList.Get(0)
	if err != nil {
		t.Error(err)
	}
	if *v != "10" {
		t.Fatal()
	}
	if linkedList.Length != 2 {
		t.Fatal("Length should be 2")
	}
}

func TestInsertAt(t *testing.T) {
	linkedList := &LinkedList[string]{}
	linkedList.Prepend("5")
	linkedList.Prepend("10")
	linkedList.Prepend("t10")
	linkedList.Prepend("19sw0")
	linkedList.InsertAt("Corinthians", 2)
	v, _ := linkedList.Get(2)
	if *v != "Corinthians" {
		t.Fatal("should be equals to Corinthians")
	}
}

func TestRemoveAt(t *testing.T) {
	linkedList := &LinkedList[string]{}
	linkedList.Append("5")
	linkedList.Append("10")
	linkedList.Append("BBB")
	linkedList.Append("AAA")
	linkedList.RemoveAt(2)
	v, _ := linkedList.Get(2)
	if *v != "AAA" {
		t.Fail()
	}
	v, _ = linkedList.Get(1)
	if *v != "10" {
		t.Fail()
	}
}

func TestRemoveAtOutOfRange(t *testing.T) {
	linkedList := &LinkedList[string]{}
	linkedList.Append("5")
	linkedList.Append("10")
	linkedList.Append("BBB")
	linkedList.Append("AAA")
	err := linkedList.RemoveAt(4)
	if err == nil {
		t.Fail()
	}
}
