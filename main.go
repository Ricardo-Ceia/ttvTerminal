package main
import (
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"github.com/joho/godotenv"
	"ttvTerminal/User"
)

func main(){
	godotenv.Load()
	
	access_token := os.Getenv("TWITCH_ACCESS_TOKEN")
	client_id := os.Getenv("TWITCH_CLIENT_ID")
	username := "ronaldomadeir" 	
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users?login=" + username, nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	userFile,	err := os.Create("userData.txt")
	if err!=nil{
		log.Fatal(err)
	} 
	//Think about this defer -> the file should be closed when its no longer needed 
	defer userFile.Close()

	bearerString := fmt.Sprintf("Bearer %s", access_token)
	req.Header.Add("Authorization", bearerString)
	req.Header.Add("Client-Id", client_id)
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//Think about this defer -> the body should be closed when its no longer needed
	defer res.Body.Close()
	
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	err = User.FillUserDataFile(body,userFile)
	if err!=nil{
		log.Fatal(err)
	}
}
