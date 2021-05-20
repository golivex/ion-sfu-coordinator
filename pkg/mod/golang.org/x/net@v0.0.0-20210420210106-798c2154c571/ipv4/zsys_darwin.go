// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs_darwin.go

package ipv4

const (
	sizeofSockaddrStorage = 0x80
	sizeofSockaddrInet    = 0x10
	sizeofInetPktinfo     = 0xc

	sizeofIPMreq         = 0x8
	sizeofIPMreqn        = 0xc
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x84
	sizeofGroupSourceReq = 0x104
)

type sockaddrStorage struct {
	Len         uint8
	Family      uint8
	X__ss_pad1  [6]int8
	X__ss_align int64
	X__ss_pad2  [112]int8
}

type sockaddrInet struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type inetPktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type ipMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

type ipMreqSource struct {
	Multiaddr  [4]byte /* in_addr */
	Sourceaddr [4]byte /* in_addr */
	Interface  [4]byte /* in_addr */
}

type groupReq struct {
	Interface uint32
	Pad_cgo_0 [128]byte
}

type groupSourceReq struct {
	Interface uint32
	Pad_cgo_0 [128]byte
	Pad_cgo_1 [128]byte
}
