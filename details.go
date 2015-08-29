package degreespkg

import (
	"degreespkg/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//FetchActorDetails function to fetch actor details and umarshals json data
func FetchActorDetails(actor1 string, actor2 string) error {
	var actTogether bool
	noerr := 1
	var Actor1Obj structs.ActorMoviesList
	var Actor2Obj structs.ActorMoviesList
	for i := 0; i < 2; i++ {
		if i == 0 {
			data := FetchDetails(actor1)
			err := json.Unmarshal(data, &Actor1Obj)
			if err != nil {
				noerr = 0
				fmt.Println("Please enter valid url for actor name")
				return nil
			}
		} else {
			data := FetchDetails(actor2)
			err := json.Unmarshal(data, &Actor2Obj)
			if err != nil {
				noerr = 0
				fmt.Println("Please enter valid url for actor name")
				return nil
			}
		}
	}
	if noerr == 1 {
		//var wg sync.WaitGroup
		//wg.Add(1)
		actTogether = CheckTogetherAct(Actor1Obj, Actor2Obj)
		//wg.Wait()
	}
	if actTogether == false {
		FetchMovieDetails(Actor1Obj, Actor2Obj)
	}
	return nil
}

//FetchDetails function to fetch actor details via moviebuff url and returns byte data
func FetchDetails(actor string) []byte {
	baseUrl := "http://data.moviebuff.com/"
	baseUrl += actor
	cli := &http.Client{}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	res, err := cli.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	re, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return re
}
