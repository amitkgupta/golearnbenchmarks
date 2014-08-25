package sharedbehaviors

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
)

type ClassifiesAccuratelyAndQuicklyBehaviorInputs struct {
	Classifier              base.Classifier
	TrainingData            base.FixedDataGrid
	TestData                base.FixedDataGrid
	ExpectedAccuracy        float64
	MinAccuracyThreshold    float64
	MaxSecondsTimeThreshold float64
}

var ClassifiesWithoutError = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func() {
	return func() {
		err := inputs.Classifier.Fit(inputs.TrainingData)
		Ω(err).ShouldNot(HaveOccurred())

		_, err = inputs.Classifier.Predict(inputs.TestData)
		Ω(err).ShouldNot(HaveOccurred())
	}
}

var ClassifiesWithDeterministicAccuracy = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func() {
	return func() {
		inputs.Classifier.Fit(inputs.TrainingData)
		predictions, _ := inputs.Classifier.Predict(inputs.TestData)

		confusionMatrix, err := evaluation.GetConfusionMatrix(inputs.TestData, predictions)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(evaluation.GetAccuracy(confusionMatrix)).Should(BeNumerically("~", inputs.ExpectedAccuracy, 0.001))
	}
}

var ClassifiesSufficientlyAccurately = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func(Benchmarker) {
	return func(b Benchmarker) {
		inputs.Classifier.Fit(inputs.TrainingData)
		predictions, _ := inputs.Classifier.Predict(inputs.TestData)

		confusionMatrix, err := evaluation.GetConfusionMatrix(inputs.TestData, predictions)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(evaluation.GetAccuracy(confusionMatrix)).Should(BeNumerically(">", inputs.MinAccuracyThreshold))
	}
}

var ClassifiesSufficientlyQuickly = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func(Benchmarker) {
	return func(b Benchmarker) {
		fitAndPredictTime := b.Time("fit and predict", func() {
			inputs.Classifier.Fit(inputs.TrainingData)
			inputs.Classifier.Predict(inputs.TestData)
		})

		Ω(fitAndPredictTime.Seconds()).Should(BeNumerically("<", inputs.MaxSecondsTimeThreshold))
	}
}
