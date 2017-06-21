package depdemoraven

import (
	"fmt"

	raven "github.com/getsentry/raven-go"
)

func RavenInfo() {
	fmt.Println(raven.INFO)
}
