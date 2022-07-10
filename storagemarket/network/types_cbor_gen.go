// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package network

import (
	"fmt"
	"io"

	storagemarket "github.com/filecoin-project/go-fil-markets/storagemarket"
	market "github.com/filecoin-project/specs-actors/actors/builtin/market"
	crypto "github.com/filecoin-project/specs-actors/actors/crypto"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf

var lengthBufAskRequest = []byte{129}

func (t *AskRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufAskRequest); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskRequest) UnmarshalCBOR(r io.Reader) error {
	*t = AskRequest{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Miner: %w", err)
		}

	}
	return nil
}

var lengthBufAskResponse = []byte{129}

func (t *AskResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufAskResponse); err != nil {
		return err
	}

	// t.Ask (storagemarket.SignedStorageAsk) (struct)
	if err := t.Ask.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskResponse) UnmarshalCBOR(r io.Reader) error {
	*t = AskResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Ask (storagemarket.SignedStorageAsk) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.Ask = new(storagemarket.SignedStorageAsk)
			if err := t.Ask.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Ask pointer: %w", err)
			}
		}

	}
	return nil
}

var lengthBufProposal = []byte{131}

func (t *Proposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufProposal); err != nil {
		return err
	}

	// t.DealProposal (market.ClientDealProposal) (struct)
	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Piece (storagemarket.DataRef) (struct)
	if err := t.Piece.MarshalCBOR(w); err != nil {
		return err
	}

	// t.FastRetrieval (bool) (bool)
	if err := cbg.WriteBool(w, t.FastRetrieval); err != nil {
		return err
	}
	return nil
}

func (t *Proposal) UnmarshalCBOR(r io.Reader) error {
	*t = Proposal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal (market.ClientDealProposal) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.DealProposal = new(market.ClientDealProposal)
			if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
			}
		}

	}
	// t.Piece (storagemarket.DataRef) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.Piece = new(storagemarket.DataRef)
			if err := t.Piece.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Piece pointer: %w", err)
			}
		}

	}
	// t.FastRetrieval (bool) (bool)

	maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.FastRetrieval = false
	case 21:
		t.FastRetrieval = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	return nil
}

var lengthBufResponse = []byte{132}

func (t *Response) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufResponse); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.State (uint64) (uint64)

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.Proposal (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.PublishMessage (cid.Cid) (struct)

	if t.PublishMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PublishMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishMessage: %w", err)
		}
	}

	return nil
}

func (t *Response) UnmarshalCBOR(r io.Reader) error {
	*t = Response{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.State (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.State = uint64(extra)

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadStringBuf(br, scratch)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.Proposal (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
		}

		t.Proposal = c

	}
	// t.PublishMessage (cid.Cid) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PublishMessage: %w", err)
			}

			t.PublishMessage = &c
		}

	}
	return nil
}

var lengthBufSignedResponse = []byte{130}

func (t *SignedResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufSignedResponse); err != nil {
		return err
	}

	// t.Response (network.Response) (struct)
	if err := t.Response.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (crypto.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedResponse) UnmarshalCBOR(r io.Reader) error {
	*t = SignedResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Response (network.Response) (struct)

	{

		if err := t.Response.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Response: %w", err)
		}

	}
	// t.Signature (crypto.Signature) (struct)

	{

		b, err := br.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := br.UnreadByte(); err != nil {
				return err
			}
			t.Signature = new(crypto.Signature)
			if err := t.Signature.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Signature pointer: %w", err)
			}
		}

	}
	return nil
}

var lengthBufDealStatusRequest = []byte{130}

func (t *DealStatusRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealStatusRequest); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Proposal (cid.Cid) (struct)

	if err := cbg.WriteCidBuf(scratch, w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.Signature (crypto.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealStatusRequest) UnmarshalCBOR(r io.Reader) error {
	*t = DealStatusRequest{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Proposal (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
		}

		t.Proposal = c

	}
	// t.Signature (crypto.Signature) (struct)

	{

		if err := t.Signature.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Signature: %w", err)
		}

	}
	return nil
}

var lengthBufDealStatusResponse = []byte{130}

func (t *DealStatusResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write(lengthBufDealStatusResponse); err != nil {
		return err
	}

	// t.DealState (storagemarket.ProviderDealState) (struct)
	if err := t.DealState.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (crypto.Signature) (struct)
	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealStatusResponse) UnmarshalCBOR(r io.Reader) error {
	*t = DealStatusResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealState (storagemarket.ProviderDealState) (struct)

	{

		if err := t.DealState.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DealState: %w", err)
		}

	}
	// t.Signature (crypto.Signature) (struct)

	{

		if err := t.Signature.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Signature: %w", err)
		}

	}
	return nil
}
