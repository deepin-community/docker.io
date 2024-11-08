package plugin

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/client"
)

type fakeClient struct {
	client.Client
	pluginCreateFunc  func(createContext io.Reader, createOptions types.PluginCreateOptions) error
	pluginDisableFunc func(name string, disableOptions types.PluginDisableOptions) error
	pluginEnableFunc  func(name string, options types.PluginEnableOptions) error
	pluginRemoveFunc  func(name string, options types.PluginRemoveOptions) error
	pluginInstallFunc func(name string, options types.PluginInstallOptions) (io.ReadCloser, error)
	pluginListFunc    func(filter filters.Args) (types.PluginsListResponse, error)
	pluginInspectFunc func(name string) (*types.Plugin, []byte, error)
	pluginUpgradeFunc func(name string, options types.PluginInstallOptions) (io.ReadCloser, error)
}

func (c *fakeClient) PluginCreate(_ context.Context, createContext io.Reader, createOptions types.PluginCreateOptions) error {
	if c.pluginCreateFunc != nil {
		return c.pluginCreateFunc(createContext, createOptions)
	}
	return nil
}

func (c *fakeClient) PluginEnable(_ context.Context, name string, enableOptions types.PluginEnableOptions) error {
	if c.pluginEnableFunc != nil {
		return c.pluginEnableFunc(name, enableOptions)
	}
	return nil
}

func (c *fakeClient) PluginDisable(_ context.Context, name string, disableOptions types.PluginDisableOptions) error {
	if c.pluginDisableFunc != nil {
		return c.pluginDisableFunc(name, disableOptions)
	}
	return nil
}

func (c *fakeClient) PluginRemove(_ context.Context, name string, removeOptions types.PluginRemoveOptions) error {
	if c.pluginRemoveFunc != nil {
		return c.pluginRemoveFunc(name, removeOptions)
	}
	return nil
}

func (c *fakeClient) PluginInstall(_ context.Context, name string, installOptions types.PluginInstallOptions) (io.ReadCloser, error) {
	if c.pluginInstallFunc != nil {
		return c.pluginInstallFunc(name, installOptions)
	}
	return nil, nil
}

func (c *fakeClient) PluginList(_ context.Context, filter filters.Args) (types.PluginsListResponse, error) {
	if c.pluginListFunc != nil {
		return c.pluginListFunc(filter)
	}

	return types.PluginsListResponse{}, nil
}

func (c *fakeClient) PluginInspectWithRaw(_ context.Context, name string) (*types.Plugin, []byte, error) {
	if c.pluginInspectFunc != nil {
		return c.pluginInspectFunc(name)
	}

	return nil, nil, nil
}

func (c *fakeClient) Info(context.Context) (system.Info, error) {
	return system.Info{}, nil
}

func (c *fakeClient) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (io.ReadCloser, error) {
	if c.pluginUpgradeFunc != nil {
		return c.pluginUpgradeFunc(name, options)
	}
	return nil, nil
}
