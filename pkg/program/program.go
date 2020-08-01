package program

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xrash/gonel/pkg/runner"
	"os"
)

type Program struct {
}

func NewProgram() *Program {
	return &Program{}
}

func (p *Program) Run() {
	var imports []string

	rootCmd := &cobra.Command{
		Use:   "gonel",
		Short: "Gonel is the Go one-liner",
		Long:  `Gonel is the Go one-liner. Run 'gonel --help' to read more.`,
		Run: func(cmd *cobra.Command, args []string) {

			if len(args) < 1 {
				cmd.Usage()
				p.Exit(1)
			}

			r := &runner.Runner{
				Stdin:   os.Stdin,
				Stdout:  os.Stdout,
				Stderr:  os.Stderr,
				Tempdir: os.TempDir(),
			}

			if err := r.Run(imports, args); err != nil {
				fmt.Println(err)
				return
			}
		},
	}

	rootCmd.Flags().StringArrayVarP(&imports, "import", "i", []string{}, "package name to import")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		p.Exit(1)
	}

	p.Exit(0)
}

func (p *Program) Exit(code int) {
	os.Exit(code)
}
