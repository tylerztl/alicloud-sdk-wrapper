package helpers

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
)

func ParseError(err interface{}) error {
	switch err.(type) {
	case *errors.ClientError:
		originError := err.(*errors.ClientError).OriginError()
		if nil == originError {
			return err.(*errors.ClientError)
		} else {
			return originError
		}
	case *errors.ServerError:
		return err.(*errors.ServerError)
	case error:
		return err.(error)
	default:
	}

	return nil
}
