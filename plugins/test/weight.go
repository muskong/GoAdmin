package main

import (
	"log"
	"time"
)

type Node struct {
	Name    string
	Current int
	Weight  int
}

func main() {
	nodes := []*Node{
		&Node{"a", 0, 4},
		&Node{"b", 0, 1},
		// &Node{"c", 0, 1},
	}

	for i := 0; i < 7; i++ {
		best := SmoothWeight(nodes)
		time.Sleep(1 * time.Second)
		log.Printf("best2 %#v \n", best)
	}
}

func printNode(nodes []*Node) {
	for _, node := range nodes {
		log.Printf(">>>>%#+v<<<<\n", node)
	}
}
func SmoothWeight(nodes []*Node) (best *Node) {
	if len(nodes) == 0 {
		return
	}

	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}

		total += node.Weight
		node.Current += node.Weight

		log.Printf("%#v", total)
		log.Printf("%#v", node)
		if best == nil || node.Current > best.Current {
			best = node
			log.Printf("best %#v \n", best)
		}
	}

	if best == nil {
		return
	}

	best.Current -= total
	log.Printf("best 2 %#v \n", best)
	printNode(nodes)

	return
}
