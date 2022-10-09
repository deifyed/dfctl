package storage

func (s *Store) upsert(el Path) {
	s.remove(el)

	s.paths = append(s.paths, el)
}

func (s *Store) remove(el Path) {
	for i, path := range s.paths {
		if path.OriginalPath == el.OriginalPath {
			s.paths = append(s.paths[:i], s.paths[i+1:]...)

			break
		}
	}
}
