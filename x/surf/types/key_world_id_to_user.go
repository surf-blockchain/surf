package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WorldIdToUserKeyPrefix is the prefix to retrieve all WorldIdToUser
	WorldIdToUserKeyPrefix = "WorldIdToUser/value/"
)

// WorldIdToUserKey returns the store key to retrieve a WorldIdToUser from the index fields
func WorldIdToUserKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
