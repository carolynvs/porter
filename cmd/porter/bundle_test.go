package main

import (
	"strings"
	"testing"

	"get.porter.sh/porter/pkg/porter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuildBundleActionCommand_Flags(t *testing.T) {
	commonFlags := []string{
		// action flags
		"allow-docker-host-access",
		"file",
		"cnab-file",
		"parameter-set",
		"param",
		"cred",
		"driver",
		// pull flags
		"tag",
		"reference",
		"insecure-registry",
		"force",
	}

	t.Run("install", func(t *testing.T) {
		p := porter.NewTestPorter(t)
		cmd := buildBundleInstallCommand(p.Porter)
		assert.Equal(t, "install", cmd.Name())

		for _, f := range commonFlags {
			assert.NotNil(t, cmd.Flag(f), "expected flag %s to be defined on install", f)
		}
	})

	t.Run("upgrade", func(t *testing.T) {
		p := porter.NewTestPorter(t)
		cmd := buildBundleUpgradeCommand(p.Porter)
		assert.Equal(t, "upgrade", cmd.Name())

		for _, f := range commonFlags {
			assert.NotNil(t, cmd.Flag(f), "expected flag %s to be defined on upgrade", f)
		}
	})

	t.Run("invoke", func(t *testing.T) {
		invokeFlags := make([]string, len(commonFlags))
		copy(invokeFlags, commonFlags)
		invokeFlags = append(invokeFlags, "action")

		p := porter.NewTestPorter(t)
		cmd := buildBundleInvokeCommand(p.Porter)
		assert.Equal(t, "invoke", cmd.Name())

		for _, f := range invokeFlags {
			assert.NotNil(t, cmd.Flag(f), "expected flag %s to be defined on invoke", f)
		}
	})

	t.Run("uninstall", func(t *testing.T) {
		uninstallFlags := make([]string, len(commonFlags))
		copy(uninstallFlags, commonFlags)
		uninstallFlags = append(uninstallFlags, "delete", "force-delete")

		p := porter.NewTestPorter(t)
		cmd := buildBundleUninstallCommand(p.Porter)
		assert.Equal(t, "uninstall", cmd.Name())

		for _, f := range uninstallFlags {
			assert.NotNil(t, cmd.Flag(f), "expected flag %s to be defined on invoke", f)
		}
	})
}

func TestValidateInstallCommand(t *testing.T) {
	testcases := []struct {
		name      string
		args      string
		wantError string
	}{
		{"no args", "install", ""},
		{"invalid param", "install --param A:B", "invalid parameter (A:B), must be in name=value format"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			p := buildRootCommand()
			osargs := strings.Split(tc.args, " ")
			cmd, args, err := p.Find(osargs)
			require.NoError(t, err)

			err = cmd.ParseFlags(args)
			require.NoError(t, err)

			err = cmd.PreRunE(cmd, cmd.Flags().Args())
			if tc.wantError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantError)
			}
		})
	}
}

func TestValidateUninstallCommand(t *testing.T) {
	testcases := []struct {
		name      string
		args      string
		wantError string
	}{
		{"no args", "uninstall", ""},
		{"invalid param", "uninstall --param A:B", "invalid parameter (A:B), must be in name=value format"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			p := buildRootCommand()
			osargs := strings.Split(tc.args, " ")
			cmd, args, err := p.Find(osargs)
			require.NoError(t, err)

			err = cmd.ParseFlags(args)
			require.NoError(t, err)

			err = cmd.PreRunE(cmd, cmd.Flags().Args())
			if tc.wantError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantError)
			}
		})
	}
}

func TestValidateInvokeCommand(t *testing.T) {
	testcases := []struct {
		name      string
		args      string
		wantError string
	}{
		{"no args", "invoke", "--action is required"},
		{"action specified", "invoke --action status", ""},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			p := buildRootCommand()
			osargs := strings.Split(tc.args, " ")
			cmd, args, err := p.Find(osargs)
			require.NoError(t, err)

			err = cmd.ParseFlags(args)
			require.NoError(t, err)

			err = cmd.PreRunE(cmd, cmd.Flags().Args())
			if tc.wantError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantError)
			}
		})
	}
}

func TestValidateInstallationListCommand(t *testing.T) {
	testcases := []struct {
		name      string
		args      string
		wantError string
	}{
		{"no args", "installation list", ""},
		{"output json", "installation list -o json", ""},
		{"invalid format", "installation list -o wingdings", "invalid format: wingdings"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			p := buildRootCommand()
			osargs := strings.Split(tc.args, " ")
			cmd, args, err := p.Find(osargs)
			require.NoError(t, err)

			err = cmd.ParseFlags(args)
			require.NoError(t, err)

			err = cmd.PreRunE(cmd, cmd.Flags().Args())
			if tc.wantError == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tc.wantError)
			}
		})
	}
}
