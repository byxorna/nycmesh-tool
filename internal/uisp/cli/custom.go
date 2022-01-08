package cli

import (
	"fmt"
	"os"

	"github.com/byxorna/nycmesh-tool/generated/go/uisp/client"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewClientCustom(cmd *cobra.Command, args []string) (*client.UISPAPI, error) {
	return makeClient(cmd, args)
}

func MakeRootCmdCustom(use string, short string) (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   use,
		Short: short,
	}
	var err error

	cmd := bindFlags(rootCmd)

	if cmd, err = registerOperations(rootCmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func initViperConfigsCustom() {

	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".nycmesh-tool")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using (uisp) config file:", viper.ConfigFileUsed())
	}
}

func bindFlags(rootCmd *cobra.Command) *cobra.Command {
	//cobra.OnInitialize(initViperConfigsCustom)

	// register basic flags
	rootCmd.PersistentFlags().String("hostame", "uisp.mesh", "hostname of the service")
	viper.BindPFlag("uisp.hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("uisp.scheme", rootCmd.PersistentFlags().Lookup("scheme"))

	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	rootCmd.PersistentFlags().Bool("skip-verify-tls", false, fmt.Sprintf("sets &tls.Config{InsecureSkipVerify: true} in UISP HTTP Client"))
	viper.BindPFlag("uisp.skip-verify-tls", rootCmd.PersistentFlags().Lookup("skip-verify-tls"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	viper.BindPFlag("uisp.debug", rootCmd.PersistentFlags().Lookup("debug"))

	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	return rootCmd
}

// registerAuthInoWriterFlagsCustom registers all flags needed to perform authentication
func registerAuthInoWriterFlagsCustom(cmd *cobra.Command) error {
	/*x-auth-token User authorization token*/
	cmd.PersistentFlags().String("x-auth-token", "", `User authorization token`)
	viper.BindPFlag("uisp.x-auth-token", cmd.PersistentFlags().Lookup("x-auth-token"))
	return nil
}

// makeAuthInfoWriterCustom retrieves cmd flags and construct an auth info writer
func makeAuthInfoWriterCustom(cmd *cobra.Command) (runtime.ClientAuthInfoWriter, error) {
	auths := []runtime.ClientAuthInfoWriter{}
	/*x-auth-token User authorization token*/
	if viper.IsSet("uisp.x-auth-token") {
		XAuthTokenKey := viper.GetString("uisp.x-auth-token")
		auths = append(auths, httptransport.APIKeyAuth("x-auth-token", "header", XAuthTokenKey))
	}
	if len(auths) == 0 {
		logDebugf("Warning: No auth params detected.")
		return nil, nil
	}
	// compose all auths together
	return httptransport.Compose(auths...), nil
}

func registerOperations(rootCmd *cobra.Command) (*cobra.Command, error) {
	// register security flags
	if err := registerAuthInoWriterFlagsCustom(rootCmd); err != nil {
		return nil, err
	}
	// add all operation groups
	operationGroupAuthorizationCmd, err := makeOperationGroupAuthorizationCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupAuthorizationCmd)

	operationGroupBackupsCmd, err := makeOperationGroupBackupsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupBackupsCmd)

	operationGroupCrmCmd, err := makeOperationGroupCrmCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupCrmCmd)

	operationGroupDataLinksCmd, err := makeOperationGroupDataLinksCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupDataLinksCmd)

	operationGroupDevicesCmd, err := makeOperationGroupDevicesCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupDevicesCmd)

	operationGroupDiscoveryCmd, err := makeOperationGroupDiscoveryCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupDiscoveryCmd)

	operationGroupExportCmd, err := makeOperationGroupExportCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupExportCmd)

	operationGroupFirmwareCmd, err := makeOperationGroupFirmwareCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupFirmwareCmd)

	operationGroupGatewaysCmd, err := makeOperationGroupGatewaysCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupGatewaysCmd)

	operationGroupInstallationsCmd, err := makeOperationGroupInstallationsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupInstallationsCmd)

	operationGroupLogsCmd, err := makeOperationGroupLogsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupLogsCmd)

	operationGroupOutagesCmd, err := makeOperationGroupOutagesCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupOutagesCmd)

	operationGroupServerCmd, err := makeOperationGroupServerCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupServerCmd)

	operationGroupSimulationCmd, err := makeOperationGroupSimulationCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupSimulationCmd)

	operationGroupSitesCmd, err := makeOperationGroupSitesCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupSitesCmd)

	operationGroupSpeedTestCmd, err := makeOperationGroupSpeedTestCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupSpeedTestCmd)

	operationGroupTasksCmd, err := makeOperationGroupTasksCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupTasksCmd)

	operationGroupTokenCmd, err := makeOperationGroupTokenCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupTokenCmd)

	operationGroupTrafficCmd, err := makeOperationGroupTrafficCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupTrafficCmd)

	operationGroupUsersCmd, err := makeOperationGroupUsersCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupUsersCmd)

	operationGroupVaultCmd, err := makeOperationGroupVaultCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupVaultCmd)

	operationGroupWebSocketCmd, err := makeOperationGroupWebSocketCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupWebSocketCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.UISPAPI, error) {
	hostname := viper.GetString("uisp.hostname")
	scheme := viper.GetString("uisp.scheme")

	httpc, err := httptransport.TLSClient(httptransport.TLSClientOptions{
		InsecureSkipVerify: viper.GetBool("uisp.skip-verify-tls"),
	})
	if err != nil {
		return nil, err
	}

	r := httptransport.NewWithClient(hostname, client.DefaultBasePath, []string{scheme}, httpc)
	r.SetDebug(viper.GetBool("uisp.debug"))

	r.Consumers["application/json"] = runtime.JSONConsumer()

	r.Producers["application/json"] = runtime.JSONProducer()

	auth, err := makeAuthInfoWriterCustom(cmd)
	if err != nil {
		return nil, err
	}
	r.DefaultAuthentication = auth

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}
