package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/evm/tests/integration/rpc/backend"
)

func TestBackend(t *testing.T) {
	s := backend.NewTestSuite(CreateApp)
	suite.Run(t, s)
}
