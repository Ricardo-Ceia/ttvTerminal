package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"github.com/joho/godotenv"
)

type User struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

func main(){
	godotenv.Load()
	
	access_token := os.Getenv("TWITCH_ACCESS_TOKEN")
	client_id := os.Getenv("TWITCH_CLIENT_ID")
	username := "ronaldomadeir" 	
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users?login=" + username, nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	bearerString := fmt.Sprintf("Bearer %s", access_token)
	req.Header.Add("Authorization", bearerString)
	req.Header.Add("Client-Id", client_id)
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user User

	err = json.Unmarshal(body,&user)
	if err!=nil{
		log.Fatal(err)
	}

	user_id := user.Data[0].ID
	fmt.Println("id:",user_id)
}
