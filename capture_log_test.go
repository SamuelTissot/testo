package testo_test

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"testo"
)

func TestCaptureStdout(t *testing.T) {
	ch := make(chan bool, 1)
	r := []int{1, 2, 3, 4, 5}
	format := "<< %d >>"
	output := testo.CaptureLog(ch, func() {
		for _, v := range r {
			log.Printf(format, v)
		}
		ch <- true
	})

	for _, v := range r {
		w := fmt.Sprintf(format, v)
		if !strings.Contains(output, w) {
			t.Errorf("CaptureLog() error could not find: %s, in: %s", w, output)
		}
	}
}
