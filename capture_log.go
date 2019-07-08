package testo

import (
	"bytes"
	"log"
	"os"
)

func CaptureLog(stop chan bool, f func()) string {
	var buf bytes.Buffer
	output := ""
	log.SetOutput(&buf)
	f()
	for {
		output += buf.String()

		select {
		case <-stop:
			log.SetOutput(os.Stderr)
			return output
		}

	}
}
