package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s : %v\n", url, err)
		// 	os.Exit(1)
		// }

		bytes, err_copy := io.Copy(os.Stdout, resp.Body)
		if err_copy != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy errors :  %v\n", err)
			os.Exit(1)
		}

		resp.Body.Close()

		fmt.Printf("Bytes copied : %v", bytes)
	}
}
