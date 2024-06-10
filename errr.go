package errr

// we use the charm.sh log library
import "github.com/charmbracelet/log"

// Fatal takes in the potential error, and handles the error with a fatal if it exists
func Fatal(error error) {
	if error != nil {
		log.Fatal(error.Error())
	}
}

// Warn takes in the potential error, and handles the error with a warn if it exists
func Warn(error error) {
	if error != nil {
		log.Warn(error.Error())
	}
}
