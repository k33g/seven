package scripts

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"seven/types"
)

// RunAfterScriptIfExists runs the after script if it exists.
//
// It takes a manifest of type types.Manifest as a parameter.
// The function checks if the after script in the manifest is not empty.
// If it is not empty, it executes the after script using the exec.Command function.
// The stdout of the executed script is captured and printed.
// If there is an error running the after script, it logs a fatal error.
// If the file "tmp.seven.env" exists, it is removed.
// If there is an error removing the file, it logs a fatal error.
func RunAfterScriptIfExists(manifest types.Manifest) {
	if manifest.AfterScript != "" {

		//fmt.Println("🚀 after script:")
		cmd := exec.Command("sh", "-c", manifest.AfterScript)
		stdout, err := cmd.Output()
		//err = cmd.Run()
		if err != nil {
			log.Fatal("😡 error running after script: ", err)
		}
		fmt.Println(string(stdout))

		// remove the file if it exists
		if _, err := os.Stat("tmp.seven.env"); err == nil {
			err = os.Remove("tmp.seven.env")
			if err != nil {
				log.Fatal("😡 error removing tmp.seven.env: ", err)
			}
		}
	}
}
