package main

import (
	"fmt"
)

func ValidateResponse(restyErr error, UnMarshallErr error) error {
	if restyErr != nil {
		fmt.Printf("\n UnMarsRestyhallErr error: %v \n", restyErr)
		return restyErr
	}

	if UnMarshallErr != nil {
		fmt.Printf("\n UnMarshallErr error: %v \n", UnMarshallErr)
		return UnMarshallErr
	}

	return nil

}
