package types

var KrbDictionary1 = struct {
	MsgTypesByID      map[int]string
	MsgTypesByName    map[string]int
	NameTypesByID     map[int]string
	ErrorCodesByID    map[int]string
	ADTypesByID       map[int]string
	PADataTypesByID   map[int]string
	PADataTypesByName map[string]int
	ETypesByID        map[int]string
	ETypesByName      map[string]int
}{
	MsgTypesByID: map[int]string{
		10: "KRB_AS_REQ",
		11: "KRB_AS_REP",
		12: "KRB_TGS_REQ",
		13: "KRB_TGS_REP",
		14: "KRB_AP_REQ",
		15: "KRB_AP_REP",
		16: "KRB_RESERVED16",
		17: "KRB_RESERVED17",
		20: "KRB_SAFE",
		21: "KRB_PRIV",
		22: "KRB_CRED",
		30: "KRB_ERROR",
	},
	MsgTypesByName: map[string]int{
		"KRB_AS_REQ":     10,
		"KRB_AS_REP":     11,
		"KRB_TGS_REQ":    12,
		"KRB_TGS_REP":    13,
		"KRB_AP_REQ":     14,
		"KRB_AP_REP":     15,
		"KRB_RESERVED16": 16,
		"KRB_RESERVED17": 17,
		"KRB_SAFE":       20,
		"KRB_PRIV":       21,
		"KRB_CRED":       22,
		"KRB_ERROR":      30,
	},
	NameTypesByID: map[int]string{
		0:  "KRB_NT_UNKNOWN",
		1:  "KRB_NT_PRINCIPAL",
		2:  "KRB_NT_SRV_INST",
		3:  "KRB_NT_SRV_HST",
		4:  "KRB_NT_SRV_XHST",
		5:  "KRB_NT_UID",
		6:  "KRB_NT_X500_PRINCIPAL",
		7:  "KRB_NT_SMTP_NAME",
		10: "KRB_NT_ENTERPRISE",
	},
	ErrorCodesByID: map[int]string{
		0:  "KDC_ERR_NONE",
		1:  "KDC_ERR_NAME_EXP",
		2:  "KDC_ERR_SERVICE_EXP",
		3:  "KDC_ERR_BAD_PVNO",
		4:  "KDC_ERR_C_OLD_MAST_KVNO",
		5:  "KDC_ERR_S_OLD_MAST_KVNO",
		6:  "KDC_ERR_C_PRINCIPAL_UNKNOWN",
		7:  "KDC_ERR_S_PRINCIPAL_UNKNOWN",
		8:  "KDC_ERR_PRINCIPAL_NOT_UNIQUE",
		9:  "KDC_ERR_NULL_KEY",
		10: "KDC_ERR_CANNOT_POSTDATE",
		11: "KDC_ERR_NEVER_VALID",
		12: "KDC_ERR_POLICY",
		13: "KDC_ERR_BADOPTION",
		14: "KDC_ERR_ETYPE_NOSUPP",
		15: "KDC_ERR_SUMTYPE_NOSUPP",
		16: "KDC_ERR_PADATA_TYPE_NOSUPP",
		17: "KDC_ERR_TRTYPE_NOSUPP",
		18: "KDC_ERR_CLIENT_REVOKED",
		19: "KDC_ERR_SERVICE_REVOKED",
		20: "KDC_ERR_TGT_REVOKED",
		21: "KDC_ERR_CLIENT_NOTYET",
		22: "KDC_ERR_SERVICE_NOTYET",
		23: "KDC_ERR_KEY_EXPIRED",
		24: "KDC_ERR_PREAUTH_FAILED",
		25: "KDC_ERR_PREAUTH_REQUIRED",
		26: "KDC_ERR_SERVER_NOMATCH",
		27: "KDC_ERR_MUST_USE_USER2USER",
		28: "KDC_ERR_PATH_NOT_ACCEPTED",
		29: "KDC_ERR_SVC_UNAVAILABLE",
		31: "KRB_AP_ERR_BAD_INTEGRITY",
		32: "KRB_AP_ERR_TKT_EXPIRED",
		33: "KRB_AP_ERR_TKT_NYV",
		34: "KRB_AP_ERR_REPEAT",
		35: "KRB_AP_ERR_NOT_US",
		36: "KRB_AP_ERR_BADMATCH",
		37: "KRB_AP_ERR_SKEW",
		38: "KRB_AP_ERR_BADADDR",
		39: "KRB_AP_ERR_BADVERSION",
		40: "KRB_AP_ERR_MSG_TYPE",
		41: "KRB_AP_ERR_MODIFIED",
		42: "KRB_AP_ERR_BADORDER",
		44: "KRB_AP_ERR_BADKEYVER",
		45: "KRB_AP_ERR_NOKEY",
		46: "KRB_AP_ERR_MUT_FAIL",
		47: "KRB_AP_ERR_BADDIRECTION",
		48: "KRB_AP_ERR_METHOD",
		49: "KRB_AP_ERR_BADSEQ",
		50: "KRB_AP_ERR_INAPP_CKSUM",
		51: "KRB_AP_PATH_NOT_ACCEPTED",
		52: "KRB_ERR_RESPONSE_TOO_BIG",
		60: "KRB_ERR_GENERIC",
		61: "KRB_ERR_FIELD_TOOLONG",
		62: "KDC_ERROR_CLIENT_NOT_TRUSTED",
		63: "KDC_ERROR_KDC_NOT_TRUSTED",
		64: "KDC_ERROR_INVALID_SIG",
		65: "KDC_ERR_KEY_TOO_WEAK",
		66: "KDC_ERR_CERTIFICATE_MISMATCH",
		67: "KRB_AP_ERR_NO_TGT",
		68: "KDC_ERR_WRONG_REALM",
		69: "KRB_AP_ERR_USER_TO_USER_REQUIRED",
		70: "KDC_ERR_CANT_VERIFY_CERTIFICATE",
		71: "KDC_ERR_INVALID_CERTIFICATE",
		72: "KDC_ERR_REVOKED_CERTIFICATE",
		73: "KDC_ERR_REVOCATION_STATUS_UNKNOWN",
		74: "KDC_ERR_REVOCATION_STATUS_UNAVAILABLE",
		75: "KDC_ERR_CLIENT_NAME_MISMATCH",
		76: "KDC_ERR_KDC_NAME_MISMATCH",
	},
	ADTypesByID: map[int]string{
		1: "AD-IF-RELEVANT",
		4: "AD-KDCIssued",
		5: "AD-AND-OR",
		8: "AD-MANDATORY-FOR-KDC",
	},
	PADataTypesByID: map[int]string{
		1:  "pa-tgs-req",
		2:  "pa-enc-timestamp",
		3:  "pa-pw-salt",
		11: "pa-etype-info",
		19: "pa-etype-info2",
	},
	PADataTypesByName: map[string]int{
		"pa-tgs-req":       1,
		"pa-enc-timestamp": 2,
		"pa-pw-salt":       3,
		"pa-etype-info":    11,
		"pa-etype-info2":   19,
	},
	ETypesByID: map[int]string{
		1:  "des-cbc-crc",
		2:  "des-cbc-md4",
		3:  "des-cbc-md5",
		4:  "reserved4",
		5:  "des3-cbc-md5",
		6:  "reserved6",
		7:  "des3-cbc-sha1",
		9:  "dsaWithSHA1-CmsOID",
		10: "md5WithRSAEncryption-CmsOID",
		11: "sha1WithRSAEncryption-CmsOID",
		12: "rc2CBC-EnvOID",
		13: "rsaEncryption-EnvOID",
		14: "rsaES-OAEP-ENV-OID",
		15: "des-ede3-cbc-Env-OID",
		16: "des3-cbc-sha1-kd",
		17: "aes128-cts-hmac-sha1-96",
		18: "aes256-cts-hmac-sha1-96",
		23: "rc4-hmac",
		24: "rc4-hmac-exp",
		65: "subkey-keymaterial",
	},
	ETypesByName: map[string]int{
		"des-cbc-crc":                  1,
		"des-cbc-md4":                  2,
		"des-cbc-md5":                  3,
		"reserved4":                    4,
		"des3-cbc-md5":                 5,
		"reserved6":                    6,
		"des3-cbc-sha1":                7,
		"dsaWithSHA1-CmsOID":           9,
		"md5WithRSAEncryption-CmsOID":  10,
		"sha1WithRSAEncryption-CmsOID": 11,
		"rc2CBC-EnvOID":                12,
		"rsaEncryption-EnvOID":         13,
		"rsaES-OAEP-ENV-OID":           14,
		"des-ede3-cbc-Env-OID":         15,
		"des3-cbc-sha1-kd":             16,
		"aes128-cts-hmac-sha1-96":      17,
		"aes256-cts-hmac-sha1-96":      18,
		"rc4-hmac":                     23,
		"rc4-hmac-exp":                 23,
		"subkey-keymaterial":           65,
	},
}

// TODO I think we should have a map of message type map[int]string or message interface???
