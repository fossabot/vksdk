package vksdk

import "encoding/json"

// UtilsCheckLinkResponse struct
type UtilsCheckLinkResponse struct{}

// TODO utils.checkLink Checks whether a link is blocked in VK.

// UtilsDeleteFromLastShortenedResponse struct
type UtilsDeleteFromLastShortenedResponse struct{}

// TODO utils.deleteFromLastShortened Deletes shortened link from user's list.

// UtilsGetLastShortenedLinksResponse struct
type UtilsGetLastShortenedLinksResponse struct{}

// TODO utils.getLastShortenedLinks Returns a list of user's shortened links.

// UtilsGetLinkStatsResponse struct
type UtilsGetLinkStatsResponse struct{}

// TODO utils.getLinkStats Returns stats data for shortened link.

// UtilsGetServerTimeResponse struct
type UtilsGetServerTimeResponse struct{}

// TODO utils.getServerTime Returns the current time of the VK server.

// UtilsGetShortLinkResponse struct
type UtilsGetShortLinkResponse struct{}

// TODO utils.getShortLink Allows to receive a link shortened via vk.cc.

// UtilsResolveScreenNameResponse struct
type UtilsResolveScreenNameResponse struct {
	Type     string `json:"type,omitempty"`
	ObjectID int    `json:"object_id,omitempty"`
}

// UtilsResolveScreenName detects a type of object (e.g., user, community, application) and its ID by screen name.
// TODO testing utilx.resolveScreenName
func (vk *VK) UtilsResolveScreenName(params map[string]string) (UtilsResolveScreenNameResponse, error) {
	var response UtilsResolveScreenNameResponse

	rawResponse, err := vk.Request("utils.resolveScreenName", params)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}