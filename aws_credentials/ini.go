package ini

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/aws/smithy-go"
)

// OpenFile takes a path to a given file, and will open  and parse
// that file.
func OpenFile(path string) (Sections, error) {
	f, err := os.Open(path)
	if err != nil {
		var oe *smithy.OperationError
		if errors.As(err, &oe) {
			fmt.Printf("Unable to Read file: %s, operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap())
		}
		return Sections{}, err
	}
	defer f.Close()

	return Parse(f)
}

// Parse will parse the given file using the shared config
// visitor.
func Parse(f io.Reader) (Sections, error) {
	tree, err := ParseAST(f)
	if err != nil {
		return Sections{}, err
	}

	v := NewDefaultVisitor()
	if err = Walk(tree, v); err != nil {
		return Sections{}, err
	}

	return v.Sections, nil
}

// ParseBytes will parse the given bytes and return the parsed sections.
func ParseBytes(b []byte) (Sections, error) {
	tree, err := ParseASTBytes(b)
	if err != nil {
		return Sections{}, err
	}

	v := NewDefaultVisitor()
	if err = Walk(tree, v); err != nil {
		return Sections{}, err
	}

	return v.Sections, nil
}
