package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAccessResource{}

// MsgAccessResource defines actions taken by an agent
type MsgAccessResource struct {
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`   // address of the principal
	Client sdk.AccAddress `json:"client" yaml:"client"` // address of the agent
	Action string         `json:"action" yaml:"action"` // maybe other actions in future, but now just transfer coins
	Amount sdk.Coins      `json:"amount" yaml:"amount"`
	Sig    []byte         `json:"sig" yaml:"sig"`
}

// NewMsgAccessResource creates a new MsgAccessResource instance
func NewMsgAccessResource(owner sdk.AccAddress, client sdk.AccAddress, action string, amount sdk.Coins, sig []byte) MsgAccessResource {
	return MsgAccessResource{
		Owner:  owner,
		Client: client,
		Action: action,
		Amount: amount,
		Sig:    sig,
	}
}

// nolint
func (msg MsgAccessResource) Route() string { return RouterKey }
func (msg MsgAccessResource) Type() string  { return "AccessResource" }
func (msg MsgAccessResource) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Client)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAccessResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgAccessResource) ValidateBasic() error {
	if msg.Client.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "client can't be empty")
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner can't be empty")
	}
	return nil
}
