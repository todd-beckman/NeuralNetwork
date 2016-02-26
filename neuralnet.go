package main

import (
    "fmt"
    "math"
    "math/rand"
)

//  So GO doesn't yell that fmt wasn't used
func unused() {
    fmt.Println("unused")
}

//  Sigmoid function as activation function
func activation(x float64) float64 {
    return 1 / (1 + math.Exp(-x))
}

func deractivation(x float64) float64 {
    return x * (1 - x)
}



type NNetwork struct {
    //  List of layers
        //  list of nodes (this)
            //  list of edges (where weight comes from)
    net         [][][]float64  // the network
    //  The values each node output during the last run
    values      [][]float64
    //  How fast to learn
    LearnRate   float64
}

//  Helper method. Calculates the values of all nodes
//  in an entire layer.
func (net *NNetwork) feedforward(input []float64) {
    //  Set the starting layer
    for i := range input {
        net.values[0][i] = input[i]
    }
    for layer := 1; layer < len(net.net); layer++ {
        //  Populate results node by node
        for i := 0; i < len(net.values[layer]); i++ {
            //  Sum edges of ith node by jth edge
            sum := 0.0
            for j := 0; j < len(net.values[layer - 1]); j++ {
                sum += net.values[layer - 1][j] * net.net[layer][i][j]
            }
            //  Store the results for learning
            net.values[layer][i] = activation(sum)
        }
    }
}

//  Helper method. Updates edge weights based on their
//  assumed contributions to the error
func (net *NNetwork) backprop(input, answer []float64) {
    //  the real challenge
}

//  Learn will cast blame and henceforth reduce error in future runs
func (net *NNetwork) Learn(input, ans []float64) {
    net.feedforward(input)
    net.backprop(input, ans)
}

//  Guess attempts to calculate what the function's output is
func (net *NNetwork) Guess(input []float64) []float64 {
    net.feedforward(input)
    lastlayer := len(net.values) - 1
    lastlayersize := len(net.values[lastlayer]) - 1
    results := make([]float64, lastlayersize)
    copy(results, net.values[lastlayer])
    return results;
}

func NewNetwork(learn float64, nodes ...int) *NNetwork {
    //  allocate network
    net := make([][][]float64, len(nodes))
    values := make([][]float64, len(nodes))
    //  Input layer does not have input edges
    net[0] = make([][]float64, nodes[0])
    values[0] = make([]float64, nodes[0])
    //  for each layer
    for i := 1; i < len(net); i++ {
        //  make the layer of nodes
        net[i] = make([][]float64, nodes[i])
        values[i] = make([]float64, nodes[i])
        //  for each node
        for j := range net[i] {
            //  make the nodes' input edges (num nodes in previous layer)
            net[i][j] = make([]float64, nodes[i - 1])
            //  for each edge in the node
            for k := range net[i][j] {
                //  populate weight randomly
                net[i][j][k] = rand.Float64() - 0.5
            }
        }
    }
    return &NNetwork{net: net, LearnRate: learn}
}
