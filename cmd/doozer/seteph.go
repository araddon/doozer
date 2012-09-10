package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	cmds["seteph"] = cmd{seteph, "<path> <rev>", "write a file ephemerally"}
	cmdHelp["seteph"] = `Sets the body of the file at <path>.

The body is read from stdin. If <rev> is not greater than or equal to
the revision of the file, no change will be made.

Prints the new revision on stdout, or an error message on stderr.

Once the client that created this path/file goes away this also goes away
and anyclients with watches on it will be notified.
`
}

func seteph(path, rev string) {
	oldRev := mustAtoi64(rev)

	c := dial()

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		bail(err)
	}

	newRev, err := c.SetEph(path, oldRev, body)
	if err != nil {
		bail(err)
	}

	fmt.Println(newRev)
}
