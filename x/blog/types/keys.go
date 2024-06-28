package types

import "cosmossdk.io/collections"

const (
	// ModuleName defines the module name
	ModuleName = "blog"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blog"

	// Version defines the current version the IBC module supports
	Version = "blog-1"

	// PortID is the default port id that module binds to
	PortID = "blog"
)

var (
	ParamsKey = collections.NewPrefix("p_blog")
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = collections.NewPrefix("blog-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	PostKey      = collections.NewPrefix("post/value/")
	PostCountKey = collections.NewPrefix("post/count/")
)

var (
	SentPostKey      = collections.NewPrefix("sentpost/value/")
	SentPostCountKey = collections.NewPrefix("sentpost/count/")
)

var (
	TimeoutPostKey      = collections.NewPrefix("timeoutpost/value/")
	TimeoutPostCountKey = collections.NewPrefix("timeoutpost/count/")
)
