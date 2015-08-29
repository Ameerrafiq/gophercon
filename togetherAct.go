package degreespkg

import (
	"degreespkg/structs"
	"fmt"
)

//CheckTogetherAct function to check whether they have relation within a film and returns bool value
func CheckTogetherAct(Actor1Obj structs.ActorMoviesList, Actor2Obj structs.ActorMoviesList) bool {
	//defer wg.Done()
	degreeCount := 1
	var k int
	var l int
	togetherAct := false
	for i := range Actor1Obj.Movies {
		for j := range Actor2Obj.Movies {
			if Actor1Obj.Movies[i].Name == Actor2Obj.Movies[j].Name {
				togetherAct = true
				k = i
				l = j
				break
			}
		}
	}
	if togetherAct == true {
		fmt.Println("\n\t#####################################################")
		fmt.Println("\n\t\tThey have worked together in a film\n")
		fmt.Println("\t\tDegrees of Separation: ", degreeCount, "\n")
		fmt.Println("\t\t", degreeCount, ".Movie: ", Actor1Obj.Movies[k].Name)
		fmt.Println("\t\t", Actor1Obj.Movies[k].Role, ": ", Actor1Obj.Name)
		fmt.Println("\t\t", Actor2Obj.Movies[l].Role, ": ", Actor2Obj.Name, "\n")
		fmt.Println("\n\t#####################################################")
	}
	return togetherAct
}
