package User

import (
	"encoding/json"
)

type UserData struct {
		ID              string `json:"id"`
		Login           string `json:"login"`
		DisplayName     string `json:"display_name"`
		Type            string `json:"type"`
		BroadcasterType string `json:"broadcaster_type"`
		Description     string `json:"description"`
		ProfileImageUrl string `json:"profile_image_url"`
		OfflineImageUrl string `json:"offline_image_url"`
		ViewCount       int    `json:"view_count"`
		CreatedAt       string `json:"created_at"`
}

type User struct {
	Data	[]UserData `json:"data"`
}

func ParseUser(data []byte) (UserData,error){
	var response User 
	err := json.Unmarshal(data,&response)

	if err!=nil{
		return UserData{},err
	}
	return response.Data[0], nil
}
