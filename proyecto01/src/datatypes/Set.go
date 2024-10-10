package datatypes

type Set map[interface{}]bool

func (s Set) Add(element interface{}) {
	if !s.Contains(element) {
		s[element] = true
	}
}

func (s Set) Contains(element interface{}) bool {
	return s[element]
}

func (s Set) Remove(element interface{}) {
	delete(s, element)
}
