/*
Package registry is an expirmental package to facillitate altering the otto runtime via import.

This interface can change at any time.
*/
package registry

import "sync"

var registry []*Entry = make([]*Entry, 0)
var mutex sync.Mutex

type Entry struct {
	active bool
	source func() string
}

func newEntry(source func() string) *Entry {
	return &Entry{
		active: true,
		source: source,
	}
}

func (self *Entry) Enable() {
	self.active = true
}

func (self *Entry) Disable() {
	self.active = false
}

func (self Entry) Source() string {
	return self.source()
}

func Apply(callback func(Entry)) {

	// TODO try catch...

	for _, entry := range registry {
		if !entry.active {
			continue
		}
		callback(*entry)
	}
}

// Lock thread safety
func Lock() {
	mutex.Lock()
}

// Unlock thread safety
func UnLock() {
	mutex.Unlock()
}

func Register(source func() string) *Entry {
	Lock()
	entry := newEntry(source)
	registry = append(registry, entry)
	UnLock()
	return entry
}
