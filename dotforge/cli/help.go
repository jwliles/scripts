package cli

import "fmt"

func PrintCommandHelp(command string) {
	switch command {
	case "set-default-target":
		fmt.Println("Usage: --set-default-target, -spt <path>\nSets the default target directory for symlinks.")
	case "new-filetypes":
		fmt.Println("Usage: --new-filetypes, -nft <extensions>\nAdds new file extensions to the approved list.")
	// Add cases for other commands...
	default:
		PrintHelp() // Print general help if command is not recognized
	}
}

// PrintHelp Print help functions (implementations assumed to exist)
func PrintHelp() {
	fmt.Println(`Usage of Dotforge:
  -help
        Show help
  -ignore-paths string
        Add paths to be ignored
  -init
  	  Initialize a dotfile repository at the specified path (default is current directory)	
  -list-default-path
        List the current default path
  -list-filetypes
        List all approved file types
  -list-ignored
        List all ignored paths
  -max-lines int
        Set max lines to display before opening in editor (default 20)
  -new-filetypes string
        Add new file types
  -remove-filetypes string
        Remove file types from the approved list
  -remove-ignored string
        Remove an ignored path
  -scan
        Scan the current directory for blocked or unaccepted files
  -sdp string
        Shorthand for --set-default-path
  -set-default-path string
        Set the default path`)
}

func PrintFlagHelp(flag string) {
	flagHelp := map[string]string{
		"set-default-path":  "Set the default path to the given directory",
		"sdp":               "Shorthand for --set-default-path",
		"list-default-path": "List the current default path",
		"new-filetypes":     "Comma-separated list of new file types to add to the approved list",
		"remove-filetypes":  "Comma-separated list of file types to remove from the approved list",
		"list-filetypes":    "List all approved file types",
		"lft":               "Shorthand for --list-filetypes",
		"ignore-paths":      "Comma-separated list of paths to ignore",
		"remove-ignored":    "Comma-separated list of paths to remove from the ignored list",
		"list-ignored":      "List all ignored paths",
		"scan":              "Scan the current directory for blocked or unaccepted files",
		"max-lines":         "Set the maximum number of lines to display before opening in editor",
		"help":              "Show help",
	}

	if helpText, exists := flagHelp[flag]; exists {
		fmt.Printf("%s: %s\n", flag, helpText)
	} else {
		fmt.Printf("No help available for flag: %s\n", flag)
	}
}
