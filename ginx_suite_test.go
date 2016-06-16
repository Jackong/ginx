package ginx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGinx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginx Suite")
}
