package smuggle

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/fatih/color"
)

// Compiles the resulting WASM file with the regular go compiler, or TinyGo if specified.
func compileWasm(smugglefilename, outdir string, tinygo bool) error {

	var Bin string

	if tinygo {
		Bin = "tinygo"
		fmt.Println("Using tinygo for executable. Ensure tinygo is in your PATH and installed correctly.")
	} else {
		Bin = "go"
		fmt.Println("Using go for executable. Ensure go is in your PATH and installed.")

	}

	Env := []string{
		fmt.Sprintf("CC=%s", os.Getenv("CC")),
		fmt.Sprintf("CGO_ENABLED=%s", "0"),
		fmt.Sprintf("GOPRIVATE=%s", os.Getenv("GOPRIVATE")),
		fmt.Sprintf("PATH=%s:%s", path.Join(os.Getenv("GOVERSION"), "bin"), os.Getenv("PATH")),
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
	}

	Env = append(Env, fmt.Sprintf("GOARCH=%s", "wasm"))
	Env = append(Env, fmt.Sprintf("GOOS=%s", "js"))

	var command []string
	var err error

	if err != nil {
		log.Fatalf("Error changing to output dir %v.\n", err)
	}

	goBinPath := Bin

	if err != nil {
		log.Fatalf("Error getting working dir %v.\n", err)
	}

	fmt.Printf("Getting go env and saving to %s\n", outdir+"/goenv.txt")
	os.Remove("go.mod")
	os.Remove("go.sum")
	if err != nil {
		log.Printf("Tried to clean any previous go.mod files but failed:\n%v.\n", err)
	}

	//go mod init the directory

	var out bytes.Buffer
	var stderr bytes.Buffer

	initCmd := exec.Command("go", "mod", "init", "securetransfer")
	initCmd.Env = Env

	initCmd.Stdout = &out
	initCmd.Stderr = &stderr
	initCmd.Dir = outdir
	err = initCmd.Run()
	if err != nil {
		color.Red("[smuggle/gen.go] Woops, something went wrong with go mod init, soz.\n")
		color.Red(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	initCmd.Wait()
	//go mod tidy to get depencies
	tidyCmd := exec.Command("go", "mod", "tidy")
	tidyCmd.Env = Env
	tidyCmd.Stdout = &out
	tidyCmd.Stderr = &stderr
	tidyCmd.Dir = outdir
	err = tidyCmd.Run()
	if err != nil {
		color.Red("[smuggle/gen.go] Woops, something went wrong with go mod tidy:\n")
		color.Red(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	tidyCmd.Wait()

	if err != nil {
		log.Fatalf("Error changing to output dir %v.\n", err)
	}

	if err != nil {
		log.Fatalf("Error getting working dir %v.\n", err)
	}

	outname := fmt.Sprintf("%s", strings.ReplaceAll(smugglefilename, ".go", ""))

	outname = outname + ".wasm"

	if tinygo {
		command = []string{"build", "-o", outname, "-target", "WASM", "--no-debug", smugglefilename}

	} else {
		command = []string{"build", "-o", outname, "-trimpath", `-ldflags=-w -s -H=windowsgui`, smugglefilename}
	}

	cmd := exec.Command(goBinPath, command...)
	cmd.Env = Env
	cmd.Dir = outdir
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		color.Red("[smuggle/gen.go] Woops, something went wrong with compiling, soz.\n")
		color.Red(fmt.Sprint(err) + ": " + stderr.String())
		return err
	}

	cmd.Wait()

	if tinygo {
		color.Green("Smuggler compiled with TinyGo, find it at %s/%s.\n", outdir, outname)
		return nil
	} else {
		color.Green("Smuggler compiled with regular go compiler, find it at %s/%s.\n", outdir, outname)

	}

	return nil

}
