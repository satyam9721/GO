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
 package main

import "fmt"

var taskone = " Watch Nana's good boy "
var tasktwo = " Watch go crash course"
var taskthree = " Watch Naruto"

// var maxItemsInGroup = 20

// below line is slice example
var taskItems = []string{taskone, tasktwo, taskthree}

func main() {
	fmt.Println("### Welcome to Task Manager ###")
	//printTasks() // call the function so it actually prints tasks
}

func printTasks() {
	fmt.Println("Printing tasks")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

--------variables scope & Global scope------
package main

import "fmt"

var taskone = " Watch Nana's good boy "
var tasktwo = " Watch go crash course"
var taskthree = " Watch Naruto"

// var maxItemsInGroup = 20

// below line is slice example
var taskItems = []string{taskone, tasktwo, taskthree}

func main() {
	fmt.Println("### Welcome to Task Manager ###")
	printTasks() // call the function so it actually prints tasks
}

func printTasks() {
	fmt.Println("Printing tasks")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

//Passing values of one functions to another functions, so we can acces the variables and array

//ex-1 passing parameter
package main

import "fmt"



func main() {
	fmt.Println("### Welcome to Task Manager ###")
	var taskone = " Watch Nana's good boy "
var tasktwo = " Watch go crash course"
var taskthree = " Watch Naruto"

// var maxItemsInGroup = 20

// below line is slice example
var taskItems = []string{taskone, tasktwo, taskthree}
	printTasks(taskItems) // call the function so it actually prints tasks
}

func printTasks(taskItems[]string) {
	fmt.Println("Printing tasks")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}
//working with listsitems , adding functions / calling functions inside the functions

package main

import "fmt"



func main() {
	fmt.Println("### Welcome to Task Manager ###")
	var taskone = " Watch Nana's good boy "
var tasktwo = " Watch go crash course"
var taskthree = " Watch Naruto"

// var maxItemsInGroup = 20

// below line is slice example
var taskItems = []string{taskone, tasktwo, taskthree}
	printTasks(taskItems) // call the function so it actually prints tasks
	fmt.Println()
	addTask(taskItems," Let's Run")
}

func printTasks(taskItems[]string,) {
	fmt.Println("Printing tasks")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems[]string,newTask string){
var updatedTaskItems = append(taskItems,newTask)
 printTasks(updatedTaskItems)

}

//Return value from functions executions and stored in var/list

package main

import "fmt"



func main() {
	fmt.Println("### Welcome to Task Manager ###")
	var taskone = " Watch Nana's good boy "
var tasktwo = " Watch go crash course"
var taskthree = " Watch Naruto"

// var maxItemsInGroup = 20

// below line is slice example
var taskItems = []string{taskone, tasktwo, taskthree}
	printTasks(taskItems) // call the function so it actually prints tasks
	fmt.Println()
	//saving in variables returned updated slice
	taskItems = addTask(taskItems," Let's Run")
	fmt.Println()
	printTasks(taskItems) // call the function so it actually prints tasks
}

func printTasks(taskItems[]string,) {
	fmt.Println("Printing tasks")
	for index, task := range taskItems {
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

//returns updated slice of tasks
func addTask(taskItems[]string,newTask string) []string{ 
var updatedTaskItems = append(taskItems,newTask)
  return updatedTaskItems;

}
//WebApi code in Go


package main

import (
	"fmt"
	"net/http"
)

func main() {
	// define route and handler
	http.HandleFunc("/", helloUser)

	// start web server (moved after route definition for best practice)
	fmt.Println("Server starting at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// define handler function
func helloUser(writer http.ResponseWriter, r *http.Request) {
	greeting := "Hello User, Welcome to Task Manager"
	fmt.Fprintln(writer, greeting)
}

//go in the browser type http://localhost:8080/

