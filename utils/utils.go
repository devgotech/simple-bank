package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random interger between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomOwner generates a random  owner name
func RandomOwner() string {
	return RandomString(6)
}

//RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
//RandonCurrency gereates a random currency name
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "GBP", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomAccountID generates a random account ID (assuming you've created enough accounts)
func RandomAccountID() int64 {
	return RandomInt(1, 30)  // You can adjust the range as per your database setup
}