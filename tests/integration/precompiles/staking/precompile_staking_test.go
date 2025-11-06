package staking

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MANTRA-Chain/mantrachain/v7/tests/integration"
	"github.com/cosmos/evm/tests/integration/precompiles/staking"
)

func TestStakingPrecompileTestSuite(t *testing.T) {
	s := staking.NewPrecompileTestSuite(integration.CreateApp)
	suite.Run(t, s)
}

func TestStakingPrecompileIntegrationTestSuite(t *testing.T) {
	staking.TestPrecompileIntegrationTestSuite(t, integration.CreateApp)
}
