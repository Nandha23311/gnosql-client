package gnosql_client

import (
	"errors"
	"fmt"
)

func ValidateResponse(UnMarshallErr error, gRPCErr error) error {
	if UnMarshallErr != nil {
		fmt.Printf("\n UnMarshallErr error: %v \n", UnMarshallErr)
		return errors.New(UnMarshallErr.Error())
	}

	if gRPCErr != nil {
		fmt.Printf("\n gRestErr error: %v \n", gRPCErr)
		return errors.New(gRPCErr.Error())
	}

	return nil

}
