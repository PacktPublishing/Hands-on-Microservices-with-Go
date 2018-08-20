package entities

import "testing"

func Test_User_GetAccountType(t *testing.T) {
	u := new(User)
	u.Account = 500

	accountType := u.GetAccountType()

	if accountType != NormalAccount {
		t.Errorf("Test failed. Expected Normal Account.")
	}

	u.Account = 1600

	accountType = u.GetAccountType()

	if accountType != SilverAccount {
		t.Errorf("Test failed. Expected Silver Account.")
	}

	//....
}
