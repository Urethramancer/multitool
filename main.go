package main

import (
	"os"
	"path"

	"github.com/Urethramancer/multitool/md5cmd"
	"github.com/Urethramancer/multitool/pwgen"
	"github.com/Urethramancer/multitool/rn"
	"github.com/Urethramancer/multitool/sha1cmd"
	"github.com/Urethramancer/multitool/sha512cmd"
	"github.com/Urethramancer/signor/log"
)

func main() {
	app := path.Base(os.Args[0])

	switch app {
	case "Multitool.exe":
		fallthrough
	case "multitool.exe":
		fallthrough
	case "Multitool":
		fallthrough
	case "multitool":
		usage()
	case "pwgen.exe":
		fallthrough
	case "pwgen":
		pwgen.Run()
	case "sha1.exe":
		fallthrough
	case "sha1":
		sha1cmd.Run()
	case "sha3.exe":
		fallthrough
	case "sha512":
		sha512cmd.Run()
	case "md5.exe":
		fallthrough
	case "md5":
		md5cmd.Run()
	case "rn.exe":
		fallthrough
	case "rn":
		rn.Run()
	}
}

func usage() {
	m := log.Default.Msg
	m("Multitool\n")
	m("Do not call directly. Link program names to this binary.\n\n")
	m("Admin utilities:\n")
	m("\tpwgen\tPassword generator\n")
	m("\nChecksumming:\n")
	m("\tmd5\tMD5 checksummer\n")
	m("\tsha1\tSHA-1 checksummer\n")
	m("\tsha3\tSHA-3 Keccak 512 checksummer\n")
	m("\nFile utilities:\n")
	m("\trn\tBulk renamer/name stripper\n")
}
