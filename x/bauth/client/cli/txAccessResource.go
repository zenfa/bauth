package cli

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/zenfa/bauth/x/bauth/types"
)

func GetCmdAccessResource(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "access-resource [owner] [action] [amount]",
		Short: "Access resource on behalf of a user",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			var action = args[1]

			file, err := os.Open("accessToken.txt")
			if err != nil {
				return err
			}
			defer file.Close()

			sig, err := ioutil.ReadAll(file)

			amount, err := sdk.ParseCoins(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgAccessResource(owner, cliCtx.GetFromAddress(), action, amount, sig)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
