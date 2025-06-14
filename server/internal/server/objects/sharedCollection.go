package objects

import "sync"

// A generic thread safe map of objects with auto-incrementing IDs
type SharedCollection[T any] struct {
	objectsMap map[uint64]T
	nextId     uint64
	mapMux     sync.Mutex
}

func NewSharedCollection[T any](capacity ...int) *SharedCollection[T] {
	var newObjMap map[uint64]T
	if len(capacity) > 0 {
		newObjMap = make(map[uint64]T, capacity[0])
	} else {
		newObjMap = make(map[uint64]T)
	}
	return &SharedCollection[T]{
		objectsMap: newObjMap,
		nextId:     1,
	}
}

// Add an object to the map with a given ID (if provided) or the next available ID.

func (s *SharedCollection[T]) Add(obj T, id ...uint64) uint64 {
	s.mapMux.Lock()
	defer s.mapMux.Unlock()
	thisId := s.nextId

	if len(id) > 0 {
		thisId = id[0]
	}

	s.objectsMap[thisId] = obj
	s.nextId++
	return thisId
}

// Remove an object from the map by ID
func (s *SharedCollection[T]) Remove(id uint64) {
	s.mapMux.Lock()
	defer s.mapMux.Unlock()
	delete(s.objectsMap, id)
}

// Call the callback function for each object in the map
func (s *SharedCollection[T]) ForEach(callback func(uint64, T)) {
	// Create a local copy while holding the lock
	s.mapMux.Lock()
	LocalCopy := make(map[uint64]T, len(s.objectsMap))
	for id, obj := range s.objectsMap {
		LocalCopy[id] = obj
	}
	s.mapMux.Unlock()

	// Iterate over local copy without holding the lock
	for id, obj := range LocalCopy {
		callback(id, obj)
	}
}

// Get an object with the given ID, if it exists, otherwise nil.
// Also returns a boolean indicating whether the object was found.
func (s *SharedCollection[T]) Get(id uint64) (T, bool) {
	s.mapMux.Lock()
	defer s.mapMux.Unlock()
	obj, found := s.objectsMap[id]
	return obj, found
}

// Get the number of objects in the map
func (s *SharedCollection[T]) Count() int {
	s.mapMux.Lock()
	defer s.mapMux.Unlock()
	return len(s.objectsMap)
}
