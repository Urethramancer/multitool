package rn

import (
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/mgutz/str"
)

var Options struct {
	opt.DefaultHelp
	Count             bool     `short:"c" long:"count" help:"Put a count followed by '.' before each filename. This won't be added if the name already starts with it."`
	Prepend           string   `short:"p" long:"prepend" help:"String to insert before each filename. This won't be added if the name already starts with it."`
	Remove            string   `short:"R" long:"remove" help:"String to remove from each filename. Removes first instance."`
	Replace           string   `short:"r" long:"replace" help:"String to replace removed section with. 'remove' must also be specified."`
	All               bool     `short:"a" long:"all" help:"Replace/remove all instances of specified string. Stop after the first one otherwise."`
	SpaceToUnderscore bool     `short:"u" long:"tounderscore" help:"Convert all spaces in filenames to underscores."`
	UnderscoreToSpace bool     `short:"s" long:"tospace" help:"Convert all underscores in filenames to spaces."`
	Camel             bool     `short:"C" long:"camel" help:"Remove all spaces and capitalise each part."`
	Quiet             bool     `short:"q" long:"quiet" help:"Quiet (don't print progress)."`
	Files             []string `placeholder:"FILE" help:"File to rename."`
}

// Run the rename command.
func Run() {
	a := opt.Parse(&Options)
	if Options.Help || len(Options.Files) == 0 {
		a.Usage()
		return
	}

	for c, oldname := range Options.Files {
		m := log.Default.Msg
		if _, err := os.Stat(oldname); os.IsNotExist(err) {
			m("No such file or directory: %s", oldname)
			continue
		}
		name := oldname
		if Options.Count {
			name = str.EnsurePrefix(name, fmt.Sprintf("%d.", c+1))
		}
		if Options.Prepend != "" {
			name = str.EnsurePrefix(name, Options.Prepend)
		}
		if Options.Remove != "" {
			n := 1
			if Options.All {
				n = -1
			}
			if Options.Replace != "" {
				name = strings.Replace(name, Options.Remove, Options.Replace, n)
			} else {
				name = strings.Replace(name, Options.Remove, "", n)
			}
		}
		if Options.SpaceToUnderscore {
			name = strings.Replace(name, " ", "_", -1)
		}
		if Options.UnderscoreToSpace {
			name = strings.Replace(name, "_", " ", -1)
		}
		if Options.Camel {
			names := strings.Split(name, " ")
			name = ""
			for _, e := range names {
				c, n := utf8.DecodeRuneInString(e)
				c = unicode.ToUpper(c)
				e = string(c) + e[n:]
				name = fmt.Sprintf("%s%s", name, e)
			}
		}
		if oldname != name {
			err := os.Rename(oldname, name)
			if err != nil {
				log.Default.Err("Error:%s could not be renamed to %s.\n", oldname, name)
				return
			}
			if !Options.Quiet {
				m("%s renamed to %s.\n", oldname, name)
			}
		} else {
			if !Options.Quiet {
				m("%s was not renamed because there were no changes.\n", oldname)
			}
		}
	}
}
