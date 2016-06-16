package ginx_test

import (
	. "github.com/Jackong/ginx"
	"github.com/gin-gonic/gin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ginx", func() {
	Describe("config", func() {
		Describe("GetString", func() {
			Context("When then value is not found", func() {
				It("should be return empty", func() {
					var c = &gin.Context{}
					Expect(GetString(c, "test")).Should(BeEmpty())
				})
			})
			Context("When the value is not a string type", func() {
				It("should be return empty", func() {
					var c = &gin.Context{}
					c.Set("test", 123)
					Expect(GetString(c, "test")).Should(BeEmpty())
				})
			})
			Context("When the value is a string type", func() {
				It("should be return string", func() {
					var c = &gin.Context{}
					c.Set("test", "123")
					Expect(GetString(c, "test")).Should(Equal("123"))
				})
			})
		})

		Describe("GetObject", func() {
			Context("When the value is not a valid json string", func() {
				It("should be return not nil", func() {
					var c = &gin.Context{}
					c.Set("test", "1.0.0")
					obj := map[string]interface{}{}
					err := GetObject(c, "test", &obj)
					Expect(obj).Should(BeEmpty())
					Expect(err).ShouldNot(BeNil())
				})
			})

			Context("When the value is a valid json string", func() {
				It("should be return not nil", func() {
					var c = &gin.Context{}
					c.Set("test", `{"A":1}`)
					type Test struct {
						A int
						B []string
					}
					obj := Test{}
					err := GetObject(c, "test", &obj)
					Expect(obj.A).Should(Equal(1))
					Expect(obj.B).Should(BeEmpty())
					Expect(err).Should(BeNil())
				})
			})
		})
	})
})
