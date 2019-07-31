package pwgen

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

var Options struct {
	opt.DefaultHelp
	Complex bool `short:"c" long:"complex" help:"Make the password really complex. May not be accepted everywhere."`
	Length  int  `short:"l" long:"length" help:"Password length." placeholder:"N" default:"12"`
}

func genString(size int, complex bool) string {
	valid := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")
	if complex {
		valid = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!'#$%&/()=?@*^<>-_.:,;|[]{}")
	}
	pw := make([]byte, size)
	for i := 0; i < size; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(valid))))
		c := valid[n.Int64()]
		pw[i] = c
	}
	return string(pw)
}

// Run the generator.
func Run() {
	a := opt.Parse(&Options)
	if Options.Help {
		a.Usage()
		return
	}

	m := log.Default.Msg
	for i := 0; i < 20; i++ {
		for j := 0; j < 7; j++ {
			fmt.Printf(" %s", genString(Options.Length, Options.Complex))
		}
		m("")
	}
}
