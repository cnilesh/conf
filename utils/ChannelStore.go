package utils

import (
	"github.com/pion/webrtc"
	"sync"
)

var lock = &sync.Mutex{}

// type global
type channelSingleton map[string]chan *webrtc.Track
type apiSingleton map[string]*webrtc.API

var (
	channelInstance channelSingleton
	apiInstance apiSingleton
)

func NewChannelInstance() channelSingleton {

	lock.Lock()
	defer lock.Unlock()

	if channelInstance == nil {

		channelInstance = make(channelSingleton) // <-- thread safe
	}

	return channelInstance
}

func NewApiInstance() apiSingleton {

	lock.Lock()
	defer lock.Unlock()

	if apiInstance == nil {

		apiInstance = make(apiSingleton) // <-- thread safe
	}

	return apiInstance
}

func PutChannel(id string, track chan *webrtc.Track) {
	if channelInstance == nil {
		channelInstance = NewChannelInstance()
	}
	channelInstance[id] = track
}

func GetChannel(id string) chan *webrtc.Track{
	return channelInstance[id]
}

func PutApi(id string, api *webrtc.API) {
	if apiInstance == nil {
		apiInstance = NewApiInstance()
	}
	apiInstance[id] = api
}

func GetApi(id string) *webrtc.API {
	return apiInstance[id]
}
