package cli

import (
	"bufio"
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

func GetCmdGetToken(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-token [client] [action]",
		Short: "Grant permission to Client to access resource , i.e. transfering tokens from users' wallet",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			keybase := txBldr.Keybase()
			// _, _ = fmt.Fprintf(os.Stderr, "%+v\n\n\n", keybase)
			if keybase == nil {
				return nil
			}

			fromName := cliCtx.GetFromName()
			passphrase := keys.DefaultKeyPass
			msg2sign := "Client:" + args[0] + "Scope" + args[1]

			// _, _ = fmt.Fprintf(os.Stderr, "before sign: %+v\n\n\n", msg2sign)
			sigBytes, _, err := keybase.Sign(fromName, passphrase, []byte(msg2sign))
			if err != nil {
				return err
			}

			file, err := os.OpenFile(
				"accessToken.txt",
				os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
				0666,
			)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = file.Write(sigBytes)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
