package countdown

import (
	"fmt"
	"io"
)

func CountDown(out io.Writer) {
  fmt.Fprintf(out, "3")
}
