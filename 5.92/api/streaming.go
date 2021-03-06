package api // import "github.com/severecloud/vksdk/5.92/api"

import "encoding/json"

// StreamingGetServerURLResponse struct
type StreamingGetServerURLResponse struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
}

// StreamingGetServerURL allows to receive data for the connection to Streaming API.
// https://vk.com/dev/streaming.getServerUrl
func (vk VK) StreamingGetServerURL() (response StreamingGetServerURLResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("streaming.getServerUrl", map[string]string{})
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StreamingGetSettingsResponse struct
type StreamingGetSettingsResponse struct {
	MonthlyLimit string `json:"monthly_limit"`
}

// StreamingGetSettings allows to receive monthly tier for Streaming API.
// https://vk.com/dev/streaming.getSettings
func (vk VK) StreamingGetSettings() (response StreamingGetSettingsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("streaming.getSettings", map[string]string{})
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StreamingGetStatsResponse struct
type StreamingGetStatsResponse []struct {
	EventType string `json:"event_type"`
	Stats     []struct {
		Timestamp int `json:"timestamp"`
		Value     int `json:"value"`
	} `json:"stats"`
}

// StreamingGetStats allows to receive statistics for prepared and received events in Streaming API.
// https://vk.com/dev/streaming.getStats
func (vk VK) StreamingGetStats(params map[string]string) (response StreamingGetStatsResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("streaming.getStats", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StreamingGetStemResponse struct
type StreamingGetStemResponse struct {
	Stem string `json:"stem"`
}

// StreamingGetStem allows to receive the stem of the word.
// https://vk.com/dev/streaming.getStem
func (vk VK) StreamingGetStem(params map[string]string) (response StreamingGetStemResponse, vkErr Error) {
	rawResponse, vkErr := vk.Request("streaming.getStem", params)
	if vkErr.Code != 0 {
		return
	}

	err := json.Unmarshal(rawResponse, &response)
	if err != nil {
		panic(err)
	}

	return
}

// StreamingSetSettings allows to set monthly tier for Streaming API.
// https://vk.com/dev/streaming.setSettings
func (vk VK) StreamingSetSettings(params map[string]string) (vkErr Error) {
	_, vkErr = vk.Request("streaming.setSettings", params)

	return
}
