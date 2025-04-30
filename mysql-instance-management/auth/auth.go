package auth

import _ "embed"

//go:embed my.crt
var CertFile []byte

//go:embed my.key
var KeyFile []byte
