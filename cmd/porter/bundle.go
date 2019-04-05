package main

import (
	"github.com/deislabs/porter/pkg/porter"
	"github.com/spf13/cobra"
)

func buildBundleCommands(p *porter.Porter) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bundle",
		Short: "bundle commands",
	}

	cmd.AddCommand(buildBundleInstallCommand(p))
	cmd.AddCommand(buildBundleUninstallCommand(p))

	return cmd
}

func buildBundleAliasCommands(p *porter.Porter) []*cobra.Command {
	return []*cobra.Command{
		buildInstallCommand(p),
		buildUninstallCommand(p),
	}
}

func buildBundleInstallCommand(p *porter.Porter) *cobra.Command {
	opts := porter.InstallOptions{}
	cmd := &cobra.Command{
		Use:   "install [CLAIM]",
		Short: "Install a bundle",
		Long: `Install a bundle.

The first argument is the name of the claim to create for the installation. The claim name defaults to the name of the bundle.`,
		Example: `  porter install
  porter install --insecure
  porter install MyAppInDev --file myapp/bundle.json
  porter install --param-file base-values.txt --param-file dev-values.txt --param test-mode=true --param header-color=blue
  porter install --cred azure --cred kubernetes
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.InstallBundle(opts)
		},
	}

	f := cmd.Flags()
	f.BoolVar(&opts.Insecure, "insecure", false,
		"Allow working with untrusted bundles")
	f.StringVarP(&opts.File, "file", "f", "",
		"Path to the CNAB definition to install. Defaults to the bundle in the current directory.")
	f.StringSliceVar(&opts.ParamFiles, "param-file", nil,
		"Path to a parameters definition file for the bundle, each line in the form of NAME=VALUE. May be specified multiple times.")
	f.StringSliceVar(&opts.Params, "param", nil,
		"Define an individual parameter in the form NAME=VALUE. Overrides parameters set with the same name using --param-file. May be specified multiple times.")
	f.StringSliceVarP(&opts.CredentialIdentifiers, "cred", "c", nil,
		"Credential to use when installing the bundle. May be either a named set of credentials or a filepath, and specified multiple times.")

	return cmd
}

func buildInstallCommand(p *porter.Porter) *cobra.Command {
	return buildBundleInstallCommand(p)
}

func buildBundleUninstallCommand(p *porter.Porter) *cobra.Command {
	opts := porter.UninstallOptions{}
	cmd := &cobra.Command{
		Use:   "uninstall [CLAIM]",
		Short: "Uninstall a bundle",
		Long: `Uninstall a bundle

The first argument is the name of the claim to uninstall. The claim name defaults to the name of the bundle.`,
		Example: `  porter uninstall
  porter uninstall --insecure
  porter uninstall MyAppInDev --file myapp/bundle.json
  porter uninstall --param-file base-values.txt --param-file dev-values.txt --param test-mode=true --param header-color=blue
  porter uninstall --cred azure --cred kubernetes
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.Validate(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return p.UninstallBundle(opts)
		},
	}

	f := cmd.Flags()
	f.BoolVar(&opts.Insecure, "insecure", false,
		"Allow working with untrusted bundles")
	f.StringVarP(&opts.File, "file", "f", "",
		"Path to the CNAB definition to uninstall. Defaults to the bundle in the current directory. Optional unless a newer version of the bundle should be used to uninstall the bundle.")
	f.StringSliceVar(&opts.ParamFiles, "param-file", nil,
		"Path to a parameters definition file for the bundle, each line in the form of NAME=VALUE. May be specified multiple times.")
	f.StringSliceVar(&opts.Params, "param", nil,
		"Define an individual parameter in the form NAME=VALUE. Overrides parameters set with the same name using --param-file. May be specified multiple times.")
	f.StringSliceVarP(&opts.CredentialIdentifiers, "cred", "c", nil,
		"Credential to use when uninstalling the bundle. May be either a named set of credentials or a filepath, and specified multiple times.")

	return cmd
}

func buildUninstallCommand(p *porter.Porter) *cobra.Command {
	return buildBundleUninstallCommand(p)
}