package service

import (
	"fmt"
	"github.com/pion/signal/signal"
	"github.com/pion/webrtc"
)

func ProcessConsumer(sdp string, api *webrtc.API, localTrack chan *webrtc.Track) string {
	fmt.Println("")
	fmt.Println("Curl an base64 SDP to start sendonly peer connection")

	peerConnectionConfig := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	recvOnlyOffer := webrtc.SessionDescription{}
	fmt.Println("Offer received")
	signal.Decode(sdp, &recvOnlyOffer)
	fmt.Println("Signal Decoded as well and waiting for new connection to join")
	// Create a new PeerConnection
	peerConnection, err := api.NewPeerConnection(peerConnectionConfig)
	fmt.Println("New connection joined")
	if err != nil {
		panic(err)
	}
	fmt.Println("Peerconnection established")

	_, err = peerConnection.AddTrack(<-localTrack)
	if err != nil {
		panic(err)
	}

	// Set the remote SessionDescription
	err = peerConnection.SetRemoteDescription(recvOnlyOffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Remote description for receiver set")
	// Create answer
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Answer for the receiver created")
	// Sets the LocalDescription, and starts our UDP listeners
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		panic(err)
	}
	fmt.Println("Local description for the receiver set successfully!")
	// Get the LocalDescription and take it to base64 so we can paste in browser
	fmt.Println(signal.Encode(answer))
	return signal.Encode(answer)
}

