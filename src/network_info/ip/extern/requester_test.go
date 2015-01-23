package extern

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestCanGetIPv4(test *testing.T) {
	ip := GetExternalIPAddress()

	matched, _ := regexp.MatchString("^\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}$", ip)
	assert.True(test, matched)
}
