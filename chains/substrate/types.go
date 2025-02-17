// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package substrate

import (
	"math/big"

	"github.com/Phala-Network/chainbridge-utils/msg"
	"github.com/Phala-Network/go-substrate-rpc-client/v3/scale"
	"github.com/Phala-Network/go-substrate-rpc-client/v3/types"
)

type voteState struct {
	VotesFor     []types.AccountID
	VotesAgainst []types.AccountID
	Status       voteStatus
}

type voteStatus struct {
	IsActive   bool
	IsApproved bool
	IsRejected bool
}

var BridgeTransfer string = "ChainBridge.handle_fungible_transfer"

func (m *voteStatus) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsActive = true
	} else if b == 1 {
		m.IsApproved = true
	} else if b == 2 {
		m.IsRejected = true
	}

	return nil
}

type FungibleTransferItem struct {
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Amount       types.U256
	Recipient    types.Bytes
}

type NonFungibleTransferItem struct {
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	TokenId      types.Bytes
	Recipient    types.Bytes
	Metadata     types.Bytes
}

type GenericTransferItem struct {
	Destination  types.U8
	DepositNonce types.U64
	ResourceId   types.Bytes32
	Metadata     types.Bytes
}

type BridgeEvents []BridgeEvent

type BridgeEvent struct {
	IsFungible          bool
	FungibleTransfer    FungibleTransferItem
	IsNonFungible       bool
	NonFungibleTransfer NonFungibleTransferItem
	IsGeneric           bool
	GenericTransfer     GenericTransferItem
}

func (m *BridgeEvent) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()

	if err != nil {
		return err
	}

	if b == 0 {
		m.IsFungible = true
		dfungible := FungibleTransferItem{}
		err = decoder.Decode(&dfungible)
		if err != nil {
			return err
		}
		m.FungibleTransfer = dfungible
	} else if b == 1 {
		m.IsNonFungible = true
		dnonfungible := NonFungibleTransferItem{}
		err = decoder.Decode(&dnonfungible)
		if err != nil {
			return err
		}
		m.NonFungibleTransfer = dnonfungible
	} else if b == 2 {
		m.IsGeneric = true
		dgeneric := GenericTransferItem{}
		err = decoder.Decode(&dgeneric)
		if err != nil {
			return err
		}
		m.GenericTransfer = dgeneric
	}

	return nil
}

// proposal represents an on-chain proposal
type proposal struct {
	depositNonce types.U64
	call         types.Call
	sourceId     types.U8
	resourceId   types.Bytes32
	method       string
}

// encode takes only nonce and call and encodes them for storage queries
func (p *proposal) encode() ([]byte, error) {
	return types.EncodeToBytes(struct {
		types.U64
		types.Call
	}{p.depositNonce, p.call})
}

func (w *writer) createFungibleProposal(m msg.Message) (*proposal, error) {
	bigAmt := big.NewInt(0).SetBytes(m.Payload[0].([]byte))
	amount := types.NewU128(*bigAmt)
	recipient := m.Payload[1].([]byte)
	depositNonce := types.U64(m.DepositNonce)

	meta := w.conn.getMetadata()
	call, err := types.NewCall(
		&meta,
		BridgeTransfer,
		recipient,
		amount,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := types.EncodeToBytes(m.ResourceId)
		if err != nil {
			return nil, err
		}
		call.Args = append(call.Args, eRID...)
	}

	return &proposal{
		depositNonce: depositNonce,
		call:         call,
		sourceId:     types.U8(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       BridgeTransfer,
	}, nil
}

func (w *writer) createNonFungibleProposal(m msg.Message) (*proposal, error) {
	tokenId := types.NewU256(*big.NewInt(0).SetBytes(m.Payload[0].([]byte)))
	recipient := types.NewAccountID(m.Payload[1].([]byte))
	metadata := types.Bytes(m.Payload[2].([]byte))
	depositNonce := types.U64(m.DepositNonce)

	meta := w.conn.getMetadata()
	call, err := types.NewCall(
		&meta,
		BridgeTransfer,
		recipient,
		tokenId,
		metadata,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := types.EncodeToBytes(m.ResourceId)
		if err != nil {
			return nil, err
		}
		call.Args = append(call.Args, eRID...)
	}

	return &proposal{
		depositNonce: depositNonce,
		call:         call,
		sourceId:     types.U8(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       BridgeTransfer,
	}, nil
}

func (w *writer) createGenericProposal(m msg.Message) (*proposal, error) {
	metadata := types.Bytes(m.Payload[0].([]byte))

	meta := w.conn.getMetadata()
	call, err := types.NewCall(
		&meta,
		BridgeTransfer,
		metadata,
	)
	if err != nil {
		return nil, err
	}
	if w.extendCall {
		eRID, err := types.EncodeToBytes(m.ResourceId)
		if err != nil {
			return nil, err
		}

		call.Args = append(call.Args, eRID...)
	}
	return &proposal{
		depositNonce: types.U64(m.DepositNonce),
		call:         call,
		sourceId:     types.U8(m.Source),
		resourceId:   types.NewBytes32(m.ResourceId),
		method:       BridgeTransfer,
	}, nil
}
