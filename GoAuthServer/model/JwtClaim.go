package model

type JwtClaim struct {
	UserName    string
	Authorities string
	JTI         string
}
