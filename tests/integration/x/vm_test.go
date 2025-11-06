package x

import (
	"testing"

	"github.com/MANTRA-Chain/mantrachain/v7/tests/integration"

	"github.com/cosmos/evm/tests/integration/x/vm"

	"github.com/stretchr/testify/suite"
)

// func TestKeeperTestSuite(t *testing.T) {
// 	s := vm.NewKeeperTestSuite(integration.CreateApp)
// 	s.EnableFeemarket = false
// 	s.EnableLondonHF = true
// 	suite.Run(t, s)
// }

func TestNestedEVMExtensionCallSuite(t *testing.T) {
	s := vm.NewNestedEVMExtensionCallSuite(integration.CreateApp)
	suite.Run(t, s)
}

// func TestGenesisTestSuite(t *testing.T) {
// 	s := vm.NewGenesisTestSuite(integration.CreateApp)
// 	suite.Run(t, s)
// }

func TestVmAnteTestSuite(t *testing.T) {
	s := vm.NewEvmAnteTestSuite(integration.CreateApp)
	suite.Run(t, s)
}

// func TestIterateContracts(t *testing.T) {
// 	vm.TestIterateContracts(t, integration.CreateApp)
// }
