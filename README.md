# Coyote
User and IP Traffic Throttling for Go APIs and Services

## Introduction
Coyote enables dynamic per-user and per-ip address throttling for Go services. Here's a quick example API to show how easy it is to enable traffic management with coyote:

```go

import(
	"fmt"
	"net/http"
	"github.com/tjw0051/coyote"
)

func main() {

	coyote.Connect("localhost:6379")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err := coyote.IPEvent(r.RemoteAddr)
		if err != nil {
			fmt.Fprintf(w, "You've sent too many requests recently!")
		} else {
			fmt.Fprintf(w, "Welcome!")
		}

	})
	http.ListenAndServe(":8080", nil)
}
```