package client

import (
	"fmt"
	/*"strings"
	"text/tabwriter"*/

	Cli "github.com/daolinet/daolictl/cli"
)

const (
	CONNECTED    = "ACCEPT"
	DISCONNECTED = "DROP"
)

// Usage: daolictl policy
/*func (cli *DaoliCli) CmdPolicy(args ...string) error {
	cmd := Cli.Subcmd("policy", []string{"COMMAND [OPTIONS]"}, policyUsage(), false)
	err := ParseFlags(cmd, args, true)
	cmd.Usage()
	return err
}*/

//Usage: daolictl show <CONTAINER:CONTAINER>
func (cli *DaoliCli) CmdShow(args ...string) error {
	cmd := Cli.Subcmd("show", []string{"CONTAINER:CONTAINER"}, Cli.GlobalCommands["show"].Description, true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	status, err := cli.client.PolicyShow(cmd.Arg(0))
	if err != nil {
		return err
	}

	if status == CONNECTED {
		fmt.Fprintf(cli.out, "%s\n", "CONNECTED")
	} else if status == DISCONNECTED {
		fmt.Fprintf(cli.out, "%s\n", "DISCONNECTED")
	}

	return nil
}

// Usage: daolictl clear <CONTAINER:CONTAINER>
func (cli *DaoliCli) CmdClear(args ...string) error {
	cmd := Cli.Subcmd("clear", []string{"CONTAINER:CONTAINER"}, Cli.GlobalCommands["clear"].Description, true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	err := cli.client.PolicyDelete(cmd.Arg(0))
	if err != nil {
		return err
	}

	fmt.Fprintf(cli.out, "%s\n", "CLEARED")
	return nil
}

// Usage: daolictl connect <CONTAINER:CONTAINER>
func (cli *DaoliCli) CmdConnect(args ...string) error {
	cmd := Cli.Subcmd("connect", []string{"CONTAINER:CONTAINER"}, Cli.GlobalCommands["connect"].Description, true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	err := cli.client.PolicyUpdate(cmd.Arg(0), CONNECTED)
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", "CONNECTED")
	return nil
}

// Usage: daolictl disconnect <CONTAINER:CONTAINER>
func (cli *DaoliCli) CmdDisconnect(args ...string) error {
	cmd := Cli.Subcmd("disconnect", []string{"CONTAINER:CONTAINER"}, Cli.GlobalCommands["disconnect"].Description, true)
	if err := ParseFlags(cmd, args, true); err != nil {
		return err
	}

	if len(cmd.Args()) <= 0 {
		cmd.Usage()
		return nil
	}

	err := cli.client.PolicyUpdate(cmd.Arg(0), DISCONNECTED)
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "%s\n", "DISCONNECTED")
	return nil
}

/*func policyUsage() string {
	policyCommands := map[string]string{
		"list":   "List all policy",
		"create": "Create a rule",
		"delete": "Delete a rule",
	}

	help := "Commands:\n"

	for cmd, description := range policyCommands {
		help += fmt.Sprintf("  %-25.25s%s\n", cmd, description)
	}

	help += fmt.Sprintf("\nRun 'daolictl policy COMMAND --help' for more information on a command.")
	return help
}*/
