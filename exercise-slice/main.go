package main

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	newBooks := [5]string{"ulysses", "fire"}
	if books == newBooks {

	}
	books = newBooks

	games := []string{"pokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}
	newGames = games
	games = nil

}
