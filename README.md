# The Silk Wasm

Obfuscate your HTML Smuggling with web assembly!

## Prerequisites

To use SilkWASM, you must have [golang](https://go.dev/doc/install) installed as this is used to compile the WASM files. 

In addition, for smaller payloads, you can use [tinygo](https://tinygo.org/getting-started/)

## Installing SilkWasm

First, build the tool:

```sh
go build -o silkwasm silkwasm.go
```

Now ensure you move the binary to your path, e.g. ~/go/bin/

**Run silkwasm in a folder that does not contain go-code**
Running silkwasm in the same folder as another go project will not work, and will confuse any go.mod/go.sum files.

Please note that 

## Using SilkWasm
Once you have a payload to smuggle, executing silkwasm is simple:

```sh
silkwasm gen -i <smuggling payload.txt> -f <wasmfuncname>
```

This will create several files, the two key ones are your .wasm file, and the html example file. 

Place these in your webroot, along with the correct wasm_exec.js for your build of go. It's usually here: `$(go env GOROOT)/misc/wasm/wasm_exec.js)`

Please note, tinygo uses a different wasm_exec.js, usually found here: `$(tinygo env TINYGOROOT)/targets/wasm_exec.js`

And that is it! You'll need to modify your html to suit your pretext, the example html contains only the bare minimum JavaScript to run the html smuggle.
