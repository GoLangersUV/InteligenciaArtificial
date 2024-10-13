package datatypes

type Set map[interface{}]bool

func (s Set) Add(element interface{}) {
	if !s.Contains(element) {
		s[element] = true
	}
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Contains(element interface{}) bool {
	return s[element]
}

func (s Set) Remove(element interface{}) {
	delete(s, element)
}

// GetAndRemove retrieves the element from the set and removes it.
// It returns the element if it exists, otherwise it returns nil.
func (s *Set) GetAndRemove(element interface{}) interface{} {
	if s.Contains(element) {
		// Remove the element from the set
		s.Remove(element)
		// Return the element
		return element
	}
	// Return nil if the element does not exist
	return nil
}
