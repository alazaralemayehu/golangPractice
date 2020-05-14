package main

type Set struct {
	integerMap map[int]bool
}

func (s *Set) New() {
	s.integerMap = make(map[int]bool)
}

func (s *Set) Contains(val int) bool {
	_, exists := s.integerMap[val]
	return exists
}

func (s *Set) AddElement(val int) {
	if !(s.Contains(val)) {
		s.integerMap[val] = true
	}
}

func (s *Set) deleteElement(val int) {
	delete(s.integerMap, val)
}

func (s *Set) Intersect(anotherSet *Set) *Set {
	var intersectSet *Set = &Set{}
	intersectSet.New()
	for val, _ := range s.integerMap {
		if anotherSet.Contains(val) {
			intersectSet.AddElement(val)
		}
	}
	return intersectSet
}

func (s *Set) Union(anotherSet *Set) *Set {
	var unionSet *Set = &Set{}
	unionSet.New()
	for val, _ := range s.integerMap {
		if !anotherSet.Contains(val) {
			unionSet.AddElement(val)
		}
	}

	for val2, _ := range anotherSet.integerMap {
		if !s.Contains(val2) {
			unionSet.AddElement(val2)
		}
	}
	for val := range s.Intersect(anotherSet).integerMap {
		unionSet.AddElement(val)
	}
	return unionSet
}
