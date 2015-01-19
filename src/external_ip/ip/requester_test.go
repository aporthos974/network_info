package ip

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestCanGetIP(test *testing.T) {
	ip := GetActualIPAddress()

	matched, _ := regexp.MatchString("^\\d{1,3}.\\d{1,3}.\\d{1,3}.\\d{1,3}$", ip)
	assert.True(test, matched)
}
