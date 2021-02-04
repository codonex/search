package search

// Search super bir arama fonksiyonudur.
func Search(names []string) func(name string) bool {
	return func(name string) bool {
		for _, n := range names {
			if n == name {
				return true
			}
		}
		return false

	}

}
