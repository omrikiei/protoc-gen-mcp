package generator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// PathParam represents a path parameter
type PathParam struct {
	Name  string
	Value string
}

// ParsePathParams parses path parameters from the request URL
func ParsePathParams(pathTemplate, requestPath string) (map[string]string, error) {
	// Convert path template to regex pattern
	pattern := regexp.QuoteMeta(pathTemplate)
	pattern = strings.ReplaceAll(pattern, `\{`, `(?P<`)
	pattern = strings.ReplaceAll(pattern, `\}`, `>[^/]+)`)
	pattern = "^" + pattern + "$"

	// Compile regex
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("failed to compile path pattern: %w", err)
	}

	// Match request path against pattern
	matches := re.FindStringSubmatch(requestPath)
	if matches == nil {
		return nil, fmt.Errorf("path does not match template")
	}

	// Extract parameter values
	params := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && name != "" {
			params[name] = matches[i]
		}
	}

	return params, nil
}

// ApplyPathParams applies path parameters to a request message
func ApplyPathParams(req interface{}, params map[string]string) error {
	// Get the value of the request
	val := reflect.ValueOf(req)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("request must be a non-nil pointer")
	}

	// Get the type of the request
	typ := val.Type().Elem()

	// Iterate over the fields of the request
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Elem().Field(i)

		// Check if the field has a path parameter tag
		if paramName, ok := field.Tag.Lookup("mcp_path_param"); ok {
			// Use the tag value if provided, otherwise use the field name
			if paramName == "" {
				paramName = field.Name
			}

			// Get the parameter value
			paramValue, ok := params[paramName]
			if !ok {
				return fmt.Errorf("missing path parameter: %s", paramName)
			}

			// Set the field value
			switch fieldVal.Kind() {
			case reflect.String:
				fieldVal.SetString(paramValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				intVal, err := strconv.ParseInt(paramValue, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid integer value for parameter %s: %w", paramName, err)
				}
				fieldVal.SetInt(intVal)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				uintVal, err := strconv.ParseUint(paramValue, 10, 64)
				if err != nil {
					return fmt.Errorf("invalid unsigned integer value for parameter %s: %w", paramName, err)
				}
				fieldVal.SetUint(uintVal)
			case reflect.Float32, reflect.Float64:
				floatVal, err := strconv.ParseFloat(paramValue, 64)
				if err != nil {
					return fmt.Errorf("invalid float value for parameter %s: %w", paramName, err)
				}
				fieldVal.SetFloat(floatVal)
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(paramValue)
				if err != nil {
					return fmt.Errorf("invalid boolean value for parameter %s: %w", paramName, err)
				}
				fieldVal.SetBool(boolVal)
			default:
				return fmt.Errorf("unsupported field type for path parameter: %s", fieldVal.Kind())
			}
		}
	}

	return nil
}

// ApplyQueryParams applies query parameters to a request message
func ApplyQueryParams(req interface{}, queryParams map[string][]string) error {
	// Get the value of the request
	val := reflect.ValueOf(req)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return fmt.Errorf("request must be a non-nil pointer")
	}

	// Get the type of the request
	typ := val.Type().Elem()

	// Iterate over the fields of the request
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Elem().Field(i)

		// Check if the field has a query parameter tag
		if paramName, ok := field.Tag.Lookup("mcp_query_param"); ok {
			// Use the tag value if provided, otherwise use the field name
			if paramName == "" {
				paramName = field.Name
			}

			// Get the parameter values
			paramValues, ok := queryParams[paramName]
			if !ok {
				continue // Skip if parameter is not present
			}

			// Handle array/slice types
			if fieldVal.Kind() == reflect.Slice || fieldVal.Kind() == reflect.Array {
				sliceType := fieldVal.Type().Elem()
				slice := reflect.MakeSlice(fieldVal.Type(), len(paramValues), len(paramValues))
				for j, value := range paramValues {
					elem := reflect.New(sliceType).Elem()
					if err := setValue(elem, value); err != nil {
						return fmt.Errorf("failed to set array element for parameter %s: %w", paramName, err)
					}
					slice.Index(j).Set(elem)
				}
				fieldVal.Set(slice)
				continue
			}

			// Use the first value for non-array types
			if len(paramValues) > 0 {
				if err := setValue(fieldVal, paramValues[0]); err != nil {
					return fmt.Errorf("failed to set value for parameter %s: %w", paramName, err)
				}
			}
		}
	}

	return nil
}

// setValue sets a string value to a field of any supported type
func setValue(fieldVal reflect.Value, value string) error {
	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid integer value: %w", err)
		}
		fieldVal.SetInt(intVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintVal, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid unsigned integer value: %w", err)
		}
		fieldVal.SetUint(uintVal)
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Errorf("invalid float value: %w", err)
		}
		fieldVal.SetFloat(floatVal)
	case reflect.Bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Errorf("invalid boolean value: %w", err)
		}
		fieldVal.SetBool(boolVal)
	default:
		return fmt.Errorf("unsupported field type: %s", fieldVal.Kind())
	}
	return nil
}
