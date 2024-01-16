package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const (
	ringDegrees  = 360
	ringMaxNodes = 180
)

type HashRing struct {
	nodes map[string]int
}

func (h *HashRing) getRandRingPosition() (int, error) {
	var newPosition int

	if len(h.nodes) == 0 {
		return rand.Intn(ringDegrees), nil
	}

	if len(h.nodes) == ringMaxNodes {
		return 0, errors.New("maximum number of nodes reached")
	}

	for {
		newPosition = rand.Intn(ringDegrees)
		if !h.checkPosition(newPosition) {
			break
		}
	}
	return newPosition, nil
}

func (h *HashRing) AddNode(nodeName string) {
	ringPosition, err := h.getRandRingPosition()
	if err != nil {
		panic(err)
	}

	h.nodes[nodeName] = ringPosition
}

func (h *HashRing) checkPosition(position int) bool {
	for _, v := range h.nodes {
		if v == position {
			return true
		}
	}
	return false
}

func balanceNodes(numPoints int) {
	incDegree := float64(ringDegrees) / float64(numPoints)

	for i := 0; i < numPoints; i++ {
		degree := float64(i) * incDegree
		degree = math.Mod(degree, float64(ringDegrees))
		fmt.Println(degree)
	}
}

func main() {
	fmt.Println("starting")
}
