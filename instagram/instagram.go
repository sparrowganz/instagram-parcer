package instagram

import (
	"fmt"
	"strings"
)

const (
	//Validation Template
	instagram = "instagram.com"

	//Error
	ErrIsNotInstagram = "Url domain is not instagram.com"
	InvalidLink       = "Invalid Link"
)

//Get Account Struct from username
func GetAccountByUsername(username string) (*Account, error) {
	accountUrl := GetAccountUrl(username)
	return GetAccountByUrl(accountUrl)
}

// GetAccountByUrl try to find account by url.
func GetAccountByUrl(accountUrl string) (*Account, error) {
	if !strings.Contains(accountUrl, instagram) {
		return &Account{}, fmt.Errorf(ErrIsNotInstagram)
	}
	htmlBytes, err := newRequest(accountUrl).send()
	if err != nil {
		return &Account{}, err
	}
	scriptBytes, err := findJsonOnHtml(htmlBytes)
	if err != nil {
		return &Account{}, err
	}

	accountRoot := new(accountRoot)
	err = accountRoot.parceJSON(scriptBytes)
	if err != nil {
		return &Account{}, err
	}

	return accountRoot.account()
}

// GetMediaByUrl try to find media by url.
// URL should be like https://www.instagram.com/p/XXXXXXXXXX/
func GetMediaByUrl(url string) (*Media, error) {
	if !strings.Contains(url, instagram) {
		return &Media{}, fmt.Errorf(ErrIsNotInstagram)
	}
	params := strings.Split(url, "/")
	if len(params) < 5 {
		return &Media{}, fmt.Errorf(InvalidLink)
	}
	code := strings.Split(url, "/")[4]
	return GetMediaByShortCode(code)
}

// GetMediaByShortCode try to find media by code.
// Code can be find in URL to media, after p/.
// If URL to media is https://www.instagram.com/p/XXXXXXXXXX/,
// then code of the media is XXXXXXXXXX.
func GetMediaByShortCode(code string) (*Media, error) {
	parceUrl := fmt.Sprintf(mediaInfoURL, code)
	htmlBytes, err := newRequest(parceUrl).send()
	if err != nil {
		return &Media{}, err
	}
	/*scriptBytes, err := findJsonOnHtml(htmlBytes)
	if err != nil {
		return &Media{}, err
	}*/

	mediaRoot := new(mediaRoot)
	err = mediaRoot.readeJSON(htmlBytes)
	if err != nil {
		return &Media{}, err
	}

	return mediaRoot.media()
}

// GetAccountMedia try to get slice of user's media from the main page by username.
func GetLastMediasByUsername(username string) ([]Media, error) {
	accountUrl := GetAccountUrl(username)
	return GetLastMediasByUrl(accountUrl)
}

// GetAccountMedia try to get slice of user's media from the main page by url.
func GetLastMediasByUrl(accountUrl string) ([]Media, error) {
	if !strings.Contains(accountUrl, instagram) {
		return []Media{}, fmt.Errorf(ErrIsNotInstagram)
	}
	htmlBytes, err := newRequest(accountUrl).send()
	if err != nil {
		return []Media{}, err
	}
	scriptBytes, err := findJsonOnHtml(htmlBytes)
	if err != nil {
		return []Media{}, err
	}

	accountRoot := new(accountRoot)
	err = accountRoot.parceJSON(scriptBytes)
	if err != nil {
		return []Media{}, err
	}
	return accountRoot.medias()
}

//Substitute username in link template
func GetAccountUrl(username string) string {
	return fmt.Sprintf(accountInfoURL, username)
}

//Parce link and get username
// URL should be like https://www.instagram.com/username/
func GetUsernameFromUrl(url string) (string, error) {
	if !strings.Contains(url, instagram) {
		return "", fmt.Errorf(ErrIsNotInstagram)
	}
	params := strings.Split(url, "/")
	if len(params) < 5 {
		return "", fmt.Errorf(InvalidLink)
	}
	return params[4], nil
}
