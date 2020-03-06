package commonUtils

import (
	"log"
	"testing"
)

func _TestTimeStringToUnix(t *testing.T) {
	unixTS, err := TimeStringToUnix("2018-12-17T09:14:00Z")
	if err != nil {
		return
	}
	log.Println("unix timestamp:", unixTS)
}
