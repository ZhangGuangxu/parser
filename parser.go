package parser

type Packer interface {
	Pack(msg interface{}) (data []byte)
}

type Parser interface {
	Packer
	Unpack(data []byte) (msg interface{})
}
