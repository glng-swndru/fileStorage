package p2p

type HandshakeFunc func(Peer) error

func NOPHanshakeFunc(Peer) error { return nil }
