package main

import (
	"fmt"
)

const QueueSize = 5

type Node struct{
	 Data string
	 Left *Node
	 Right *Node
}

type Queue struct{
    Head *Node
	Tail *Node
	Length int
}

type Cache struct{
	Queue Queue
	Hash Hash
}

type Hash map[string]*Node

func NewCache() Cache{
	return Cache{Queue:NewQueue(),Hash: Hash{}}
}

func NewQueue() Queue{
   head := &Node{}
   tail := &Node{}

   head.Right = tail
   tail.Left = head

   return Queue{Head: head,Tail:tail}
}

func (c *Cache) Check(key string){
	  
	node := &Node{}

	if val,ok := c.Hash[key];ok{
		node = c.Remove(val)

	} else{
		node =  &Node{Data: key}
		
	}

	c.Add(node)
	c.Hash[key] = node 
	 
	 
}

func (c *Cache) Add(node *Node){
	 fmt.Printf("Adding: %s\n",node.Data)

	 temp:=c.Queue.Head.Right

	 c.Queue.Head.Right = node

	 node.Right = temp

	 temp.Left = node

	 node.Left = c.Queue.Head

	 c.Queue.Length+=1

	 if c.Queue.Length > QueueSize{
		c.Remove(c.Queue.Tail.Left)
	 }

	// c.Hash[node.Data] = node

}

func (c *Cache) Display(key string){
	  
	if c.Queue.Length <=0{
		 fmt.Println("No Elelment present in the Queue!!")
		 return 
	}

	fmt.Printf("Current Cache Length %d\n",c.Queue.Length)
	currNode := c.Queue.Head.Right

	fmt.Print("[")
	for currNode!=c.Queue.Tail{
		 fmt.Printf("{%s}",currNode.Data)
		 if currNode.Right.Right !=c.Queue.Tail{
			fmt.Printf("::")
		 }

		 currNode = currNode.Right
	}

	fmt.Print("]\n")

	//fmt.Println("Not Found the Key in the Cache!!")

}

func (c *Cache) Remove(node *Node) *Node{
 
	 fmt.Printf("Removing: %s\n",node.Data)

	 left := node.Left
	 right :=node.Right

	 left.Right = right

	 right.Left  = left

	 c.Queue.Length-=1

	 delete(c.Hash,node.Data)

	 return node

}
func main(){
	 fmt.Println("Starting Cache")

	 cache := NewCache()

	 for _,word := range []string{"Alex","Bob","James","Lee","Sophie","Devine","James","Sophie","Annie","Sansa","Vanessa"}{
		cache.Check(word)
		cache.Display(word)

	 }
}
