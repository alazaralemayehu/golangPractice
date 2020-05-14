package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Tree struct {
	leftNode  *Tree
	value     int
	rightNode *Tree
}

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

var database *sql.DB

func main() {

	/*
			fmt.Println("Demonstrating List")
			var inList list.List
			inList.PushBack(11)
			inList.PushBack(133)

			inList.PushBack(9)

			for element := inList.Front(); element != nil; element = element.Next() {
				fmt.Println(element)
			}

			fmt.Println("Demonstrating Doubli Linked List")

			l1 := DoubliLinkedList{}
			l1.AddToHead(12)

			l1.AddToHead(2)
			l1.AddAfter(2, 9)
			l1.PrintAll()
			fmt.Println(l1.NodeBetweenValues(2, 12))

			fmt.Println("Demonstrating Set")

			set1 := Set{}
			set1.New()
			set1.AddElement(12)
			set1.AddElement(5)
			set1.AddElement(4)
			set1.AddElement(9)

			set2 := Set{}
			set2.New()
			set2.AddElement(12)
			set2.AddElement(53)
			set2.AddElement(9)
			set2.AddElement(78)

			fmt.Println(set1.Intersect(&set2))
			fmt.Println(set1.Union(&set2))

			fmt.Println("Demonstrating Queue")

			que := Queue{}
			order := &Order{}

			order.New(1, 12, "Abebe", "string product")
			que.Add(order)
			order.New(2, 12, "Abebe Besso ", "string product 2")
			que.Add(order)
			fmt.Println(order)
			fmt.Println(*que[1])

			fmt.Println("Demonstrating Stack")

			stack := &Stack{}
			stack.New()

			element := &Element{}
			element.NewElement(10)
			stack.Push(element)
			element.NewElement(4)
			stack.Push(element)
			element.NewElement(30)
			stack.Push(element)
			element.NewElement(70)

		fmt.Println("Demonstrating Dictionary")

			fmt.Println("Demonstrating Dictionary ")

			var dict *Dictionary = &Dictionary{}
			dict.Put("1", "1")
			dict.Put("2", "2")
			dict.Put("3", "3")
			dict.Put("4", "4")
			fmt.Println(dict)
	*/
	// t := &BinarySearchTree{}
	// t.InsertElement(2, 2)
	// t.InsertElement(1, 1)
	// t.InsertElement(3, 3)
	// t.removeNode(3)

	// fmt.Println("demonstrating Bubble sort")
	// numbers := []int{44, 27, 350, 100, 110, 88}
	// QuickSort(numbers)
	// fmt.Println(numbers)
	// string1 := "The quick brow fox jumps over the lazy dog"
	// needle := "llo"

	// i, err := StringSearch(string1, needle)
	// if err == nil {
	// 	fmt.Println(i)
	// 	fmt.Println(string([]rune(string1)[i : i+len(needle)]))
	// } else {
	// 	fmt.Println("Nothing Found")
	// }
	// InterviewQuestionSolutions.IsUniques("Hello")
	// fmt.Println(InterviewQuestionSolutions.ArePurmutation("hello", "llehl"))

	// t := &datastructure.BinarySearchTree{}
	// t.InsertElement(10, 10)
	// t.InsertElement(5, 5)
	// t.InsertElement(15, 15)
	// t.InsertElement(19, 13)
	// t.InsertElement(2, 22)
	// t.InsertElement(2, 2)
	// fmt.Println(InterviewQuestionSolutions.FindTheClosestValueINBST(t, 12))
	// fmt.Println(InterviewQuestionSolutions.URLify("hello world", 11))
	// InterviewQuestionSolutions.SmartFibonacci(6)
	// m := make(map[int]int)
	// m[1] = 0
	// m[2] = 1

	// fmt.Print(InterviewQuestionSolutions.FibonaciWithMemoize(9, m))
	// fmt.Println(InterviewQuestionSolutions.ThreeSums([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0))
	// fmt.Println(InterviewQuestionSolutions.SmallestDifference([]int{-1, 5, 10, 20, 27, 3}, []int{27, 133, 135, 15, 15}))
	// A := []int{-1, -3, -4, 7, -1}
	// fmt.Println(InterviewQuestionSolutions.Solution(A))
	// N := 1041
	// maxGap := 0

	// currentGap := 0
	// previousState := -1
	// counting := false

	// for N > 0 {
	// 	remainder := N % 2
	// 	if remainder == 1 {
	// 		if previousState == 0 && counting {
	// 			if maxGap < currentGap {
	// 				maxGap = currentGap
	// 			}
	// 			currentGap = 0
	// 		}
	// 		previousState = 1
	// 		counting = true
	// 	} else {
	// 		if counting {
	// 			currentGap += 1
	// 		}
	// 		fmt.Println(currentGap)
	// 		previousState = 0
	// 	}

	// 	N = N / 2

	// }
	// fmt.Println(maxGap)

	// fmt.Println(maxGap)

}

//  getting a connection
func GetConnection() {
	databaseDriver := "mysql"
	databseUser := "golanguser"
	databasePassword := "golang"
	databaseName := "golangdatabase"
	var err error
	database, err = sql.Open(databaseDriver, databseUser+":"+databasePassword+"@/"+databaseName)
	if err != nil {
		panic("Unable to connect " + err.Error())
	}
	err = database.Ping()
	if err != nil {
		panic("unable to connect " + err.Error())
	}

}

func GetUsers() {
	preStmt, err := database.Prepare("SELECT id, email, firstname, lastname  from users")

	user1 := User{}
	if err != nil {
		panic(err.Error())
	}
	rows, err := preStmt.Query()
	fmt.Println(rows)
	if err != nil {
		panic(err.Error())

	}
	for rows.Next() {
		rows.Scan(&user1.ID, &user1.Email, &user1.FirstName, &user1.LastName)
		fmt.Println(user1)

	}
}

func (t *Tree) insert(v int) {
	if t != nil {
		if t.leftNode == nil {
			t.leftNode = &Tree{nil, v, nil}
		} else {
			if t.rightNode == nil {
				t.rightNode = &Tree{nil, v, nil}
			}
			if t.leftNode != nil {
				t.leftNode.insert(v)
			} else {
				t.rightNode.insert(v)
			}
		}
	} else {
		t = &Tree{nil, v, nil}
	}

}

func (t *Tree) printNode() {
	if t == nil {
		return
	} else {
		fmt.Println((t.value))
	}
	t.rightNode.printNode()
}

func InsertUser(user User) {
	var insertStmt *sql.Stmt
	var err error
	insertStmt, err = database.Prepare("INSERT INTO users(email, password, firstname, lastname) VALUES (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	r, err := insertStmt.Exec(user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		panic(err.Error())
	}
	result, _ := r.LastInsertId()
	fmt.Println("The Id of last inserted item is ", result)
}
