package bauth

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/zenfa/bauth/x/bauth/keeper"
	"github.com/zenfa/bauth/x/bauth/types"
)

func handleMsgAccessResource(ctx sdk.Context, k keeper.Keeper, msg types.MsgAccessResource) (*sdk.Result, error) {

	acc := k.AccKeeper.GetAccount(ctx, msg.Owner)
	if acc == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "cannot find account")
	}

	pubKey := acc.GetPubKey()
	if pubKey == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidPubKey, "pubkey on account is not set")
	}

	// verify signature with msg sender's address and requested scope
	msg2verify := "Client:" + msg.Client.String() + "Scope" + msg.Action
	if !pubKey.VerifyBytes([]byte(msg2verify), msg.Sig) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "signature verification failed; Client doesn't have access to bank belongs to Resource Owner")
	}

	sdkError := k.CoinKeeper.SendCoins(ctx, msg.Owner, msg.Client, msg.Amount)
	if sdkError != nil {
		return nil, sdkError
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAccessResource),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Client.String()),
			sdk.NewAttribute(types.AttributeOwner, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeAction, msg.Action),
			sdk.NewAttribute(types.AttributeAmount, msg.Amount.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
