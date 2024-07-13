package types

const (
	// ModuleName defines the module name
	ModuleName = "surf"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_surf"

	// Version defines the current version the IBC module supports
	Version = "surf-1"

	// PortID is the default port id that module binds to
	PortID = "surf"
)

var (
	ParamsKey = []byte("p_surf")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("surf-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
