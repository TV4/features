// Deprecated: This package is no longer maintained.
package features

import "sync"

// Features allows for feature flagging
type Features interface {
	Enabled(key string) bool
	Enable(key string)
	Disable(key string)
}

// New creates new Features
func New(data map[string]bool) Features {
	if data == nil {
		return &features{data: map[string]bool{}}
	}

	return &features{data: data}
}

type features struct {
	sync.RWMutex
	data map[string]bool
}

func (f *features) Enabled(key string) bool {
	f.RLock()
	val := f.data[key]
	f.RUnlock()

	return val
}

func (f *features) Enable(key string) {
	f.Lock()
	f.data[key] = true
	f.Unlock()
}

func (f *features) Disable(key string) {
	f.Lock()
	delete(f.data, key)
	f.Unlock()
}
