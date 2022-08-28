package sqlite

import (
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	l := hclog.NewNullLogger()
	_, err := InitDatabase(l, "test.db")
	assert.NoError(t, err)
}
