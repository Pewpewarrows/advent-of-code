package main

import (
    advent "github.com/Pewpewarrows/advent-of-code/pkg"
    "bufio"
    "bytes"
    "fmt"
    "strings"
)

func main() {
    var caveGraph graph
    advent.Execute(scanInputData, &caveGraph)

    pathCount := caveGraph.pathCount(1, 1)
    fmt.Println("part one:", pathCount)

    pathCount = caveGraph.pathCount(1, 2)
    fmt.Println("part two:", pathCount)
}

// assumed all edges are bidirectional
type graph struct {
    nodes []*node
    start *node
    end *node
}

func newGraph(edges [][2]string) *graph {
    var nodes []*node
    var start *node
    var end *node

    for _, edge := range edges {
        leftNode := labelledNode(nodes, edge[0])
        rightNode := labelledNode(nodes, edge[1])

        if leftNode == nil {
            leftNode = &node{edge[0], []*node{}}
            nodes = append(nodes, leftNode)

            // TODO: const for start/end labels
            if leftNode.label == "start" {
                start = leftNode
            }

            if leftNode.label == "end" {
                end = leftNode
            }
        }

        if rightNode == nil {
            rightNode = &node{edge[1], []*node{}}
            nodes = append(nodes, rightNode)

            if rightNode.label == "start" {
                start = rightNode
            }

            if rightNode.label == "end" {
                end = rightNode
            }
        }

        // TODO: ensure dedup if redundant edge is provided
        leftNode.peers = append(leftNode.peers, rightNode)
        rightNode.peers = append(rightNode.peers, leftNode)
    }

    // TODO: ensure there are start/end nodes

    return &graph{nodes, start, end}
}

func (caveGraph *graph) pathCount(maxSmallCaveVisits int, maxSingleSmallCaveVisits int) (pathCount int) {
    paths := allPathsFromChain(path{caveGraph.start}, make(map[*node]int), caveGraph, maxSmallCaveVisits, maxSingleSmallCaveVisits)

    return len(paths)
}

type node struct {
    label string
    peers []*node
}

func (n node) String() string {
    b := bytes.NewBufferString(n.label)

    if len(n.peers) == 0 {
        goto end
    }

    b.WriteString(" -> {")

    for i, node := range n.peers {
        if i != 0 {
            b.WriteString(", ")
        }
        b.WriteString(node.label)
    }

    b.WriteString("}")

end:
    return b.String()
}

func (n *node) isSmallCave() bool {
    return advent.IsLower(n.label)
}

func labelledNode(nodes []*node, label string) *node {
    for _, n := range nodes {
        if n.label == label {
            return n
        }
    }

    return nil
}

// TODO: this is an ugly number of params, find a better way
func allPathsFromChain(chain path, visitCountBySmallNode map[*node]int, caveGraph *graph, maxSmallCaveVisits int, maxSingleSmallCaveVisits int) (paths []path) {
    // We assume that no two big caves are directly connected by an edge, which
    // would cause an infinite loop of potential paths

    n := chain[len(chain) - 1]

    nextVisitCountBySmallNode := make(map[*node]int)
    for k, v := range visitCountBySmallNode {
        nextVisitCountBySmallNode[k] = v
    }
    if n.isSmallCave() {
        nextVisitCountBySmallNode[n]++
    }

    chainIsAlreadyRevisitingChosenCave := false
    for _, v := range nextVisitCountBySmallNode {
        if v > maxSmallCaveVisits {
            chainIsAlreadyRevisitingChosenCave = true
            break
        }
    }

    if n == caveGraph.end {
        paths = append(paths, chain)
        return
    }

peerLoop:
    for _, peer := range n.peers {
        if peer.label == "start" {
            continue
        }
        for k, v := range nextVisitCountBySmallNode {
            if peer == k {
                // if this peer is the chosen one, v > maxSmallCaveVisits
                if chainIsAlreadyRevisitingChosenCave && ((v == maxSmallCaveVisits) || (v == maxSingleSmallCaveVisits)) {
                    continue peerLoop
                } else if (maxSmallCaveVisits == maxSingleSmallCaveVisits) && (v == maxSmallCaveVisits) {
                    continue peerLoop
                }

                break
            }
        }
        paths = append(paths, allPathsFromChain(append(chain, peer), nextVisitCountBySmallNode, caveGraph, maxSmallCaveVisits, maxSingleSmallCaveVisits)...)
    }

    return
}

func scanInputData(scanner *bufio.Scanner, inputDataPtr interface{}) {
    // caveGraph := *inputDataPtr.(*graph)
    var edges [][2]string

    for scanner.Scan() {
        var edge [2]string
        for i, text := range strings.Split(scanner.Text(), "-") {
            edge[i] = text
            // TODO: handle i > 1 warning
        }
        edges = append(edges, edge)
    }

    *inputDataPtr.(*graph) = *newGraph(edges)
}

type path []*node

func (p path) String() string {
    var b bytes.Buffer

    for i, node := range p {
        if i != 0 {
            b.WriteString(" -> ")
        }
        b.WriteString(fmt.Sprintf("%s", node.label))
    }

    return b.String()
}
