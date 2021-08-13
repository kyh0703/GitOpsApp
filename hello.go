package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `sdfkljsdalfjlasdjf
		asdlfjlasdjflajsdf
		asdfasdf;asdkf;sakasdfsadfasdfasdfdf;k;ds;fak;asdkf;ksadf;kasd;fksad;fkdkdkdkd
		alsdjflsadjflajdslfj
		asdlkfjlsadjflasjdlfj
		asldfjlasdjflsjdflj`)
	})
	http.ListenAndServe(":3000", nil)
}
