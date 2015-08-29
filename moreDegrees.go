package degreespkg

import (
	"degreespkg/structs"
	"fmt"
)

//To find more degrees between actors
func FindMoreDegree(links int, MovieObj1 structs.Movies, MovieObj2 structs.Movies) {
	fmt.Println("\n\tThere is no possible 2 degree links between them...\n")
	fmt.Println("\n\tPlease be patience while searching more possibilities...\n")

	fmt.Println("\n\t\t\tHave found", links, "links between them")
	fmt.Println("\n\t##########################################################")
}
