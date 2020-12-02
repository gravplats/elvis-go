package spinner

import (
	"github.com/theckman/yacspin"
	"time"
)

var cfg = yacspin.Config{
	Colors:            []string{"fgGreen"},
	Frequency:         100 * time.Millisecond,
	CharSet:           yacspin.CharSets[57],
	HideCursor:        true,
	Suffix:            "",
	StopCharacter:     "✓",
	StopColors:        []string{"fgGreen"},
	StopFailCharacter: "✗",
	StopFailColors:    []string{"fgRed"},
}

var spinner, err = yacspin.New(cfg)

func Start(text string) {
	if err != nil {
		return
	}

	// It's recommended that this start with an empty space.
	spinner.Suffix(" " + text)
	spinner.Start()
}

func Succeed() {
	if err != nil {
		return
	}

	spinner.Stop()
}

func Fail() {
	if err != nil {
		return
	}

	spinner.StopFail()
}

func Stop(condition bool) {
	if err != nil {
		return
	}

	if condition {
		Succeed()
	} else {
		Fail()
	}
}
