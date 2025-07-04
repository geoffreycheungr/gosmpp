package pdu

import (
	"github.com/geoffreycheungr/gosmpp/data"
)

// QuerySMResp PDU.
type QuerySMResp struct {
	base
	MessageID    string
	FinalDate    string
	MessageState byte
	ErrorCode    byte
}

// NewQuerySMResp returns new QuerySM PDU.
func NewQuerySMResp() PDU {
	c := &QuerySMResp{
		base:         newBase(),
		FinalDate:    data.DFLT_DATE,
		MessageState: data.DFLT_MSG_STATE,
		ErrorCode:    data.DFLT_ERR,
	}
	c.CommandID = data.QUERY_SM_RESP
	return c
}

// NewQuerySMRespFromReq returns new QuerySM PDU.
func NewQuerySMRespFromReq(req *QuerySM) PDU {
	c := NewQuerySMResp().(*QuerySMResp)
	if req != nil {
		c.SequenceNumber = req.SequenceNumber
	}
	return c
}

// CanResponse implements PDU interface.
func (c *QuerySMResp) CanResponse() bool {
	return false
}

// GetResponse implements PDU interface.
func (c *QuerySMResp) GetResponse() PDU {
	return nil
}

// Marshal implements PDU interface.
func (c *QuerySMResp) Marshal(b *ByteBuffer) {
	c.base.marshal(b, func(b *ByteBuffer) {
		b.Grow(len(c.MessageID) + len(c.FinalDate) + 4)

		_ = b.WriteCString(c.MessageID)
		_ = b.WriteCString(c.FinalDate)
		_ = b.WriteByte(c.MessageState)
		_ = b.WriteByte(c.ErrorCode)
	})
}

// Unmarshal implements PDU interface.
func (c *QuerySMResp) Unmarshal(b *ByteBuffer) error {
	return c.base.unmarshal(b, func(b *ByteBuffer) (err error) {
		if c.MessageID, err = b.ReadCString(); err == nil {
			if c.FinalDate, err = b.ReadCString(); err == nil {
				if c.MessageState, err = b.ReadByte(); err == nil {
					c.ErrorCode, err = b.ReadByte()
				}
			}
		}
		return
	})
}
