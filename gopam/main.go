package main

// #include <security/pam_ext.h>
// #include <security/pam_modules.h>
// #cgo CFLAGS: -Wall -fPIC
// #cgo LDFLAGS: -lpam
import "C"
import (
	"log"
)

const (
	PAMSuccess          C.int = C.PAM_SUCCESS
	PAMAuthErr          C.int = C.PAM_AUTH_ERR
	PAMAuthToken        C.int = C.PAM_AUTHTOK
	PAMAuthTokErr       C.int = C.PAM_AUTHTOK_ERR
	PAMPermissionDenied C.int = C.PAM_PERM_DENIED
	PAMUserUnknown      C.int = C.PAM_USER_UNKNOWN
	PAMServiceError     C.int = C.PAM_SERVICE_ERR
)

//PAMAuthenticate is the go implementation of pam_sm_authenticate
//export PAMAuthenticate
func PAMAuthenticate(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	var pamusername *C.char
	if err := C.pam_get_user(pamh, &pamusername, nil); err != PAMSuccess {
		log.Println("ERROR: unable to get username:", C.pam_strerror(pamh, err))
		return PAMServiceError
	}
	username := C.GoString(pamusername)

	var pamauthenticationtoken *C.char
	if err := C.pam_get_authtok(pamh, PAMAuthToken, &pamauthenticationtoken, nil); err != PAMSuccess {
		log.Println("ERROR: unable to get the authentication token:", C.pam_strerror(pamh, err))
		return PAMServiceError
	}
	password := C.GoString(pamauthenticationtoken)

	log.Println("Authenticating", username, "with authentication token", password)
	return PAMSuccess
}

//PAMSetCred is the go implementation of pam_sm_setcred
//export PAMSetCred
func PAMSetCred(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	return PAMSuccess
}

//PAMAccountManagement is the go implementation of pam_sm_acct_mgmt
//export PAMAccountManagement
func PAMAccountManagement(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	return PAMSuccess
}

//PAMOpenSession is the go implementation of pam_sm_open_session
//export PAMOpenSession
func PAMOpenSession(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	return PAMSuccess
}

//PAMCloseSession is the go implementation of pam_sm_close_session
//export PAMCloseSession
func PAMCloseSession(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	return PAMSuccess
}

//PAMChangeAuthToken is the go implementation of pam_sm_chauthtok
//export PAMChangeAuthToken
func PAMChangeAuthToken(pamh *C.pam_handle_t, flags C.int, argc C.int, argv **C.char) C.int {
	return PAMSuccess
}

func main() {}
