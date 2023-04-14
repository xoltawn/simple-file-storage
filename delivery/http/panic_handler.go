package http

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

type PanicHandler struct {
	Next http.Handler
}

func (h PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 10<<10)
			n := runtime.Stack(buf, false)
			fmt.Fprintf(os.Stderr, "panic: %v\n\n%s", err, buf[:n])

			// Uncomment to exit instead of recovering.
			// os.Exit(1)

			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}
	}()

	h.Next.ServeHTTP(w, r)
}
