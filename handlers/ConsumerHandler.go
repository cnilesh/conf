package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"service"
	"struct/handler"
	"utils"
)

func ConsumerHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	consumer := handler.ConsumePost{}
	decoder.Decode(&consumer)

	io.WriteString(w, service.ProcessConsumer(consumer.Sdp, utils.GetApi(consumer.PublisherId), utils.GetChannel(consumer.PublisherId)))
}
