package main

import "aspect.build/cli/cmd"

func main() {
	// Detect whether we are being run as a tools/bazel wrapper (look for BAZEL_REAL in the environment)
	// If so,
    //     Is this a bazel-native command? just call through to bazel without touching the arguments for now
	//     Is this an aspect-custom command? (like `outputs`) then write an implementation
	// otherwise,
	//     we are installing ourselves. Check with the user they intended to do that.
	//     then create 
	//         - a WORKSPACE file, ask the user for the repository name if interactive
	//     ask the user if they want to install for all users of the workspace, if so
	//         - tools/bazel file and put our bootstrap code in there
	//     
	cmd.Execute()
}