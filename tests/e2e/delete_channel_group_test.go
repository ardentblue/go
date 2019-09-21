package e2e

import (
	"testing"

	pubnub "github.com/ardentblue/go-pubnub"
	"github.com/ardentblue/go-pubnub/tests/stubs"
	"github.com/stretchr/testify/assert"
)

func TestRemoveChannelGroupNotStubbed(t *testing.T) {
	assert := assert.New(t)

	pn := pubnub.NewPubNub(configCopy())

	_, _, err := pn.DeleteChannelGroup().
		ChannelGroup("cg").
		Execute()

	assert.Nil(err)
}

func TestRemoveChannelGroupNotStubbedContext(t *testing.T) {
	assert := assert.New(t)

	pn := pubnub.NewPubNub(configCopy())

	_, _, err := pn.DeleteChannelGroupWithContext(backgroundContext).
		ChannelGroup("cg1").
		Execute()

	assert.Nil(err)
}

func TestRemoveChannelGroupMissingGroup(t *testing.T) {
	assert := assert.New(t)

	pn := pubnub.NewPubNub(configCopy())

	_, _, err := pn.DeleteChannelGroup().
		Execute()

	assert.Contains(err.Error(), "Missing Channel Group")
}

func TestRemoveChannelGroupSuperCall(t *testing.T) {
	assert := assert.New(t)

	config := pamConfigCopy()

	// Not allowed characters:
	// .,:*

	validCharacters := "-_~?#[]@!$&'()+;=`|"

	config.UUID = validCharacters
	//config.AuthKey = validCharacters

	pn := pubnub.NewPubNub(config)

	_, _, err := pn.DeleteChannelGroup().
		ChannelGroup(validCharacters).
		Execute()

	assert.Nil(err)
}

func TestRemoveChannelGroupSuccessRemoved(t *testing.T) {
	assert := assert.New(t)
	myChannel := "my-channel-remove"
	myGroup := "my-unique-group-remove"

	interceptor := stubs.NewInterceptor()
	interceptor.AddStub(&stubs.Stub{
		Method:             "GET",
		Path:               "/v1/channel-registration/sub-key/sub-c-e41d50d4-43ce-11e8-a433-9e6b275e7b64/channel-group/my-unique-group-remove",
		Query:              "add=my-channel-remove",
		ResponseBody:       `{"status": 200, "message": "OK", "service": "channel-registry", "error": false}`,
		IgnoreQueryKeys:    []string{"uuid", "pnsdk", "l_cg"},
		ResponseStatusCode: 200,
	})
	interceptor.AddStub(&stubs.Stub{
		Method:             "GET",
		Path:               "/v1/channel-registration/sub-key/sub-c-e41d50d4-43ce-11e8-a433-9e6b275e7b64/channel-group/my-unique-group-remove",
		Query:              "remove=my-channel-remove",
		ResponseBody:       `{"status": 200, "message": "OK", "service": "channel-registry", "error": false}`,
		IgnoreQueryKeys:    []string{"uuid", "pnsdk", "l_cg"},
		ResponseStatusCode: 200,
	})
	interceptor.AddStub(&stubs.Stub{
		Method:             "GET",
		Path:               "/v1/channel-registration/sub-key/sub-c-e41d50d4-43ce-11e8-a433-9e6b275e7b64/channel-group/my-unique-group-remove",
		Query:              "",
		ResponseBody:       `{"status": 200, "payload": {"channels": [], "group": "my-unique-group-remove"}, "service": "channel-registry", "error": false}`,
		IgnoreQueryKeys:    []string{"uuid", "pnsdk", "l_cg"},
		ResponseStatusCode: 200,
	})

	pn := pubnub.NewPubNub(configCopy())
	pn.SetClient(interceptor.GetClient())

	_, _, err := pn.AddChannelToChannelGroup().
		Channels([]string{myChannel}).
		ChannelGroup(myGroup).
		Execute()

	assert.Nil(err)

	_, _, err = pn.RemoveChannelFromChannelGroup().
		Channels([]string{myChannel}).
		ChannelGroup(myGroup).
		Execute()

	assert.Nil(err)

	res, _, err := pn.ListChannelsInChannelGroup().
		ChannelGroup(myGroup).
		Execute()

	assert.Nil(err)

	assert.Equal(0, len(res.Channels))
	assert.Equal(myGroup, res.ChannelGroup)
}
