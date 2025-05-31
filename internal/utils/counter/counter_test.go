package counter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrement(t *testing.T) {
	// Arrange
	testTable := []struct {
		numbersStr     string
		countOperation int
		expected       string
	}{
		{
			numbersStr:     "1",
			countOperation: 4,
			expected:       "5",
		},
		{
			numbersStr:     "0",
			countOperation: 1,
			expected:       "1",
		},
		{
			numbersStr:     "0",
			countOperation: 0,
			expected:       "0",
		},
		{
			numbersStr:     "14891841948194378137389535835385835783195",
			countOperation: 5,
			expected:       "14891841948194378137389535835385835783200",
		},
		{
			numbersStr:     "-1000",
			countOperation: 4,
			expected:       "-996",
		},
		{
			numbersStr:     "-0",
			countOperation: 1,
			expected:       "1",
		},
		{
			numbersStr:     "-10",
			countOperation: 20,
			expected:       "10",
		},
		{
			numbersStr:     "-14891841948194378137389535835385835783195",
			countOperation: 10,
			expected:       "-14891841948194378137389535835385835783185",
		},
	}

	//Act
	for _, testCase := range testTable {
		result := circleIncrement(testCase.numbersStr, testCase.countOperation)
		t.Logf("Calling Increment(%v), been called %d times, result %s\n",
			testCase.numbersStr, testCase.countOperation, result)

		// Assert
		assert.Equal(t, testCase.expected, result)
	}
}

func circleIncrement(str string, counter int) string {
	for i := 0; i < counter; i++ {
		str = Increment(str)
	}

	return str
}
