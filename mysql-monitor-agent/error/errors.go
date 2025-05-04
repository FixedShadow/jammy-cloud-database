package error

import "errors"

// loggedErrors define some common errors.
type loggedErrors struct {
	NoConfigFileFound         error
	ConfigFileValidationError error
	NoMatchedFileFound        error
	OpenConfigFileError       error
	CastTypeError             error
	DirPathError              error
}

// Errors is the set of errors that can occur
var Errors = loggedErrors{
	NoConfigFileFound:         errors.New("no config file found, please check if missing"),
	ConfigFileValidationError: errors.New("config file validation failed, please re-check the content on it"),
	NoMatchedFileFound:        errors.New("no matched file found with a pattern"),
	OpenConfigFileError:       errors.New("open config file error"),
	CastTypeError:             errors.New("cast type error"),
	DirPathError:              errors.New("dir is not a path"),
}
