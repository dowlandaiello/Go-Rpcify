package common

import (
	"encoding/base64"

	"golang.org/x/crypto/sha3"
)

/* BEGIN EXPORTED METHODS */

/*
	BEGIN CRYPTO METHODS:
*/

// SHA3String - generate and return SHA3 string-value hash of specified byte array
func SHA3String(b []byte) string {
	hash := sha3.Sum256(b) // Calculate hash

	return base64.StdEncoding.EncodeToString(hash[:]) // Return base64 string
}

// SHA3Bytes - generate and return SHA3 hash of specified byte array
func SHA3Bytes(b []byte) []byte {
	hash := sha3.Sum256(b) // Calculate hash

	return hash[:] // Return hash
}

/*
	END CRYPTO METHODS
*/

/* END EXPORTED METHODS */
