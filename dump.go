package toolz

import (
	"encoding/json"
	"fmt"
	"os"
)

// Dump pretty prints a value to stdout using json encoding - nice for debugging, but
// not recommended for production use.
func Dump(t string, v interface{}) {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	if t == "" {
		fmt.Println(string(j))
		return
	}
	fmt.Printf("%s: %s\n", t, string(j))
}

// DumpToFile pretty prints a value to a file using json encoding - nice for debugging, but
// not recommended for production use.
func DumpToFile(filename string, v interface{}) error {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(string(j))
	return err
}
