package sha512cmd

import (
	"fmt"
	"io"
	"os"

	"crypto/sha512"

	"github.com/Urethramancer/signor/log"
	"github.com/grimdork/climate/arg"
	ll "github.com/grimdork/loglines"
)

func sha512sum(file string) {
	hash := sha512.New()
	hash.Reset()
	in, err := os.Open(file)
	if err != nil {
		log.Default.Err("sha3: couldn't open %s.", file)
		os.Exit(2)
	}

	_, err = io.Copy(hash, in)
	if err != nil {
		ll.Err("sha512: couldn't copy %s: %s", file, err.Error())
		os.Exit(2)
	}

	log.Default.Msg("%x  %s", hash.Sum(nil), file)
}

// Run the checksummer.
func Run() {
	opt := arg.New("sha512")
	opt.SetDefaultHelp(true)
	opt.SetPositional("FILE", "Filename to process.", "", true, arg.VarStringSlice)
	err := opt.Parse(os.Args[1:])
	if err != nil {
		if err == arg.ErrNoArgs {
			opt.PrintHelp()
			return
		}

		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(2)
	}

	args := opt.GetPosStringSlice("FILE")
	for _, fn := range args {
		sha512sum(fn)
	}
}
