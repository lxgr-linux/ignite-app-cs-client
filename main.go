package main

import (
	"context"

	hplugin "github.com/hashicorp/go-plugin"
	"github.com/ignite/cli/v29/ignite/pkg/cliui"
	"github.com/ignite/cli/v29/ignite/pkg/errors"
	"github.com/ignite/cli/v29/ignite/services/plugin"
	"github.com/lxgr-linux/ignite-app-cs-client/cmd"
	"github.com/lxgr-linux/ignite-app-cs-client/deptools"
	"github.com/lxgr-linux/ignite-app-cs-client/gen"
)

type app struct{}

func (app) Manifest(context.Context) (*plugin.Manifest, error) {
	return &plugin.Manifest{
		Name:     "cs-client",
		Commands: cmd.GetCommands(),
	}, nil
}

func (app) Execute(ctx context.Context, cmd *plugin.ExecutedCommand, api plugin.ClientAPI) error {
	session := cliui.New()
	defer session.End()

	chainInfo, err := api.GetChainInfo(ctx)
	if err != nil {
		return errors.Errorf("failed to get chain info: %s", err)
	}

	g, err := gen.New(ctx, cmd, chainInfo)
	if err != nil {
		return errors.Errorf("failed to init genrator: %s", err)
	}

	session.StartSpinner("Installing dependencies...")
	err = deptools.ProvideTools(ctx, chainInfo.AppPath)
	if err != nil {
		return err
	}
	session.StopSpinner()

	err = g.GenerateClients(ctx)
	if err != nil {
		return err
	}

	session.StartSpinner("Generating...")
	err = g.GenerateCsproj()
	if err != nil {
		return err
	}

	err = g.GenerateClient(ctx)
	if err != nil {
		return err
	}

	err = g.InstallStaticFiles()
	if err != nil {
		return err
	}

	session.StopSpinner()
	return nil
}

func (app) ExecuteHookPre(context.Context, *plugin.ExecutedHook, plugin.ClientAPI) error {
	return nil
}

func (app) ExecuteHookPost(context.Context, *plugin.ExecutedHook, plugin.ClientAPI) error {
	return nil
}

func (app) ExecuteHookCleanUp(context.Context, *plugin.ExecutedHook, plugin.ClientAPI) error {
	return nil
}

func main() {
	hplugin.Serve(&hplugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig(),
		Plugins: map[string]hplugin.Plugin{
			"cs-client": plugin.NewGRPC(&app{}),
		},
		GRPCServer: hplugin.DefaultGRPCServer,
	})
}
