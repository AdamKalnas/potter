package bookstore

const (
	book_cost = 800
)

var (
	discounts map[int]int
)

func init() {
	discounts = make(map[int]int)
	discounts[0] = 0
	discounts[1] = 0
	discounts[2] = 5
}

func Cost(books map[string]int) (cost int) {
	bundles := bundleBooks(books)
	return getCostForBundles(bundles)
}

func bundleBooks(books map[string]int) (bundles []int) {
	var (
		shaved int
	)
	remaining_books := getTotalNumberOfBooks(books)
	bundles = []int{}
	for i := 0; remaining_books > 0; i++ {
		shaved, books = getBooksInRound(books)
		remaining_books -= shaved
		bundles = append(bundles, shaved)
	}
	return bundles
}

func getCostForBundles(bundles []int) (cost int) {
	cost = 0
	for _, bundle := range bundles {
		cost += bundle * book_cost * (100 - discounts[bundle]) / 100
	}
	return cost
}

func getTotalNumberOfBooks(books map[string]int) int {
	number_of_books := 0
	for _, copies := range books {
		number_of_books += copies
	}
	return number_of_books
}

func getBooksInRound(books map[string]int) (num int, booksPruned map[string]int) {
	booksPruned = make(map[string]int)
	round := 0
	for book, copies := range books {
		if copies > 0 {
			round++
			booksPruned[book] = copies - 1
		}
	}
	return round, booksPruned
}
