package keyGenerator

import (
	C "crypto/rand"
	"fmt"
	M "math/rand"
	"strings"
	"time"
)

/*
	upper :101 - 132
	lower: 141 - 172

	rand: (max-min)+min
*/
func KeyGen(ranges []uint) (newID string) {
	var keys []string
	var randomInt int

	M.Seed(time.Now().UnixNano())
	for _, value := range ranges {
		RandomCrypto, _ := C.Prime(C.Reader, 20)
		if value%2 == 0 {
			randomInt = M.Intn(90-65) + 65
			keys = append(keys, fmt.Sprintf("%s", string(randomInt)))
			keys = append(keys, string(RandomCrypto.String()))
		} else {
			randomInt = M.Intn(122-97) + 97
			keys = append(keys, fmt.Sprintf("%s", string(randomInt)))
			keys = append(keys, string(RandomCrypto.String()))
		}
	}

	newID = strings.Join(keys, "")
	return
}
