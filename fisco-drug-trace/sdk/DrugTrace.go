// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sdk

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// DrugTraceDrug is an auto generated low-level Go binding around an user-defined struct.
type DrugTraceDrug struct {
	Id                    *big.Int
	Producer              *big.Int
	ProductionDate        *big.Int
	ProductionDateStr     string
	Flow                  string
	TransportationStr     string
	Dealer                *big.Int
	DrugAcceptanceTime    *big.Int
	DrugStorageConditions string
	DrugStorageLocation   string
	Buyer                 *big.Int
	BuyTime               *big.Int
}

// DrugTraceABI is the input ABI used to generate the binding from.
const DrugTraceABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_drugStorageConditions\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_drugStorageLocation\",\"type\":\"string\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_producer\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_prouctionDataStr\",\"type\":\"string\"}],\"name\":\"addDrug\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_buyer\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"drugs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"producer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"productionDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"productionDateStr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"flow\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportationStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dealer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drugAcceptanceTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"drugStorageConditions\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"drugStorageLocation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"buyer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDrugCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_userId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDrugs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"producer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"productionDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"productionDateStr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"flow\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportationStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dealer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drugAcceptanceTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"drugStorageConditions\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"drugStorageLocation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"buyer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTime\",\"type\":\"uint256\"}],\"internalType\":\"structDrugTrace.Drug[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPayableDrugCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPayableDrugs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"producer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"productionDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"productionDateStr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"flow\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"transportationStr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"dealer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"drugAcceptanceTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"drugStorageConditions\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"drugStorageLocation\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"buyer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTime\",\"type\":\"uint256\"}],\"internalType\":\"structDrugTrace.Drug[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payableDrugIdList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dealer\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_transportationStr\",\"type\":\"string\"}],\"name\":\"sale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userIdToDrugIdList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DrugTraceBin is the compiled bytecode used for deploying new contracts.
var DrugTraceBin = "0x60806040526001805534801561001457600080fd5b506121ee806100246000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063891cde9911610071578063891cde9914610176578063897b3c74146101a65780639af82521146101d6578063d6febde814610206578063d70518c714610222578063fca5e61e1461023e576100a9565b8063154cce35146100ae5780632df74c27146100de578063327af74c146100fa57806344bd7bf61461011657806349eb063d14610146575b600080fd5b6100c860048036038101906100c39190611844565b610279565b6040516100d59190611880565b60405180910390f35b6100f860048036038101906100f391906119e1565b61031e565b005b610114600480360381019061010f9190611a6c565b610407565b005b610130600480360381019061012b9190611844565b6104e8565b60405161013d9190611880565b60405180910390f35b610160600480360381019061015b9190611adb565b61050c565b60405161016d9190611da7565b60405180910390f35b610190600480360381019061018b9190611dc9565b610a11565b60405161019d9190611880565b60405180910390f35b6101c060048036038101906101bb9190611e09565b610ad9565b6040516101cd9190611da7565b60405180910390f35b6101f060048036038101906101eb9190611dc9565b611034565b6040516101fd9190611880565b60405180910390f35b610220600480360381019061021b9190611dc9565b611065565b005b61023c60048036038101906102379190611e70565b6111c4565b005b61025860048036038101906102539190611844565b6113de565b6040516102709c9b9a99989796959493929190611f16565b60405180910390f35b6000806000905060005b6003805490508110156103145760008414806102ed57508360006001600384815481106102b3576102b2611ff3565b5b90600052602060002001546102c89190612051565b815481106102d9576102d8611ff3565b5b90600052602060002090600c020160000154145b156103015781806102fd90612085565b9250505b808061030c90612085565b915050610283565b5080915050919050565b60008060018561032e9190612051565b8154811061033f5761033e611ff3565b5b90600052602060002090600c02019050428160070181905550828160080190805190602001906103709291906116f6565b50818160090190805190602001906103899291906116f6565b506040518060400160405280600981526020017fe7bb8fe99480e5958600000000000000000000000000000000000000000000008152508160040190805190602001906103d79291906116f6565b50600384908060018154018082558091505060019003906000526020600020016000909190919091505550505050565b6000806001856104179190612051565b8154811061042857610427611ff3565b5b90600052602060002090600c020190508281600601819055506040518060400160405280601481526020017fe7949fe4baa7e595862d3ee7bb8fe99480e5958600000000000000000000000081525081600401908051906020019061048e9291906116f6565b50818160050190805190602001906104a79291906116f6565b506002600084815260200190815260200160002084908060018154018082558091505060019003906000526020600020016000909190919091505550505050565b600381815481106104f857600080fd5b906000526020600020016000915090505481565b606060008360018661051e9190612051565b61052891906120ce565b9050600081858761053991906120ce565b6105439190612051565b9050600380549050858761055791906120ce565b1115610571578160038054905061056e9190612051565b90505b60008167ffffffffffffffff81111561058d5761058c6118b6565b5b6040519080825280602002602001820160405280156105c657816020015b6105b361177c565b8152602001906001900390816105ab5790505b5090506000808490505b600380549050811080156105e357508382105b15610a02576000871480610645575086600060016003848154811061060b5761060a611ff3565b5b90600052602060002001546106209190612051565b8154811061063157610630611ff3565b5b90600052602060002090600c020160000154145b156109ef57600060016003838154811061066257610661611ff3565b5b90600052602060002001546106779190612051565b8154811061068857610687611ff3565b5b90600052602060002090600c0201604051806101800160405290816000820154815260200160018201548152602001600282015481526020016003820180546106d090612157565b80601f01602080910402602001604051908101604052809291908181526020018280546106fc90612157565b80156107495780601f1061071e57610100808354040283529160200191610749565b820191906000526020600020905b81548152906001019060200180831161072c57829003601f168201915b5050505050815260200160048201805461076290612157565b80601f016020809104026020016040519081016040528092919081815260200182805461078e90612157565b80156107db5780601f106107b0576101008083540402835291602001916107db565b820191906000526020600020905b8154815290600101906020018083116107be57829003601f168201915b505050505081526020016005820180546107f490612157565b80601f016020809104026020016040519081016040528092919081815260200182805461082090612157565b801561086d5780601f106108425761010080835404028352916020019161086d565b820191906000526020600020905b81548152906001019060200180831161085057829003601f168201915b50505050508152602001600682015481526020016007820154815260200160088201805461089a90612157565b80601f01602080910402602001604051908101604052809291908181526020018280546108c690612157565b80156109135780601f106108e857610100808354040283529160200191610913565b820191906000526020600020905b8154815290600101906020018083116108f657829003601f168201915b5050505050815260200160098201805461092c90612157565b80601f016020809104026020016040519081016040528092919081815260200182805461095890612157565b80156109a55780601f1061097a576101008083540402835291602001916109a5565b820191906000526020600020905b81548152906001019060200180831161098857829003601f168201915b50505050508152602001600a8201548152602001600b820154815250508383815181106109d5576109d4611ff3565b5b602002602001018190525081806109eb90612085565b9250505b80806109fa90612085565b9150506105d0565b82955050505050509392505050565b6000806000905060005b6002600086815260200190815260200160002080549050811015610ace576000841480610aa757508360006001600260008981526020019081526020016000208481548110610a6d57610a6c611ff3565b5b9060005260206000200154610a829190612051565b81548110610a9357610a92611ff3565b5b90600052602060002090600c020160000154145b15610abb578180610ab790612085565b9250505b8080610ac690612085565b915050610a1b565b508091505092915050565b6060600084600187610aeb9190612051565b610af591906120ce565b90506000818688610b0691906120ce565b610b109190612051565b905060026000868152602001908152602001600020805490508688610b3591906120ce565b1115610b6057816002600087815260200190815260200160002080549050610b5d9190612051565b90505b60008167ffffffffffffffff811115610b7c57610b7b6118b6565b5b604051908082528060200260200182016040528015610bb557816020015b610ba261177c565b815260200190600190039081610b9a5790505b5090506000808490505b600260008981526020019081526020016000208054905081108015610be357508382105b15611024576000871480610c5657508660006001600260008c81526020019081526020016000208481548110610c1c57610c1b611ff3565b5b9060005260206000200154610c319190612051565b81548110610c4257610c41611ff3565b5b90600052602060002090600c020160000154145b156110115760006001600260008b81526020019081526020016000208381548110610c8457610c83611ff3565b5b9060005260206000200154610c999190612051565b81548110610caa57610ca9611ff3565b5b90600052602060002090600c020160405180610180016040529081600082015481526020016001820154815260200160028201548152602001600382018054610cf290612157565b80601f0160208091040260200160405190810160405280929190818152602001828054610d1e90612157565b8015610d6b5780601f10610d4057610100808354040283529160200191610d6b565b820191906000526020600020905b815481529060010190602001808311610d4e57829003601f168201915b50505050508152602001600482018054610d8490612157565b80601f0160208091040260200160405190810160405280929190818152602001828054610db090612157565b8015610dfd5780601f10610dd257610100808354040283529160200191610dfd565b820191906000526020600020905b815481529060010190602001808311610de057829003601f168201915b50505050508152602001600582018054610e1690612157565b80601f0160208091040260200160405190810160405280929190818152602001828054610e4290612157565b8015610e8f5780601f10610e6457610100808354040283529160200191610e8f565b820191906000526020600020905b815481529060010190602001808311610e7257829003601f168201915b505050505081526020016006820154815260200160078201548152602001600882018054610ebc90612157565b80601f0160208091040260200160405190810160405280929190818152602001828054610ee890612157565b8015610f355780601f10610f0a57610100808354040283529160200191610f35565b820191906000526020600020905b815481529060010190602001808311610f1857829003601f168201915b50505050508152602001600982018054610f4e90612157565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7a90612157565b8015610fc75780601f10610f9c57610100808354040283529160200191610fc7565b820191906000526020600020905b815481529060010190602001808311610faa57829003601f168201915b50505050508152602001600a8201548152602001600b82015481525050838381518110610ff757610ff6611ff3565b5b6020026020010181905250818061100d90612085565b9250505b808061101c90612085565b915050610bbf565b8295505050505050949350505050565b6002602052816000526040600020818154811061105057600080fd5b90600052602060002001600091509150505481565b6000806001846110759190612051565b8154811061108657611085611ff3565b5b90600052602060002090600c020190508181600a01819055504281600b01819055506040518060400160405280600981526020017fe5b7b2e594aee587ba00000000000000000000000000000000000000000000008152508160040190805190602001906110f59291906116f6565b5060005b6003805490508110156111be57836003828154811061111b5761111a611ff3565b5b906000526020600020015414156111ab57600360016003805490506111409190612051565b8154811061115157611150611ff3565b5b9060005260206000200154600382815481106111705761116f611ff3565b5b906000526020600020018190555060038054806111905761118f612189565b5b600190038181906000526020600020016000905590556111be565b80806111b690612085565b9150506110f9565b50505050565b600060405180610180016040528060015481526020018481526020014281526020018381526020016040518060400160405280600981526020017fe7949fe4baa7e5958600000000000000000000000000000000000000000000008152508152602001604051806020016040528060008152508152602001600081526020016000815260200160405180602001604052806000815250815260200160405180602001604052806000815250815260200160008152602001600081525090806001815401808255809150506001900390600052602060002090600c020160009091909190915060008201518160000155602082015181600101556040820151816002015560608201518160030190805190602001906112e39291906116f6565b5060808201518160040190805190602001906113009291906116f6565b5060a082015181600501908051906020019061131d9291906116f6565b5060c0820151816006015560e0820151816007015561010082015181600801908051906020019061134f9291906116f6565b5061012082015181600901908051906020019061136d9291906116f6565b5061014082015181600a015561016082015181600b01555050600260008381526020019081526020016000206001549080600181540180825580915050600190039060005260206000200160009091909190915055600160008154809291906113d590612085565b91905055505050565b600081815481106113ee57600080fd5b90600052602060002090600c020160009150905080600001549080600101549080600201549080600301805461142390612157565b80601f016020809104026020016040519081016040528092919081815260200182805461144f90612157565b801561149c5780601f106114715761010080835404028352916020019161149c565b820191906000526020600020905b81548152906001019060200180831161147f57829003601f168201915b5050505050908060040180546114b190612157565b80601f01602080910402602001604051908101604052809291908181526020018280546114dd90612157565b801561152a5780601f106114ff5761010080835404028352916020019161152a565b820191906000526020600020905b81548152906001019060200180831161150d57829003601f168201915b50505050509080600501805461153f90612157565b80601f016020809104026020016040519081016040528092919081815260200182805461156b90612157565b80156115b85780601f1061158d576101008083540402835291602001916115b8565b820191906000526020600020905b81548152906001019060200180831161159b57829003601f168201915b5050505050908060060154908060070154908060080180546115d990612157565b80601f016020809104026020016040519081016040528092919081815260200182805461160590612157565b80156116525780601f1061162757610100808354040283529160200191611652565b820191906000526020600020905b81548152906001019060200180831161163557829003601f168201915b50505050509080600901805461166790612157565b80601f016020809104026020016040519081016040528092919081815260200182805461169390612157565b80156116e05780601f106116b5576101008083540402835291602001916116e0565b820191906000526020600020905b8154815290600101906020018083116116c357829003601f168201915b50505050509080600a01549080600b015490508c565b82805461170290612157565b90600052602060002090601f016020900481019282611724576000855561176b565b82601f1061173d57805160ff191683800117855561176b565b8280016001018555821561176b579182015b8281111561176a57825182559160200191906001019061174f565b5b50905061177891906117dd565b5090565b6040518061018001604052806000815260200160008152602001600081526020016060815260200160608152602001606081526020016000815260200160008152602001606081526020016060815260200160008152602001600081525090565b5b808211156117f65760008160009055506001016117de565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6118218161180e565b811461182c57600080fd5b50565b60008135905061183e81611818565b92915050565b60006020828403121561185a57611859611804565b5b60006118688482850161182f565b91505092915050565b61187a8161180e565b82525050565b60006020820190506118956000830184611871565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6118ee826118a5565b810181811067ffffffffffffffff8211171561190d5761190c6118b6565b5b80604052505050565b60006119206117fa565b905061192c82826118e5565b919050565b600067ffffffffffffffff82111561194c5761194b6118b6565b5b611955826118a5565b9050602081019050919050565b82818337600083830152505050565b600061198461197f84611931565b611916565b9050828152602081018484840111156119a05761199f6118a0565b5b6119ab848285611962565b509392505050565b600082601f8301126119c8576119c761189b565b5b81356119d8848260208601611971565b91505092915050565b6000806000606084860312156119fa576119f9611804565b5b6000611a088682870161182f565b935050602084013567ffffffffffffffff811115611a2957611a28611809565b5b611a35868287016119b3565b925050604084013567ffffffffffffffff811115611a5657611a55611809565b5b611a62868287016119b3565b9150509250925092565b600080600060608486031215611a8557611a84611804565b5b6000611a938682870161182f565b9350506020611aa48682870161182f565b925050604084013567ffffffffffffffff811115611ac557611ac4611809565b5b611ad1868287016119b3565b9150509250925092565b600080600060608486031215611af457611af3611804565b5b6000611b028682870161182f565b9350506020611b138682870161182f565b9250506040611b248682870161182f565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611b638161180e565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611ba3578082015181840152602081019050611b88565b83811115611bb2576000848401525b50505050565b6000611bc382611b69565b611bcd8185611b74565b9350611bdd818560208601611b85565b611be6816118a5565b840191505092915050565b600061018083016000830151611c0a6000860182611b5a565b506020830151611c1d6020860182611b5a565b506040830151611c306040860182611b5a565b5060608301518482036060860152611c488282611bb8565b91505060808301518482036080860152611c628282611bb8565b91505060a083015184820360a0860152611c7c8282611bb8565b91505060c0830151611c9160c0860182611b5a565b5060e0830151611ca460e0860182611b5a565b50610100830151848203610100860152611cbe8282611bb8565b915050610120830151848203610120860152611cda8282611bb8565b915050610140830151611cf1610140860182611b5a565b50610160830151611d06610160860182611b5a565b508091505092915050565b6000611d1d8383611bf1565b905092915050565b6000602082019050919050565b6000611d3d82611b2e565b611d478185611b39565b935083602082028501611d5985611b4a565b8060005b85811015611d955784840389528151611d768582611d11565b9450611d8183611d25565b925060208a01995050600181019050611d5d565b50829750879550505050505092915050565b60006020820190508181036000830152611dc18184611d32565b905092915050565b60008060408385031215611de057611ddf611804565b5b6000611dee8582860161182f565b9250506020611dff8582860161182f565b9150509250929050565b60008060008060808587031215611e2357611e22611804565b5b6000611e318782880161182f565b9450506020611e428782880161182f565b9350506040611e538782880161182f565b9250506060611e648782880161182f565b91505092959194509250565b60008060408385031215611e8757611e86611804565b5b6000611e958582860161182f565b925050602083013567ffffffffffffffff811115611eb657611eb5611809565b5b611ec2858286016119b3565b9150509250929050565b600082825260208201905092915050565b6000611ee882611b69565b611ef28185611ecc565b9350611f02818560208601611b85565b611f0b816118a5565b840191505092915050565b600061018082019050611f2c600083018f611871565b611f39602083018e611871565b611f46604083018d611871565b8181036060830152611f58818c611edd565b90508181036080830152611f6c818b611edd565b905081810360a0830152611f80818a611edd565b9050611f8f60c0830189611871565b611f9c60e0830188611871565b818103610100830152611faf8187611edd565b9050818103610120830152611fc48186611edd565b9050611fd4610140830185611871565b611fe2610160830184611871565b9d9c50505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061205c8261180e565b91506120678361180e565b92508282101561207a57612079612022565b5b828203905092915050565b60006120908261180e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156120c3576120c2612022565b5b600182019050919050565b60006120d98261180e565b91506120e48361180e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561211d5761211c612022565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061216f57607f821691505b6020821081141561218357612182612128565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212204df80d2ae169ef07601cbef988cecbcda6c730d58740feea70ec33bef909dcc664736f6c634300080b0033"
var DrugTraceSMBin = "0x"

// DeployDrugTrace deploys a new contract, binding an instance of DrugTrace to it.
func DeployDrugTrace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *DrugTrace, error) {
	parsed, err := abi.JSON(strings.NewReader(DrugTraceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(DrugTraceSMBin)
	} else {
		bytecode = common.FromHex(DrugTraceBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, DrugTraceABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &DrugTrace{DrugTraceCaller: DrugTraceCaller{contract: contract}, DrugTraceTransactor: DrugTraceTransactor{contract: contract}, DrugTraceFilterer: DrugTraceFilterer{contract: contract}}, nil
}

func AsyncDeployDrugTrace(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(DrugTraceABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(DrugTraceSMBin)
	} else {
		bytecode = common.FromHex(DrugTraceBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, DrugTraceABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// DrugTrace is an auto generated Go binding around a Solidity contract.
type DrugTrace struct {
	DrugTraceCaller     // Read-only binding to the contract
	DrugTraceTransactor // Write-only binding to the contract
	DrugTraceFilterer   // Log filterer for contract events
}

// DrugTraceCaller is an auto generated read-only Go binding around a Solidity contract.
type DrugTraceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrugTraceTransactor is an auto generated write-only Go binding around a Solidity contract.
type DrugTraceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrugTraceFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type DrugTraceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DrugTraceSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type DrugTraceSession struct {
	Contract     *DrugTrace        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DrugTraceCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type DrugTraceCallerSession struct {
	Contract *DrugTraceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DrugTraceTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type DrugTraceTransactorSession struct {
	Contract     *DrugTraceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DrugTraceRaw is an auto generated low-level Go binding around a Solidity contract.
type DrugTraceRaw struct {
	Contract *DrugTrace // Generic contract binding to access the raw methods on
}

// DrugTraceCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type DrugTraceCallerRaw struct {
	Contract *DrugTraceCaller // Generic read-only contract binding to access the raw methods on
}

// DrugTraceTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type DrugTraceTransactorRaw struct {
	Contract *DrugTraceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDrugTrace creates a new instance of DrugTrace, bound to a specific deployed contract.
func NewDrugTrace(address common.Address, backend bind.ContractBackend) (*DrugTrace, error) {
	contract, err := bindDrugTrace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DrugTrace{DrugTraceCaller: DrugTraceCaller{contract: contract}, DrugTraceTransactor: DrugTraceTransactor{contract: contract}, DrugTraceFilterer: DrugTraceFilterer{contract: contract}}, nil
}

// NewDrugTraceCaller creates a new read-only instance of DrugTrace, bound to a specific deployed contract.
func NewDrugTraceCaller(address common.Address, caller bind.ContractCaller) (*DrugTraceCaller, error) {
	contract, err := bindDrugTrace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DrugTraceCaller{contract: contract}, nil
}

// NewDrugTraceTransactor creates a new write-only instance of DrugTrace, bound to a specific deployed contract.
func NewDrugTraceTransactor(address common.Address, transactor bind.ContractTransactor) (*DrugTraceTransactor, error) {
	contract, err := bindDrugTrace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DrugTraceTransactor{contract: contract}, nil
}

// NewDrugTraceFilterer creates a new log filterer instance of DrugTrace, bound to a specific deployed contract.
func NewDrugTraceFilterer(address common.Address, filterer bind.ContractFilterer) (*DrugTraceFilterer, error) {
	contract, err := bindDrugTrace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DrugTraceFilterer{contract: contract}, nil
}

// bindDrugTrace binds a generic wrapper to an already deployed contract.
func bindDrugTrace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DrugTraceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrugTrace *DrugTraceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DrugTrace.Contract.DrugTraceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrugTrace *DrugTraceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.DrugTraceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrugTrace *DrugTraceRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.DrugTraceTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DrugTrace *DrugTraceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DrugTrace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DrugTrace *DrugTraceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DrugTrace *DrugTraceTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// Drugs is a free data retrieval call binding the contract method 0xfca5e61e.
//
// Solidity: function drugs(uint256 ) constant returns(uint256 id, uint256 producer, uint256 productionDate, string productionDateStr, string flow, string transportationStr, uint256 dealer, uint256 drugAcceptanceTime, string drugStorageConditions, string drugStorageLocation, uint256 buyer, uint256 buyTime)
func (_DrugTrace *DrugTraceCaller) Drugs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id                    *big.Int
	Producer              *big.Int
	ProductionDate        *big.Int
	ProductionDateStr     string
	Flow                  string
	TransportationStr     string
	Dealer                *big.Int
	DrugAcceptanceTime    *big.Int
	DrugStorageConditions string
	DrugStorageLocation   string
	Buyer                 *big.Int
	BuyTime               *big.Int
}, error) {
	ret := new(struct {
		Id                    *big.Int
		Producer              *big.Int
		ProductionDate        *big.Int
		ProductionDateStr     string
		Flow                  string
		TransportationStr     string
		Dealer                *big.Int
		DrugAcceptanceTime    *big.Int
		DrugStorageConditions string
		DrugStorageLocation   string
		Buyer                 *big.Int
		BuyTime               *big.Int
	})
	out := ret
	err := _DrugTrace.contract.Call(opts, out, "drugs", arg0)
	return *ret, err
}

// Drugs is a free data retrieval call binding the contract method 0xfca5e61e.
//
// Solidity: function drugs(uint256 ) constant returns(uint256 id, uint256 producer, uint256 productionDate, string productionDateStr, string flow, string transportationStr, uint256 dealer, uint256 drugAcceptanceTime, string drugStorageConditions, string drugStorageLocation, uint256 buyer, uint256 buyTime)
func (_DrugTrace *DrugTraceSession) Drugs(arg0 *big.Int) (struct {
	Id                    *big.Int
	Producer              *big.Int
	ProductionDate        *big.Int
	ProductionDateStr     string
	Flow                  string
	TransportationStr     string
	Dealer                *big.Int
	DrugAcceptanceTime    *big.Int
	DrugStorageConditions string
	DrugStorageLocation   string
	Buyer                 *big.Int
	BuyTime               *big.Int
}, error) {
	return _DrugTrace.Contract.Drugs(&_DrugTrace.CallOpts, arg0)
}

// Drugs is a free data retrieval call binding the contract method 0xfca5e61e.
//
// Solidity: function drugs(uint256 ) constant returns(uint256 id, uint256 producer, uint256 productionDate, string productionDateStr, string flow, string transportationStr, uint256 dealer, uint256 drugAcceptanceTime, string drugStorageConditions, string drugStorageLocation, uint256 buyer, uint256 buyTime)
func (_DrugTrace *DrugTraceCallerSession) Drugs(arg0 *big.Int) (struct {
	Id                    *big.Int
	Producer              *big.Int
	ProductionDate        *big.Int
	ProductionDateStr     string
	Flow                  string
	TransportationStr     string
	Dealer                *big.Int
	DrugAcceptanceTime    *big.Int
	DrugStorageConditions string
	DrugStorageLocation   string
	Buyer                 *big.Int
	BuyTime               *big.Int
}, error) {
	return _DrugTrace.Contract.Drugs(&_DrugTrace.CallOpts, arg0)
}

// GetDrugCount is a free data retrieval call binding the contract method 0x891cde99.
//
// Solidity: function getDrugCount(uint256 _userId, uint256 _id) constant returns(uint256)
func (_DrugTrace *DrugTraceCaller) GetDrugCount(opts *bind.CallOpts, _userId *big.Int, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DrugTrace.contract.Call(opts, out, "getDrugCount", _userId, _id)
	return *ret0, err
}

// GetDrugCount is a free data retrieval call binding the contract method 0x891cde99.
//
// Solidity: function getDrugCount(uint256 _userId, uint256 _id) constant returns(uint256)
func (_DrugTrace *DrugTraceSession) GetDrugCount(_userId *big.Int, _id *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.GetDrugCount(&_DrugTrace.CallOpts, _userId, _id)
}

// GetDrugCount is a free data retrieval call binding the contract method 0x891cde99.
//
// Solidity: function getDrugCount(uint256 _userId, uint256 _id) constant returns(uint256)
func (_DrugTrace *DrugTraceCallerSession) GetDrugCount(_userId *big.Int, _id *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.GetDrugCount(&_DrugTrace.CallOpts, _userId, _id)
}

// GetDrugs is a free data retrieval call binding the contract method 0x897b3c74.
//
// Solidity: function getDrugs(uint256 _page, uint256 _pageSize, uint256 _userId, uint256 _id) constant returns([]DrugTraceDrug)
func (_DrugTrace *DrugTraceCaller) GetDrugs(opts *bind.CallOpts, _page *big.Int, _pageSize *big.Int, _userId *big.Int, _id *big.Int) ([]DrugTraceDrug, error) {
	var (
		ret0 = new([]DrugTraceDrug)
	)
	out := ret0
	err := _DrugTrace.contract.Call(opts, out, "getDrugs", _page, _pageSize, _userId, _id)
	return *ret0, err
}

// GetDrugs is a free data retrieval call binding the contract method 0x897b3c74.
//
// Solidity: function getDrugs(uint256 _page, uint256 _pageSize, uint256 _userId, uint256 _id) constant returns([]DrugTraceDrug)
func (_DrugTrace *DrugTraceSession) GetDrugs(_page *big.Int, _pageSize *big.Int, _userId *big.Int, _id *big.Int) ([]DrugTraceDrug, error) {
	return _DrugTrace.Contract.GetDrugs(&_DrugTrace.CallOpts, _page, _pageSize, _userId, _id)
}

// GetDrugs is a free data retrieval call binding the contract method 0x897b3c74.
//
// Solidity: function getDrugs(uint256 _page, uint256 _pageSize, uint256 _userId, uint256 _id) constant returns([]DrugTraceDrug)
func (_DrugTrace *DrugTraceCallerSession) GetDrugs(_page *big.Int, _pageSize *big.Int, _userId *big.Int, _id *big.Int) ([]DrugTraceDrug, error) {
	return _DrugTrace.Contract.GetDrugs(&_DrugTrace.CallOpts, _page, _pageSize, _userId, _id)
}

// GetPayableDrugCount is a free data retrieval call binding the contract method 0x154cce35.
//
// Solidity: function getPayableDrugCount(uint256 _id) constant returns(uint256)
func (_DrugTrace *DrugTraceCaller) GetPayableDrugCount(opts *bind.CallOpts, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DrugTrace.contract.Call(opts, out, "getPayableDrugCount", _id)
	return *ret0, err
}

// GetPayableDrugCount is a free data retrieval call binding the contract method 0x154cce35.
//
// Solidity: function getPayableDrugCount(uint256 _id) constant returns(uint256)
func (_DrugTrace *DrugTraceSession) GetPayableDrugCount(_id *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.GetPayableDrugCount(&_DrugTrace.CallOpts, _id)
}

// GetPayableDrugCount is a free data retrieval call binding the contract method 0x154cce35.
//
// Solidity: function getPayableDrugCount(uint256 _id) constant returns(uint256)
func (_DrugTrace *DrugTraceCallerSession) GetPayableDrugCount(_id *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.GetPayableDrugCount(&_DrugTrace.CallOpts, _id)
}

// GetPayableDrugs is a free data retrieval call binding the contract method 0x49eb063d.
//
// Solidity: function getPayableDrugs(uint256 _page, uint256 _pageSize, uint256 _id) constant returns([]DrugTraceDrug)
func (_DrugTrace *DrugTraceCaller) GetPayableDrugs(opts *bind.CallOpts, _page *big.Int, _pageSize *big.Int, _id *big.Int) ([]DrugTraceDrug, error) {
	var (
		ret0 = new([]DrugTraceDrug)
	)
	out := ret0
	err := _DrugTrace.contract.Call(opts, out, "getPayableDrugs", _page, _pageSize, _id)
	return *ret0, err
}

// GetPayableDrugs is a free data retrieval call binding the contract method 0x49eb063d.
//
// Solidity: function getPayableDrugs(uint256 _page, uint256 _pageSize, uint256 _id) constant returns([]DrugTraceDrug)
func (_DrugTrace *DrugTraceSession) GetPayableDrugs(_page *big.Int, _pageSize *big.Int, _id *big.Int) ([]DrugTraceDrug, error) {
	return _DrugTrace.Contract.GetPayableDrugs(&_DrugTrace.CallOpts, _page, _pageSize, _id)
}

// GetPayableDrugs is a free data retrieval call binding the contract method 0x49eb063d.
//
// Solidity: function getPayableDrugs(uint256 _page, uint256 _pageSize, uint256 _id) constant returns([]DrugTraceDrug)
func (_DrugTrace *DrugTraceCallerSession) GetPayableDrugs(_page *big.Int, _pageSize *big.Int, _id *big.Int) ([]DrugTraceDrug, error) {
	return _DrugTrace.Contract.GetPayableDrugs(&_DrugTrace.CallOpts, _page, _pageSize, _id)
}

// PayableDrugIdList is a free data retrieval call binding the contract method 0x44bd7bf6.
//
// Solidity: function payableDrugIdList(uint256 ) constant returns(uint256)
func (_DrugTrace *DrugTraceCaller) PayableDrugIdList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DrugTrace.contract.Call(opts, out, "payableDrugIdList", arg0)
	return *ret0, err
}

// PayableDrugIdList is a free data retrieval call binding the contract method 0x44bd7bf6.
//
// Solidity: function payableDrugIdList(uint256 ) constant returns(uint256)
func (_DrugTrace *DrugTraceSession) PayableDrugIdList(arg0 *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.PayableDrugIdList(&_DrugTrace.CallOpts, arg0)
}

// PayableDrugIdList is a free data retrieval call binding the contract method 0x44bd7bf6.
//
// Solidity: function payableDrugIdList(uint256 ) constant returns(uint256)
func (_DrugTrace *DrugTraceCallerSession) PayableDrugIdList(arg0 *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.PayableDrugIdList(&_DrugTrace.CallOpts, arg0)
}

// UserIdToDrugIdList is a free data retrieval call binding the contract method 0x9af82521.
//
// Solidity: function userIdToDrugIdList(uint256 , uint256 ) constant returns(uint256)
func (_DrugTrace *DrugTraceCaller) UserIdToDrugIdList(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DrugTrace.contract.Call(opts, out, "userIdToDrugIdList", arg0, arg1)
	return *ret0, err
}

// UserIdToDrugIdList is a free data retrieval call binding the contract method 0x9af82521.
//
// Solidity: function userIdToDrugIdList(uint256 , uint256 ) constant returns(uint256)
func (_DrugTrace *DrugTraceSession) UserIdToDrugIdList(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.UserIdToDrugIdList(&_DrugTrace.CallOpts, arg0, arg1)
}

// UserIdToDrugIdList is a free data retrieval call binding the contract method 0x9af82521.
//
// Solidity: function userIdToDrugIdList(uint256 , uint256 ) constant returns(uint256)
func (_DrugTrace *DrugTraceCallerSession) UserIdToDrugIdList(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DrugTrace.Contract.UserIdToDrugIdList(&_DrugTrace.CallOpts, arg0, arg1)
}

// Accept is a paid mutator transaction binding the contract method 0x2df74c27.
//
// Solidity: function accept(uint256 _id, string _drugStorageConditions, string _drugStorageLocation) returns()
func (_DrugTrace *DrugTraceTransactor) Accept(opts *bind.TransactOpts, _id *big.Int, _drugStorageConditions string, _drugStorageLocation string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _DrugTrace.contract.TransactWithResult(opts, out, "accept", _id, _drugStorageConditions, _drugStorageLocation)
	return transaction, receipt, err
}

func (_DrugTrace *DrugTraceTransactor) AsyncAccept(handler func(*types.Receipt, error), opts *bind.TransactOpts, _id *big.Int, _drugStorageConditions string, _drugStorageLocation string) (*types.Transaction, error) {
	return _DrugTrace.contract.AsyncTransact(opts, handler, "accept", _id, _drugStorageConditions, _drugStorageLocation)
}

// Accept is a paid mutator transaction binding the contract method 0x2df74c27.
//
// Solidity: function accept(uint256 _id, string _drugStorageConditions, string _drugStorageLocation) returns()
func (_DrugTrace *DrugTraceSession) Accept(_id *big.Int, _drugStorageConditions string, _drugStorageLocation string) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.Accept(&_DrugTrace.TransactOpts, _id, _drugStorageConditions, _drugStorageLocation)
}

func (_DrugTrace *DrugTraceSession) AsyncAccept(handler func(*types.Receipt, error), _id *big.Int, _drugStorageConditions string, _drugStorageLocation string) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncAccept(handler, &_DrugTrace.TransactOpts, _id, _drugStorageConditions, _drugStorageLocation)
}

// Accept is a paid mutator transaction binding the contract method 0x2df74c27.
//
// Solidity: function accept(uint256 _id, string _drugStorageConditions, string _drugStorageLocation) returns()
func (_DrugTrace *DrugTraceTransactorSession) Accept(_id *big.Int, _drugStorageConditions string, _drugStorageLocation string) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.Accept(&_DrugTrace.TransactOpts, _id, _drugStorageConditions, _drugStorageLocation)
}

func (_DrugTrace *DrugTraceTransactorSession) AsyncAccept(handler func(*types.Receipt, error), _id *big.Int, _drugStorageConditions string, _drugStorageLocation string) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncAccept(handler, &_DrugTrace.TransactOpts, _id, _drugStorageConditions, _drugStorageLocation)
}

// AddDrug is a paid mutator transaction binding the contract method 0xd70518c7.
//
// Solidity: function addDrug(uint256 _producer, string _prouctionDataStr) returns()
func (_DrugTrace *DrugTraceTransactor) AddDrug(opts *bind.TransactOpts, _producer *big.Int, _prouctionDataStr string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _DrugTrace.contract.TransactWithResult(opts, out, "addDrug", _producer, _prouctionDataStr)
	return transaction, receipt, err
}

func (_DrugTrace *DrugTraceTransactor) AsyncAddDrug(handler func(*types.Receipt, error), opts *bind.TransactOpts, _producer *big.Int, _prouctionDataStr string) (*types.Transaction, error) {
	return _DrugTrace.contract.AsyncTransact(opts, handler, "addDrug", _producer, _prouctionDataStr)
}

// AddDrug is a paid mutator transaction binding the contract method 0xd70518c7.
//
// Solidity: function addDrug(uint256 _producer, string _prouctionDataStr) returns()
func (_DrugTrace *DrugTraceSession) AddDrug(_producer *big.Int, _prouctionDataStr string) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.AddDrug(&_DrugTrace.TransactOpts, _producer, _prouctionDataStr)
}

func (_DrugTrace *DrugTraceSession) AsyncAddDrug(handler func(*types.Receipt, error), _producer *big.Int, _prouctionDataStr string) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncAddDrug(handler, &_DrugTrace.TransactOpts, _producer, _prouctionDataStr)
}

// AddDrug is a paid mutator transaction binding the contract method 0xd70518c7.
//
// Solidity: function addDrug(uint256 _producer, string _prouctionDataStr) returns()
func (_DrugTrace *DrugTraceTransactorSession) AddDrug(_producer *big.Int, _prouctionDataStr string) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.AddDrug(&_DrugTrace.TransactOpts, _producer, _prouctionDataStr)
}

func (_DrugTrace *DrugTraceTransactorSession) AsyncAddDrug(handler func(*types.Receipt, error), _producer *big.Int, _prouctionDataStr string) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncAddDrug(handler, &_DrugTrace.TransactOpts, _producer, _prouctionDataStr)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 _id, uint256 _buyer) returns()
func (_DrugTrace *DrugTraceTransactor) Buy(opts *bind.TransactOpts, _id *big.Int, _buyer *big.Int) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _DrugTrace.contract.TransactWithResult(opts, out, "buy", _id, _buyer)
	return transaction, receipt, err
}

func (_DrugTrace *DrugTraceTransactor) AsyncBuy(handler func(*types.Receipt, error), opts *bind.TransactOpts, _id *big.Int, _buyer *big.Int) (*types.Transaction, error) {
	return _DrugTrace.contract.AsyncTransact(opts, handler, "buy", _id, _buyer)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 _id, uint256 _buyer) returns()
func (_DrugTrace *DrugTraceSession) Buy(_id *big.Int, _buyer *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.Buy(&_DrugTrace.TransactOpts, _id, _buyer)
}

func (_DrugTrace *DrugTraceSession) AsyncBuy(handler func(*types.Receipt, error), _id *big.Int, _buyer *big.Int) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncBuy(handler, &_DrugTrace.TransactOpts, _id, _buyer)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 _id, uint256 _buyer) returns()
func (_DrugTrace *DrugTraceTransactorSession) Buy(_id *big.Int, _buyer *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.Buy(&_DrugTrace.TransactOpts, _id, _buyer)
}

func (_DrugTrace *DrugTraceTransactorSession) AsyncBuy(handler func(*types.Receipt, error), _id *big.Int, _buyer *big.Int) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncBuy(handler, &_DrugTrace.TransactOpts, _id, _buyer)
}

// Sale is a paid mutator transaction binding the contract method 0x327af74c.
//
// Solidity: function sale(uint256 _id, uint256 _dealer, string _transportationStr) returns()
func (_DrugTrace *DrugTraceTransactor) Sale(opts *bind.TransactOpts, _id *big.Int, _dealer *big.Int, _transportationStr string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _DrugTrace.contract.TransactWithResult(opts, out, "sale", _id, _dealer, _transportationStr)
	return transaction, receipt, err
}

func (_DrugTrace *DrugTraceTransactor) AsyncSale(handler func(*types.Receipt, error), opts *bind.TransactOpts, _id *big.Int, _dealer *big.Int, _transportationStr string) (*types.Transaction, error) {
	return _DrugTrace.contract.AsyncTransact(opts, handler, "sale", _id, _dealer, _transportationStr)
}

// Sale is a paid mutator transaction binding the contract method 0x327af74c.
//
// Solidity: function sale(uint256 _id, uint256 _dealer, string _transportationStr) returns()
func (_DrugTrace *DrugTraceSession) Sale(_id *big.Int, _dealer *big.Int, _transportationStr string) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.Sale(&_DrugTrace.TransactOpts, _id, _dealer, _transportationStr)
}

func (_DrugTrace *DrugTraceSession) AsyncSale(handler func(*types.Receipt, error), _id *big.Int, _dealer *big.Int, _transportationStr string) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncSale(handler, &_DrugTrace.TransactOpts, _id, _dealer, _transportationStr)
}

// Sale is a paid mutator transaction binding the contract method 0x327af74c.
//
// Solidity: function sale(uint256 _id, uint256 _dealer, string _transportationStr) returns()
func (_DrugTrace *DrugTraceTransactorSession) Sale(_id *big.Int, _dealer *big.Int, _transportationStr string) (*types.Transaction, *types.Receipt, error) {
	return _DrugTrace.Contract.Sale(&_DrugTrace.TransactOpts, _id, _dealer, _transportationStr)
}

func (_DrugTrace *DrugTraceTransactorSession) AsyncSale(handler func(*types.Receipt, error), _id *big.Int, _dealer *big.Int, _transportationStr string) (*types.Transaction, error) {
	return _DrugTrace.Contract.AsyncSale(handler, &_DrugTrace.TransactOpts, _id, _dealer, _transportationStr)
}
