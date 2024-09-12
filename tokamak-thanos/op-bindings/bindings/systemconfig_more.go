// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"basefeeScalar\",\"offset\":8,\"slot\":\"104\",\"type\":\"t_uint32\"},{\"astId\":1010,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"blobbasefeeScalar\",\"offset\":12,\"slot\":\"104\",\"type\":\"t_uint32\"},{\"astId\":1011,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1012_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106102f45760003560e01c8063935f029e11610191578063e0e2016d116100e3578063f2fde38b11610097578063f8c68de011610071578063f8c68de014610708578063fd32aa0f14610710578063ffa1ad741461071857600080fd5b8063f2fde38b146106d8578063f45e65d8146106eb578063f68016b7146106f457600080fd5b8063e81b2c6d116100c8578063e81b2c6d146106a7578063ec707517146106b0578063f2b4e617146106d057600080fd5b8063e0e2016d14610697578063e2a3285c1461069f57600080fd5b8063bfb14fb711610145578063cc731b021161011f578063cc731b0214610553578063d844471514610687578063dac6e63a1461068f57600080fd5b8063bfb14fb714610507578063c4e8ddfa14610538578063c9b26f611461054057600080fd5b8063a711986911610176578063a7119869146104e4578063b40a817c146104ec578063bc49ce5f146104ff57600080fd5b8063935f029e146104c95780639b7d7f0a146104dc57600080fd5b806348cd4cb11161024a57806354fd4d50116101fe578063697844c6116101d8578063697844c61461049b578063715018a6146104a35780638da5cb5b146104ab57600080fd5b806354fd4d5014610449578063550fcdc91461048b5780635d73369c1461049357600080fd5b80634d0047ee1161022f5780634d0047ee146104075780634f16540b1461040f57806353d794be1461043657600080fd5b806348cd4cb1146103f75780634add321d146103ff57600080fd5b806318d13918116102ac5780632132684911610286578063213268491461039657806321d7fde5146103ae5780634397dfef146103c157600080fd5b806318d139181461037157806319f5cea8146103865780631fd19ee11461038e57600080fd5b80630a49cb03116102dd5780630a49cb03146103415780630ae14b1b146103495780630c18c1621461036857600080fd5b806306c92657146102f9578063078f29cf14610314575b600080fd5b610301610720565b6040519081526020015b60405180910390f35b61031c61074e565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161030b565b61031c610787565b630bebc2005b60405167ffffffffffffffff909116815260200161030b565b61030160655481565b61038461037f366004611e48565b6107b7565b005b6103016107cb565b61031c6107f6565b61039e610820565b604051901515815260200161030b565b6103846103bc366004611e7e565b61085f565b6103c9610875565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835260ff90911660208301520161030b565b610301610889565b61034f6108b9565b61031c6108df565b6103017f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b610384610444366004612019565b61090f565b60408051808201909152600c81527f322e332e302d626574612e32000000000000000000000000000000000000000060208201525b60405161030b91906121d3565b61047e610d44565b610301610d4e565b610301610d79565b610384610da4565b60335473ffffffffffffffffffffffffffffffffffffffff1661031c565b6103846104d73660046121e6565b610db8565b61031c610dca565b61031c610dfa565b6103846104fa366004612208565b610e2a565b610301610e3b565b6068546105239068010000000000000000900463ffffffff1681565b60405163ffffffff909116815260200161030b565b61031c610e66565b61038461054e366004612223565b610e96565b6106176040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b60405161030b9190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b61047e610ea7565b61031c610eb1565b610301610ee1565b610301610f0c565b61030160675481565b606854610523906c01000000000000000000000000900463ffffffff1681565b61031c610f37565b6103846106e6366004611e48565b610f67565b61030160665481565b60685461034f9067ffffffffffffffff1681565b61030161101b565b610301611046565b610301600081565b61074b60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d61226b565b81565b600061078261077e60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad637761226b565b5490565b905090565b600061078261077e60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad61226b565b6107bf611095565b6107c881611116565b50565b61074b60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a861226b565b60006107827f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b60008061082b610875565b5073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141592915050565b610867611095565b61087182826111d3565b5050565b600080610880611311565b90939092509050565b600061078261077e60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a061226b565b6069546000906107829063ffffffff6a0100000000000000000000820481169116612282565b600061078261077e60017fe1e3a95fb10ed56538cc130c2250de9823e7716d1142b8521655d7f7317b8ef161226b565b600054610100900460ff161580801561092f5750600054600160ff909116105b806109495750303b158015610949575060005460ff166001145b6109da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610a3857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610a4061138e565b610a498a610f67565b610a528761142d565b610a5c89896111d3565b610a6586611455565b610a8e7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08869055565b610ac1610abc60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc59861226b565b849055565b610af5610aef60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce958063761226b565b83519055565b610b2c610b2360017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a861226b565b60208401519055565b610b63610b5a60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad637761226b565b60408401519055565b610b9a610b9160017f52322a25d9f59ea17656545543306b7aef62bc0cc53a0e65ccfa0c75b97aa90761226b565b60608401519055565b610bd1610bc860017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad61226b565b60808401519055565b610c08610bff60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d61226b565b60a08401519055565b610c3f610c3660017fe1e3a95fb10ed56538cc130c2250de9823e7716d1142b8521655d7f7317b8ef161226b565b60e08401519055565b610c476115ab565b610c5084611613565b610c586108b9565b67ffffffffffffffff168667ffffffffffffffff161015610cd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016109d1565b8015610d3857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050505050565b6060610782611a87565b61074b60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce958063761226b565b61074b60017fe1e3a95fb10ed56538cc130c2250de9823e7716d1142b8521655d7f7317b8ef161226b565b610dac611095565b610db66000611b48565b565b610dc0611095565b6108718282611bbf565b600061078261077e60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d61226b565b600061078261077e60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce958063761226b565b610e32611095565b6107c881611455565b61074b60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc59861226b565b600061078261077e60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a861226b565b610e9e611095565b6107c88161142d565b6060610782611c95565b600061078261077e60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc59861226b565b61074b60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a061226b565b61074b60017f52322a25d9f59ea17656545543306b7aef62bc0cc53a0e65ccfa0c75b97aa90761226b565b600061078261077e60017f52322a25d9f59ea17656545543306b7aef62bc0cc53a0e65ccfa0c75b97aa90761226b565b610f6f611095565b73ffffffffffffffffffffffffffffffffffffffff8116611012576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016109d1565b6107c881611b48565b61074b60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad637761226b565b61074b60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad61226b565b9055565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b5490565b60335473ffffffffffffffffffffffffffffffffffffffff163314610db6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109d1565b61113f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08829055565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516111c791906121d3565b60405180910390a35050565b606880547fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff166801000000000000000063ffffffff8581169182027fffffffffffffffffffffffffffffffff00000000ffffffffffffffffffffffff16929092176c0100000000000000000000000092851692909202919091179091557f0100000000000000000000000000000000000000000000000000000000000000602083811b67ffffffff000000001690921717606681905560655460408051938401919091528201526000906060015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161130491906121d3565b60405180910390a3505050565b6000808061134361077e60017f04adb1412b2ddc16fcc0d4538d5c8f07cf9c83abecc6b41f6f69037b708fbcec61226b565b73ffffffffffffffffffffffffffffffffffffffff81169350905082611382575073eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee92601292509050565b60a081901c9150509091565b600054610100900460ff16611425576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016109d1565b610db6611d4b565b6067819055604080516020808201849052825180830390910181529082019091526000611196565b61145d6108b9565b67ffffffffffffffff168167ffffffffffffffff1610156114da576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016109d1565b630bebc20067ffffffffffffffff82161115611552576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206869676860448201526064016109d1565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002611196565b6115d961077e60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a061226b565b600003610db657610db661160e60017fa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb1a061226b565b439055565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff1611156116c3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016109d1565b6001816040015160ff161161175a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016109d1565b6068546080820151825167ffffffffffffffff9092169161177b91906122ae565b63ffffffff1611156117e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016109d1565b6000816020015160ff1611611880576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016109d1565b8051602082015163ffffffff82169160ff909116906118a09082906122cd565b6118aa9190612317565b63ffffffff161461193d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016109d1565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b60606000611a93611311565b5090507fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff821601611b0c57505060408051808201909152600381527f4554480000000000000000000000000000000000000000000000000000000000602082015290565b611b42611b3d61077e60017fa48b38a4b44951360fbdcbfaaeae5ed6ae92585412e9841b70ec72ed8cd0576461226b565b611deb565b91505090565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b7fff00000000000000000000000000000000000000000000000000000000000000811615611c6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f53797374656d436f6e6669673a207363616c61722065786365656473206d617860448201527f2e0000000000000000000000000000000000000000000000000000000000000060648201526084016109d1565b6065829055606681905560408051602081018490529081018290526000906060016112a1565b60606000611ca1611311565b5090507fffffffffffffffffffffffff111111111111111111111111111111111111111273ffffffffffffffffffffffffffffffffffffffff821601611d1a57505060408051808201909152600581527f4574686572000000000000000000000000000000000000000000000000000000602082015290565b611b42611b3d61077e60017f657c3582c29b3176614e3a33ddd1ec48352696a04e92b3c0566d72010fa8863d61226b565b600054610100900460ff16611de2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016109d1565b610db633611b48565b60405160005b82811a15611e0157600101611df1565b80825260208201838152600082820152505060408101604052919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611e4357600080fd5b919050565b600060208284031215611e5a57600080fd5b611e6382611e1f565b9392505050565b803563ffffffff81168114611e4357600080fd5b60008060408385031215611e9157600080fd5b611e9a83611e6a565b9150611ea860208401611e6a565b90509250929050565b803567ffffffffffffffff81168114611e4357600080fd5b60405160c0810167ffffffffffffffff81118282101715611f13577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803560ff81168114611e4357600080fd5b6000610100808385031215611f3e57600080fd5b6040519081019067ffffffffffffffff82118183101715611f88577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b81604052809250611f9884611e1f565b8152611fa660208501611e1f565b6020820152611fb760408501611e1f565b6040820152611fc860608501611e1f565b6060820152611fd960808501611e1f565b6080820152611fea60a08501611e1f565b60a0820152611ffb60c08501611e1f565b60c082015261200c60e08501611e1f565b60e0820152505092915050565b6000806000806000806000806000898b036102a081121561203957600080fd5b6120428b611e1f565b995061205060208c01611e6a565b985061205e60408c01611e6a565b975060608b0135965061207360808c01611eb1565b955061208160a08c01611e1f565b945060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff40820112156120b357600080fd5b506120bc611ec9565b6120c860c08c01611e6a565b81526120d660e08c01611f19565b60208201526120e86101008c01611f19565b60408201526120fa6101208c01611e6a565b606082015261210c6101408c01611e6a565b60808201526101608b01356fffffffffffffffffffffffffffffffff8116811461213557600080fd5b60a082015292506121496101808b01611e1f565b91506121598b6101a08c01611f2a565b90509295985092959850929598565b6000815180845260005b8181101561218e57602081850181015186830182015201612172565b818111156121a0576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e636020830184612168565b600080604083850312156121f957600080fd5b50508035926020909101359150565b60006020828403121561221a57600080fd5b611e6382611eb1565b60006020828403121561223557600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561227d5761227d61223c565b500390565b600067ffffffffffffffff8083168185168083038211156122a5576122a561223c565b01949350505050565b600063ffffffff8083168185168083038211156122a5576122a561223c565b600063ffffffff8084168061230b577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff8083168185168183048111821515161561233a5761233a61223c565b0294935050505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
	immutableReferences["SystemConfig"] = false
}
