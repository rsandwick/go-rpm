package rpm

//go:generate stringer -type SigTag

type SigTag Tag

const RPMTAG_SIG_BASE SigTag = 256

const (
	RPMSIGTAG_SIZE SigTag = iota + 1000
	RPMSIGTAG_LEMD5_1
	RPMSIGTAG_PGP
	RPMSIGTAG_LEMD5_2
	RPMSIGTAG_MD5
	RPMSIGTAG_GPG
	RPMSIGTAG_PGP5
	RPMSIGTAG_PAYLOADSIZE
	RPMSIGTAG_RESERVEDSPACE
	RPMSIGTAG_BADSHA1_1
	RPMSIGTAG_BADSHA1_2
	RPMSIGTAG_DSA
	RPMSIGTAG_RSA
	RPMSIGTAG_SHA1
	RPMSIGTAG_LONGSIZE
	RPMSIGTAG_LONGARCHIVESIZE
	RPMSIGTAG_SHA256

	RPMSIGTAG_FILESIGNATURES SigTag = RPMTAG_SIG_BASE + 18 + iota
	RPMSIGTAG_FILESIGNATURESLENGTH
)
