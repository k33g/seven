package scripts

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"seven/types"

	"github.com/joho/godotenv"
)

// RunBeforeScriptIfExists runs the before script if it exists.
//
// It takes a manifest of type types.Manifest as a parameter.
// The function checks if the before script in the manifest is not empty.
// If it is not empty, it executes the before script using the exec.Command function.
// The stdout of the executed script is captured and printed.
// If there is an error running the before script, it logs a fatal error.
// If the file "tmp.seven.env" exists, it is loaded using the godotenv.Load function.
// If there is an error loading the file, it logs a fatal error.
func RunBeforeScriptIfExists(manifest types.Manifest) {
	if manifest.BeforeScript != "" {
		cmd := exec.Command("sh", "-c", manifest.BeforeScript)
		stdout, err := cmd.Output()
		//err = cmd.Run()
		if err != nil {
			log.Fatal("😡 error running before script: ", err)
		}
		fmt.Println(string(stdout))

		// test if the tmp.seven.env file exists

		if _, err := os.Stat("tmp.seven.env"); err == nil {
			// if the file exists, load it
			err = godotenv.Load("tmp.seven.env")
			if err != nil {
				log.Fatal("😡 error running before script (dotenv): ", err)
			}
		}
	}
}

/*
### Use the exported variables in `before-script` and `after-script`

To make available an environement variable created in the `before-script` scripts, you can use the **`tmp.seven.env`** file:

```yaml
before-script: |
  echo "GREETING=🎉" > tmp.seven.env
```
Then, you can use the `GREETING` environment variable with the variables substitution and in the `after-script` scripts:
```yaml
after-script: |
  echo "GREETING: ${GREETING}"
```

**Remarks**:
- `tmp.seven.env` will be deleted after the `after-script` script.
- `tmp.seven.env` is located in the execution directory.
> 🚧 this could change in the future.
*/