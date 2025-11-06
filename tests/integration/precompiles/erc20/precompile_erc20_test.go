package erc20

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MANTRA-Chain/mantrachain/v7/tests/integration"
	erc21 "github.com/cosmos/evm/tests/integration/precompiles/erc20"
)

func TestErc20PrecompileTestSuite(t *testing.T) {
	s := erc21.NewPrecompileTestSuite(integration.CreateApp)
	suite.Run(t, s)
}

func TestErc20IntegrationTestSuite(t *testing.T) {
	erc21.TestIntegrationTestSuite(t, integration.CreateApp)
}
