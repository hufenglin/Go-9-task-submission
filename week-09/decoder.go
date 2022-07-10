package week09

import (
	"encoding/binary"
	"errors"
)

const (
	// MaxBodySize max body size
	MaxBodySize = uint32(1 << 12) // 4096
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + uint32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_bodyOffset   = _seqOffset + _seqSize
)

type Idecoder interface {
	PacketLen() uint32
	HeaderLen() uint16
	Version() uint16
	Operation() uint32
	Sequence() uint32
	Body() []byte
}

type Decoder struct {
	packetLen uint32
	headerLen uint16
	version   uint16
	operation uint32
	sequence  uint32
	body      []byte
}

func Decode(buf []byte) (Idecoder, error) {
	decoder := &Decoder{}

	decoder.packetLen = binary.BigEndian.Uint32(buf[_packOffset : _packOffset+_packSize])
	decoder.headerLen = binary.BigEndian.Uint16(buf[_headerOffset : _headerOffset+_headerSize])
	decoder.version = binary.BigEndian.Uint16(buf[_verOffset : _verOffset+_verSize])
	decoder.operation = binary.BigEndian.Uint32(buf[_opOffset : _opOffset+_opSize])
	decoder.sequence = binary.BigEndian.Uint32(buf[_seqOffset : _seqOffset+_seqSize])

	if decoder.packetLen > _maxPackSize {
		return nil, errors.New("error package length")
	}

	if _bodyLen := int(decoder.packetLen - uint32(decoder.headerLen)); _bodyLen > 0 {
		decoder.body = buf[_bodyOffset : _bodyOffset+_bodyLen]
	}

	return decoder, nil
}

func (d *Decoder) PacketLen() uint32 {
	return d.packetLen
}

func (d *Decoder) HeaderLen() uint16 {
	return d.headerLen
}

func (d *Decoder) Version() uint16 {
	return d.version
}

func (d *Decoder) Operation() uint32 {
	return d.operation
}

func (d *Decoder) Sequence() uint32 {
	return d.sequence
}

func (d *Decoder) Body() []byte {
	if d.body == nil {
		return []byte{}
	} else {
		return d.body
	}
}
