// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cid "github.com/ipfs/go-cid"
import context "context"
import io "io"
import ipld "github.com/ipld/go-ipld-prime"
import mock "github.com/stretchr/testify/mock"
import pieceio "github.com/filecoin-project/go-fil-markets/pieceio"

// CarIO is an autogenerated mock type for the CarIO type
type CarIO struct {
	mock.Mock
}

// LoadCar provides a mock function with given fields: bs, r
func (_m *CarIO) LoadCar(bs pieceio.WriteStore, r io.Reader) (cid.Cid, error) {
	ret := _m.Called(bs, r)

	var r0 cid.Cid
	if rf, ok := ret.Get(0).(func(pieceio.WriteStore, io.Reader) cid.Cid); ok {
		r0 = rf(bs, r)
	} else {
		r0 = ret.Get(0).(cid.Cid)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pieceio.WriteStore, io.Reader) error); ok {
		r1 = rf(bs, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrepareCar provides a mock function with given fields: ctx, bs, payloadCid, node
func (_m *CarIO) PrepareCar(ctx context.Context, bs pieceio.ReadStore, payloadCid cid.Cid, node ipld.Node) (pieceio.PreparedCar, error) {
	ret := _m.Called(ctx, bs, payloadCid, node)

	var r0 pieceio.PreparedCar
	if rf, ok := ret.Get(0).(func(context.Context, pieceio.ReadStore, cid.Cid, ipld.Node) pieceio.PreparedCar); ok {
		r0 = rf(ctx, bs, payloadCid, node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pieceio.PreparedCar)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, pieceio.ReadStore, cid.Cid, ipld.Node) error); ok {
		r1 = rf(ctx, bs, payloadCid, node)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteCar provides a mock function with given fields: ctx, bs, payloadCid, node, w
func (_m *CarIO) WriteCar(ctx context.Context, bs pieceio.ReadStore, payloadCid cid.Cid, node ipld.Node, w io.Writer) error {
	ret := _m.Called(ctx, bs, payloadCid, node, w)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pieceio.ReadStore, cid.Cid, ipld.Node, io.Writer) error); ok {
		r0 = rf(ctx, bs, payloadCid, node, w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
