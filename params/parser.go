package params

import (
	"strings"

	"github.com/be-true/config-go/errors"
	"github.com/be-true/config-go/utils"
)

func parseParams(tag string) (*Params, error) {
	paramsMap, err := getParamsMap(tag)
	if err != nil {
		return nil, err
	}
	required, err := utils.ParseBoolean(paramsMap["required"], true)
	if err != nil {
		return nil, err
	}
	format := strings.ToLower(paramsMap["format"])
	doc := paramsMap["doc"]
	return &Params{
		Required: required,
		Format:   format,
		Doc:      doc,
	}, nil
}

func getParamsMap(tag string) (map[string]string, error) {
	result := make(map[string]string)
	params := strings.Split(tag, ",")
	for _, param := range params {
		split := strings.Split(param, "=")
		if len(split) == 0 {
			return nil, errors.ErrorEmptyParams
		}
		value := ""
		if len(split) > 1 {
			value = split[1]
		}
		name := split[0]
		result[name] = utils.TrimEscape(value)
	}
	return result, nil
}
