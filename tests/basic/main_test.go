package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("Actual(%d), input(%d), output(%d)", actual, input, output)
	// }

	assert.Equal(t, AddOne(2), 4, "Error Equal")
	assert.NotEqual(t, AddOne(2), 3, "Error NotEqual")
	assert.Nil(t, nil, 4, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Not executing")
}
