package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserToWorldIdKeyPrefix is the prefix to retrieve all UserToWorldId
	UserToWorldIdKeyPrefix = "UserToWorldId/value/"
)

// UserToWorldIdKey returns the store key to retrieve a UserToWorldId from the index fields
func UserToWorldIdKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
