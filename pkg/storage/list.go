package storage

func upsert(db *store, el Path) {
	remove(db, el)

	db.Paths = append(db.Paths, el)
}

func remove(db *store, el Path) {
	for i, path := range db.Paths {
		if path.OriginalPath == el.OriginalPath {
			db.Paths = append(db.Paths[:i], db.Paths[i+1:]...)

			break
		}
	}
}
