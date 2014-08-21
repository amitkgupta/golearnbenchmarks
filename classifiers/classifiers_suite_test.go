package classifiers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestClassifiers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Classifiers Suite")
}
