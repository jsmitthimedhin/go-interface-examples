package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockLogger struct{}

func (m mockLogger) log(message string) {
	fmt.Println("Fake logger implementation!")
}
func TestDetermineLogger_UnknownLogger(t *testing.T) {
	m := mockLogger{}
	// My expected result of the test is that it would return the string below:
	expected := "It's an unknown logger!"
	// I call the function and pass in the mockLogger object:
	result := determineLogger(m)
	// and assert that what’s expected will equal the result:
	assert.Equal(t, expected, result)
}
