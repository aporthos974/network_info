package ip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanGetIP(test *testing.T) {
	ip := GetIPAddress()

	assert.Equal(test, "83.200.184.130", ip)
}
