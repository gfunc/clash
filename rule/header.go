package rules

import (
	"strings"

	C "github.com/Dreamacro/clash/constant"
)

type Header struct {
	value   string
	adapter string
}

func (d *Header) RuleType() C.RuleType {
	return C.Domain
}

func (d *Header) Match(metadata *C.Metadata) bool {
	return strings.ToLower(metadata.ClashHeader) == d.value
}

func (d *Header) Adapter() string {
	return d.adapter
}

func (d *Header) Payload() string {
	return d.value
}

func (d *Header) ShouldResolveIP() bool {
	return false
}

func (d *Header) ShouldFindProcess() bool {
	return false
}

func NewHeader(header string, adapter string) *Header {
	return &Header{
		value:   strings.ToLower(header),
		adapter: adapter,
	}
}
