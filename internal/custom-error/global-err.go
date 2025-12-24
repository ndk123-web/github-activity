package customerror

import (
	"fmt"
)

func GlobalError(err *error) {
	if *err != nil {
		fmt.Println("\nError: ", (*err).Error())
	}
}
