package degreespkg

import (
	"fmt"
	"os"
)

//StartServer function to fetch Actor details and checking cmd line arguments
func StartServer() {
	var actor1 string
	var actor2 string
	if len(os.Args) > 3 && (os.Args[1] == "degrees") {
		actor1 = os.Args[2]
		actor2 = os.Args[3]
		err := FetchActorDetails(actor1, actor2)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Println("Error in arguments")
		return
	}
}
