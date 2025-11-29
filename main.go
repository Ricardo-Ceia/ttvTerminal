package main

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	access_token := os.Getenv("TWITCH_ACCESS_TOKEN")
	client_id := os.Getenv("TWITCH_CLIENT_ID")
	fmt.Println(access_token)
	fmt.Println(client_id)
	
	req,	err := http.NewRequest("GET","https://api.twitch.tv/helix/users?login=twitchdev",nil)
	
	if err!=nil{
		log.Fatal(err)
	}

	bearerString := fmt.Sprintf("Bearer %s",access_token)
	req.Header.Add("Authorization",bearerString)
	req.Header.Add("Client-Id",client_id)
	client := &http.Client{}
	res,	err := client.Do(req)

	if err!=nil{
		log.Fatal(err)
	}

	body,	err := io.ReadAll(res.Body)
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("res:",string(body))
	defer res.Body.Close()
}
