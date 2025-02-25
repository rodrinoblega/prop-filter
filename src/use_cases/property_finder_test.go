package use_cases

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleErrors_PrintsErrors(t *testing.T) {
	errorChan := make(chan error, 2)
	errorChan <- fmt.Errorf("custom error")
	close(errorChan)

	propertyFinder := &PropertyFinder{}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	propertyFinder.handleErrors(errorChan)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()
	assert.Contains(t, output, "Error: custom error")
}
