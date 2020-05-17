package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"service"
	"struct/handler"
	"utils"
)


func PublishHandler(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	sdp := handler.SdpPost{}
	decoder.Decode(&sdp)
	w.WriteHeader(http.StatusOK)
	//localTrack := make(chan *webrtc.Track)
	answer, localTrack, api:= service.ProcessPublisher(sdp.Sdp)

	utils.PutChannel(sdp.Id, localTrack)
	utils.PutApi(sdp.Id, api)

	io.WriteString(w, answer)
}

func registerPublisher(id string) {

}