package gmail_imap

import (
	"context"
	"encoding/json"
	"net/url"
	"os"

	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-sasl"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type gmailService struct {
	client *client.Client
}

const (
	clientID     = "136883535673-o8s6kttf6sge7ks19brl0q9n44pnltqq.apps.googleusercontent.com"
	clientSecret = "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"
	redirectURL  = "https://google.com"
	scope        = "https://mail.google.com/"
)

// const path = "."
const file = "token.json"

func NewIMAPService() *gmailService {
	return &gmailService{}
}

func (s *gmailService) CreateIMAPClient(addr string) error {
	var err error
	s.client, err = client.DialTLS(addr, nil)
	return err
}

func (s *gmailService) Login(email, password string) error {
	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       []string{scope},
	}
	return s.authenticate(cfg, email)
}

func (s *gmailService) authenticate(cfg oauth2.Config, email string) error {
	token, err := tokenFromFile(file)
	if err != nil {
		token, err = tokenFromWeb(cfg)
		if err != nil {
			return err
		}
	}
	// Login to the IMAP server with XOAUTH2
	saslClient := sasl.NewOAuthBearerClient(&sasl.OAuthBearerOptions{
		Username: email,
		Token:    token.AccessToken,
	})
	return s.client.Authenticate(saslClient)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func tokenFromWeb(cfg oauth2.Config) (*oauth2.Token, error) {
	// // Ask for the user to login with his Google account
	// authUrl := cfg.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	// fmt.Printf("Go to %v", authUrl)

	const authCode = "4%2F0AdQt8qiMGSOI3UiKh3Fi6MV3Ut7qNf7fY-0NMaz9aiZxkFYbkcqY_ZyCynxXu8KuaE9I8Q"
	code, _ := url.QueryUnescape(authCode)
	//Get a token from the returned code
	//This token can be saved in a secure store to be reused later
	token, err := cfg.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}
	err = saveToken(file, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func saveToken(path string, token *oauth2.Token) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func (s *gmailService) Logout() error {
	return s.client.Logout()
}

func (s *gmailService) GetUnreadMessageCount() {

}

func (s *gmailService) SubsribeToInbox() error {
	if _, err := s.client.Select("INBOX", false); err != nil {
		return err
	}
	return nil
}
