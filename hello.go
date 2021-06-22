package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `sdfkljsdalfjlasdjf
		asdlfjlasdjflajsdf
		alsdfjlsadjflasdjf
		alsdjflsadjflajdslfj
		asdlkfjlsadjflasjdlfj
		asldfjlasdjflsjdflj`)
	})
	http.ListenAndServe(":3000", nil)
}
