package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// set up heap
type NextHeap []string

func (h NextHeap) Len() int           { return len(h) }
func (h NextHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h NextHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NextHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}
func (h *NextHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	inputFile, _ := filepath.Abs("./testInput")
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// inboundNodes will have {node: [incoming edge nodes]}
	inboundNodes := map[string][]string{}
	outboundNodes := map[string][]string{}
	var startNode, endNode string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curString := scanner.Text()
		// parse string into key value
		splitString := strings.Split(curString, " ")
		parentChar := splitString[7]
		childChar := splitString[1]
		// add to inboundNodes
		if val, ok := inboundNodes[parentChar]; ok {
			val = append(val, childChar)
			inboundNodes[parentChar] = val
		} else {
			deps := make([]string, 0)
			deps = append(deps, childChar)
			inboundNodes[parentChar] = deps
		}
		// add to outboundNodes
		if val, ok := outboundNodes[childChar]; ok {
			val = append(val, parentChar)
			outboundNodes[childChar] = val
		} else {
			deps := make([]string, 0)
			deps = append(deps, parentChar)
			outboundNodes[childChar] = deps
		}
		// add final edge (with no deps) to inboundNodes
		if _, found := (inboundNodes[childChar]); !found {
			deps := make([]string, 0)
			inboundNodes[childChar] = deps
			startNode = childChar
		}
		// add final edge (with no deps) to outboundNodes
		if _, found := (outboundNodes[parentChar]); !found {
			deps := make([]string, 0)
			outboundNodes[parentChar] = deps
			endNode = parentChar
		}

	}
	fmt.Printf("ingraph: %v, startnode: %s\n", inboundNodes, startNode)
	fmt.Printf("outgraph: %v, endnode: %s\n", outboundNodes, endNode)
	result := getPriorityOrder(startNode, inboundNodes, outboundNodes)
	fmt.Printf("%v\n", result)
}

func getPriorityOrder(start string, edgesIn map[string][]string, edgesOut map[string][]string) []string {
	var result []string
	nextHeap := &NextHeap{}
	heap.Init(nextHeap)
	heap.Push(nextHeap, start)
	for nextHeap.Len() > 0 {
		thisNode := heap.Pop(nextHeap).(string)
		result = append(result, thisNode)
		for _, neighbor := range edgesOut[thisNode] {
			fmt.Printf("%s\n", neighbor)
			//remove from ingraph array
			edgesIn[neighbor] = removeNode(edgesIn[neighbor], thisNode)
			if len(edgesIn[neighbor]) == 0 {
				heap.Push(nextHeap, neighbor)
			}
		}
	}
	return result
}

func removeNode(nodes []string, node string) []string {
	nodeIndex := -1
	result := nodes
	for i, n := range nodes {
		if node == n {
			nodeIndex = i
		}
	}
	if nodeIndex > -1 {
		result = append(nodes[:nodeIndex], nodes[nodeIndex+1:]...)
	}
	return result
}
