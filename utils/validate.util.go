package utils

import "unicode/utf8"

func ValidateRequestValue(params string) bool {
	isCategoryNull := utf8.ValidString(params)

	// 값이 NULL인지 체크
	if !isCategoryNull {
		return false
	}

	// 요청 파라미터 값이 빈 문자열이 아닌지 체크
	if len(params) > 1 {
		return true
	}

	return false
}