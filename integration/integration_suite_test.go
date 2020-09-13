package integration_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("BeforeSuite..")
})

var _ = AfterSuite(func() {
	fmt.Println("AfterSuite..")
})
