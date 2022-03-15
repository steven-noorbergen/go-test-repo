package countdown

import (
	"reflect"
	"testing"
)

type ObservableCountdownOperations struct {
	Calls []string
}

func (s *ObservableCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *ObservableCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("sleep before every print", func(t *testing.T) {
		observableSleepPrinter := &ObservableCountdownOperations{}
		Countdown(observableSleepPrinter, observableSleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, observableSleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, observableSleepPrinter.Calls)
		}
	})
}
