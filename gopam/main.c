
#include <security/pam_modules.h>
#include "_cgo_export.h"

PAM_EXTERN int pam_sm_authenticate(pam_handle_t *_pamh, int _flags, int _argc, const char **_argv){
   return PAMAuthenticate(_pamh,_flags,_argc,(char**)_argv);
}

PAM_EXTERN int pam_sm_setcred(pam_handle_t *_pamh, int _flags, int _argc, const char **_argv){
    return PAMSetCred(_pamh,_flags,_argc,(char**)_argv);
}

PAM_EXTERN int pam_sm_acct_mgmt(pam_handle_t *_pamh, int _flags, int _argc, const char **_argv){
    return PAMAccountManagement(_pamh,_flags,_argc,(char**)_argv);
}

PAM_EXTERN int pam_sm_open_session(pam_handle_t *_pamh, int _flags, int _argc, const char **_argv){
    return PAMOpenSession(_pamh,_flags,_argc,(char**)_argv);
}

PAM_EXTERN int pam_sm_close_session(pam_handle_t *_pamh, int _flags, int _args, const char **_argv){
    return PAMCloseSession(_pamh,_flags,_args,(char**)_argv);
}

PAM_EXTERN int pam_sm_chauthtok(pam_handle_t *_pamh, int _flags, int _argc, const char **_argv){
    return PAMChangeAuthToken(_pamh, _flags, _argc, (char**)_argv);
}
