package main
import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"ttvTerminal/User"
	"ttvTerminal/Client"
	"ttvTerminal/Streams"
)

func main(){
	godotenv.Load()
	access_token := os.Getenv("TWITCH_ACCESS_TOKEN")
	client_id := os.Getenv("TWITCH_CLIENT_ID")
	username := "ronaldomadeir"

	ttvClient := Client.NewTwitchClient(access_token,client_id)
	userUrl := "/users?login=" + username
	
	userData,	err := ttvClient.Get(userUrl)	
	
	if err != nil {
		log.Fatal(err)
	}

	user, err := User.ParseUser(userData)
	if err!=nil{
		log.Fatal(err)
	}
	log.Printf("%+v",user)

	streamsData,err := ttvClient.Get("/streams/followed?user_id="+user.ID)
	if err!=nil{
		log.Fatal(err)
	}
	log.Printf("Raw respons: %s",streamsData)
	streamsParsedArr, err := Streams.ParseStreams(streamsData)
	if err!=nil{
		log.Fatal(err)
	}
	
	log.Printf("%+v",streamsParsedArr)
}
