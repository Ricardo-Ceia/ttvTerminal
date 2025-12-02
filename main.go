package main
import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"

	"ttvTerminal/Client"
	"ttvTerminal/Streams"
	"ttvTerminal/User"
)

type Model struct{
	streams	[] Streams.Stream
	cursor 	int
	err			error
}

func initialModel(streams []Streams.Stream) Model{
	return Model{streams: streams}
}

func (m Model) Init() tea.Cmd{
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model,tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:
		switch msg.String(){
		case "q","ctrl+c":
			return m,tea.Quit
		case "up","k":
			if m.cursor>0{
				m.cursor--
			}
		case "down","j":
			if m.cursor < len(m.streams)-1 {
				m.cursor++
			}
		}
	}
	return m,nil
}

func (m Model) View() string {
	s := "Followed Streams\n\n"

	if len(m.streams) == 0 {
		return s + "No streams online.\n\nPress q to quit."
	}

	for i, stream := range m.streams {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s - %s (%d viewers)\n",
			cursor,
			stream.UserName,
			stream.GameName,
			stream.ViewerCount,
		)
	}

	s += "\n↑/k up • ↓/j down • q quit"
	return s
}

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
	if err != nil {
		log.Fatal(err)
	}

	streamsData, err := ttvClient.Get("/streams/followed?user_id=" + user.ID)
	if err != nil {
		log.Fatal(err)
	}

	streamsParsedArr, err := Streams.ParseStreams(streamsData)

	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(initialModel(streamsParsedArr))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
