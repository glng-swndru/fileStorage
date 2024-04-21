package p2p

// Peer adalah sebuah interface yang merepresentasikan sebuah remote node.
type Peer interface{}

// Transport adalah untuk handle semua komunikasi
// antara nodes di network. ini bisa jadi sebuah
// form (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
