package utils

import "fmt"

// RemoveDuplicateElement slice去重
func RemoveDuplicateElement(ori any) (any, error) {
	temp := map[any]struct{}{}

	switch sType := ori.(type) {
	case []string:
		result := make([]string, 0, len(ori.([]string)))

		for _, item := range sType {
			if _, ok := temp[item]; !ok {
				temp[item] = struct{}{}
				result = append(result, item)
			}
		}

		return result, nil
	case []int64:
		result := make([]int64, 0, len(ori.([]int64)))

		for _, item := range sType {
			if _, ok := temp[item]; !ok {
				temp[item] = struct{}{}
				result = append(result, item)
			}
		}
		return result, nil
	default:
		err := fmt.Errorf("unknown type: %T", sType)
		return nil, err
	}
}
