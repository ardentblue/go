package e2e

import (
	//"fmt"
	//"log"
	//"os"
	"testing"

	pubnub "github.com/ardentblue/go-pubnub"
	"github.com/stretchr/testify/assert"
)

func TestAddChannelToPushNotStubbed(t *testing.T) {
	assert := assert.New(t)

	pn := pubnub.NewPubNub(configCopy())
	//pn.Config.Log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	_, _, err := pn.AddPushNotificationsOnChannels().
		Channels([]string{"ch"}).
		DeviceIDForPush("cg").
		PushType(pubnub.PNPushTypeGCM).
		Execute()
	//fmt.Println(err.Error())
	assert.Nil(err)
}

func TestAddChannelToPushNotStubbedContext(t *testing.T) {
	assert := assert.New(t)

	pn := pubnub.NewPubNub(configCopy())
	//pn.Config.Log = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	_, _, err := pn.AddPushNotificationsOnChannelsWithContext(backgroundContext).
		Channels([]string{"ch1"}).
		DeviceIDForPush("cg1").
		PushType(pubnub.PNPushTypeGCM).
		Execute()
	//fmt.Println(err.Error())
	assert.Nil(err)
}
