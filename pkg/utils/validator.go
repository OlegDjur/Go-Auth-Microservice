package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/OlegDjur/Go-Auth-Microservice/internal/user/dto"
)

// const alphaNumeric = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
const onlyAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ValidateUser(user *dto.CreateUserRequest) error {
	err := CheckNull(user.Email)
	if err != nil {
		return err
	}
	err = CheckNull(user.Password)
	if err != nil {
		return err
	}
	err = CheckNull(user.Username)
	if err != nil {
		return err
	}

	err = CheckEmail(user.Email)
	if err != nil {
		return err
	}

	err = CheckLen(user.Password, 6, 20)
	if err != nil {
		return err
	}
}

func CheckEmail(email string) error {
	emailRegex := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("email: %v is not correct", email)
	}

	return nil
}

func CheckLen(arg string, lenMin, lenMax int) error {
	if len(arg) < lenMin {
		return fmt.Errorf("%v should contain at lest %v characters", arg, lenMin)
	}

	if len(arg) > lenMax {
		return fmt.Errorf("%v should contain more than %v characters", arg, lenMin)
	}

	return nil
}

func OnlyAlphabet(arg string) error {
	for _, v := range arg {
		if !strings.ContainsAny(string(v), onlyAlphabet) {
			return fmt.Errorf("%s should contain only alphabet characters", arg)
		}
	}

	return nil
}

func CheckNull(arg string) error {
	if len(arg) == 0 {
		return fmt.Errorf("field should not be empty")
	}

	return nil
}
