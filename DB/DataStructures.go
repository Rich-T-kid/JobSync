package DB
// IMplement needed Data strucutres for the Program here
// Implement a way to load messages from database into a queue (order in which they originally came)
// 	

type linkedList struct{
	data node
	next *linkedList
}

type node struct{
	data interface{}

}

type queue struct{

}

