package commonUtils

import (
	"log"
	"testing"
)

func TestTimeStringToUnix(t *testing.T) {
	unixTS := TimeStringToUnix("2018-12-17T09:14:00Z")
	log.Println("unix timestamp:", unixTS)
}
