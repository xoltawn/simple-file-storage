package http

import (
	"fmt"
	"log"
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

			log.Println(err)

			if newErr, ok := err.(error); ok {
				if newErr.Error() == "multipart: NextPart: EOF" || newErr.Error() == "no such file" {
					http.Error(w, "required fields are not provided", http.StatusBadRequest)
				}
				if newErr.Error() == "unsupported content type" {
					w.WriteHeader(http.StatusUnsupportedMediaType)
					http.Error(w, "unaccepted content type", http.StatusUnsupportedMediaType)

				}
			} else {
				http.Error(w, "Internal error", http.StatusInternalServerError)
			}

		}
	}()

	h.Next.ServeHTTP(w, r)
}
