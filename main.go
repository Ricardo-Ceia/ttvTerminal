package main
import (
	"log"
	"os"
	"os/exec"
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
	userFileName := "userData.txt"

	ttvClient := Client.NewTwitchClient(access_token,client_id)
	userUrl := "/users?login=" + username
	
	
	userFile,	err := os.Create(userFileName)
	if err!=nil{
		log.Fatal(err)
	} 
	defer userFile.Close()

	userData,	err := ttvClient.Get(userUrl)	
	
	if err != nil {
		log.Fatal(err)
	}
	
	err = User.FillUserDataFile(userData, userFile)
	
	if err != nil {
		log.Fatal(err)
	}
	
	userFile, err = os.Open(userFileName)
	
	if err != nil {
		log.Fatal(err)
	}

	defer userFile.Close()
	
	user := User.GetUserInfo(userFile)
	user_id := user.Data[0].ID
	
	streamsData,err := ttvClient.Get("/streams/followed?user_id="+user_id)
	if err!=nil{
		log.Fatal(err)
	}
	log.Printf("Raw respons: %s",streamsData)
	streamsParsedArr, err := Streams.ParseStreams(streamsData)
	if err!=nil{
		log.Fatal(err)
	}
	
	log.Printf("%+v",streamsParsedArr)
	cmd := exec.Command("rm",userFileName)
	err = cmd.Run()
	
	if err!=nil{
		log.Printf("Error (%v) executing rm %s",err,userFileName)
	}
}
