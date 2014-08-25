package classifiers_test

import (
	. "github.com/amitkgupta/golearnbenchmarks/classifiers/sharedbehaviors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/knn"
)

var _ = Describe("KNN Classifier", func() {
	inputs := ClassifiesAccuratelyAndQuicklyBehaviorInputs{}

	const (
		basicDatasetNumNeighbours           = 1
		basicDatasetExpectedAccuracy        = 0.932
		basicDatasetMaxSecondsTimeThreshold = 0.02
		basicDatasetNumRepetitions          = 10

		manyFeaturesDatasetNumNeighbours           = 1
		manyFeaturesDatasetExpectedAccuracy        = 0.948
		manyFeaturesDatasetMaxSecondsTimeThreshold = 60
		manyFeaturesDatasetNumRepetitions          = 2
	)

	Context("When given a basic dataset", func() {
		BeforeEach(func() {
			trainingData, err := base.ParseCSVToInstances("datasets/basic_training.csv", true)
			Ω(err).ShouldNot(HaveOccurred())
			testData, err := base.ParseCSVToInstances("datasets/basic_test.csv", true)
			Ω(err).ShouldNot(HaveOccurred())

			inputs.TrainingData = trainingData
			inputs.TestData = testData
			inputs.ExpectedAccuracy = basicDatasetExpectedAccuracy
			inputs.MaxSecondsTimeThreshold = basicDatasetMaxSecondsTimeThreshold
		})

		Context("When using Manhattan metric", func() {
			BeforeEach(func() {
				inputs.Classifier = knn.NewKnnClassifier("manhattan", basicDatasetNumNeighbours)
			})

			It("classifies without error", ClassifiesWithoutError(&inputs))
			It("classifies with deterministic accuracy", ClassifiesWithDeterministicAccuracy(&inputs))
			Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), basicDatasetNumRepetitions)
		})

		Context("When using Euclidean metric", func() {
			BeforeEach(func() {
				inputs.Classifier = knn.NewKnnClassifier("euclidean", basicDatasetNumNeighbours)
			})

			It("classifies without error", ClassifiesWithoutError(&inputs))
			It("classifies with deterministic accuracy", ClassifiesWithDeterministicAccuracy(&inputs))
			Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), basicDatasetNumRepetitions)
		})
	})

	XContext("When given a dataset with many features", func() {
		BeforeEach(func() {
			trainingData, err := base.ParseCSVToInstances("datasets/many_features_training.csv", true)
			Ω(err).ShouldNot(HaveOccurred())
			testData, err := base.ParseCSVToInstances("datasets/many_features_test.csv", true)
			Ω(err).ShouldNot(HaveOccurred())

			inputs.TrainingData = trainingData
			inputs.TestData = testData
			inputs.ExpectedAccuracy = manyFeaturesDatasetExpectedAccuracy
			inputs.MaxSecondsTimeThreshold = manyFeaturesDatasetMaxSecondsTimeThreshold
		})

		Context("When using Manhattan metric", func() {
			BeforeEach(func() {
				inputs.Classifier = knn.NewKnnClassifier("manhattan", manyFeaturesDatasetNumNeighbours)
			})

			It("classifies without error", ClassifiesWithoutError(&inputs))
			It("classifies with deterministic accuracy", ClassifiesWithDeterministicAccuracy(&inputs))
			Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), manyFeaturesDatasetNumRepetitions)
		})

		Context("When using Euclidean metric", func() {
			BeforeEach(func() {
				inputs.Classifier = knn.NewKnnClassifier("euclidean", manyFeaturesDatasetNumNeighbours)
			})

			It("classifies without error", ClassifiesWithoutError(&inputs))
			It("classifies with deterministic accuracy", ClassifiesWithDeterministicAccuracy(&inputs))
			Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), manyFeaturesDatasetNumRepetitions)
		})
	})
})
