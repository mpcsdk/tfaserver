package model

const (
	Kind_RiskTx  string = "riskTx"
	Kind_RiskTfa string = "riskTfa"
)
const (
	///
	Type_TfaBindPhone   string = "bindPhone"
	Type_TfaBindMail    string = "bindMail"
	Type_TfaUpdatePhone string = "updatePhone"
	Type_TfaUpdateMail  string = "updateMail"
)

type RiskPenndingKey string

func RiskPenddingKey(userId, riskSerial string) RiskPenndingKey {

	return RiskPenndingKey("riskPendding:" + userId + ":" + riskSerial)
}

type RiskKind string
type VerifyKind string

const (
	RiskKind_Nil       = "RiskKind_Nil"
	RiskKind_Tx        = "RiskKind_Tx"
	RiskKind_BindPhone = "RiskKind_BindPhone"
	RiskKind_UpPhone   = "RiskKind_UpPhone"
	RiskKind_BindMail  = "RiskKind_BindMail"
	RiskKind_UpMail    = "RiskKind_UpMail"
)

type RiskTfa struct {
	UserId    string `json:"userId"`
	UserToken string `json:"token"`
	Type      string `json:"type"`
	///
	Mail  string `json:"mail"`
	Phone string `json:"phone"`
}
type IVerifier interface {
	Verify(verifierCode *VerifyCode) (RiskKind, error)
	SetCode(string)
	RiskKind() RiskKind
	VerifyKind() VerifyKind
	IsDone() bool
	///
	SendVerificationCode() (string, error)
	SendCompletion() error
	//
	Destination() string
}
