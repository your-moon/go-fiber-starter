package ierror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
)

var (
	CIError = errors.New("err test")
)

func NewIError(ierror error, message ...error) error {
	var result error
	result = multierror.Append(result, ierror)
	result = multierror.Append(result, message...)

	merr, ok := result.(*multierror.Error)
	if ok {
		merr.ErrorFormat = func(errs []error) string {
			return ListInLine(errs)
		}
	}
	return merr
}

func ListInLine(es []error) string {
	var errString string
	for i, e := range es {
		if i == 0 {
			errString += "Err: "
		}
		if e != nil {
			errString += e.Error()
			if i < len(es)-1 {
				errString += "; Msg: "
			}
		}
	}
	return errString
}

func ListNewLine(es []error) string {
	if len(es) == 1 {
		return fmt.Sprintf("1 error occurred:\n\t* %s\n\n", es[0])
	}

	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf(
		"%d errors occurred:\n\t%s\n\n",
		len(es), strings.Join(points, "\n\t"))

}
