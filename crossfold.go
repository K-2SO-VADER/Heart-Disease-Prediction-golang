package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/evaluation"
)

func crossfold(path string) {

	var tree base.Classifier

	// Load in the iris dataset
	iris, err := base.ParseCSVToInstances(path, true)
	if err != nil {
		panic(err)
	}

	for i := 1; i < 60; i += 2 {
		// Demonstrate the effect of adding more trees to the forest
		// and also how much better it is without discretisation.
		rand.Seed(44111342)

		tree = ensemble.NewRandomForest(i, 4)
		cfs, err := evaluation.GenerateCrossFoldValidationConfusionMatrices(iris, tree, 6)
		if err != nil {
			panic(err)
		}

		mean, variance := evaluation.GetCrossValidatedMetric(cfs, evaluation.GetAccuracy)
		stdev := math.Sqrt(variance)

		fmt.Printf("%d\t%.2f\t(+/- %.2f)\n", i, mean, stdev*2)
	}
}
