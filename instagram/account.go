package instagram

import (
	"fmt"
)

const (
	//Main webpage
	accountInfoURL = "https://instagram.com/%v/"

	//ERROR
	ErrMedias404 = "Medias not found"
)

// An Account describes an Instagram account info.
type Account struct {
	Id          string
	Username    string
	FullName    string
	Biography   string
	ExternalUrl string
	Followers   uint32
	Follows     uint32
	IsPrivate   bool
	IsVerified  bool
	MediaCount  uint32
	Image       struct {
		ProfilePicUrl   string
		ProfilePicUrlHd string
	}
}

//Filling Account struct from JSON struct
func (root *accountRoot) account() (*Account, error) {

	if len(root.EntryData.ProfilePage) == 0 {
		return &Account{}, fmt.Errorf(InvalidJson)
	}
	account := new(Account)
	account.Id = root.EntryData.ProfilePage[0].Qraphql.User.Id
	account.Biography = root.EntryData.ProfilePage[0].Qraphql.User.Biography
	account.Username = root.EntryData.ProfilePage[0].Qraphql.User.Username
	account.IsPrivate = root.EntryData.ProfilePage[0].Qraphql.User.IsPrivate
	account.Image.ProfilePicUrl = root.EntryData.ProfilePage[0].Qraphql.User.ProfilePicUrl
	account.Image.ProfilePicUrlHd = root.EntryData.ProfilePage[0].Qraphql.User.ProfilePicUrlHd
	account.FullName = root.EntryData.ProfilePage[0].Qraphql.User.FullName
	account.IsVerified = root.EntryData.ProfilePage[0].Qraphql.User.IsVerified
	account.MediaCount = root.EntryData.ProfilePage[0].Qraphql.User.EdgeOwnerToTimelineMedia.Count
	account.Follows = root.EntryData.ProfilePage[0].Qraphql.User.EdgeFollow.Count
	account.Followers = root.EntryData.ProfilePage[0].Qraphql.User.EdgeFollowedBy.Count
	account.ExternalUrl = root.EntryData.ProfilePage[0].Qraphql.User.ExternalUrl
	return account, nil
}
