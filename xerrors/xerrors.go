package xerrors

import "fmt"

func PanicIfError(err error) {
	if err != nil {
		message := fmt.Sprintf("%+v\n", err)
		panic(message)
	}
}
