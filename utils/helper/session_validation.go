package helper

import (
	"errors"
	"let-you-cook/domain/model"
)

func ValidateSession(session model.PomodoroSession) error {
	if session.Name == "" {
		return errors.New("name cannot be empty")
	}
	if session.FocusDuration <= 0 {
		return errors.New("focus duration should be greater than 0")
	}
	if session.BreakDuration < 0 {
		return errors.New("break duration cannot be negative")
	}
	return nil
}
