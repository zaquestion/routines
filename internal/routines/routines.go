package routines

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/VojtechVitek/go-trello"
	"github.com/garyburd/go-oauth/oauth"
)

var (
	oauthClient = oauth.Client{
		TemporaryCredentialRequestURI: "https://trello.com/1/OAuthGetRequestToken",
		ResourceOwnerAuthorizationURI: "https://trello.com/1/OAuthAuthorizeToken",
		TokenRequestURI:               "https://trello.com/1/OAuthGetAccessToken",
	}
	clientKey      = os.Getenv("TRELLO_DEV_KEY")
	clientSecret   = os.Getenv("TRELLO_DEV_SECRET")
	trelloAuthPort = os.Getenv("TRELLO_AUTH_PORT")
	tempCreds      *oauth.Credentials

	trelloClient *trello.Client
	token        = make(chan *oauth.Credentials)
)

func init() {
	go func() {
		oauthClient.Credentials = oauth.Credentials{
			Token:  clientKey,
			Secret: clientSecret,
		}
		http.HandleFunc("/scrap", Scrap)
		http.HandleFunc("/TrelloAuth", TrelloAuth)
		http.HandleFunc("/oauth", OAuthCallback)
		if err := http.ListenAndServe(":"+trelloAuthPort, nil); err != nil {
			log.Fatalf("Error listening, %v", err)
		}

	}()
	go func() {
		creds := <-token
		var err error
		trelloClient, err = trello.NewAuthClient(clientKey, &creds.Token)
		if err != nil {
			log.Fatal(err)
		}
	}()

}

type oauthClientInfo struct {
	clientKey    string
	clientSecret string
}

type oauthToken struct {
	oauthToken  string
	oauthSecret string
}

func TrelloAuth(w http.ResponseWriter, r *http.Request) {
	callback := fmt.Sprintf("http://%s/oauth", r.Host)
	query := url.Values{
		"expiration": {"never"},
		"name":       {"RoutinesScrapper"},
	}
	var err error
	tempCreds, err = oauthClient.RequestTemporaryCredentials(nil, callback, query)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	http.Redirect(w, r, oauthClient.AuthorizationURL(tempCreds, nil), 302)
}

func OAuthCallback(w http.ResponseWriter, r *http.Request) {
	tokenCred, _, err := oauthClient.RequestToken(nil, tempCreds, r.FormValue("oauth_verifier"))
	if err != nil {
		http.Error(w, "Error getting request token, "+err.Error(), 500)
		return
	}
	if tokenCred != nil {
		fmt.Fprintln(w, "success")
		token <- tokenCred
	}
}

func Scrap(w http.ResponseWriter, r *http.Request) {
	if trelloClient == nil {
		fmt.Fprintln(w, "Need to authorized with trello")
		return
	}
	fmt.Println("HERE 1:")

	// User @trelloClient
	user, err := trelloClient.Member("zaquestion")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user.FullName)

	// @trelloClient Boards
	boards, err := user.Boards()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("HERE 2:", len(boards))

	if len(boards) > 0 {
		board := boards[0]
		fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)

		// @trelloClient Board Lists
		lists, err := board.Lists()
		if err != nil {
			log.Fatal(err)
		}

		for _, list := range lists {
			fmt.Println("   - ", list.Name)

			// @trelloClient Board List Cards
			cards, _ := list.Cards()
			for _, card := range cards {
				fmt.Println("      + ", card.Name)
			}
		}
	}
}

func List(start, end string) {
	return
}
