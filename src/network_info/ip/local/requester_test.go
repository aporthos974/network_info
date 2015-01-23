package local

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestCanGetIPv4(test *testing.T) {
	ip := GetIPv4()

	matched, _ := regexp.MatchString("^\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}$", ip)
	assert.True(test, matched)
}

func TestCanGetIPv6(test *testing.T) {
	ip := GetIPv6()

	assert.NotEmpty(test, ip)
}
