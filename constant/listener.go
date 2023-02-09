package constant

const ClashHeader = "Clash"

type Listener interface {
	RawAddress() string
	Address() string
	Close() error
}
