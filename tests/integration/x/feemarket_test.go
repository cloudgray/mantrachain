package x

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/MANTRA-Chain/mantrachain/v7/tests/integration"

	"github.com/cosmos/evm/tests/integration/x/feemarket"
)

func TestFeeMarketKeeperTestSuite(t *testing.T) {
	s := feemarket.NewTestKeeperTestSuite(integration.CreateApp)
	suite.Run(t, s)
}
