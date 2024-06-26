package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WorldIdtoUserKeyPrefix is the prefix to retrieve all WorldIdtoUser
	WorldIdtoUserKeyPrefix = "WorldIdtoUser/value/"
)

// WorldIdtoUserKey returns the store key to retrieve a WorldIdtoUser from the index fields
func WorldIdtoUserKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
