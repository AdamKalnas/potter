package bookstore

func Cost(books map[string]int) (cost int) {
	number_of_books := 0
	for _, copies := range books {
		number_of_books += copies
	}
	return number_of_books * 800
}
