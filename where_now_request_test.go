package pubnub

import (
	"net/url"
	"testing"

	h "github.com/ardentblue/go-pubnub/tests/helpers"
	"github.com/stretchr/testify/assert"
)

func init() {
	pnconfig = NewConfig()

	pnconfig.PublishKey = "pub_key"
	pnconfig.SubscribeKey = "sub_key"
	pnconfig.SecretKey = "secret_key"

	pubnub = NewPubNub(pnconfig)
}

func TestWhereNowBasicRequest(t *testing.T) {
	assert := assert.New(t)

	opts := &whereNowOpts{
		UUID:   "my-custom-uuid",
		pubnub: pubnub,
	}
	path, err := opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}
	h.AssertPathsEqual(t,
		"/v2/presence/sub-key/sub_key/uuid/my-custom-uuid",
		u.EscapedPath(), []int{})

	query, err := opts.buildQuery()
	assert.Nil(err)

	expected := &url.Values{}
	h.AssertQueriesEqual(t, expected, query, []string{"pnsdk", "uuid"}, []string{})

	body, err := opts.buildBody()

	assert.Nil(err)
	assert.Equal([]byte{}, body)
}

func TestNewWhereNowBuilder(t *testing.T) {
	assert := assert.New(t)

	o := newWhereNowBuilder(pubnub)
	o.UUID("my-custom-uuid")

	path, err := o.opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}
	h.AssertPathsEqual(t,
		"/v2/presence/sub-key/sub_key/uuid/my-custom-uuid",
		u.EscapedPath(), []int{})
}

func TestNewWhereNowBuilderContext(t *testing.T) {
	assert := assert.New(t)

	o := newWhereNowBuilderWithContext(pubnub, backgroundContext)
	o.UUID("my-custom-uuid")

	path, err := o.opts.buildPath()
	assert.Nil(err)
	u := &url.URL{
		Path: path,
	}
	h.AssertPathsEqual(t,
		"/v2/presence/sub-key/sub_key/uuid/my-custom-uuid",
		u.EscapedPath(), []int{})
}

func TestNewWhereNowResponserrorUnmarshalling(t *testing.T) {
	assert := assert.New(t)
	jsonBytes := []byte(`s`)

	_, _, err := newWhereNowResponse(jsonBytes, StatusResponse{})
	assert.Equal("pubnub/parsing: Error unmarshalling response: {s}", err.Error())
}

func TestWhereNowValidateSubscribeKey(t *testing.T) {
	assert := assert.New(t)
	pn := NewPubNub(NewDemoConfig())
	pn.Config.SubscribeKey = ""
	opts := &whereNowOpts{
		UUID:   "my-custom-uuid",
		pubnub: pn,
	}

	assert.Equal("pubnub/validation: pubnub: \a: Missing Subscribe Key", opts.validate().Error())
}
