package User

import(
	"os"
)

type User struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
}

func FillUserDataFile(data []byte,file *os.File) error{
	_,err := file.Write(data)
	return err
}

/*
func GetUserId(user User){
	return 
}
*/

