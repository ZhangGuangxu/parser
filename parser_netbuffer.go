package parser

import (
	"github.com/ZhangGuangxu/netbuffer"
)

type NetbufferParser interface {
	Packer
	Unpack(buf *netbuffer.Buffer) (msg interface{})
}
