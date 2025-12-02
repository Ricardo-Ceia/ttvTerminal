package Client

import (
	"fmt"
	"io"
	"net/http"
)

type TwitchClient struct{
	accessToken string
	clientID		string
	httpClient  *http.Client
}

const baseURL = "https://api.twitch.tv/helix"

func  NewTwitchClient(accessToken,clientID string) *TwitchClient{
	return &TwitchClient{
		accessToken:	accessToken,
		clientID:			clientID,
		httpClient:		&http.Client{},
	}
}

func (c* TwitchClient) Get(endpoint string) ([] byte,error){
	req,	err := http.NewRequest("GET",baseURL+endpoint,nil)
	
	if err!=nil{
		return nil,err	
	}
	req.Header.Add("Authorization",fmt.Sprintf("Bearer %s",c.accessToken))
	req.Header.Add("Client-Id",c.clientID)

	res,	err := c.httpClient.Do(req)
	
	if err!=nil{
		return nil,err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}
