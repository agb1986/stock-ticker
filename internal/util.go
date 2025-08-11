package internal

import (
	"fmt"

	"github.com/yalp/jsonpath"
)

func handleError(err error, function string) {
	if err != nil {
		panic(fmt.Errorf("error with %s: %w", function, err))
	}
}

func getDataByJsonPath(path string, data any) any {
	preparedPath, err := jsonpath.Prepare(path)
	handleError(err, "getDataByJsonPath")
	parsedData, err := preparedPath(data)
	handleError(err, "getDataByJsonPath")
	return parsedData
}
