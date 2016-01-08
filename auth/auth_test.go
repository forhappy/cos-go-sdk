package auth

import (
	"testing"
)

func TestSign(t *testing.T) {
	signer := NewSignature("200001",
		"newbucket",
		"AKIDUfLUEUigQiXqm7CVSspKJnuaiIKtxqAv",
		"1438669115",
		"1436077115",
		"11162",
		"tencentyunSignTest")
	expected := "Ijp2RvogJLTzpLwn3gqplPdNYwxhPTIwMDAwMSZrPUFLSURVZkxVRVVpZ1FpWHFtN0NWU3NwS0pudWFpSUt0eHFBdiZlPTE0Mzg2NjkxMTUmdD0xNDM2MDc3MTE1JnI9MTExNjImZj0mYj1uZXdidWNrZXQ="
	actual := signer.Sign("bLcPnl88WU30VY57ipRhSePfPdOfSruK")
	if expected != actual {
		t.Errorf("Should match [EXPECTED:%s]:[ACTUAL:%s]", expected, actual)
	}
}

func TestSignOnce(t *testing.T) {
	signer := NewSignature("200001",
		"newbucket",
		"AKIDUfLUEUigQiXqm7CVSspKJnuaiIKtxqAv",
		"1438669115",
		"1436077115",
		"11162",
		"tencentyunSignTest")
	expected := "ZdjX8GBgSMlgHzYNAj7QTQ8GjeJhPTIwMDAwMSZrPUFLSURVZkxVRVVpZ1FpWHFtN0NWU3NwS0pudWFpSUt0eHFBdiZlPTAmdD0xNDM2MDc3MTE1JnI9MTExNjImZj10ZW5jZW50eXVuU2lnblRlc3QmYj1uZXdidWNrZXQ="
	actual := signer.SignOnce("bLcPnl88WU30VY57ipRhSePfPdOfSruK")
	if expected != actual {
		t.Errorf("Should match [EXPECTED:%s]:[ACTUAL:%s]", expected, actual)
	}
}
