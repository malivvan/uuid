package uuid

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var genTest = Func("TEST", 8)

func TestHostID(t *testing.T) {
	id, err := HostID()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	t.Log(id)
}

func TestStringParse(t *testing.T) {
	id1, err := genTest()
	assert.NoError(t, err)
	id2, err := Parse(id1.String())
	assert.NoError(t, err)
	assert.Equal(t, id1, id2)
	assert.Equal(t, id1.String(), id2.String())
	assert.Equal(t, id1.Type(), id2.Type())
	assert.Equal(t, id1.Host(), id2.Host())
	assert.Equal(t, id1.Time(), id2.Time())
	assert.Equal(t, id1.Data(), id2.Data())
	assert.Equal(t, id1.CRC(), id2.CRC())
	t.Log(id1.String())
	t.Log(fmt.Sprintf("Type: %s", id1.Type()))
	t.Log(fmt.Sprintf("Host: %X", id1.Host()))
	t.Log(fmt.Sprintf("Time: %d", uint32(id1.Time().Unix())))
	t.Log(fmt.Sprintf("Data: %X", id1.Data()))
	t.Log(fmt.Sprintf("CRC: %X", id1.CRC()))
}

func TestB64Encoding(t *testing.T) {
	id1, err := genTest()
	assert.NoError(t, err)
	id2, err := Decode(id1.Encode())
	assert.NoError(t, err)
	assert.Equal(t, id1, id2)
	assert.Equal(t, id1.String(), id2.String())
	assert.Equal(t, id1.Type(), id2.Type())
	assert.Equal(t, id1.Host(), id2.Host())
	assert.Equal(t, id1.Time(), id2.Time())
	assert.Equal(t, id1.Data(), id2.Data())
	assert.Equal(t, id1.CRC(), id2.CRC())
	t.Log(id1.String())
	t.Log(fmt.Sprintf("Type: %s", id1.Type()))
	t.Log(fmt.Sprintf("Host: %X", id1.Host()))
	t.Log(fmt.Sprintf("Time: %d", uint32(id1.Time().Unix())))
	t.Log(fmt.Sprintf("Data: %X", id1.Data()))
	t.Log(fmt.Sprintf("CRC: %X", id1.CRC()))
}
