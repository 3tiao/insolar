package message

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/insolar/insolar/core"
	"github.com/insolar/insolar/cryptohelpers/hash"
	"github.com/insolar/insolar/testutils"
	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	ref := testutils.RandomRef()

	tmp := core.Message(&GenesisRequest{})
	msg, err := NewParcel(context.TODO(), tmp, ref, key, 1234, nil)
	assert.NoError(t, err)

	serialized, err := ToBytes(msg.Message())
	assert.NoError(t, err)
	msgHash := hash.SHA3Bytes256(serialized)

	err = ValidateToken(&key.PublicKey, msg.GetToken(), msgHash)
	assert.NoError(t, err)
}
