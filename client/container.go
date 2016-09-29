package client

import (
	"fmt"
	//"strings"
	"text/tabwriter"

	Cli "github.com/daolinet/daolictl/cli"
)

// Usage: daolictl container
func (cli *DaoliCli) CmdContainer(args ...string) error {
	cmd := Cli.Subcmd("container", []string{"COMMAND [OPTIONS]"}, containerUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}

//Usage: daolictl container move <CONTAINER>
func (cli *DaoliCli) CmdContainerMove(args ...string) error {
	cmd := Cli.Subcmd("container move", []string{"[OPTIONS] CONTAINER"}, "Rescheduler container to new container.", false)
        node := cmd.String("node", "", "Node id or name")
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	container, err := cli.client.ResetContainer(cmd.Arg(0), *node)
	if err != nil {
		return err
	}

        fmt.Fprintf(cli.out, "%s\n", container)
	return nil
}


//Usage: daolictl container shownet <CONTAINER>
func (cli *DaoliCli) CmdContainerShownet(args ...string) error {
	cmd := Cli.Subcmd("container shownet", []string{"CONTAINER"}, "Show container network details.", false)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	containerNetworks, err := cli.client.ShowContainer(cmd.Arg(0))
	if err != nil {
		return err
	}

        wr := tabwriter.NewWriter(cli.out, 20, 1, 3, ' ', 0)
        fmt.Fprintln(wr, "IPADDRESS\tMACADDRESS\tGATEWAY\tNETWORKNAME\tVIPADDRESS")
        for _, net := range containerNetworks {
            fmt.Fprintf(wr, "%s\t%s\t%s\t%s\t%s",
                        net.IPAddress,
                        net.MacAddress,
                        net.Gateway,
                        net.NetworkName,
                        net.VIPAddress,
            )
            fmt.Fprintf(wr, "\n")
        }
        wr.Flush()
	return nil
}


func containerUsage() string {
	containerCommands := map[string]string{
		"move":   "Rescheduler container.",
		"shownet":    "Show container network info.",
	}

	help := "Commands:\n"

	for cmd, description := range containerCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl container COMMAND --help' for more information on a command.")
	return help
}
