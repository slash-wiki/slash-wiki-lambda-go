package handler

import (
	"github.com/gorilla/schema"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	var decoder = schema.NewDecoder()
	var message Message
	err := request.ParseForm()
	
	var parameters Parameters
	_ = decoder.Decode(&parameters, request.PostForm)

	if nil != err {
		message = Message{
			ResponseType: "ephemeral",
			Text:         "The form could not be parsed.",
		}
	} else {
		message = Message{
			ResponseType: "in_channel",
			Text:         fmt.Sprintf("Hello %s, the date and time is %s.", parameters.UserId, time.Now().Format(time.RFC850)),
		}
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(message)
}

type Parameters struct {
	Token       string `schema:"token"`
	TeamId      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelId   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserId      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Command     string `schema:"command"`
	Text        string `schema:"text"`
	ResponseUrl string `schema:"response_url"`
	TriggerId   string `schema:"trigger_id"`
	ApiAppId    string `schema:"api_app_id"`
}

type Message struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}