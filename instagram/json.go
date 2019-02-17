package instagram

import "encoding/json"

//Resource struct contains part information about media
type Resource struct {
	Src          string `json:"src"`
	ConfigWidth  int    `json:"config_width"`
	ConfigHeight int    `json:"config_height"`
}

//Time structure in which Account JSON is parsed
type accountRoot struct {
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	Hostname     string `json:"hostname"`
	EntryData    struct {
		ProfilePage []struct {
			Qraphql struct {
				User struct {
					Id                     string `json:"id"`
					Username               string `json:"username"`
					Biography              string `json:"biography"`
					BlockedByViewer        bool   `json:"blocked_by_viewer"`
					CountryBlock           bool   `json:"country_block"`
					ExternalUrl            string `json:"external_url"`
					ExternalUrlLinkshimmed string `json:"external_url_linkshimmed"`
					EdgeFollowedBy         struct {
						Count uint32 `json:"count"`
					} `json:"edge_followed_by"`
					EdgeFollow struct {
						Count uint32 `json:"count"`
					} `json:"edge_follow"`
					FullName                 string `json:"full_name"`
					HasChannel               bool   `json:"has_channel"`
					HasBlockedViewer         bool   `json:"has_blocked_viewer"`
					HighlightReelCount       int    `json:"highlight_reel_count"`
					IsBusinessAccount        bool   `json:"is_business_account"`
					IsPrivate                bool   `json:"is_private"`
					IsVerified               bool   `json:"is_verified"`
					ProfilePicUrl            string `json:"profile_pic_url"`
					ProfilePicUrlHd          string `json:"profile_pic_url_hd"`
					EdgeOwnerToTimelineMedia struct {
						Count    uint32 `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Id                 string `json:"id"`
								Typename           string `json:"__typename"`
								EdgeMediaToCaption struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								Shortcode          string `json:"shortcode"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								Dimensions       struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayUrl  string `json:"display_url"`
								EdgeLikedBy struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								Location     interface{} `json:"location"`
								GatingInfo   string      `json:"gating_info"`
								MediaPreview string      `json:"media_preview"`
								Owner        struct {
									Id       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								ThumbnailSrc       string     `json:"thumbnail_src"`
								ThumbnailResources []Resource `json:"thumbnail_resources"`
								IsVideo            bool       `json:"is_video"`
								VideoViewCount     int        `json:"video_view_count"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}

//Parcing JSON to account Struct.
func (root *accountRoot) parceJSON(scriptBytes []byte) error {

	err := json.Unmarshal(scriptBytes, &root)
	if err != nil {
		return err
	}
	return nil
}

//Time structure in which media JSON is parsed
type mediaRoot struct {
	Graphql struct {
		ShortcodeMedia struct {
			Id         string `json:"id"`
			Typename   string `json:"__typename"`
			Shortcode  string `json:"shortcode"`
			Dimensions struct {
				Height int `json:"height"`
				Width  int `json:"width"`
			} `json:"dimensions"`
			DisplayUrl         string     `json:"display_url"`
			DisplayResources   []Resource `json:"display_resources"`
			IsVideo            bool       `json:"is_video"`
			EdgeMediaToCaption struct {
				Edges []struct {
					Node struct {
						Text string `json:"text"`
					} `json:"node"`
				} `json:"edges"`
			} `json:"edge_media_to_caption"`
			EdgeMediaToComment struct {
				Count int `json:"count"`
			} `json:"edge_media_to_comment"`
			TakenAtTimestamp int `json:"taken_at_timestamp"`
			Owner            struct {
				Id            string    `json:"id"`
				ProfilePicUrl string `json:"profile_pic_url"`
				Username      string `json:"username"`
				FullName      string `json:"full_name"`
				IsPrivate     bool   `json:"is_private"`
			} `json:"owner"`
		} `json:"shortcode_media"`
	} `json:"graphql"`
}

//JSON to media Struct.
func (root *mediaRoot) readeJSON(scriptBytes []byte) error {

	err := json.Unmarshal(scriptBytes, &root)
	if err != nil {
		return err
	}
	return nil
}
