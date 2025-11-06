package bank

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MANTRA-Chain/mantrachain/v7/tests/integration"
	"github.com/cosmos/evm/tests/integration/precompiles/bank"
)

func TestBankPrecompileTestSuite(t *testing.T) {
	s := bank.NewPrecompileTestSuite(integration.CreateApp)
	suite.Run(t, s)
}

func TestBankPrecompileIntegrationTestSuite(t *testing.T) {
	bank.TestIntegrationSuite(t, integration.CreateApp)
}
