package Streams

import (
	"encoding/json"
)

type Stream struct {
  StreamId    string `json:"id"`
  UserId      string `json:"user_id"`
  UserLogin   string `json:"user_login"`
	UserName    string `json:"user_name"`
	GameId      string `json:"game_id"`
	GameName    string `json:"game_name"`
	Title       string `json:"title"`
	ViewerCount int    `json:"viewer_count"`
	StartedAt   string `json:"started_at"`
	IsMature    bool   `json:"is_mature"`
}

type StreamsResponse struct {
	Data []Stream `json:"data"`
}

func ParseStreams(data []byte) ([]Stream,error){
	var response StreamsResponse
	err := json.Unmarshal(data,&response)

	if err!=nil{
		return nil,err
	}
	return response.Data, nil
}


