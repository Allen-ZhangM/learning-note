package demorpc

import "github.com/pkg/errors"

type CalculateService struct {
}

type Args struct {
	A, B int
}

func (CalculateService) Div(args Args, result *float64) error {

	if args.B == 0 {
		return errors.New("除数为0")
	}

	*result = float64(args.A) / float64(args.B)

	return nil

}
