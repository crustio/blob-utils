// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zkblob

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ZkBlobMetaData contains all meta data concerning the ZkBlob contract.
var ZkBlobMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zkblobAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostBatchCountNotMatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"}],\"name\":\"RepeatedBatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"}],\"name\":\"WrongBatchNum\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"WrongSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZkBlobAddressNotSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZKBLOB_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"batchNum\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"root\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"postBatchs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zkblobAddress\",\"type\":\"address\"}],\"name\":\"setZkBlobAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"verifyBlob\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zkblobAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620021e4380380620021e483398181016040528101906200003691906200027d565b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200008a5f801b33620000c360201b60201c565b620000bc7fbb8b631c09334a4c0dcfa286c85022656a6c62b488a8c6c6b66f2b96106d830082620000c360201b60201c565b50620002ad565b620000d58282620001ae60201b60201c565b620001aa5760015f808481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506200014f6200021160201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b5f805f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000247826200021c565b9050919050565b62000259816200023b565b811462000264575f80fd5b50565b5f8151905062000277816200024e565b92915050565b5f6020828403121562000295576200029462000218565b5b5f620002a48482850162000267565b91505092915050565b611f2980620002bb5f395ff3fe608060405234801561000f575f80fd5b50600436106100cd575f3560e01c806391d148541161008a578063c2b40ae411610064578063c2b40ae41461021f578063c3bf40d71461024f578063d547741f1461026d578063ec1cb4dc14610289576100cd565b806391d14854146101a15780639755fd6c146101d1578063a217fddf14610201576100cd565b806301ffc9a7146100d157806320dc787914610101578063248a9ca31461011d5780632f2ff15d1461014d57806336568abe146101695780636b8d130914610185575b5f80fd5b6100eb60048036038101906100e6919061112f565b6102a7565b6040516100f89190611174565b60405180910390f35b61011b600480360381019061011691906111e7565b610320565b005b61013760048036038101906101329190611245565b6103e5565b604051610144919061127f565b60405180910390f35b61016760048036038101906101629190611298565b610401565b005b610183600480360381019061017e9190611298565b610422565b005b61019f600480360381019061019a91906115c9565b6104a5565b005b6101bb60048036038101906101b69190611298565b6106ee565b6040516101c89190611174565b60405180910390f35b6101eb60048036038101906101e691906116c6565b610751565b6040516101f89190611174565b60405180910390f35b6102096107d0565b604051610216919061127f565b60405180910390f35b61023960048036038101906102349190611737565b6107d6565b604051610246919061127f565b60405180910390f35b6102576107eb565b604051610264919061127f565b60405180910390f35b61028760048036038101906102829190611298565b61080f565b005b610291610830565b60405161029e9190611771565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610319575061031882610855565b5b9050919050565b5f801b61032c816108be565b8160015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103b77fbb8b631c09334a4c0dcfa286c85022656a6c62b488a8c6c6b66f2b96106d830060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166108d2565b6103e17fbb8b631c09334a4c0dcfa286c85022656a6c62b488a8c6c6b66f2b96106d8300836109ac565b5050565b5f805f8381526020019081526020015f20600101549050919050565b61040a826103e5565b610413816108be565b61041d83836109ac565b505050565b61042a610a86565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610497576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048e9061180a565b60405180910390fd5b6104a182826108d2565b5050565b7fbb8b631c09334a4c0dcfa286c85022656a6c62b488a8c6c6b66f2b96106d83006104cf816108be565b825184511461050a576040517f8c6797b300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b84518110156105b3575f801b60025f87848151811061052e5761052d611828565b5b602002602001015181526020019081526020015f2054146105a05784818151811061055c5761055b611828565b5b60200260200101516040517f276215340000000000000000000000000000000000000000000000000000000081526004016105979190611864565b60405180910390fd5b80806105ab906118aa565b91505061050c565b505f84846040516020016105c8929190611a53565b6040516020818303038152906040528051906020012090505f6105eb8285610a8d565b905060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461067d57806040517e414a420000000000000000000000000000000000000000000000000000000081526004016106749190611771565b60405180910390fd5b5f5b86518110156106e55785818151811061069b5761069a611828565b5b602002602001015160025f8984815181106106b9576106b8611828565b5b602002602001015181526020019081526020015f208190555080806106dd906118aa565b91505061067f565b50505050505050565b5f805f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f8060025f8781526020019081526020015f205490505f801b8103610779575f9150506107c8565b6107c48484808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f820116905080830192505050505050508287610ab2565b9150505b949350505050565b5f801b81565b6002602052805f5260405f205f915090505481565b7fbb8b631c09334a4c0dcfa286c85022656a6c62b488a8c6c6b66f2b96106d830081565b610818826103e5565b610821816108be565b61082b83836108d2565b505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6108cf816108ca610a86565b610ac8565b50565b6108dc82826106ee565b156109a8575f805f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff02191690831515021790555061094d610a86565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6109b682826106ee565b610a825760015f808481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610a27610a86565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b5f33905090565b5f805f610a9a8585610b4c565b91509150610aa781610b98565b819250505092915050565b5f82610abe8584610cfd565b1490509392505050565b610ad282826106ee565b610b4857610adf81610d51565b610aec835f1c6020610d7e565b604051602001610afd929190611b76565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3f9190611be7565b60405180910390fd5b5050565b5f806041835103610b89575f805f602086015192506040860151915060608601515f1a9050610b7d87828585610fb3565b94509450505050610b91565b5f6002915091505b9250929050565b5f6004811115610bab57610baa611c07565b5b816004811115610bbe57610bbd611c07565b5b0315610cfa5760016004811115610bd857610bd7611c07565b5b816004811115610beb57610bea611c07565b5b03610c2b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2290611c7e565b60405180910390fd5b60026004811115610c3f57610c3e611c07565b5b816004811115610c5257610c51611c07565b5b03610c92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8990611ce6565b60405180910390fd5b60036004811115610ca657610ca5611c07565b5b816004811115610cb957610cb8611c07565b5b03610cf9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf090611d74565b60405180910390fd5b5b50565b5f808290505f5b8451811015610d4657610d3182868381518110610d2457610d23611828565b5b602002602001015161108b565b91508080610d3e906118aa565b915050610d04565b508091505092915050565b6060610d778273ffffffffffffffffffffffffffffffffffffffff16601460ff16610d7e565b9050919050565b60605f6002836002610d909190611d92565b610d9a9190611dd3565b67ffffffffffffffff811115610db357610db26112ea565b5b6040519080825280601f01601f191660200182016040528015610de55781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000815f81518110610e1c57610e1b611828565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610e7f57610e7e611828565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a9053505f6001846002610ebd9190611d92565b610ec79190611dd3565b90505b6001811115610f66577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110610f0957610f08611828565b5b1a60f81b828281518110610f2057610f1f611828565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a905350600485901c945080610f5f90611e06565b9050610eca565b505f8414610fa9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa090611e77565b60405180910390fd5b8091505092915050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0835f1c1115610feb575f600391509150611082565b5f6001878787876040515f815260200160405260405161100e9493929190611eb0565b6020604051602081039080840390855afa15801561102e573d5f803e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361107a575f60019250925050611082565b805f92509250505b94509492505050565b5f8183106110a25761109d82846110b5565b6110ad565b6110ac83836110b5565b5b905092915050565b5f825f528160205260405f20905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61110e816110da565b8114611118575f80fd5b50565b5f8135905061112981611105565b92915050565b5f60208284031215611144576111436110d2565b5b5f6111518482850161111b565b91505092915050565b5f8115159050919050565b61116e8161115a565b82525050565b5f6020820190506111875f830184611165565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6111b68261118d565b9050919050565b6111c6816111ac565b81146111d0575f80fd5b50565b5f813590506111e1816111bd565b92915050565b5f602082840312156111fc576111fb6110d2565b5b5f611209848285016111d3565b91505092915050565b5f819050919050565b61122481611212565b811461122e575f80fd5b50565b5f8135905061123f8161121b565b92915050565b5f6020828403121561125a576112596110d2565b5b5f61126784828501611231565b91505092915050565b61127981611212565b82525050565b5f6020820190506112925f830184611270565b92915050565b5f80604083850312156112ae576112ad6110d2565b5b5f6112bb85828601611231565b92505060206112cc858286016111d3565b9150509250929050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611320826112da565b810181811067ffffffffffffffff8211171561133f5761133e6112ea565b5b80604052505050565b5f6113516110c9565b905061135d8282611317565b919050565b5f67ffffffffffffffff82111561137c5761137b6112ea565b5b602082029050602081019050919050565b5f80fd5b5f819050919050565b6113a381611391565b81146113ad575f80fd5b50565b5f813590506113be8161139a565b92915050565b5f6113d66113d184611362565b611348565b905080838252602082019050602084028301858111156113f9576113f861138d565b5b835b81811015611422578061140e88826113b0565b8452602084019350506020810190506113fb565b5050509392505050565b5f82601f8301126114405761143f6112d6565b5b81356114508482602086016113c4565b91505092915050565b5f67ffffffffffffffff821115611473576114726112ea565b5b602082029050602081019050919050565b5f61149661149184611459565b611348565b905080838252602082019050602084028301858111156114b9576114b861138d565b5b835b818110156114e257806114ce8882611231565b8452602084019350506020810190506114bb565b5050509392505050565b5f82601f830112611500576114ff6112d6565b5b8135611510848260208601611484565b91505092915050565b5f80fd5b5f67ffffffffffffffff821115611537576115366112ea565b5b611540826112da565b9050602081019050919050565b828183375f83830152505050565b5f61156d6115688461151d565b611348565b90508281526020810184848401111561158957611588611519565b5b61159484828561154d565b509392505050565b5f82601f8301126115b0576115af6112d6565b5b81356115c084826020860161155b565b91505092915050565b5f805f606084860312156115e0576115df6110d2565b5b5f84013567ffffffffffffffff8111156115fd576115fc6110d6565b5b6116098682870161142c565b935050602084013567ffffffffffffffff81111561162a576116296110d6565b5b611636868287016114ec565b925050604084013567ffffffffffffffff811115611657576116566110d6565b5b6116638682870161159c565b9150509250925092565b5f80fd5b5f8083601f840112611686576116856112d6565b5b8235905067ffffffffffffffff8111156116a3576116a261166d565b5b6020830191508360208202830111156116bf576116be61138d565b5b9250929050565b5f805f80606085870312156116de576116dd6110d2565b5b5f6116eb878288016113b0565b94505060206116fc87828801611231565b935050604085013567ffffffffffffffff81111561171d5761171c6110d6565b5b61172987828801611671565b925092505092959194509250565b5f6020828403121561174c5761174b6110d2565b5b5f611759848285016113b0565b91505092915050565b61176b816111ac565b82525050565b5f6020820190506117845f830184611762565b92915050565b5f82825260208201905092915050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e63655f8201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b5f6117f4602f8361178a565b91506117ff8261179a565b604082019050919050565b5f6020820190508181035f830152611821816117e8565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b61185e81611391565b82525050565b5f6020820190506118775f830184611855565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6118b482611391565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118e6576118e561187d565b5b600182019050919050565b5f81519050919050565b5f81905092915050565b5f819050602082019050919050565b61191d81611391565b82525050565b5f61192e8383611914565b60208301905092915050565b5f602082019050919050565b5f611950826118f1565b61195a81856118fb565b935061196583611905565b805f5b8381101561199557815161197c8882611923565b97506119878361193a565b925050600181019050611968565b5085935050505092915050565b5f81519050919050565b5f81905092915050565b5f819050602082019050919050565b6119ce81611212565b82525050565b5f6119df83836119c5565b60208301905092915050565b5f602082019050919050565b5f611a01826119a2565b611a0b81856119ac565b9350611a16836119b6565b805f5b83811015611a46578151611a2d88826119d4565b9750611a38836119eb565b925050600181019050611a19565b5085935050505092915050565b5f611a5e8285611946565b9150611a6a82846119f7565b91508190509392505050565b5f81905092915050565b7f416363657373436f6e74726f6c3a206163636f756e74200000000000000000005f82015250565b5f611ab4601783611a76565b9150611abf82611a80565b601782019050919050565b5f81519050919050565b5f5b83811015611af1578082015181840152602081019050611ad6565b5f8484015250505050565b5f611b0682611aca565b611b108185611a76565b9350611b20818560208601611ad4565b80840191505092915050565b7f206973206d697373696e6720726f6c65200000000000000000000000000000005f82015250565b5f611b60601183611a76565b9150611b6b82611b2c565b601182019050919050565b5f611b8082611aa8565b9150611b8c8285611afc565b9150611b9782611b54565b9150611ba38284611afc565b91508190509392505050565b5f611bb982611aca565b611bc3818561178a565b9350611bd3818560208601611ad4565b611bdc816112da565b840191505092915050565b5f6020820190508181035f830152611bff8184611baf565b905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b7f45434453413a20696e76616c6964207369676e617475726500000000000000005f82015250565b5f611c6860188361178a565b9150611c7382611c34565b602082019050919050565b5f6020820190508181035f830152611c9581611c5c565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e677468005f82015250565b5f611cd0601f8361178a565b9150611cdb82611c9c565b602082019050919050565b5f6020820190508181035f830152611cfd81611cc4565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c5f8201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b5f611d5e60228361178a565b9150611d6982611d04565b604082019050919050565b5f6020820190508181035f830152611d8b81611d52565b9050919050565b5f611d9c82611391565b9150611da783611391565b9250828202611db581611391565b91508282048414831517611dcc57611dcb61187d565b5b5092915050565b5f611ddd82611391565b9150611de883611391565b9250828201905080821115611e0057611dff61187d565b5b92915050565b5f611e1082611391565b91505f8203611e2257611e2161187d565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e745f82015250565b5f611e6160208361178a565b9150611e6c82611e2d565b602082019050919050565b5f6020820190508181035f830152611e8e81611e55565b9050919050565b5f60ff82169050919050565b611eaa81611e95565b82525050565b5f608082019050611ec35f830187611270565b611ed06020830186611ea1565b611edd6040830185611270565b611eea6060830184611270565b9594505050505056fea2646970667358221220b863d447dd072c3fd2b443dacdec953ffe4e047150851f1fdb0af8ea1a4d6a9d64736f6c63430008140033",
}

// ZkBlobABI is the input ABI used to generate the binding from.
// Deprecated: Use ZkBlobMetaData.ABI instead.
var ZkBlobABI = ZkBlobMetaData.ABI

// ZkBlobBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ZkBlobMetaData.Bin instead.
var ZkBlobBin = ZkBlobMetaData.Bin

// DeployZkBlob deploys a new Ethereum contract, binding an instance of ZkBlob to it.
func DeployZkBlob(auth *bind.TransactOpts, backend bind.ContractBackend, _zkblobAddress common.Address) (common.Address, *types.Transaction, *ZkBlob, error) {
	parsed, err := ZkBlobMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ZkBlobBin), backend, _zkblobAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ZkBlob{ZkBlobCaller: ZkBlobCaller{contract: contract}, ZkBlobTransactor: ZkBlobTransactor{contract: contract}, ZkBlobFilterer: ZkBlobFilterer{contract: contract}}, nil
}

// ZkBlob is an auto generated Go binding around an Ethereum contract.
type ZkBlob struct {
	ZkBlobCaller     // Read-only binding to the contract
	ZkBlobTransactor // Write-only binding to the contract
	ZkBlobFilterer   // Log filterer for contract events
}

// ZkBlobCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZkBlobCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBlobTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZkBlobTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBlobFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZkBlobFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZkBlobSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZkBlobSession struct {
	Contract     *ZkBlob           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBlobCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZkBlobCallerSession struct {
	Contract *ZkBlobCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ZkBlobTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZkBlobTransactorSession struct {
	Contract     *ZkBlobTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZkBlobRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZkBlobRaw struct {
	Contract *ZkBlob // Generic contract binding to access the raw methods on
}

// ZkBlobCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZkBlobCallerRaw struct {
	Contract *ZkBlobCaller // Generic read-only contract binding to access the raw methods on
}

// ZkBlobTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZkBlobTransactorRaw struct {
	Contract *ZkBlobTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZkBlob creates a new instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlob(address common.Address, backend bind.ContractBackend) (*ZkBlob, error) {
	contract, err := bindZkBlob(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZkBlob{ZkBlobCaller: ZkBlobCaller{contract: contract}, ZkBlobTransactor: ZkBlobTransactor{contract: contract}, ZkBlobFilterer: ZkBlobFilterer{contract: contract}}, nil
}

// NewZkBlobCaller creates a new read-only instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlobCaller(address common.Address, caller bind.ContractCaller) (*ZkBlobCaller, error) {
	contract, err := bindZkBlob(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBlobCaller{contract: contract}, nil
}

// NewZkBlobTransactor creates a new write-only instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlobTransactor(address common.Address, transactor bind.ContractTransactor) (*ZkBlobTransactor, error) {
	contract, err := bindZkBlob(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZkBlobTransactor{contract: contract}, nil
}

// NewZkBlobFilterer creates a new log filterer instance of ZkBlob, bound to a specific deployed contract.
func NewZkBlobFilterer(address common.Address, filterer bind.ContractFilterer) (*ZkBlobFilterer, error) {
	contract, err := bindZkBlob(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZkBlobFilterer{contract: contract}, nil
}

// bindZkBlob binds a generic wrapper to an already deployed contract.
func bindZkBlob(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZkBlobMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBlob *ZkBlobRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBlob.Contract.ZkBlobCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBlob *ZkBlobRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBlob.Contract.ZkBlobTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBlob *ZkBlobRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBlob.Contract.ZkBlobTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZkBlob *ZkBlobCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZkBlob.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZkBlob *ZkBlobTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZkBlob.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZkBlob *ZkBlobTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZkBlob.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZkBlob.Contract.DEFAULTADMINROLE(&_ZkBlob.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZkBlob.Contract.DEFAULTADMINROLE(&_ZkBlob.CallOpts)
}

// ZKBLOBROLE is a free data retrieval call binding the contract method 0xc3bf40d7.
//
// Solidity: function ZKBLOB_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) ZKBLOBROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "ZKBLOB_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ZKBLOBROLE is a free data retrieval call binding the contract method 0xc3bf40d7.
//
// Solidity: function ZKBLOB_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobSession) ZKBLOBROLE() ([32]byte, error) {
	return _ZkBlob.Contract.ZKBLOBROLE(&_ZkBlob.CallOpts)
}

// ZKBLOBROLE is a free data retrieval call binding the contract method 0xc3bf40d7.
//
// Solidity: function ZKBLOB_ROLE() view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) ZKBLOBROLE() ([32]byte, error) {
	return _ZkBlob.Contract.ZKBLOBROLE(&_ZkBlob.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkBlob *ZkBlobSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZkBlob.Contract.GetRoleAdmin(&_ZkBlob.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZkBlob.Contract.GetRoleAdmin(&_ZkBlob.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkBlob *ZkBlobCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkBlob *ZkBlobSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZkBlob.Contract.HasRole(&_ZkBlob.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZkBlob *ZkBlobCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZkBlob.Contract.HasRole(&_ZkBlob.CallOpts, role, account)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ZkBlob *ZkBlobCaller) Roots(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "roots", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ZkBlob *ZkBlobSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _ZkBlob.Contract.Roots(&_ZkBlob.CallOpts, arg0)
}

// Roots is a free data retrieval call binding the contract method 0xc2b40ae4.
//
// Solidity: function roots(uint256 ) view returns(bytes32)
func (_ZkBlob *ZkBlobCallerSession) Roots(arg0 *big.Int) ([32]byte, error) {
	return _ZkBlob.Contract.Roots(&_ZkBlob.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkBlob *ZkBlobCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkBlob *ZkBlobSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZkBlob.Contract.SupportsInterface(&_ZkBlob.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZkBlob *ZkBlobCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZkBlob.Contract.SupportsInterface(&_ZkBlob.CallOpts, interfaceId)
}

// VerifyBlob is a free data retrieval call binding the contract method 0x9755fd6c.
//
// Solidity: function verifyBlob(uint256 batchNum, bytes32 hash, bytes32[] merkleProof) view returns(bool)
func (_ZkBlob *ZkBlobCaller) VerifyBlob(opts *bind.CallOpts, batchNum *big.Int, hash [32]byte, merkleProof [][32]byte) (bool, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "verifyBlob", batchNum, hash, merkleProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyBlob is a free data retrieval call binding the contract method 0x9755fd6c.
//
// Solidity: function verifyBlob(uint256 batchNum, bytes32 hash, bytes32[] merkleProof) view returns(bool)
func (_ZkBlob *ZkBlobSession) VerifyBlob(batchNum *big.Int, hash [32]byte, merkleProof [][32]byte) (bool, error) {
	return _ZkBlob.Contract.VerifyBlob(&_ZkBlob.CallOpts, batchNum, hash, merkleProof)
}

// VerifyBlob is a free data retrieval call binding the contract method 0x9755fd6c.
//
// Solidity: function verifyBlob(uint256 batchNum, bytes32 hash, bytes32[] merkleProof) view returns(bool)
func (_ZkBlob *ZkBlobCallerSession) VerifyBlob(batchNum *big.Int, hash [32]byte, merkleProof [][32]byte) (bool, error) {
	return _ZkBlob.Contract.VerifyBlob(&_ZkBlob.CallOpts, batchNum, hash, merkleProof)
}

// ZkblobAddress is a free data retrieval call binding the contract method 0xec1cb4dc.
//
// Solidity: function zkblobAddress() view returns(address)
func (_ZkBlob *ZkBlobCaller) ZkblobAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZkBlob.contract.Call(opts, &out, "zkblobAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZkblobAddress is a free data retrieval call binding the contract method 0xec1cb4dc.
//
// Solidity: function zkblobAddress() view returns(address)
func (_ZkBlob *ZkBlobSession) ZkblobAddress() (common.Address, error) {
	return _ZkBlob.Contract.ZkblobAddress(&_ZkBlob.CallOpts)
}

// ZkblobAddress is a free data retrieval call binding the contract method 0xec1cb4dc.
//
// Solidity: function zkblobAddress() view returns(address)
func (_ZkBlob *ZkBlobCallerSession) ZkblobAddress() (common.Address, error) {
	return _ZkBlob.Contract.ZkblobAddress(&_ZkBlob.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.GrantRole(&_ZkBlob.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.GrantRole(&_ZkBlob.TransactOpts, role, account)
}

// PostBatchs is a paid mutator transaction binding the contract method 0x6b8d1309.
//
// Solidity: function postBatchs(uint256[] batchNum, bytes32[] root, bytes signature) returns()
func (_ZkBlob *ZkBlobTransactor) PostBatchs(opts *bind.TransactOpts, batchNum []*big.Int, root [][32]byte, signature []byte) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "postBatchs", batchNum, root, signature)
}

// PostBatchs is a paid mutator transaction binding the contract method 0x6b8d1309.
//
// Solidity: function postBatchs(uint256[] batchNum, bytes32[] root, bytes signature) returns()
func (_ZkBlob *ZkBlobSession) PostBatchs(batchNum []*big.Int, root [][32]byte, signature []byte) (*types.Transaction, error) {
	return _ZkBlob.Contract.PostBatchs(&_ZkBlob.TransactOpts, batchNum, root, signature)
}

// PostBatchs is a paid mutator transaction binding the contract method 0x6b8d1309.
//
// Solidity: function postBatchs(uint256[] batchNum, bytes32[] root, bytes signature) returns()
func (_ZkBlob *ZkBlobTransactorSession) PostBatchs(batchNum []*big.Int, root [][32]byte, signature []byte) (*types.Transaction, error) {
	return _ZkBlob.Contract.PostBatchs(&_ZkBlob.TransactOpts, batchNum, root, signature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RenounceRole(&_ZkBlob.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RenounceRole(&_ZkBlob.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RevokeRole(&_ZkBlob.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZkBlob *ZkBlobTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.RevokeRole(&_ZkBlob.TransactOpts, role, account)
}

// SetZkBlobAddress is a paid mutator transaction binding the contract method 0x20dc7879.
//
// Solidity: function setZkBlobAddress(address _zkblobAddress) returns()
func (_ZkBlob *ZkBlobTransactor) SetZkBlobAddress(opts *bind.TransactOpts, _zkblobAddress common.Address) (*types.Transaction, error) {
	return _ZkBlob.contract.Transact(opts, "setZkBlobAddress", _zkblobAddress)
}

// SetZkBlobAddress is a paid mutator transaction binding the contract method 0x20dc7879.
//
// Solidity: function setZkBlobAddress(address _zkblobAddress) returns()
func (_ZkBlob *ZkBlobSession) SetZkBlobAddress(_zkblobAddress common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.SetZkBlobAddress(&_ZkBlob.TransactOpts, _zkblobAddress)
}

// SetZkBlobAddress is a paid mutator transaction binding the contract method 0x20dc7879.
//
// Solidity: function setZkBlobAddress(address _zkblobAddress) returns()
func (_ZkBlob *ZkBlobTransactorSession) SetZkBlobAddress(_zkblobAddress common.Address) (*types.Transaction, error) {
	return _ZkBlob.Contract.SetZkBlobAddress(&_ZkBlob.TransactOpts, _zkblobAddress)
}

// ZkBlobRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZkBlob contract.
type ZkBlobRoleAdminChangedIterator struct {
	Event *ZkBlobRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBlobRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBlobRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBlobRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBlobRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBlobRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBlobRoleAdminChanged represents a RoleAdminChanged event raised by the ZkBlob contract.
type ZkBlobRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkBlob *ZkBlobFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZkBlobRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZkBlob.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZkBlobRoleAdminChangedIterator{contract: _ZkBlob.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkBlob *ZkBlobFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZkBlobRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ZkBlob.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBlobRoleAdminChanged)
				if err := _ZkBlob.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZkBlob *ZkBlobFilterer) ParseRoleAdminChanged(log types.Log) (*ZkBlobRoleAdminChanged, error) {
	event := new(ZkBlobRoleAdminChanged)
	if err := _ZkBlob.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBlobRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZkBlob contract.
type ZkBlobRoleGrantedIterator struct {
	Event *ZkBlobRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBlobRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBlobRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBlobRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBlobRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBlobRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBlobRoleGranted represents a RoleGranted event raised by the ZkBlob contract.
type ZkBlobRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZkBlobRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkBlobRoleGrantedIterator{contract: _ZkBlob.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZkBlobRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBlobRoleGranted)
				if err := _ZkBlob.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) ParseRoleGranted(log types.Log) (*ZkBlobRoleGranted, error) {
	event := new(ZkBlobRoleGranted)
	if err := _ZkBlob.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZkBlobRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZkBlob contract.
type ZkBlobRoleRevokedIterator struct {
	Event *ZkBlobRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ZkBlobRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZkBlobRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ZkBlobRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ZkBlobRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZkBlobRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZkBlobRoleRevoked represents a RoleRevoked event raised by the ZkBlob contract.
type ZkBlobRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZkBlobRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZkBlobRoleRevokedIterator{contract: _ZkBlob.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZkBlobRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ZkBlob.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZkBlobRoleRevoked)
				if err := _ZkBlob.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZkBlob *ZkBlobFilterer) ParseRoleRevoked(log types.Log) (*ZkBlobRoleRevoked, error) {
	event := new(ZkBlobRoleRevoked)
	if err := _ZkBlob.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
