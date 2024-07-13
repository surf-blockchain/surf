package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountToUserKeyPrefix is the prefix to retrieve all AccountToUser
	AccountToUserKeyPrefix = "AccountToUser/value/"
)

// AccountToUserKey returns the store key to retrieve a AccountToUser from the index fields
func AccountToUserKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
