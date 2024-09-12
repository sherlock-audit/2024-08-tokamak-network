package genesis_test

import (
	"context"
	"encoding/json"
	"flag"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/bindings"
	"github.com/tokamak-network/tokamak-thanos/op-bindings/predeploys"
	"github.com/tokamak-network/tokamak-thanos/op-chain-ops/genesis"
	"github.com/tokamak-network/tokamak-thanos/op-service/eth"
)

var writeFile bool

func init() {
	flag.BoolVar(&writeFile, "write-file", false, "write the genesis file to disk")
}

var testKey, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

// Tests the BuildL2MainnetGenesis factory with the provided config.
func testBuildL2Genesis(t *testing.T, config *genesis.DeployConfig) *core.Genesis {
	backend := simulated.NewBackend(
		types.GenesisAlloc{
			crypto.PubkeyToAddress(testKey.PublicKey): {Balance: big.NewInt(10000000000000000)},
		},
	)
	defer backend.Close()

	client := backend.Client()
	block, err := client.BlockByNumber(context.Background(), nil)
	require.NoError(t, err)

	gen, err := genesis.BuildL2Genesis(config, block)
	require.Nil(t, err)
	require.NotNil(t, gen)

	proxyBytecode, err := bindings.GetDeployedBytecode("Proxy")
	require.NoError(t, err)

	transparentUpgradeableProxyBytecode, err := bindings.GetDeployedBytecode("TransparentUpgradeableProxy")
	require.NoError(t, err)

	l2UsdcBridgeProxyBytecode, err := bindings.GetDeployedBytecode("L2UsdcBridgeProxy")
	require.NoError(t, err)

	fiatTokenProxyBytecode, err := bindings.GetDeployedBytecode("FiatTokenProxy")
	require.NoError(t, err)

	for name, predeploy := range predeploys.Predeploys {
		addr := predeploy.Address

		account, ok := gen.Alloc[addr]
		require.Equal(t, true, ok, name)
		require.Greater(t, len(account.Code), 0)

		adminSlot, ok := account.Storage[genesis.AdminSlot]
		isProxy := !predeploy.ProxyDisabled ||
			(!config.EnableGovernance && addr == predeploys.GovernanceTokenAddr)

		if isProxy {
			switch addr {
			case predeploys.L2UsdcBridgeAddr:
				require.Equal(t, true, ok, name)
				require.Equal(t, eth.AddressAsLeftPaddedHash(predeploys.ProxyAdminAddr), adminSlot)
				require.Equal(t, l2UsdcBridgeProxyBytecode, account.Code)
			case predeploys.FiatTokenV2_2Addr:
				adminSlotForZepplin, ok := account.Storage[genesis.AdminSlotForZepplin]
				require.Equal(t, true, ok, name)
				require.Equal(t, eth.AddressAsLeftPaddedHash(predeploys.ProxyAdminAddr), adminSlotForZepplin)
				require.Equal(t, fiatTokenProxyBytecode, account.Code)
			case predeploys.NonfungibleTokenPositionDescriptorAddr:
				require.Equal(t, true, ok, name)
				require.Equal(t, eth.AddressAsLeftPaddedHash(predeploys.ProxyAdminAddr), adminSlot)
				require.Equal(t, transparentUpgradeableProxyBytecode, account.Code)

			default:
				require.Equal(t, true, ok, name)
				require.Equal(t, eth.AddressAsLeftPaddedHash(predeploys.ProxyAdminAddr), adminSlot)
				require.Equal(t, proxyBytecode, account.Code)
			}
		} else {
			require.Equal(t, false, ok, name)
			require.NotEqual(t, proxyBytecode, account.Code, name)
		}
	}

	// All of the precompile addresses should be funded with a single wei
	if config.SetPrecompileBalances {
		for i := 0; i < genesis.PrecompileCount; i++ {
			addr := common.BytesToAddress([]byte{byte(i)})
			require.Equal(t, common.Big1, gen.Alloc[addr].Balance)
		}
	}

	create2Deployer := gen.Alloc[predeploys.Create2DeployerAddr]
	codeHash := crypto.Keccak256Hash(create2Deployer.Code)
	require.Equal(t, codeHash, bindings.Create2DeployerCodeHash)

	if writeFile {
		file, _ := json.MarshalIndent(gen, "", " ")
		_ = os.WriteFile("genesis.json", file, 0644)
	}
	return gen
}

func TestBuildL2MainnetGenesis(t *testing.T) {
	config, err := genesis.NewDeployConfig("./testdata/test-deploy-config-devnet-l1.json")
	require.Nil(t, err)
	config.EnableGovernance = true
	config.FundDevAccounts = false
	gen := testBuildL2Genesis(t, config)
	require.Equal(t, 2081, len(gen.Alloc))
}

func TestBuildL2MainnetNoGovernanceGenesis(t *testing.T) {
	config, err := genesis.NewDeployConfig("./testdata/test-deploy-config-devnet-l1.json")
	require.Nil(t, err)
	config.EnableGovernance = false
	config.FundDevAccounts = false
	gen := testBuildL2Genesis(t, config)
	require.Equal(t, 2081, len(gen.Alloc))
}
