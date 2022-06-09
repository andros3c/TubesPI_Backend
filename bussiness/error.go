package bussiness

import "errors"

var (
	ErrInternalServer = errors.New("something's gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrThreadsIDResource = errors.New("(ThreadsID) not found or empty")

	ErrThreadsTitleResource = errors.New("(ThreadsTitle) not found or empty")

	ErrRolesIDResource = errors.New("(RolesID) not found or empty")

	ErrIDResource = errors.New("(ID) not found or empty")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrDuplicateEmail = errors.New("duplicate email, try to use other email")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")

	ErrWrongFormat = errors.New("Wrong Email Format")

	ErrWrongPass = errors.New("Wrong Password")

	ErrEmailEmpty = errors.New("Email Empty")

	ErrPassEmpty = errors.New("Password Empty")

	ErrNameEmpty = errors.New("Name Empty")

	ErrAccNotFound = errors.New("Account Not Found")

)