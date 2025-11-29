package main
import (
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"os/exec"
	"github.com/joho/godotenv"
	"ttvTerminal/User"
)

func main(){
	godotenv.Load()
	access_token := os.Getenv("TWITCH_ACCESS_TOKEN")
	client_id := os.Getenv("TWITCH_CLIENT_ID")
	username := "ronaldomadeir"
	userFileName := "userData.txt"
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users?login=" + username, nil)
	
	if err != nil {
		log.Fatal(err)
	}
	
	userFile,	err := os.Create(userFileName)
	if err!=nil{
		log.Fatal(err)
	} 
	defer userFile.Close()

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
	
	err = User.FillUserDataFile(body, userFile)
	if err != nil {
		log.Fatal(err)
	}
	
	userFile, err = os.Open(userFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer userFile.Close()
	
	user := User.GetUserInfo(userFile)
	fmt.Printf("%+v",user)

	cmd := exec.Command("rm",userFileName)
	err = cmd.Run()

	if err!=nil{
		log.Printf("Error (%v) executing rm %s",err,userFileName)
	}
}
