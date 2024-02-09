package pwgen

import (
	"fmt"

	"github.com/grimdork/climate/arg"
)

// Run the generator.
func Run() {
	opt := arg.New("pwgen")
	opt.SetDefaultHelp(true)
	opt.SetOption("", "l", "length", "Length of passwords.", 16, false, arg.VarInt, nil)
	opt.SetOption("", "c", "count", "Number of passwords to generate.", 140, false, arg.VarInt, nil)
	opt.SetOption("", "n", "nonce", "Generate a binary string more suitable for cryptography. Sets count to 1. Should be piped to file.", false, false, arg.VarBool, nil)
	opt.SetOption("", "w", "words", "Generate word-based passwords.", false, false, arg.VarBool, nil)
	opt.SetOption("", "W", "wordcount", "Number of words in word-based passwords.", 6, false, arg.VarInt, nil)
	opt.HelpOrFail()

	if opt.GetBool("nonce") && !opt.GetBool("words") {
		fmt.Printf("%s", RandNonce(opt.GetInt("length")))
		return
	}

	if opt.GetBool("words") {
		wc := opt.GetInt("wordcount")
		for i := 0; i < opt.GetInt("count"); i++ {
			fmt.Printf("\t%s\n", RandWords(wc))
		}
		return
	}

	w := 120
	length := opt.GetInt("length")
	maxcount := opt.GetInt("count")
	maxw := w / (length + 2)
	count := 0
	for i := 0; i < maxcount; i++ {
		fmt.Printf("%s  ", RandString(length))
		count++
		if count == maxw {
			println("")
			count = 0
		}
	}

	// Breather
	println("")

}
