package main

import "fmt"

type Node struct {
	id      int
	isAlive bool
}

//bullyAlgorithm finds the leader in the distributed system
func bullyAlgorithm(nodes []Node, initiatorID int) int {
	var electedLeader int

	for _, node := range nodes {
		if node.isAlive && node.id > initiatorID {
			fmt.Printf("Node %d challanges node %d\n", initiatorID, node.id)
			electedLeader = node.id
			break
		}
	}

	if electedLeader == 0 {
		electedLeader = initiatorID
	}
	return electedLeader
}

func main() {
	nodes := []Node{
		{id: 1, isAlive: true},
		{id: 2, isAlive: true},
		{id: 3, isAlive: true},
		{id: 4, isAlive: true},
		{id: 5, isAlive: false},
	}

	initiatorID := 2
	leader := bullyAlgorithm(nodes, initiatorID)
	fmt.Printf("Leader elected: Node %d\n", leader)

}
