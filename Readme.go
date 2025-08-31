Start go
hello world in go 
package main

import "fmt"

func main() {
    fmt.Println("Hello Go World!!")
}

------------------------------------

to run go code in terminal:- @Satyams-MacBook-Air GO Language % go run main.go 
Hello Go World!!

------------------------------------

//to print multiple lines 
package main

import "fmt"

func main() {
    fmt.Println("Welcome to do list\n")
	fmt.Println("Watch go crash course")
	fmt.Println("Watch Nana's good boy ")
	fmt.Println("Rewards myself with a cheekcake")
}



------------------------------------

Data types in GO:-  Strings,Booleans,Integer,Maps,Arrays.

------------------------------------

package main

import "fmt"

func main() {
    fmt.Println("###Welcome to do list!###")
	fmt.Println("1. Watch go crash course")
	fmt.Println("2. Watch Nana's good boy ")
	fmt.Println("3. Rewards myself with a cheekcake")
	fmt.Println(4)
	fmt.Println(5)
}

------------------------------------
storing values in variables
package main

import "fmt"

func main() {


var taskone = "2. Watch Nana's good boy "

    fmt.Println("###Welcome to do list!###")
	fmt.Println("1. Watch go crash course")
	fmt.Println(taskone)
	fmt.Println("3. Rewards myself with a cheekcake")
	fmt.Println(4)
	fmt.Println(5)
}
------------------------------------

List data types in GO:- 
1. Arrays  2. Slice
package main

import "fmt"

func main() {


// var taskone = "2. Watch Nana's good boy "
// var tasktwo="1. Watch go crash course"
// var maxItensInGroup=20

//below line is slice
var taskitem =[ ]string{"Watch Nana's good boy ","Watch go crash course", "Learn Go", "Learn Python", "Learn JavaScript"}


    fmt.Println("###Welcome to do list!###")
	fmt.Println()
	fmt.Println(taskitem)
	
}
 

------------------------------------

package main

import "fmt"

func main() {


var taskone = "2. Watch Nana's good boy "
var tasktwo="1. Watch go crash course"
var maxItensInGroup=20

//below line is slice example
var taskitem =[ ]string{taskone,tasktwo}


    fmt.Println("###Welcome to do list!###")
	fmt.Println()
	fmt.Println(taskitem)
	fmt.Println(maxItensInGroup)
	
}
------------------------------------

Array vs Slice
//this is slice where size is not defined
var taskitem =[]string{taskone,tasktwo}

//this is array where size is  defined
var taskitem =[2]string{taskone,tasktwo}


------------------------------------
Loops in Go:-

1. Basic for loop

2. for loop with condition

3. for range loop

example:- 
package main

import "fmt"

func main() {
var taskone = "2. Watch Nana's good boy "
var tasktwo="1. Watch go crash course"
var taskthree="3. Watch Naruto"
//var maxItensInGroup=20

//below line is slice example
var taskitem =[]string{taskone,tasktwo,taskthree}

	for _,task := range taskitem{
		fmt.Println(task)
	}
}

example with index:-

package main

import "fmt"

func main() {
var taskone = " Watch Nana's good boy "
var tasktwo=" Watch go crash course"
var taskthree=" Watch Naruto"
//var maxItensInGroup=20

//below line is slice example
var taskitem =[]string{taskone,tasktwo,taskthree}

	
	for index,task := range taskitem{
		fmt.Println(index,task)
	}
}
//formated output for above

package main

import "fmt"

func main() {
var taskone = " Watch Nana's good boy "
var tasktwo=" Watch go crash course"
var taskthree=" Watch Naruto"
//var maxItensInGroup=20

//below line is slice example
var taskitem =[]string{taskone,tasktwo,taskthree}
	for index,task := range taskitem{
		//fmt.Println(index + 1,".",task)
		fmt.Printf("%d. %s\n",index +1,task)
	}
}

--------fuctions-----
 



