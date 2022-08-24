//array -> collection of the values of same datatype

package main

import (
	"fmt"
	"reflect"
)

func main() {

	//declare array
	var scores = [8]int{4, 6, 8, 2, 0, 3, 5, 8}
	fmt.Println(scores)

	var friends = [3]string{"aoob", "klia", "cwg"}
	fmt.Println(friends)

	var speficicPostion = [9]int{0: 7, 4: 22, 8: 18}
	fmt.Println(speficicPostion)

	fmt.Printf("Length of array specificPosition is %v \n", len(speficicPostion))
	fmt.Printf("Datatype specificPosition is %T \n", speficicPostion)
	fmt.Println("Datatype specificPosition is", reflect.TypeOf(speficicPostion).Kind())

}
