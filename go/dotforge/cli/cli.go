// cli/cli.go

package cli

import (
	"dotforge/config"
	"dotforge/symlink"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Define and initialize flags
var (
	setDefaultPath = flag.String("set-default-path", "", "Set the default path")
	sdp            = flag.String("sdp", "", "Shorthand for --set-default-path")

	listDefaultPath = flag.Bool("list-default-path", false, "List the current default path")
	ldp             = flag.Bool("ldp", false, "Shorthand for --list-default-path")

	newFiletypes = flag.String("new-filetypes", "", "Comma-separated list of new file types to add")
	nft          = flag.String("nft", "", "Shorthand for --new-filetypes")

	removeFiletypes = flag.String("remove-filetypes", "", "Comma-separated list of file types to remove")
	rft             = flag.String("rft", "", "Shorthand for --remove-filetypes")

	listFiletypes = flag.Bool("list-filetypes", false, "List all approved file types")
	lft           = flag.Bool("lft", false, "Shorthand for --list-filetypes")

	ignorePaths = flag.String("ignore-paths", "", "Comma-separated list of paths to ignore")
	ip          = flag.String("ip", "", "Shorthand for --ignore-paths")

	removeIgnored = flag.String("remove-ignored", "", "Comma-separated list of paths to remove from ignored paths")
	ri            = flag.String("ri", "", "Shorthand for --remove-ignored")

	listIgnored = flag.Bool("list-ignored", false, "List all ignored paths")
	li          = flag.Bool("li", false, "Shorthand for --list-ignored")

	scan = flag.Bool("scan", false, "Scan the current directory for blocked or unaccepted files")
	s    = flag.Bool("s", false, "Shorthand for --scan")

	maxLines = flag.Int("max-lines", 20, "Set max lines to display before opening in editor")
	ml       = flag.Int("ml", 20, "Shorthand for --max-lines")

	help = flag.Bool("help", false, "Show help")
	h    = flag.Bool("h", false, "Shorthand for --help")

	initRepo = flag.String("init", "", "Initialize a dotfile repository")
	ir       = flag.String("ir", "", "Shorthand for --init")
)

// ParseFlags handles command-line arguments
func ParseFlags() {
	fmt.Println("Parsing flags...") // Debug print
	flag.Parse()

	// Handle init flag
	if *initRepo != "" || *ir != "" {
		path := *initRepo
		if path == "" {
			path = *ir
		}
		err := config.InitializeRepo(path)
		if err != nil {
			fmt.Println("Initialization error:", err)
		}
		return
	}

	fmt.Printf("Flags: %+v\n", os.Args)                            // Debug print
	fmt.Printf("NFlag: %d, NArg: %d\n", flag.NFlag(), flag.NArg()) // Debug print

	// Check if any flags are set
	if flag.NFlag() > 0 {
		handleFlags()
		return
	}

	// Handle positional arguments if no flags are provided
	handlePositionalArgs(flag.Args())
}

// handleFlags processes all provided flags
func handleFlags() {
	fmt.Println("Handling flags...") // Debug print

	switch {
	case *help || *h:
		fmt.Println("Help flag detected") // Debug print
		if len(flag.Args()) > 0 {
			PrintFlagHelp(flag.Args()[0]) // Provide detailed help for a specific flag
		} else {
			PrintHelp()
		}
	case *setDefaultPath != "" || *sdp != "":
		fmt.Println("Set default path flag detected") // Debug print
		handleSetDefaultPath()
	case *listDefaultPath || *ldp:
		fmt.Println("List default path flag detected") // Debug print
		handleListDefaultPath()
	case *newFiletypes != "" || *nft != "":
		fmt.Println("New file types flag detected") // Debug print
		handleNewFiletypes(*newFiletypes)
	case *removeFiletypes != "" || *rft != "":
		fmt.Println("Remove file types flag detected") // Debug print
		handleRemoveFiletypes(*removeFiletypes)
	case *listFiletypes || *lft:
		fmt.Println("List file types flag detected") // Debug print
		handleListFiletypes()
	case *ignorePaths != "" || *ip != "":
		fmt.Println("Ignore paths flag detected") // Debug print
		handleIgnorePaths(*ignorePaths)
	case *removeIgnored != "" || *ri != "":
		fmt.Println("Remove ignored path flag detected") // Debug print
		handleRemoveIgnored(*removeIgnored)
	case *listIgnored || *li:
		fmt.Println("List ignored paths flag detected") // Debug print
		handleListIgnored()
	case *scan || *s:
		fmt.Println("Scan flag detected") // Debug print
		handleScan()
	default:
		fmt.Println("Invalid or unrecognized flag provided. Use --help for usage details.")
	}
}

// handlePositionalArgs processes positional arguments
func handlePositionalArgs(args []string) {
	fmt.Printf("Handling positional args: %v\n", args) // Debug print

	switch len(args) {
	case 0:
		fmt.Println("No positional args, using defaults") // Debug print
		sourceDir := "."
		targetDir := config.ReadDefaultPath()
		if targetDir == "" {
			fmt.Println("Error: No default target set. Please use --set-default-path to set one.")
			return
		}
		fmt.Printf("Source: %s\n", sourceDir) // Debug print
		fmt.Printf("Target: %s\n", targetDir) // Debug print
		symlink.CreateSymlinks(sourceDir, targetDir)
	case 1:
		fmt.Printf("Single positional argument provided, target is %s\n", args[0]) // Debug print
		fmt.Printf("Source: .\n")
		fmt.Printf("Target: %s\n", args[0])
	case 2:
		fmt.Printf("Two positional arguments provided: source %s, target %s\n", args[0], args[1]) // Debug print
		fmt.Printf("Source: %s\n", args[0])
		fmt.Printf("Target: %s\n", args[1])
		symlink.CreateSymlinks(args[0], args[1])
	default:
		fmt.Println("Too many arguments. Please provide up to two arguments: [source] [target].")
	}
}

// Individual flag handlers
func handleSetDefaultPath() {
	fmt.Println("Executing handleSetDefaultPath") // Debug print
	target := *setDefaultPath
	if target == "" {
		target = *sdp
	}
	// Validate target path
	if _, err := os.Stat(target); os.IsNotExist(err) {
		fmt.Printf("Error: Target path '%s' does not exist.\n", target)
		return
	}
	if err := config.SetDefaultPath(target); err != nil {
		fmt.Printf("Error setting default path: %s\n", err)
		return
	}
	fmt.Printf("Default target directory set to %s\n", target)
}

func handleListDefaultPath() {
	fmt.Println("Executing handleListDefaultPath") // Debug print
	defaultPath := config.ReadDefaultPath()
	fmt.Printf("Current default path: %s\n", defaultPath)
}

func handleNewFiletypes(filetypes string) {
	fmt.Println("Executing handleNewFiletypes") // Debug print
	extensions := strings.Split(filetypes, ",")
	for i := range extensions {
		extensions[i] = strings.TrimSpace(extensions[i])
	}
	config.AddFiletypes(extensions)
}

func handleRemoveFiletypes(filetypes string) {
	fmt.Println("Executing handleRemoveFiletypes") // Debug print
	extensions := strings.Split(filetypes, ",")
	for i := range extensions {
		extensions[i] = strings.TrimSpace(extensions[i])
	}
	config.RemoveFiletypes(extensions)
}

func handleListFiletypes() {
	fmt.Println("Executing handleListFiletypes") // Debug print
	config.ListFiletypes()
}

func handleIgnorePaths(paths string) {
	fmt.Println("Executing handleIgnorePaths") // Debug print
	pathList := strings.Split(paths, ",")
	for i := range pathList {
		pathList[i] = strings.TrimSpace(pathList[i])
	}
	config.AddIgnoredPaths(pathList)
}

func handleRemoveIgnored(paths string) {
	fmt.Println("Executing handleRemoveIgnored") // Debug print
	pathList := strings.Split(paths, ",")
	for i := range pathList {
		pathList[i] = strings.TrimSpace(pathList[i])
	}
	config.RemoveIgnoredPaths(pathList)
}

func handleListIgnored() {
	fmt.Println("Executing handleListIgnored") // Debug print
	config.ListIgnoredPaths()
}

func handleScan() {
	fmt.Println("Executing handleScan") // Debug print
	// Placeholder for scanning functionality
	// Integrate the scanning logic here if applicable
}
