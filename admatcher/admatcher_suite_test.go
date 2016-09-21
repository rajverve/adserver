package admatcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAdmatcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Admatcher Suite")
}
