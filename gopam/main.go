package main

// #include <security/pam_ext.h>
// #cgo CFLAGS: -Wall -fPIC
// #cgo LDFLAGS: -lpam
import "C"

const (
	PAMSuccess C.int = C.PAM_SUCCESS
	PAMAuthErr C.int = C.PAM_AUTH_ERR
	PAMAuthOK  C.int = C.PAM_AUTHTOK
)

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags C.int, argc C.int, argv unsafe.Pointer) C.int {

	return PAMSuccess
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags C.int, argc C.int, argv unsafe.Pointer) C.int {
	return PAMSuccess
}

//export pam_sm_acct_mgmt
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags C.int, argc C.int, argv unsafe.Pointer) C.int {
	return PAMSuccess
}

//expose pam_sm_open_session
func pam_sm_open_session(pamh *C.pam_handle_t, flags C.int, argc C.int, argv unsafe.Pointer) C.int {
	return PAMSuccess
}

//expose pam_sm_close_session
func pam_sm_close_session(pamh *C.pam_handle_t, flags C.int, argc C.int, argv unsafe.Pointer) C.int {
	return PAMSuccess
}

func main() {}
