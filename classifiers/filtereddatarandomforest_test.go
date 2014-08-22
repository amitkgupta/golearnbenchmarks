package classifiers_test

import (
	. "github.com/amitkgupta/golearnbenchmarks/classifiers/sharedbehaviors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/ensemble"
	"github.com/sjwhitworth/golearn/filters"
)

const (
	basicDatasetFilterSignificance      = 0.6
	basicDatasetTrainTestSplit          = 0.6
	basicDatasetForestSize              = 50
	basicDatasetMinAccuracyThreshold    = 0.7
	basicDatasetMaxSecondsTimeThreshold = 0.45
	basicDatasetNumRepitions            = 10
)

var _ = Describe("Random Forest on Filtered Data", func() {
	inputs := ClassifiesAccuratelyAndQuicklyBehaviorInputs{}

	Context("When given a basic dataset", func() {
		BeforeEach(func() {
			instances, err := base.ParseCSVToInstances("datasets/iris_headers.csv", true)
			Ω(err).ShouldNot(HaveOccurred())

			filter := filters.NewChiMergeFilter(instances, basicDatasetFilterSignificance)
			for _, a := range base.NonClassFloatAttributes(instances) {
				filter.AddAttribute(a)
			}
			filter.Train()
			filteredInstances := base.NewLazilyFilteredInstances(instances, filter)
			numNonClassAttributes := len(base.NonClassAttributes(filteredInstances))

			inputs.TrainingData, inputs.TestData = base.InstancesTrainTestSplit(filteredInstances, basicDatasetTrainTestSplit)
			inputs.Classifier = ensemble.NewRandomForest(basicDatasetForestSize, numNonClassAttributes)
			inputs.MinAccuracyThreshold = basicDatasetMinAccuracyThreshold
			inputs.MaxSecondsTimeThreshold = basicDatasetMaxSecondsTimeThreshold
		})

		It("classifies without error", ClassifiesWithoutError(&inputs))
		Measure("consistently classifies sufficiently accurately", ClassifiesSufficientlyAccurately(&inputs), basicDatasetNumRepitions)
		Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), basicDatasetNumRepitions)
	})
})
