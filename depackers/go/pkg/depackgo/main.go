package depackgo


import (
	"os"
	"path"
	"log"
	"fmt"
	"flag"
)



func getCommandPackages(
	arguments []string,
) []string {
	// the first and second element of arguments are the commands
	// make an array counting up from the third element
	commandPackages := make([]string, len(arguments) - 2)
	if len(arguments) >= 3 {
		for i := 2; i < len(arguments); i++ {
			commandPackages = append(commandPackages, arguments[i])
		}
	}
	return commandPackages
}


func Run() {
	// Subcommands
	loginCommand := flag.NewFlagSet("login", flag.ExitOnError)
	logoutCommand := flag.NewFlagSet("logout", flag.ExitOnError)
	statusCommand := flag.NewFlagSet("status", flag.ExitOnError)
	initCommand := flag.NewFlagSet("init", flag.ExitOnError)
	initializeCommand := flag.NewFlagSet("initialize", flag.ExitOnError)
	installCommand := flag.NewFlagSet("install", flag.ExitOnError)
	uninstallCommand := flag.NewFlagSet("uninstall", flag.ExitOnError)
	publishCommand := flag.NewFlagSet("publish", flag.ExitOnError)

	help :=`
	depack-go <command>

	commands:
	login					logs into depack.plurid.com
	logout					logs out of depack.plurid.com
	status					shows cli application status
	init					initialize depack-go package ('initialize' can also be used as command)
	install <package-name>			install package or multiple space-separated packages
	uninstall <package-name>		uninstall package or multiple space-separated packages
	publish					publish package to depack.plurid.com/go
	`



    // Verify that a subcommand has been provided
    // os.Arg[0] is the main command
    // os.Arg[1] will be the subcommand
    if len(os.Args) < 2 {
        fmt.Println(help)
        os.Exit(1)
    }


    // Switch on the subcommand
    // Parse the flags for appropriate FlagSet
    // FlagSet.Parse() requires a set of arguments to parse as input
    // os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
    switch os.Args[1] {
	case "login":
        loginCommand.Parse(os.Args[2:])
	case "logout":
        logoutCommand.Parse(os.Args[2:])
	case "status":
        statusCommand.Parse(os.Args[2:])
	case "init":
        initCommand.Parse(os.Args[2:])
	case "initialize":
        initializeCommand.Parse(os.Args[2:])
    case "install":
        installCommand.Parse(os.Args[2:])
    case "uninstall":
        uninstallCommand.Parse(os.Args[2:])
    case "publish":
        publishCommand.Parse(os.Args[2:])
    default:
        flag.PrintDefaults()
        os.Exit(1)
    }


	if loginCommand.Parsed() {
		fmt.Println("login account")
	}


	if logoutCommand.Parsed() {
		fmt.Println("logout account")
	}


	if statusCommand.Parsed() {
		fmt.Println("account status")
	}


	if initCommand.Parsed() || initializeCommand.Parsed() {
		directory := "."
		directoryArgument := ""
		if len(os.Args) == 3 {
			directoryArgument = os.Args[2]
		}
		if directoryArgument != "" {
			directory = directoryArgument
		}

		workingDirectory, error := os.Getwd()
		if error != nil {
			log.Fatal(error)
		}

		directoryPath := path.Join(workingDirectory, directory)


		fmt.Printf("directory: %s", directory)
		fmt.Printf("\n")

		fmt.Printf("directoryPath: %s", directoryPath)
		fmt.Printf("\n")
	}


	if installCommand.Parsed() {
		installPackages := getCommandPackages(os.Args)
		fmt.Println(installPackages)
	}


	if uninstallCommand.Parsed() {
		uninstallPackages := getCommandPackages(os.Args)
		fmt.Println(uninstallPackages)
	}


	if publishCommand.Parsed() {
		fmt.Println("publish package")
	}
}
