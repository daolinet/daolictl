package client

import (
	"fmt"
	/*"strings"
	"text/tabwriter"*/

	Cli "github.com/daolinet/daolictl/cli"
)

// Usage: daolictl container
func (cli *DaoliCli) CmdContainer(args ...string) error {
	cmd := Cli.Subcmd("container", []string{"COMMAND [OPTIONS]"}, containerUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}

//Usage: daolictl container reset <CONTAINER>
func (cli *DaoliCli) CmdContainerReset(args ...string) error {
	cmd := Cli.Subcmd("container reset", []string{"CONTAINER"}, "Rescheduler container to new container.", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	container, err := cli.client.ResetContainer(cmd.Arg(0))
	if err != nil {
		return err
	}

        fmt.Fprintf(cli.out, "%s\n", container)
	return nil
}


func containerUsage() string {
	containerCommands := map[string]string{
		"reset":   "Rescheduler container.",
		"show":    "Show container info.",
	}

	help := "Commands:\n"

	for cmd, description := range containerCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl container COMMAND --help' for more information on a command.")
	return help
}
