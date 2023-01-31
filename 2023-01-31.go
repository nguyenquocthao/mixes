package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ImageText(url string) (text string, err error) {
	var getByteContent = func(url string) []byte {
		i := strings.Index(url, ",")
		dec, err := base64.StdEncoding.DecodeString(url[i+1:])
		if err != nil {
			panic(err)
		}
		return dec
	}
	err = nil
	if strings.HasPrefix(url, "data:image") {
		subProcess := exec.Command("tesseract", "stdin", "stdout", "-l", "vie")
		var stdin io.WriteCloser
		stdin, err = subProcess.StdinPipe()
		if err != nil {
			return
		}
		defer stdin.Close()
		var stdout bytes.Buffer
		subProcess.Stdout = &stdout
		subProcess.Stderr = os.Stderr
		if err = subProcess.Start(); err != nil { //Use start, not run
			return
		}
		stdin.Write(getByteContent(url))
		stdin.Close()
		stdin.Close()
		subProcess.Wait()
		text = strings.TrimSpace(stdout.String())
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, "tesseract", url, "stdout", "-l", "vie").Output()
	text = strings.TrimSpace(string(output))
	return
}

func main() {
	a, b := ImageText("data:image/png;base64,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")
	fmt.Println(19, a, b)
	c, d := ImageText("https://shub-storage.sgp1.cdn.digitaloceanspaces.com/file/file_images/2smazwysQLSbP5DrN/1675154812866_99")
	fmt.Println(23, c, d)
}
