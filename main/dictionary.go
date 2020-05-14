package main

import "sync"

type DictKey string
type DictVal string

//Dictionary struct
type Dictionary struct {
	elements map[DictKey]DictVal
	lock     sync.RWMutex
}

func (dict *Dictionary) Put(key DictKey, value DictVal) {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	if dict.elements == nil {
		dict.elements = make(map[DictKey]DictVal)
	}
	dict.elements[key] = value
}

func (dict *Dictionary) Remove(key DictKey) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	_, exists := dict.elements[key]
	if exists {
		delete(dict.elements, key)
	}

	return exists
}

func (dict *Dictionary) Contains(key DictKey) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	_, exists := dict.elements[key]
	return exists
}

func (dict *Dictionary) Find(key DictKey) DictVal {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	var value DictVal
	value, _ = dict.elements[key]
	return value
}

func (dict *Dictionary) Reset() {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	dict.elements = make(map[DictKey]DictVal)
}

func (dict *Dictionary) NumberOfElements() int {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	return len(dict.elements)
}

func (dict *Dictionary) GetKeys() (dictKeys []DictKey) {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	for key := range dict.elements {
		dictKeys = append(dictKeys, key)
	}
	return
}

func (dict *Dictionary) GetValues() (dictValues []DictVal) {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	for _, value := range dict.elements {
		dictValues = append(dictValues, value)
	}
	return
}
