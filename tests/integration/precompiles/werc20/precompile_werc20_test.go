package werc20

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MANTRA-Chain/mantrachain/v7/tests/integration"
	"github.com/cosmos/evm/tests/integration/precompiles/werc20"
)

func TestWERC20PrecompileUnitTestSuite(t *testing.T) {
	s := werc20.NewPrecompileUnitTestSuite(integration.CreateApp)
	suite.Run(t, s)
}

func TestWERC20PrecompileIntegrationTestSuite(t *testing.T) {
	werc20.TestPrecompileIntegrationTestSuite(t, integration.CreateApp)
}
