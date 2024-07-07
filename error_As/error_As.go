// errors.As 함수는 Go에서 오류 체인에서 특정 오류 타입을 찾아내는 데 사용된다.
// 오류가 래핑되어 있더라도 체인을 따라가며 해당 타입을 찾을 수 있다.

package main

import (
	"errors"
	"fmt"
)

// 사용자 정의 오류 타입
type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

// 오류를 래핑하는 함수
func findItem(id int) error {
	if id != 1 {
		// NotFoundError 오류를 래핑하여 반환
		return fmt.Errorf("findItem: %w", &NotFoundError{Message: "item not found"})
	}
	return nil
}

func main() {
	err := findItem(2)
	if err != nil {
		var notFoundErr *NotFoundError
		// errors.As를 사용하여 NotFoundError 타입인지 확인하고 변수에 저장
		if errors.As(err, &notFoundErr) {
			fmt.Println("Item not found using errors.As:", notFoundErr.Message)
		} else {
			fmt.Println("An unexpected error occurred:", err)
		}
	} else {
		fmt.Println("Item found")
	}
}
