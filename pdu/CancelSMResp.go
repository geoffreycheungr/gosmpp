package pdu

import (
	"github.com/geoffreycheungr/gosmpp/data"
)

// CancelSMResp PDU.
type CancelSMResp struct {
	base
}

// NewCancelSMResp returns CancelSMResp.
func NewCancelSMResp() PDU {
	c := &CancelSMResp{
		base: newBase(),
	}
	c.CommandID = data.CANCEL_SM_RESP
	return c
}

// NewCancelSMRespFromReq returns CancelSMResp.
func NewCancelSMRespFromReq(req *CancelSM) PDU {
	c := NewCancelSMResp().(*CancelSMResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *CancelSMResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *CancelSMResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *CancelSMResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, nil)
}

// Unmarshal implements PDU interface.
func (c *CancelSMResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, nil)
}
