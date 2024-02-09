package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/Urethramancer/multitool/md5cmd"
	pwgen "github.com/Urethramancer/multitool/pwgencmd"
	"github.com/Urethramancer/multitool/rn"
	"github.com/Urethramancer/multitool/sha1cmd"
	"github.com/Urethramancer/multitool/sha512cmd"
	"github.com/Urethramancer/signor/files"
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
	case "sha512.exe":
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
	e := log.Default.Err

	bin, err := os.Executable()
	if err != nil {
		m("Error getting path: %s", err.Error())
		return
	}

	ret := false
	path := filepath.Dir(bin)
	md5path := filepath.Join(path, "md5")
	if !files.FileExists(md5path) {
		m("Creating symlink '%s'", md5path)
		err = os.Symlink(bin, md5path)
		if err != nil {
			e("Error creating symlink for %s: %s", md5path, err.Error())
			os.Exit(2)
		}
		ret = true
	}

	sha1path := filepath.Join(path, "sha1")
	if !files.FileExists(sha1path) {
		m("Creating symlink '%s'", sha1path)
		err = os.Symlink(bin, sha1path)
		if err != nil {
			e("Error creating symlink for %s: %s", sha1path, err.Error())
			os.Exit(2)
		}
		ret = true
	}

	sha512path := filepath.Join(path, "sha512")
	if !files.FileExists(sha512path) {
		m("Creating symlink '%s'", sha512path)
		err = os.Symlink(bin, sha512path)
		if err != nil {
			e("Error creating symlink for %s: %s", sha512path, err.Error())
			os.Exit(2)
		}
		ret = true
	}

	pwgenpath := filepath.Join(path, "pwgen")
	if !files.FileExists(pwgenpath) {
		m("Creating symlink '%s'", pwgenpath)
		err = os.Symlink(bin, pwgenpath)
		if err != nil {
			e("Error creating symlink for %s: %s", pwgenpath, err.Error())
			os.Exit(2)
		}
		ret = true
	}

	rnpath := filepath.Join(path, "rn")
	if !files.FileExists(rnpath) {
		m("Creating symlink '%s'", rnpath)
		err = os.Symlink(bin, rnpath)
		if err != nil {
			e("Error creating symlink for %s: %s", rnpath, err.Error())
			os.Exit(2)
		}
		ret = true
	}

	if ret {
		return
	}

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
