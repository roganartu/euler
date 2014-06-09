package helpers

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestStringify(t *testing.T) {
	var tmp string
	_, tmp = Stringify(76, nil)
	assert.Equal(t, "seventy six", tmp)
	_, tmp = Stringify(50, nil)
	assert.Equal(t, "fifty", tmp)
	_, tmp = Stringify(14, nil)
	assert.Equal(t, "fourteen", tmp)
	_, tmp = Stringify(123, nil)
	assert.Equal(t, "one hundred and twenty three", tmp)
	_, tmp = Stringify(100, nil)
	assert.Equal(t, "one hundred", tmp)
	_, tmp = Stringify(1305, nil)
	assert.Equal(t, "one thousand three hundred and five", tmp)
	_, tmp = Stringify(987, nil)
	assert.Equal(t, "nine hundred and eighty seven", tmp)
}

func TestIsPandigital(t *testing.T) {
	assert.Equal(t, true, IsPandigital("15234"))
	assert.Equal(t, false, IsPandigital("15534"))
}
