package archives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplesOfThreeAndFive(t *testing.T) {
	assert.Equal(t, 233168, MultiplesOfThreeAndFive())
}
