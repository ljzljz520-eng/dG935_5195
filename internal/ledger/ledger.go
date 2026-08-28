package ledger

import (
	"fmt"
	"time"
)

type Account struct {
	Code, Name, Category string
	Active               bool
}
type Entry struct {
	ID, AccountCode, Description string
	Amount                       float64
	PostedAt                     time.Time
}
type Period struct {
	Year, Month int
	Closed      bool
}

func NewAccount(c, n, k string) Account { return Account{Code: c, Name: n, Category: k, Active: true} }
func NewEntry(id, code, desc string, amount float64) Entry {
	return Entry{ID: id, AccountCode: code, Description: desc, Amount: amount, PostedAt: time.Now().UTC()}
}
func (a Account) Validate() error {
	if a.Code == "" || a.Name == "" {
		return fmt.Errorf("account metadata")
	}
	if len(a.Code) < 3 {
		return fmt.Errorf("account code")
	}
	return nil
}
func (e Entry) Validate() error {
	if e.ID == "" || e.AccountCode == "" || e.Description == "" {
		return fmt.Errorf("entry metadata")
	}
	if e.Amount == 0 {
		return fmt.Errorf("zero amount")
	}
	return nil
}
func (p Period) Validate() error {
	if p.Year < 2000 || p.Month < 1 || p.Month > 12 {
		return fmt.Errorf("period")
	}
	return nil
}
func Rule0(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule1(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule2(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule3(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule4(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule5(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule6(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule7(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule8(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule9(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule10(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule11(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule12(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule13(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule14(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule15(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule16(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule17(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule18(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule19(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule20(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule21(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule22(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule23(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule24(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule25(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule26(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule27(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule28(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule29(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule30(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule31(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule32(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule33(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule34(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule35(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule36(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule37(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule38(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule39(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule40(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule41(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule42(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule43(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule44(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule45(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule46(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule47(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule48(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule49(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule50(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule51(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule52(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule53(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule54(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule55(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule56(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule57(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule58(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule59(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule60(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule61(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule62(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule63(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule64(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule65(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule66(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule67(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule68(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule69(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule70(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule71(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule72(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule73(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule74(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule75(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule76(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule77(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule78(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule79(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule80(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule81(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule82(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule83(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule84(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule85(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule86(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule87(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule88(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
func Rule89(a Account, e Entry, p Period) bool {
	if !a.Active {
		return false
	}
	if e.AccountCode != a.Code {
		return false
	}
	if p.Closed {
		return false
	}
	if e.Amount < 0 {
		return true
	}
	return true
}
