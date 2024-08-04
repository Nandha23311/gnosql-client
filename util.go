package gnosql_client

import (
	"errors"
	"fmt"
)

func ValidateResponse(restyErr error, UnMarshallErr error, gRPCErr error, gRestResponseErr string) error {
	if restyErr != nil {
		fmt.Printf("\n UnMarsRestyhallErr error: %v \n", restyErr)
		return errors.New(restyErr.Error())
	}

	if UnMarshallErr != nil {
		fmt.Printf("\n UnMarshallErr error: %v \n", UnMarshallErr)
		return errors.New(UnMarshallErr.Error())
	}

	if gRPCErr != nil {
		fmt.Printf("\n gRestErr error: %v \n", UnMarshallErr)
		return errors.New(gRPCErr.Error())
	}

	if gRestResponseErr != "" {
		return errors.New(gRestResponseErr)
	}

	return nil

}
