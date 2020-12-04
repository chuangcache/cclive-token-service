package rtctoken

import (
	"fmt"
	"testing"
)

func Test_RtcTokenCreate(t *testing.T) {

	appId := "YLNyK4afgGSd753L"
	appToken := "YLiGj0NO5fun83bF"

	// channelName := ""
	// uidZero := uint32(0)
	// expiredTs := uint32(0)
	token, err := CreateToken(appId, appToken)

	if err != nil {
		t.Error(err)
	}

	fmt.Printf("token is %s\n", token)
}
