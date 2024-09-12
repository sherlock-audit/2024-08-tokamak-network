// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const UniversalRouterStorageLayoutJSON = "{\"storage\":[{\"astId\":4790,\"contract\":\"contracts/UniversalRouter.sol:UniversalRouter\",\"label\":\"maxAmountInCached\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"},{\"astId\":2489,\"contract\":\"contracts/UniversalRouter.sol:UniversalRouter\",\"label\":\"lockedBy\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var UniversalRouterStorageLayout = new(solc.StorageLayout)

var UniversalRouterDeployedBin = "0x60a0604081815260049081361015610022575b505050361561002057600080fd5b005b600092833560e01c90816301ffc9a71461093d57508063150b7a02146108af57806324856bc3146107e85780633593564c146106b1578063709a1cc21461044f578063bc197c811461038a578063f23a6e61146102f95763fa461e330361001257346102f55760607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102f557813590602435926044359067ffffffffffffffff918281116102f1576100db9036908301610a97565b919092878613908115806102e7575b6102bf5783850186868203126102bb5785359182116102bb5761010e9186016136d0565b5060208401359373ffffffffffffffffffffffffffffffffffffffff938486168096036102bb5761013e9161415a565b959097602b89106102935786359260178460601c98019561016d62ffffff883560601c9660481c16868b614365565b3391160361026b571561026157508186105b15610197575050505061019493503391613ac2565b80f35b9395945091929091906042871061021b5750505083601711610217577f8000000000000000000000000000000000000000000000000000000000000000821015610217577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe961021194019161020c33916141b5565b6141e2565b50505080f35b8480fd5b91969550929391508454841161023957506101949394503391613ac2565b8590517f739dbe52000000000000000000000000000000000000000000000000000000008152fd5b965085821061017f565b8483517f32b13d91000000000000000000000000000000000000000000000000000000008152fd5b8382517f3b99b53d000000000000000000000000000000000000000000000000000000008152fd5b8980fd5b8286517f316cf0eb000000000000000000000000000000000000000000000000000000008152fd5b50888813156100ea565b8680fd5b8280fd5b5082346103875760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261038757610332610a2b565b5061033b610a53565b506084359067ffffffffffffffff8211610387575060209261035f91369101610a97565b5050517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b80fd5b5082346103875760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610387576103c3610a2b565b506103cc610a53565b5067ffffffffffffffff9060443582811161044b576103ee9036908601610ac5565b505060643582811161044b576104079036908601610ac5565b5050608435918211610387575060209261042391369101610a97565b5050517fbc197c81000000000000000000000000000000000000000000000000000000008152f35b5080fd5b50346102f557602090817ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126106ad5783833567ffffffffffffffff811161044b576104a1829136908701610a97565b90818551928392833781018381520390827f00000000000000000000000000000000000000000000000000000000000000005af16104dd613675565b50156106855780517f70a082310000000000000000000000000000000000000000000000000000000081523084820152907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168383602481845afa92831561067b578693610646575b5081517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169581019586526020860184905294849186918290899082906040015b03925af193841561063c577f1e8f03f716bc104bf7d728131967a0c771e85ab54d09c1e2d6ed9e0bc4e2a16c9461060f575b5051908152a180f35b61062e90843d8611610635575b61062681836135fa565b81019061388d565b5038610606565b503d61061c565b81513d87823e3d90fd5b9092508381813d8311610674575b61065e81836135fa565b810103126106705751916105d461055b565b8580fd5b503d610654565b82513d88823e3d90fd5b9050517f7d529919000000000000000000000000000000000000000000000000000000008152fd5b8380fd5b5060607ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102f55767ffffffffffffffff8235818111610217576106fb9036908501610a97565b91602435908111610670576107139036908601610ac5565b92909160443542116107c0573330146107b1576001958654958773ffffffffffffffffffffffffffffffffffffffff88160361078b5750509185949391610782937fffffffffffffffffffffffff00000000000000000000000000000000000000009586339116178755610b54565b81541617905580f35b517f6f5ffb7e000000000000000000000000000000000000000000000000000000008152fd5b90919293506101949450610b54565b8585517f5bf6f916000000000000000000000000000000000000000000000000000000008152fd5b50807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102f55767ffffffffffffffff8235818111610217576108319036908501610a97565b91602435908111610670576108499036908601610ac5565b9290913330146107b1576001958654958773ffffffffffffffffffffffffffffffffffffffff88160361078b5750509185949391610782937fffffffffffffffffffffffff00000000000000000000000000000000000000009586339116178755610b54565b5082346103875760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610387576108e8610a2b565b506108f1610a53565b506064359067ffffffffffffffff8211610387575060209261091591369101610a97565b5050517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b849084346102f55760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102f557357fffffffff0000000000000000000000000000000000000000000000000000000081168091036102f557602092507f4e2312e0000000000000000000000000000000000000000000000000000000008114908115610a01575b81156109d7575b5015158152f35b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836109d0565b7f150b7a0200000000000000000000000000000000000000000000000000000000811491506109c9565b6004359073ffffffffffffffffffffffffffffffffffffffff82168203610a4e57565b600080fd5b6024359073ffffffffffffffffffffffffffffffffffffffff82168203610a4e57565b359073ffffffffffffffffffffffffffffffffffffffff82168203610a4e57565b9181601f84011215610a4e5782359167ffffffffffffffff8311610a4e5760208381860195010111610a4e57565b9181601f84011215610a4e5782359167ffffffffffffffff8311610a4e576020808501948460051b010111610a4e57565b919082519283825260005b848110610b405750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b602081830181015184830182015201610b01565b9192909260805282810361350d5791906000905b828210610b755750505050565b8382959394951015611b4c5760059282841b60805101357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe19182608051360301821215610a4e578160805101359767ffffffffffffffff8911610a4e576020836080510101988036038a13610a4e57606097603f90818989013560f81c166001976020821060001461317157506010808210156127b4575060088082101561187e57508061109157505050610c2a908a614198565b92909860a08560805101013560001461108757610c6173ffffffffffffffffffffffffffffffffffffffff600154169b5b35613854565b9960408660805101013585829d927f80000000000000000000000000000000000000000000000000000000000000008314610fcf575b50959c95505b7f8000000000000000000000000000000000000000000000000000000000000000811015610a4e5760428610610fc85730915b86602b11610a4e578d91601783013560601c9083359462ffffff8660601c96610d1573ffffffffffffffffffffffffffffffffffffffff92839260481c16868a614365565b169084881015610fac57806401000276a4965b602b60405199604060208c01528160608c015260808b0137600060ab8a015216604088015260a0875260c087019587871067ffffffffffffffff881117610f7d576040948288958688527f128acb080000000000000000000000000000000000000000000000000000000087521660c48a0152868a1060e48a01526101048901521661012487015260a06101448701528160007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff4088610deb610164820182610af6565b0301925af1928315610f71576000928394610f2f575b5050610e159310600014610f2857506141b5565b9a60428510610e5657309085601711610a4e5760177fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe991019501949b610c9d565b50985098606091969597949392509160805101013511610efe575b1580610ed1575b610e8a57506001019291929092610b68565b90610ecd60409283519384937f2c4029e9000000000000000000000000000000000000000000000000000000008552600485015260248401526044830190610af6565b0390fd5b507f8000000000000000000000000000000000000000000000000000000000000000828501351615610e78565b60046040517f39d35496000000000000000000000000000000000000000000000000000000008152fd5b90506141b5565b91929093506040843d604011610f69575b81610f4d604093866135fa565b8101031261038757505160e092909201519190610e1538610e01565b3d9150610f40565b6040513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8073fffd8963efd1fc6a506488495d951d5263988d2596610d28565b8b91610cd0565b60149192501061105d576020602491604051928380927f70a082310000000000000000000000000000000000000000000000000000000082523060048301523560601c5afa908115610f715760009161102b575b503880610c97565b906020823d602011611055575b81611045602093836135fa565b8101031261038757505138611023565b3d9150611038565b60046040517f3b99b53d000000000000000000000000000000000000000000000000000000008152fd5b610c61309b610c5b565b6001819d969d9b989794959a999b146000146111b7575050506040926110bf84836080510101359382614198565b608051840160a00135156111ab5760606110f273ffffffffffffffffffffffffffffffffffffffff600154169435613854565b946080510101356000557f8000000000000000000000000000000000000000000000000000000000000000851015610a4e576111319361020c866141b5565b9091901561119c5750611143906141b5565b0361117357507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6000555b610e71565b600490517fd4e0248e000000000000000000000000000000000000000000000000000000008152fd5b6111a691506141b5565b611143565b60606110f23094610c5b565b9194929391600281036112065750505061116e925073ffffffffffffffffffffffffffffffffffffffff600154166111ff604060608560805101013594608051010135613854565b91356139d0565b9193916003810361157857505060805181018084019390604090850312610a4e57823567ffffffffffffffff8111610a4e5782608051010192606084860312610a4e57604051946060860186811067ffffffffffffffff821117610f7d57604052602085013567ffffffffffffffff8111610a4e57850160208201809882011215610a4e5760208101359061129a826136a5565b926112a860405194856135fa565b8284526040602085019360071b830101918a8311610a4e57604001925b828410611513575050505085526112de60408501610a76565b956020860196875260606040870195013585526040846080510101359067ffffffffffffffff8211610a4e57602061131f92611325966080510101016136d0565b5061417b565b909173ffffffffffffffffffffffffffffffffffffffff600154169473ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163b15610a4e5794929391906040519586947f2a2d80d100000000000000000000000000000000000000000000000000000000865260048601526060602486015260c48501935193606060648701528451809152602060e487019501906000905b80821061149a575050509461143e9285949273ffffffffffffffffffffffffffffffffffffffff600098511660848701525160a48601527ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc858403016044860152613537565b03818373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165af18015610f715761148b575b50610e71565b61149490613576565b38611485565b9197965091929394602060806001928a5173ffffffffffffffffffffffffffffffffffffffff815116825273ffffffffffffffffffffffffffffffffffffffff848201511684830152606065ffffffffffff918260408201511660408501520151166060820152019801920188969795949392916113d8565b608060208584030112610a4e5760206080916040516115318161358a565b61153a87610a76565b8152611547838801610a76565b83820152611557604088016136bd565b6040820152611568606088016136bd565b60608201528152019301926112c5565b600495509193508482036116e757505090916040606061159e8286608051010135613854565b608051909501013573ffffffffffffffffffffffffffffffffffffffff908116933516806116145750479283106115ee575050806115de575b5050610e71565b6115e7916144d1565b38806115d7565b517f6a12f104000000000000000000000000000000000000000000000000000000008152fd5b9391908051937f70a082310000000000000000000000000000000000000000000000000000000085523083860152602085602481895afa9485156116dc576000956116a8575b50841061168257505081611671575b505050610e71565b61167a9261453f565b388080611669565b517f675cae38000000000000000000000000000000000000000000000000000000008152fd5b90946020823d6020116116d4575b816116c3602093836135fa565b81010312610387575051933861165a565b3d91506116b6565b82513d6000823e3d90fd5b8103611714575061116e925061170d604060608460805101013593608051010135613854565b90356138a5565b9091906006810361184e57506080510160608101359060409061173990820135613854565b9282158015611843575b61181b573573ffffffffffffffffffffffffffffffffffffffff16938461177f57505061116e92506117786127109147613984565b04906144d1565b8151907f70a082310000000000000000000000000000000000000000000000000000000082523090820152602081602481885afa91821561181157506000916117dd575b506117d661116e94939261271092613984565b049161453f565b906020823d602011611809575b816117f7602093836135fa565b810103126103875750516117d66117c3565b3d91506117ea565b513d6000823e3d90fd5b8482517fdeaa01e6000000000000000000000000000000000000000000000000000000008152fd5b506127108311611743565b83602491604051917fd76a1e9e000000000000000000000000000000000000000000000000000000008352820152fd5b819d969d9b989794959a999b93929314600014611b85575050506040916118ad83836080510101359185614198565b92909460a082608051010135600014611b7b576118e373ffffffffffffffffffffffffffffffffffffffff600154169135613854565b908615611b4c576118f385613a94565b8760011015611b4c5761191561195d9161190f60208901613a94565b90613c34565b907f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000613b2c565b938481611b32575b5050507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff860193868511611b03576119b9946119be73ffffffffffffffffffffffffffffffffffffffff9687928a85613a84565b613a94565b16948651947f70a082310000000000000000000000000000000000000000000000000000000091828752841693600499858b89015260249460208987818d5afa988915611af857600099611ac3575b509160209695949391611a1f93613cad565b8751968793849283528a8301525afa928315611ab857600093611a83575b50906060611a519260805101013592613ab5565b10611a5d575050610e71565b517f849eaf98000000000000000000000000000000000000000000000000000000008152fd5b90926020823d602011611ab0575b81611a9e602093836135fa565b81010312610387575051916060611a3d565b3d9150611a91565b84513d6000823e3d90fd5b90986020823d602011611af0575b81611ade602093836135fa565b81010312610387575051976020611a0d565b3d9150611ad1565b8b513d6000823e3d90fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b611b4492611b3f88613a94565b613ac2565b388084611965565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6118e33091610c5b565b919492939160098103611f66575050611b9e9082614198565b608051840160a0013515611f5c57611bcf73ffffffffffffffffffffffffffffffffffffffff600154169335613854565b92611bd9836136a5565b95611be760405197886135fa565b83875283901b820160208701368211610a4e5783905b828210611f44575050506000946002875110611f1a576040816080510101359680517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111611b035790815b611ca757505060805101606001358611611c7d578215611b4c5761116e9585611c7892611b3f85613a94565b613cad565b60046040517f8ab0bc16000000000000000000000000000000000000000000000000000000008152fd5b90977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff89019750888811611b035773ffffffffffffffffffffffffffffffffffffffff611cf7611d6d9984613a70565b5116611d2373ffffffffffffffffffffffffffffffffffffffff611d1b8c86613a70565b511682613c34565b819a917f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000613b2c565b90604051907f0902f1ac00000000000000000000000000000000000000000000000000000000825260608260048173ffffffffffffffffffffffffffffffffffffffff87165afa9a8b15610f7157600092839c611ed1575b5073ffffffffffffffffffffffffffffffffffffffff1603611eb7576dffffffffffffffffffffffffffff8091169916905b9880158015611eaf575b611e855782611e0f91613984565b916103e892838102938185041490151715611b0357611e2d91613ab5565b6103e590818102918183041490151715611b0357611e4a91613997565b60018101809111611b0357978015611b03577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff019081611c4c565b60046040517f7b9c8916000000000000000000000000000000000000000000000000000000008152fd5b508115611e01565b6dffffffffffffffffffffffffffff998a16991690611df7565b611f0a919c5073ffffffffffffffffffffffffffffffffffffffff935060603d8111611f13575b611f0281836135fa565b810190613c77565b509b9092611dc5565b503d611ef8565b60046040517f20db8267000000000000000000000000000000000000000000000000000000008152fd5b60208091611f5184610a76565b815201910190611bfd565b611bcf3093610c5b565b92945091600a81036120cc5750608051830160e08101358101946020808701359450909291611f9991908703018461414d565b1161105d5773ffffffffffffffffffffffffffffffffffffffff93847f00000000000000000000000000000000000000000000000000000000000000001692856001541691843b15610a4e5760409587875198899687967f2b67b570000000000000000000000000000000000000000000000000000000008852600488015261202190610a76565b166024860152808883608051010161203890610a76565b16604486015265ffffffffffff808360805101606001612057906136bd565b166064870152826080510160800161206e906136bd565b166084860152816080510160a00161208590610a76565b1660a48501526080510160c0013560c484015261010060e48401526120b1916101048401918701613537565b03815a6000948591f1908115611811575061148b5750610e71565b600b8103612296575050506120eb604080926080510101359235613854565b91807f80000000000000000000000000000000000000000000000000000000000000008103612266575050475b8061212557505050610e71565b73ffffffffffffffffffffffffffffffffffffffff90817f000000000000000000000000000000000000000000000000000000000000000016803b15610a4e578351927fd0e30db0000000000000000000000000000000000000000000000000000000008452600493600081868187875af1801561225b5761224c575b5030908616036121b4575b5050611669565b6122139460006020948651978895869485937fa9059cbb00000000000000000000000000000000000000000000000000000000855284016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b03925af1908115611811575061222d575b808080806121ad565b6122459060203d6020116106355761062681836135fa565b5038612224565b61225590613576565b386121a2565b86513d6000823e3d90fd5b47101561211857600482517f6a12f104000000000000000000000000000000000000000000000000000000008152fd5b600c810361242657505050906122ac9035613854565b9073ffffffffffffffffffffffffffffffffffffffff807f00000000000000000000000000000000000000000000000000000000000000001660408051937f70a08231000000000000000000000000000000000000000000000000000000008552600430818701526024916020878481885afa968715611ab8576000976123f2575b506080510183013586106123cb578561234e575b50505050505050610e71565b833b15610a4e57600091869183855196879485937f2e1a7d4d0000000000000000000000000000000000000000000000000000000085528401525af190811561181157506123bc575b5030908316036123ac575b8080808080612342565b6123b5916144d1565b38806123a2565b6123c590613576565b38612397565b82517f6a12f104000000000000000000000000000000000000000000000000000000008152fd5b90966020823d60201161241e575b8161240d602093836135fa565b81010312610387575051958361232e565b3d9150612400565b600d8103612681575082608051010191602083019360208260805101850312610a4e573567ffffffffffffffff8111610a4e57849160805101019182011215610a4e57602081013590612478826136a5565b93604093612488855196876135fa565b838652602086019285849560071b820101928311610a4e578501925b82841061261f575050505073ffffffffffffffffffffffffffffffffffffffff90816001541684519060005b8281106125b357505050817f00000000000000000000000000000000000000000000000000000000000000001691823b15610a4e5783517f0d58b1db000000000000000000000000000000000000000000000000000000008152602060048201529451602486018190528592604484019290916000915b81831061256f57505050509181600081819503925af1908115611811575061148b5750610e71565b91938395506080602091846060600195975182815116845282868201511686850152828d820151168d85015201511660608201520195019301909187949392612547565b81856125bf838a613a70565b515116036125f6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114611b03576001016124d0565b600486517fe7002877000000000000000000000000000000000000000000000000000000008152fd5b608060208584030112610a4e576020608091875161263c8161358a565b61264587610a76565b8152612652838801610a76565b83820152612661898801610a76565b8982015261267160608801610a76565b60608201528152019301926124a4565b9294505050600e810361278357506040918251907f70a0823100000000000000000000000000000000000000000000000000000000825260208260248173ffffffffffffffffffffffffffffffffffffffff806004983516888301528886608051010135165afa918215611ab85760009261274e575b5060805101606001351180159290612710575050610e71565b517fa3281672000000000000000000000000000000000000000000000000000000006020820152908152909150612746816135c2565b9038806115d7565b90916020823d60201161277b575b81612769602093836135fa565b810103126103875750519060606126f7565b3d915061275c565b602490604051907fd76a1e9e0000000000000000000000000000000000000000000000000000000082526004820152fd5b9150915060189b95939897999692949b808310600014612d435750810361282a5750505060009250906127e883928261417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b90610e71565b6011810361288157505050600092509061284583928261417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b601281036128d857505050600092509061289c83928261417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b919392509060138103612a3e575050909150357f0000000000000000000000000000000000000000000000000000000000000000916040600080825160208101907f8264fe9800000000000000000000000000000000000000000000000000000000825260248781830152815261294e816135de565b5190606086608051010135885af192612965613675565b948415612a04578273ffffffffffffffffffffffffffffffffffffffff612993921694608051010135613854565b90833b15610a4e5782517f8b72a2ec00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9290921660048301526024820152916000908390604490829084905af1908115611811575061148b5750610e71565b505091925050517fae9bdf0000000000000000000000000000000000000000000000000000000000602082015260048152612824816135c2565b60158103612b4f57505090604091828051917f6352211e0000000000000000000000000000000000000000000000000000000083526020836024816004976060816080510101358983015273ffffffffffffffffffffffffffffffffffffffff968791608051010135165afa928315612b4457600093612b05575b5081903516911614918215612acf575050610e71565b517f7dbe7e89000000000000000000000000000000000000000000000000000000006020820152908152909150612746816135c2565b6020939193813d602011612b3c575b81612b21602093836135fa565b8101031261044b575190828216820361038757509181612ab9565b3d9150612b14565b85513d6000823e3d90fd5b60168103612c765750506040918251907efdd58e00000000000000000000000000000000000000000000000000000000825260208280612bc160049660608660805101013590358884016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b038173ffffffffffffffffffffffffffffffffffffffff8886608051010135165afa918215611ab857600092612c41575b5060809081510101351191821592612c0b575050610e71565b517f483a6929000000000000000000000000000000000000000000000000000000006020820152908152909150612746816135c2565b90916020823d602011612c6e575b81612c5c602093836135fa565b81010312610387575051906080612bf2565b3d9150612c4f565b909290601714612c87575050610e71565b60409073ffffffffffffffffffffffffffffffffffffffff612caf8383608051010135613854565b93351692833b15610a4e5782517f42842e0e00000000000000000000000000000000000000000000000000000000815260805130600483015273ffffffffffffffffffffffffffffffffffffffff909216602482015291016060013560448201529160009083908183816064810103925af19081156118115750612d34575b806115d7565b612d3d90613576565b38612d2e565b9396938214159050612d7e5750505061282492507f000000000000000000000000000000000000000000000000000000000000000091613717565b60198103612dd5575050506000925090612d9983928261417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b601a8103612e2c575050506000925090612df083928261417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b601b8103612f53575050506000612e4481928461417b565b9390604094818651928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1918291612e87613675565b92612e95575b505090610e71565b73ffffffffffffffffffffffffffffffffffffffff608083815101013516612ec4606084608051010135613854565b90825190612ed1826135a6565b60008252803b15610a4e57612f2d94600080948651978895869485937ff242432a00000000000000000000000000000000000000000000000000000000855260a060c0836080510101359260805101013590306004870161380f565b03925af19081156118115750612f44575b80612e8d565b612f4d90613576565b38612f3e565b91949091601c8103612f8e5750505061282492507f000000000000000000000000000000000000000000000000000000000000000091613717565b9193929091601d81036131175750506060816080510101359060409173ffffffffffffffffffffffffffffffffffffffff612fcf8484608051010135613854565b9435168351947efdd58e0000000000000000000000000000000000000000000000000000000086526004936020878061302e87308a84016020909392919373ffffffffffffffffffffffffffffffffffffffff60408201951681520152565b0381865afa96871561225b576000976130e2575b50608090815101013586106130ba57845161305c816135a6565b60008152823b15610a4e576000946130a486928851998a97889687957ff242432a0000000000000000000000000000000000000000000000000000000087523090870161380f565b03925af1908115611811575061148b5750610e71565b8385517f675cae38000000000000000000000000000000000000000000000000000000008152fd5b90966020823d60201161310f575b816130fd602093836135fa565b81010312610387575051956080613042565b3d91506130f0565b929450925050601e810361278357508161313560009392849361417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b9499989a92506020819d9792969d989498146000146131da575050505050508061319e600093849361417b565b81604051928392833781018481520391357f00000000000000000000000000000000000000000000000000000000000000005af1612824613675565b602190808203613351575050505090916131ff6131f7868661415a565b96909561417b565b929061324160409788519760208901997f24856bc3000000000000000000000000000000000000000000000000000000008b5260248a01526064890191613537565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc878203016044880152818152602082818301951b82010195856000915b8483106132d357505050505050505091816132c5600094938594037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826135fa565b519082305af1612824613675565b90919293949596977fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085820301885288358284360301811215610a4e578301906020823592019167ffffffffffffffff8111610a4e578036038313610a4e5761334160209283928b95613537565b9a0198019695949301919061327f565b929750929593509350602281146000146127835750604080936080510101359060009060028310156134e1575050808491156000146134895750506000907f0000000000000000000000000000000000000000000000000000000000000000925b6020838251937f095ea7b3000000000000000000000000000000000000000000000000000000008552600496878601526024947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff868201526044968792355af13d15601f3d1187600051141617161561342e5750505050610e71565b91600e7f415050524f56455f4641494c45440000000000000000000000000000000000009260206064969551957f08c379a0000000000000000000000000000000000000000000000000000000008752860152840152820152fd5b036134b8576000907f0000000000000000000000000000000000000000000000000000000000000000926133b2565b600482517f5461585f000000000000000000000000000000000000000000000000000000008152fd5b602492507f4e487b71000000000000000000000000000000000000000000000000000000008252600452fd5b60046040517fff633a38000000000000000000000000000000000000000000000000000000008152fd5b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b67ffffffffffffffff8111610f7d57604052565b6080810190811067ffffffffffffffff821117610f7d57604052565b6020810190811067ffffffffffffffff821117610f7d57604052565b6040810190811067ffffffffffffffff821117610f7d57604052565b6060810190811067ffffffffffffffff821117610f7d57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610f7d57604052565b67ffffffffffffffff8111610f7d57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b3d156136a0573d906136868261363b565b9161369460405193846135fa565b82523d6000602084013e565b606090565b67ffffffffffffffff8111610f7d5760051b60200190565b359065ffffffffffff82168203610a4e57565b81601f82011215610a4e578035906136e78261363b565b926136f560405194856135fa565b82845260208383010111610a4e57816000926020809301838601378301015290565b919290613724908361417b565b90938460405195869384378201906000958693838580955203918635905af19261374c613675565b9284613756575050565b73ffffffffffffffffffffffffffffffffffffffff60608201351661377e6040830135613854565b91813b156106ad576040517f42842e0e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff93909316602484015260800135604483015290919081908390606490829084905af190811561380357506137f85750565b61380190613576565b565b604051903d90823e3d90fd5b919261385195949160a09473ffffffffffffffffffffffffffffffffffffffff8092168552166020840152604083015260608201528160808201520190610af6565b90565b73ffffffffffffffffffffffffffffffffffffffff908082166001810361387e5750506001541690565b90915060020361385157503090565b90816020910312610a4e57518015158103610a4e5790565b9092919073ffffffffffffffffffffffffffffffffffffffff16806138cf575061380191926144d1565b7f80000000000000000000000000000000000000000000000000000000000000008214613902575b92613801929361453f565b9050604051927f70a08231000000000000000000000000000000000000000000000000000000008452306004850152602084602481855afa938415610f7157600094613951575b5092906138f7565b6020813d821161397c575b81613969602093836135fa565b8101031261021757519350613801613949565b3d915061395c565b81810292918115918404141715611b0357565b81156139a1570490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b919273ffffffffffffffffffffffffffffffffffffffff91827f00000000000000000000000000000000000000000000000000000000000000001693843b15610a4e5760009484869281608496816040519b8c9a8b997f36c78516000000000000000000000000000000000000000000000000000000008b521660048a01521660248801521660448601521660648401525af18015610f71576137f85750565b8051821015611b4c5760209160051b010190565b9190811015611b4c5760051b0190565b3573ffffffffffffffffffffffffffffffffffffffff81168103610a4e5790565b91908203918211611b0357565b92919073ffffffffffffffffffffffffffffffffffffffff8082163003613aee575050613801926138a5565b8084959411613b02576138019416926139d0565b60046040517fc4bd89a9000000000000000000000000000000000000000000000000000000008152fd5b9173ffffffffffffffffffffffffffffffffffffffff93613c2d916040519060208201927fffffffffffffffffffffffffffffffffffffffff000000000000000000000000809260601b16845260601b16603482015260288152613b8f816135de565b519020613c01604051938492602084019687917fffffffffffffffffffffffffffffffffffffffff000000000000000000000000605594927fff00000000000000000000000000000000000000000000000000000000000000855260601b166001840152601583015260358201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826135fa565b5190201690565b73ffffffffffffffffffffffffffffffffffffffff8281169082161015613c585791565b9091565b51906dffffffffffffffffffffffffffff82168203610a4e57565b90816060910312610a4e57613c8b81613c5c565b916040613c9a60208401613c5c565b92015163ffffffff81168103610a4e5790565b9260028210614123578115611b4c57613cc584613a94565b9160019481861015611b4c5791613ce360209461190f868601613a94565b50926000935b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84018510613d1c575050505050505050565b613d2a6119b9868685613a84565b92613d3b6119b98a88018786613a84565b936040908151957f0902f1ac00000000000000000000000000000000000000000000000000000000875273ffffffffffffffffffffffffffffffffffffffff80941694606092600493808a86818b5afa998a1561225b57908d9594939291600091829c6140fd575b50508780916dffffffffffffffffffffffffffff8091169c16921692168214998a6000146140f7575b8651958680947f70a082310000000000000000000000000000000000000000000000000000000082528b8883015260249889915afa9283156140ec578e6000946140bb575b5050808303918115938480156140b3575b61408b57826103e5808602958604149114171561405e57613e439083613984565b926103e880830292830414171561403157613e689291613e629161414d565b90613997565b971561402957600097905b898b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe820181101561401d579161190f6119b9613eb9936002613f039c9601908d613a84565b8198917f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000613b2c565b965b988551918d83019367ffffffffffffffff9484811086821117613ff057885260008452813b15610a4e5760008a93613f8382968b519c8d97889687957f022c0d9f0000000000000000000000000000000000000000000000000000000087528d8701528d860152166044840152608060648401526084830190610af6565b03925af18015611ab857908d969594939291613fa8575b505050505094019391613ce9565b909192938095965011613fc45750505287903880808080613f9a565b6041907f4e487b7100000000000000000000000000000000000000000000000000000000600052526000fd5b876041887f4e487b7100000000000000000000000000000000000000000000000000000000600052526000fd5b5050508b956000613f05565b600090613e73565b856011867f4e487b7100000000000000000000000000000000000000000000000000000000600052526000fd5b866011877f4e487b7100000000000000000000000000000000000000000000000000000000600052526000fd5b8689517f7b9c8916000000000000000000000000000000000000000000000000000000008152fd5b508115613e22565b8181959293953d83116140e5575b6140d381836135fa565b8101031261038757505191388e613e11565b503d6140c9565b87513d6000823e3d90fd5b90613dcc565b899c50899250908161411a92903d10611f1357611f0281836135fa565b509b9091613da3565b60046040517fae52ad0c000000000000000000000000000000000000000000000000000000008152fd5b91908201809211611b0357565b91823583019161417460208435958186019503018561414d565b1161105d57565b91602083013583019161417460208435958186019503018561414d565b91606083013583019161417460208435958186019503018561414d565b7f80000000000000000000000000000000000000000000000000000000000000008114611b035760000390565b939193602b841061105d578462ffffff6000614267946142ee6142999935988960601c9a8b9a61423b601789013560601c9d8e109c73ffffffffffffffffffffffffffffffffffffffff9e8f998a9460481c1691614365565b16968b861461434a576401000276a49a5b60409d8e9b8c93845196879560208701526060860191613537565b91168b830152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826135fa565b848851998a98899788967f128acb080000000000000000000000000000000000000000000000000000000088521660048701528c6024870152604486015216606484015260a0608484015260a4830190610af6565b03925af190811561433f576000938492614309575b50509192565b9080949250813d8311614338575b61432181836135fa565b810103126103875750602082519201513880614303565b503d614317565b83513d6000823e3d90fd5b73fffd8963efd1fc6a506488495d951d5263988d259a61424c565b73ffffffffffffffffffffffffffffffffffffffff92838316848316116144c9575b62ffffff90846040519481602087019516855216604085015216606083015260608252608082019082821067ffffffffffffffff831117610f7d577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80613c2d9183604052845190209361449c60a08201957f0000000000000000000000000000000000000000000000000000000000000000907f000000000000000000000000000000000000000000000000000000000000000088917fffffffffffffffffffffffffffffffffffffffff000000000000000000000000605594927fff00000000000000000000000000000000000000000000000000000000000000855260601b166001840152601583015260358201520190565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff608101845201826135fa565b909190614387565b600080809381935af1156144e157565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4554485f5452414e534645525f4641494c4544000000000000000000000000006044820152fd5b60009182604492602095604051937fa9059cbb000000000000000000000000000000000000000000000000000000008552600485015260248401525af13d15601f3d116001600051141617161561459257565b60646040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600f60248201527f5452414e534645525f4641494c454400000000000000000000000000000000006044820152fdfea164736f6c6343000811000a"


func init() {
	if err := json.Unmarshal([]byte(UniversalRouterStorageLayoutJSON), UniversalRouterStorageLayout); err != nil {
		panic(err)
	}

	layouts["UniversalRouter"] = UniversalRouterStorageLayout
	deployedBytecodes["UniversalRouter"] = UniversalRouterDeployedBin
	immutableReferences["UniversalRouter"] = true
}
