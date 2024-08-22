package log

import (
	"log"
	"runtime"
)

func getStackFrames() int {
	slicePointerFrames := make([]uintptr, 30)
	return runtime.Callers(0, slicePointerFrames)
}

func LogErrorWithLine(err error) {
	if err == nil {
		return
	}

	// get stack frames: lay so luong frame trong stack trace
	frames := getStackFrames()
	if frames < 1 {
		return
	}

	pc, file, line, ok := runtime.Caller(frames - 1)

	if !ok {
		log.Print("Could not get current line information")
		return
	}

	log.Printf("Error in %s:%d: %v", file, line, err)
	log.Printf("PC: %v", pc)
}
