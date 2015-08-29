package degreespkg

import (
	"degreespkg/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//FetchMovieDetails function to fetch movie details for an actor and unmarshals json data
func FetchMovieDetails(Actor1Obj structs.ActorMoviesList, Actor2Obj structs.ActorMoviesList) {
	fmt.Println("\n\t##########################################################")
	fmt.Println("\n\t\tThey have not worked together in a film yet\n")
	fmt.Println("\tPlease wait while searching possible links between them...\n")
	var MovieObj1 structs.Movies
	var MovieObj2 structs.Movies
	actor1name := Actor1Obj.Name
	actor2name := Actor2Obj.Name
	links := 0
	degreeCount := 2
	for i := range Actor1Obj.Movies {
		data := MovieDetails(Actor1Obj.Movies[i].Url)
		err := json.Unmarshal(data, &MovieObj1)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		actor1_movie1_role := Actor1Obj.Movies[i].Role
		for j := range Actor2Obj.Movies {
			data1 := MovieDetails(Actor2Obj.Movies[j].Url)
			err := json.Unmarshal(data1, &MovieObj2)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			actor2_movie2_role := Actor2Obj.Movies[j].Role
			for k := range MovieObj1.Cast {
				for l := range MovieObj2.Cast {
					if MovieObj1.Cast[k].Name == MovieObj2.Cast[l].Name {
						midactor_movie1_role := ""
						midactor_movie2_role := ""
						links += 1
						for m := range MovieObj1.Cast {
							if MovieObj1.Cast[m].Name == MovieObj1.Cast[k].Name {
								midactor_movie1_role = MovieObj1.Cast[m].Role
							}
						}
						for n := range MovieObj2.Cast {
							if MovieObj2.Cast[n].Name == MovieObj1.Cast[k].Name {
								midactor_movie2_role = MovieObj2.Cast[n].Role
							}
						}
						fmt.Println("\n\t\tDegrees of Separation: ", degreeCount, "\n")
						fmt.Println("\t\t", degreeCount-1, ".Movie: ", MovieObj1.Name)
						fmt.Println("\t\t", actor1_movie1_role, ": ", actor1name)
						fmt.Println("\t\t", midactor_movie1_role, ": ", MovieObj1.Cast[k].Name, "\n")
						fmt.Println("\t\t", degreeCount, ".Movie: ", MovieObj2.Name)
						fmt.Println("\t\t", actor2_movie2_role, ": ", actor2name)
						fmt.Println("\t\t", midactor_movie2_role, ": ", MovieObj1.Cast[k].Name, "\n")
						fmt.Println("\n\t\tWant to find more links between them?\n\t\tIf yes press y or press n : ")
						more := ""
						_, err := fmt.Scanln(&more)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						if more == "n" {
							more = ""
							goto endofsearch
						} else if more == "y" {
							more = ""
							fmt.Println("\n\tPlease wait while searching possible links between them...\n")
							continue
						} else {
							fmt.Println("\n\t\tPlease enter correct values")
							goto endofsearch
						}
					}
				}
			}
		}
	}
	if links == 0 {
		FindMoreDegree(links, MovieObj1, MovieObj2)
		goto end
	}

endofsearch:
	fmt.Println("\n\t\t\tHave found", links, "links between them")
	fmt.Println("\n\t##########################################################")
end:
}

//FetchDetails function to fetch movie details via moviebuff url and returns byte data
func MovieDetails(movie string) []byte {
	baseUrl := "http://data.moviebuff.com/"
	baseUrl += movie
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
