package pubnub

import (
	"net/url"
	"testing"

	h "github.com/ardentblue/go-pubnub/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSubscribeSingleChannel(t *testing.T) {
	assert := assert.New(t)
	opts := &subscribeOpts{
		Channels: []string{"ch"},
		pubnub:   pubnub,
	}

	path, err := opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}
	h.AssertPathsEqual(t,
		"/v2/subscribe/sub_key/ch/0", u.EscapedPath(), []int{})
}

func TestSubscribeMultipleChannels(t *testing.T) {
	assert := assert.New(t)
	opts := &subscribeOpts{
		Channels: []string{"ch-1", "ch-2", "ch-3"},
		pubnub:   pubnub,
	}

	path, err := opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}

	h.AssertPathsEqual(t,
		"/v2/subscribe/sub_key/ch-1,ch-2,ch-3/0", u.EscapedPath(), []int{})
}

func TestSubscribeChannelGroups(t *testing.T) {
	assert := assert.New(t)
	opts := &subscribeOpts{
		ChannelGroups: []string{"cg-1", "cg-2", "cg-3"},
		pubnub:        pubnub,
	}

	path, err := opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}
	h.AssertPathsEqual(t,
		"/v2/subscribe/sub_key/,/0", u.EscapedPath(), []int{})

	query, err := opts.buildQuery()
	assert.Nil(err)

	expected := &url.Values{}
	expected.Set("channel-group", "cg-1,cg-2,cg-3")

	h.AssertQueriesEqual(t, expected, query,
		[]string{"pnsdk", "uuid"}, []string{})
}

func TestSubscribeMixedParams(t *testing.T) {
	assert := assert.New(t)

	opts := &subscribeOpts{
		Channels:         []string{"ch"},
		ChannelGroups:    []string{"cg"},
		Region:           "us-east-1",
		Timetoken:        123,
		FilterExpression: "abc",
		pubnub:           pubnub,
	}

	path, err := opts.buildPath()
	assert.Nil(err)

	u := &url.URL{
		Path: path,
	}

	h.AssertPathsEqual(t,
		"/v2/subscribe/sub_key/ch/0", u.EscapedPath(), []int{})

	query, err := opts.buildQuery()
	assert.Nil(err)

	expected := &url.Values{}
	expected.Set("channel-group", "cg")
	expected.Set("tr", "us-east-1")
	expected.Set("filter-expr", "abc")
	expected.Set("tt", "123")

	h.AssertQueriesEqual(t, expected, query,
		[]string{"pnsdk", "uuid"}, []string{})
}

func TestSubscribeValidateSubscribeKey(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	pn.Config.SubscribeKey = ""
	opts := &subscribeOpts{
		Channels:         []string{"ch"},
		ChannelGroups:    []string{"cg"},
		Region:           "us-east-1",
		Timetoken:        123,
		FilterExpression: "abc",
		pubnub:           pn,
	}

	assert.Equal("pubnub/validation: pubnub: \x01: Missing Subscribe Key", opts.validate().Error())
}

func TestSubscribeValidatePublishKey(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	pn.Config.PublishKey = ""
	opts := &subscribeOpts{
		Channels:         []string{"ch"},
		ChannelGroups:    []string{"cg"},
		Region:           "us-east-1",
		Timetoken:        123,
		FilterExpression: "abc",
		pubnub:           pn,
	}

	assert.Equal("pubnub/validation: pubnub: \x01: Missing Publish Key", opts.validate().Error())
}

func TestSubscribeValidateCHAndCG(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	opts := &subscribeOpts{
		Region:           "us-east-1",
		Timetoken:        123,
		FilterExpression: "abc",
		pubnub:           pn,
	}

	assert.Equal("pubnub/validation: pubnub: \x01: Missing Channel", opts.validate().Error())
}

func TestSubscribeValidateState(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	opts := &subscribeOpts{
		Channels:         []string{"ch"},
		ChannelGroups:    []string{"cg"},
		Region:           "us-east-1",
		Timetoken:        123,
		FilterExpression: "abc",
		pubnub:           pn,
	}
	opts.State = map[string]interface{}{"a": "a"}

	assert.Nil(opts.validate())
}
