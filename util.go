package gnosql_client

import (
	"fmt"
)

func ValidateResponse(restyErr error, UnMarshallErr error, gRPCErr error, gRestResponseErr string) string {
	if restyErr != nil {
		fmt.Printf("\n UnMarsRestyhallErr error: %v \n", restyErr)
		return restyErr.Error()
	}

	if UnMarshallErr != nil {
		fmt.Printf("\n UnMarshallErr error: %v \n", UnMarshallErr)
		return UnMarshallErr.Error()
	}

	if gRPCErr != nil {
		fmt.Printf("\n gRestErr error: %v \n", UnMarshallErr)
		return gRPCErr.Error()
	}

	if gRestResponseErr != "" {
		// fmt.Printf("\n gRestResponseErr error: %v \n", gRestResponseErr)
		return gRestResponseErr
	}

	return ""

}
