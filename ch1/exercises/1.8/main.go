package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		newUrl := url
		if !strings.HasPrefix(url, "http://") {
			newUrl = "http://" + url
		}
		resp, err := http.Get(newUrl)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", newUrl, err)
			os.Exit(1)
		}

	}
}
