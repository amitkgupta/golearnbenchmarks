package classifiers_test

import (
	. "github.com/amitkgupta/golearnbenchmarks/classifiers/sharedbehaviors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
)

var _ = Describe("Random Forest", func() {
	inputs := ClassifiesAccuratelyAndQuicklyBehaviorInputs{}

	const (
		basicDatasetForestSize              = 50
		basicDatasetMinAccuracyThreshold    = 0.4
		basicDatasetMaxSecondsTimeThreshold = 0.45
		basicDatasetNumRepetitions          = 10
	)

	Context("When given a basic dataset", func() {
		BeforeEach(func() {
			trainingData, err := base.ParseCSVToInstances("datasets/basic_training.csv", true)
			Ω(err).ShouldNot(HaveOccurred())
			testData, err := base.ParseCSVToInstances("datasets/basic_test.csv", true)
			Ω(err).ShouldNot(HaveOccurred())

			inputs.TrainingData = trainingData
			inputs.TestData = testData

			numNonClassAttributes := len(base.NonClassAttributes(trainingData))
			inputs.Classifier = ensemble.NewRandomForest(basicDatasetForestSize, numNonClassAttributes)

			inputs.MinAccuracyThreshold = basicDatasetMinAccuracyThreshold
			inputs.MaxSecondsTimeThreshold = basicDatasetMaxSecondsTimeThreshold
		})

		It("classifies without error", ClassifiesWithoutError(&inputs))
		Measure("reliably classifies sufficiently accurately", ClassifiesSufficientlyAccurately(&inputs), basicDatasetNumRepetitions)
		Measure("reliably classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), basicDatasetNumRepetitions)
	})
})
