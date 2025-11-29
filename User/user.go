package User

import (
	"encoding/json"
	"bytes"
	"os"
)

type User struct {
	Data []struct {
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
	} `json:"data"`
}

func FillUserDataFile(data []byte,file *os.File) error{
	var user User
	decoder := json.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&user)

	if err!=nil{
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent(""," ")
	err = encoder.Encode(user)
	return err
}

/*
func GetUserId(user User){
	return 
}
*/

