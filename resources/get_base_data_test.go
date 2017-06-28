package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBaseData(t *testing.T) {
	data := GetBaseData()
	assert.NotEmpty(t, data)
}
