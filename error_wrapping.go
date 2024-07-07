package main

import (
	"errors"
	"fmt"
)

// 정의된 기본 오류
var ErrNotFound = errors.New("not found")

// 오류를 래핑하는 함수
func findItem(id int) error {
	if id != 1 {
		// ErrNotFound 오류를 래핑하여 반환
		return fmt.Errorf("findItem: %w", ErrNotFound)
	}
	return nil
}

func main() {
	err := findItem(2)
	if err != nil {
		// errors.Is를 사용하여 ErrNotFound 오류인지 확인
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Item not found")
		} else {
			fmt.Println("An unexpected error occurred:", err)
		}
	} else {
		fmt.Println("Item found")
	}
}
