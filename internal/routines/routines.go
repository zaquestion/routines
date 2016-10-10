package routines

import (
	"fmt"
	"github.com/VojtechVitek/go-trello"
	"log"
	"os"
)

func init() {
	appKey := os.Getenv("TRELLO_DEV_KEY")
	token := "b89aa8ed44cc86da6d1c61b8ddc9e4b06dc568beec2a9ad5dce489a55e1e7099"

	trello, err := trello.NewAuthClient(appKey, &token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HERE 1:")

	// User @trello
	user, err := trello.Member("zaquestion")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trello Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HERE 2:", len(boards))

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		// @trello Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// @trello Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}

func Scrap() {
}

func List(start, end string) {
	return
}
