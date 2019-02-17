package instagram

import (
	"fmt"
	"time"
)

const (
	//Get Medias by shortCode template
	mediaInfoURL string = "https://instagram.com/p/%v/?__a=1"
	//Types
	TypeImage   = "GraphImage"
	TypeVideo   = "GraphVideo"
	TypeSidecar = "GraphSidecar"

	//Errors
	InvalidJson = "Json is invalid"
)

// An Medias describes an Instagram media info.
type Media struct {
	Id         string
	Type       string
	Caption    string
	Shortcode  string
	Comments   int
	Dimensions struct {
		Height int
		Width  int
	}
	DisplayUrl string
	Likes      int
	IsVideo    bool
	VideoViews int
	Resources  []Resource
	Time       time.Time
}

//Checker type media (image)
func (media *Media) isImage() bool {
	return media.Type == TypeImage
}

//Checker type media (video)
func (media *Media) isVideo() bool {
	return media.Type == TypeVideo
}

//Checker type media (sidecar)
func (media *Media) isSidecar() bool {
	return media.Type == TypeSidecar
}

//Get Medias from account structure
func (root *accountRoot) medias() ([]Media, error) {
	if len(root.EntryData.ProfilePage) == 0 {
		return []Media{}, fmt.Errorf(InvalidJson)
	}

	edges := root.EntryData.ProfilePage[0].Qraphql.User.EdgeOwnerToTimelineMedia.Edges
	if len(edges) == 0 {
		return []Media{}, fmt.Errorf(ErrMedias404)
	}

	var medias []Media
	for _, edge := range edges {
		media := new(Media)
		media.Id = edge.Node.Id
		media.Shortcode = edge.Node.Shortcode
		media.Type = edge.Node.Typename
		if len(edge.Node.EdgeMediaToCaption.Edges) > 0 {
			media.Caption = edge.Node.EdgeMediaToCaption.Edges[0].Node.Text
		} else {
			media.Caption = ""
		}
		media.Comments = edge.Node.EdgeMediaToComment.Count
		media.Dimensions.Height = edge.Node.Dimensions.Height
		media.Dimensions.Width = edge.Node.Dimensions.Width
		media.DisplayUrl = edge.Node.DisplayUrl
		media.Likes = edge.Node.EdgeLikedBy.Count
		media.IsVideo = edge.Node.IsVideo
		media.VideoViews = edge.Node.VideoViewCount
		media.Time = time.Unix(int64(edge.Node.TakenAtTimestamp), 0)
		media.Resources = edge.Node.ThumbnailResources
		medias = append(medias, *media)
	}

	return medias, nil
}

//Get count Medias from media struct
func (root *accountRoot) getCountMedia() uint32 {
	if len(root.EntryData.ProfilePage) == 0 {
		return 0
	}

	return root.EntryData.ProfilePage[0].Qraphql.User.EdgeOwnerToTimelineMedia.Count
}

//Get Medias from media structure
func (root *mediaRoot) media() (*Media, error) {

	edge := root.Graphql.ShortcodeMedia

	media := new(Media)
	media.Id = edge.Id
	media.Shortcode = edge.Shortcode
	media.Type = edge.Typename
	if len(edge.EdgeMediaToCaption.Edges) > 0 {
		media.Caption = edge.EdgeMediaToCaption.Edges[0].Node.Text
	} else {
		media.Caption = ""
	}
	media.Comments = edge.EdgeMediaToComment.Count
	media.Dimensions.Height = edge.Dimensions.Height
	media.Dimensions.Width = edge.Dimensions.Width
	media.DisplayUrl = edge.DisplayUrl
	media.Likes = 0
	media.IsVideo = edge.IsVideo
	media.VideoViews = 0
	media.Time = time.Unix(int64(edge.TakenAtTimestamp), 0)
	media.Resources = edge.DisplayResources

	return media, nil
}
