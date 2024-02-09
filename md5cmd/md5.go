package md5cmd

import (
	"crypto/md5"
	"io"
	"os"

	"github.com/Urethramancer/signor/log"
	"github.com/grimdork/climate/arg"
)

func md5sum(file string) {
	hash := md5.New()
	hash.Reset()
	in, err := os.Open(file)
	if err != nil {
		log.Default.Err("md5: couldn't open %s.", file)
		os.Exit(2)
	}

	_, err = io.Copy(hash, in)
	if err != nil {
		log.Default.Err("md5: couldn't copy %s: %s", file, err.Error())
		os.Exit(2)
	}

	log.Default.Msg("%x  %s", hash.Sum(nil), file)
}

// Run the checksummer.
func Run() {
	opt := arg.New("md5")
	opt.SetDefaultHelp(true)
	opt.SetPositional("FILE", "Filename to process.", "", true, arg.VarStringSlice)
	opt.HelpOrFail()
	args := opt.GetPosStringSlice("FILE")
	for _, fn := range args {
		md5sum(fn)
	}
}
