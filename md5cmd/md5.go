package md5cmd

import (
	"crypto/md5"
	"io"
	"os"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

var Options struct {
	opt.DefaultHelp
	Files []string `placeholder:"FILE" help:"Full path to file to checksum."`
}

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
	a := opt.Parse(&Options)
	if Options.Help || len(Options.Files) == 0 {
		a.Usage()
		return
	}

	for _, fn := range Options.Files {
		md5sum(fn)
	}
}
