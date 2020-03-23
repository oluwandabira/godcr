package config_test

import (
	"fmt"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/raedahgroup/godcr-gio/config"
)

func temp(s string) string {
	t, err := ioutil.TempFile("", "godcr_test_*.conf")
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	defer t.Close()
	_, err = t.WriteString(s)
	if err != nil {
		fmt.Printf(err.Error())
		return ""
	}
	return t.Name()
}

var _ = Describe("Config", func() {
	It("parses an empty file with no effects", func() {
		t := temp("")
		Expect(t).ToNot(Equal(""))
		var a, b Config
		err := a.ParseFile(t)
		Expect(err).To(BeNil())
		Expect(a).To(Equal(b))
	})

	It("parses a file with Quiet set", func() {
		t := temp("quiet=true")
		Expect(t).ToNot(Equal(""))
		var a, b Config
		b.Quiet = true
		Expect(a.ParseFile(t)).To(Equal(b))
	})

})
