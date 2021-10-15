// Code generated by "stringer -type=Code -linecomment"; DO NOT EDIT.

package multicodec

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Identity-0]
	_ = x[Cidv1-1]
	_ = x[Cidv2-2]
	_ = x[Cidv3-3]
	_ = x[Ip4-4]
	_ = x[Tcp-6]
	_ = x[Sha1-17]
	_ = x[Sha2_256-18]
	_ = x[Sha2_512-19]
	_ = x[Sha3_512-20]
	_ = x[Sha3_384-21]
	_ = x[Sha3_256-22]
	_ = x[Sha3_224-23]
	_ = x[Shake128-24]
	_ = x[Shake256-25]
	_ = x[Keccak224-26]
	_ = x[Keccak256-27]
	_ = x[Keccak384-28]
	_ = x[Keccak512-29]
	_ = x[Blake3-30]
	_ = x[Sha2_384-32]
	_ = x[Dccp-33]
	_ = x[Murmur3X64_64-34]
	_ = x[Murmur3_32-35]
	_ = x[Ip6-41]
	_ = x[Ip6zone-42]
	_ = x[Path-47]
	_ = x[Multicodec-48]
	_ = x[Multihash-49]
	_ = x[Multiaddr-50]
	_ = x[Multibase-51]
	_ = x[Dns-53]
	_ = x[Dns4-54]
	_ = x[Dns6-55]
	_ = x[Dnsaddr-56]
	_ = x[Protobuf-80]
	_ = x[Cbor-81]
	_ = x[Raw-85]
	_ = x[DblSha2_256-86]
	_ = x[Rlp-96]
	_ = x[Bencode-99]
	_ = x[DagPb-112]
	_ = x[DagCbor-113]
	_ = x[Libp2pKey-114]
	_ = x[GitRaw-120]
	_ = x[TorrentInfo-123]
	_ = x[TorrentFile-124]
	_ = x[LeofcoinBlock-129]
	_ = x[LeofcoinTx-130]
	_ = x[LeofcoinPr-131]
	_ = x[Sctp-132]
	_ = x[DagJose-133]
	_ = x[DagCose-134]
	_ = x[EthBlock-144]
	_ = x[EthBlockList-145]
	_ = x[EthTxTrie-146]
	_ = x[EthTx-147]
	_ = x[EthTxReceiptTrie-148]
	_ = x[EthTxReceipt-149]
	_ = x[EthStateTrie-150]
	_ = x[EthAccountSnapshot-151]
	_ = x[EthStorageTrie-152]
	_ = x[EthReceiptLogTrie-153]
	_ = x[EthRecieptLog-154]
	_ = x[Aes128-160]
	_ = x[Aes192-161]
	_ = x[Aes256-162]
	_ = x[Chacha128-163]
	_ = x[Chacha256-164]
	_ = x[BitcoinBlock-176]
	_ = x[BitcoinTx-177]
	_ = x[BitcoinWitnessCommitment-178]
	_ = x[ZcashBlock-192]
	_ = x[ZcashTx-193]
	_ = x[Caip50-202]
	_ = x[Streamid-206]
	_ = x[StellarBlock-208]
	_ = x[StellarTx-209]
	_ = x[Md4-212]
	_ = x[Md5-213]
	_ = x[Bmt-214]
	_ = x[DecredBlock-224]
	_ = x[DecredTx-225]
	_ = x[IpldNs-226]
	_ = x[IpfsNs-227]
	_ = x[SwarmNs-228]
	_ = x[IpnsNs-229]
	_ = x[Zeronet-230]
	_ = x[Secp256k1Pub-231]
	_ = x[Bls12_381G1Pub-234]
	_ = x[Bls12_381G2Pub-235]
	_ = x[X25519Pub-236]
	_ = x[Ed25519Pub-237]
	_ = x[Bls12_381G1g2Pub-238]
	_ = x[DashBlock-240]
	_ = x[DashTx-241]
	_ = x[SwarmManifest-250]
	_ = x[SwarmFeed-251]
	_ = x[Udp-273]
	_ = x[P2pWebrtcStar-275]
	_ = x[P2pWebrtcDirect-276]
	_ = x[P2pStardust-277]
	_ = x[P2pCircuit-290]
	_ = x[DagJson-297]
	_ = x[Udt-301]
	_ = x[Utp-302]
	_ = x[Unix-400]
	_ = x[Thread-406]
	_ = x[P2p-421]
	_ = x[Ipfs-421]
	_ = x[Https-443]
	_ = x[Onion-444]
	_ = x[Onion3-445]
	_ = x[Garlic64-446]
	_ = x[Garlic32-447]
	_ = x[Tls-448]
	_ = x[Noise-454]
	_ = x[Quic-460]
	_ = x[Ws-477]
	_ = x[Wss-478]
	_ = x[P2pWebsocketStar-479]
	_ = x[Http-480]
	_ = x[Swhid1Snp-496]
	_ = x[Json-512]
	_ = x[Messagepack-513]
	_ = x[Libp2pPeerRecord-769]
	_ = x[Libp2pRelayRsvp-770]
	_ = x[CarIndexSorted-1024]
	_ = x[CarMultihashIndexSorted-1025]
	_ = x[Sha2_256Trunc254Padded-4114]
	_ = x[Sha2_224-4115]
	_ = x[Sha2_512_224-4116]
	_ = x[Sha2_512_256-4117]
	_ = x[Murmur3X64_128-4130]
	_ = x[Ripemd128-4178]
	_ = x[Ripemd160-4179]
	_ = x[Ripemd256-4180]
	_ = x[Ripemd320-4181]
	_ = x[X11-4352]
	_ = x[P256Pub-4608]
	_ = x[P384Pub-4609]
	_ = x[P521Pub-4610]
	_ = x[Ed448Pub-4611]
	_ = x[X448Pub-4612]
	_ = x[RsaPub-4613]
	_ = x[Ed25519Priv-4864]
	_ = x[Secp256k1Priv-4865]
	_ = x[X25519Priv-4866]
	_ = x[Kangarootwelve-7425]
	_ = x[Sm3_256-21325]
	_ = x[Blake2b8-45569]
	_ = x[Blake2b16-45570]
	_ = x[Blake2b24-45571]
	_ = x[Blake2b32-45572]
	_ = x[Blake2b40-45573]
	_ = x[Blake2b48-45574]
	_ = x[Blake2b56-45575]
	_ = x[Blake2b64-45576]
	_ = x[Blake2b72-45577]
	_ = x[Blake2b80-45578]
	_ = x[Blake2b88-45579]
	_ = x[Blake2b96-45580]
	_ = x[Blake2b104-45581]
	_ = x[Blake2b112-45582]
	_ = x[Blake2b120-45583]
	_ = x[Blake2b128-45584]
	_ = x[Blake2b136-45585]
	_ = x[Blake2b144-45586]
	_ = x[Blake2b152-45587]
	_ = x[Blake2b160-45588]
	_ = x[Blake2b168-45589]
	_ = x[Blake2b176-45590]
	_ = x[Blake2b184-45591]
	_ = x[Blake2b192-45592]
	_ = x[Blake2b200-45593]
	_ = x[Blake2b208-45594]
	_ = x[Blake2b216-45595]
	_ = x[Blake2b224-45596]
	_ = x[Blake2b232-45597]
	_ = x[Blake2b240-45598]
	_ = x[Blake2b248-45599]
	_ = x[Blake2b256-45600]
	_ = x[Blake2b264-45601]
	_ = x[Blake2b272-45602]
	_ = x[Blake2b280-45603]
	_ = x[Blake2b288-45604]
	_ = x[Blake2b296-45605]
	_ = x[Blake2b304-45606]
	_ = x[Blake2b312-45607]
	_ = x[Blake2b320-45608]
	_ = x[Blake2b328-45609]
	_ = x[Blake2b336-45610]
	_ = x[Blake2b344-45611]
	_ = x[Blake2b352-45612]
	_ = x[Blake2b360-45613]
	_ = x[Blake2b368-45614]
	_ = x[Blake2b376-45615]
	_ = x[Blake2b384-45616]
	_ = x[Blake2b392-45617]
	_ = x[Blake2b400-45618]
	_ = x[Blake2b408-45619]
	_ = x[Blake2b416-45620]
	_ = x[Blake2b424-45621]
	_ = x[Blake2b432-45622]
	_ = x[Blake2b440-45623]
	_ = x[Blake2b448-45624]
	_ = x[Blake2b456-45625]
	_ = x[Blake2b464-45626]
	_ = x[Blake2b472-45627]
	_ = x[Blake2b480-45628]
	_ = x[Blake2b488-45629]
	_ = x[Blake2b496-45630]
	_ = x[Blake2b504-45631]
	_ = x[Blake2b512-45632]
	_ = x[Blake2s8-45633]
	_ = x[Blake2s16-45634]
	_ = x[Blake2s24-45635]
	_ = x[Blake2s32-45636]
	_ = x[Blake2s40-45637]
	_ = x[Blake2s48-45638]
	_ = x[Blake2s56-45639]
	_ = x[Blake2s64-45640]
	_ = x[Blake2s72-45641]
	_ = x[Blake2s80-45642]
	_ = x[Blake2s88-45643]
	_ = x[Blake2s96-45644]
	_ = x[Blake2s104-45645]
	_ = x[Blake2s112-45646]
	_ = x[Blake2s120-45647]
	_ = x[Blake2s128-45648]
	_ = x[Blake2s136-45649]
	_ = x[Blake2s144-45650]
	_ = x[Blake2s152-45651]
	_ = x[Blake2s160-45652]
	_ = x[Blake2s168-45653]
	_ = x[Blake2s176-45654]
	_ = x[Blake2s184-45655]
	_ = x[Blake2s192-45656]
	_ = x[Blake2s200-45657]
	_ = x[Blake2s208-45658]
	_ = x[Blake2s216-45659]
	_ = x[Blake2s224-45660]
	_ = x[Blake2s232-45661]
	_ = x[Blake2s240-45662]
	_ = x[Blake2s248-45663]
	_ = x[Blake2s256-45664]
	_ = x[Skein256_8-45825]
	_ = x[Skein256_16-45826]
	_ = x[Skein256_24-45827]
	_ = x[Skein256_32-45828]
	_ = x[Skein256_40-45829]
	_ = x[Skein256_48-45830]
	_ = x[Skein256_56-45831]
	_ = x[Skein256_64-45832]
	_ = x[Skein256_72-45833]
	_ = x[Skein256_80-45834]
	_ = x[Skein256_88-45835]
	_ = x[Skein256_96-45836]
	_ = x[Skein256_104-45837]
	_ = x[Skein256_112-45838]
	_ = x[Skein256_120-45839]
	_ = x[Skein256_128-45840]
	_ = x[Skein256_136-45841]
	_ = x[Skein256_144-45842]
	_ = x[Skein256_152-45843]
	_ = x[Skein256_160-45844]
	_ = x[Skein256_168-45845]
	_ = x[Skein256_176-45846]
	_ = x[Skein256_184-45847]
	_ = x[Skein256_192-45848]
	_ = x[Skein256_200-45849]
	_ = x[Skein256_208-45850]
	_ = x[Skein256_216-45851]
	_ = x[Skein256_224-45852]
	_ = x[Skein256_232-45853]
	_ = x[Skein256_240-45854]
	_ = x[Skein256_248-45855]
	_ = x[Skein256_256-45856]
	_ = x[Skein512_8-45857]
	_ = x[Skein512_16-45858]
	_ = x[Skein512_24-45859]
	_ = x[Skein512_32-45860]
	_ = x[Skein512_40-45861]
	_ = x[Skein512_48-45862]
	_ = x[Skein512_56-45863]
	_ = x[Skein512_64-45864]
	_ = x[Skein512_72-45865]
	_ = x[Skein512_80-45866]
	_ = x[Skein512_88-45867]
	_ = x[Skein512_96-45868]
	_ = x[Skein512_104-45869]
	_ = x[Skein512_112-45870]
	_ = x[Skein512_120-45871]
	_ = x[Skein512_128-45872]
	_ = x[Skein512_136-45873]
	_ = x[Skein512_144-45874]
	_ = x[Skein512_152-45875]
	_ = x[Skein512_160-45876]
	_ = x[Skein512_168-45877]
	_ = x[Skein512_176-45878]
	_ = x[Skein512_184-45879]
	_ = x[Skein512_192-45880]
	_ = x[Skein512_200-45881]
	_ = x[Skein512_208-45882]
	_ = x[Skein512_216-45883]
	_ = x[Skein512_224-45884]
	_ = x[Skein512_232-45885]
	_ = x[Skein512_240-45886]
	_ = x[Skein512_248-45887]
	_ = x[Skein512_256-45888]
	_ = x[Skein512_264-45889]
	_ = x[Skein512_272-45890]
	_ = x[Skein512_280-45891]
	_ = x[Skein512_288-45892]
	_ = x[Skein512_296-45893]
	_ = x[Skein512_304-45894]
	_ = x[Skein512_312-45895]
	_ = x[Skein512_320-45896]
	_ = x[Skein512_328-45897]
	_ = x[Skein512_336-45898]
	_ = x[Skein512_344-45899]
	_ = x[Skein512_352-45900]
	_ = x[Skein512_360-45901]
	_ = x[Skein512_368-45902]
	_ = x[Skein512_376-45903]
	_ = x[Skein512_384-45904]
	_ = x[Skein512_392-45905]
	_ = x[Skein512_400-45906]
	_ = x[Skein512_408-45907]
	_ = x[Skein512_416-45908]
	_ = x[Skein512_424-45909]
	_ = x[Skein512_432-45910]
	_ = x[Skein512_440-45911]
	_ = x[Skein512_448-45912]
	_ = x[Skein512_456-45913]
	_ = x[Skein512_464-45914]
	_ = x[Skein512_472-45915]
	_ = x[Skein512_480-45916]
	_ = x[Skein512_488-45917]
	_ = x[Skein512_496-45918]
	_ = x[Skein512_504-45919]
	_ = x[Skein512_512-45920]
	_ = x[Skein1024_8-45921]
	_ = x[Skein1024_16-45922]
	_ = x[Skein1024_24-45923]
	_ = x[Skein1024_32-45924]
	_ = x[Skein1024_40-45925]
	_ = x[Skein1024_48-45926]
	_ = x[Skein1024_56-45927]
	_ = x[Skein1024_64-45928]
	_ = x[Skein1024_72-45929]
	_ = x[Skein1024_80-45930]
	_ = x[Skein1024_88-45931]
	_ = x[Skein1024_96-45932]
	_ = x[Skein1024_104-45933]
	_ = x[Skein1024_112-45934]
	_ = x[Skein1024_120-45935]
	_ = x[Skein1024_128-45936]
	_ = x[Skein1024_136-45937]
	_ = x[Skein1024_144-45938]
	_ = x[Skein1024_152-45939]
	_ = x[Skein1024_160-45940]
	_ = x[Skein1024_168-45941]
	_ = x[Skein1024_176-45942]
	_ = x[Skein1024_184-45943]
	_ = x[Skein1024_192-45944]
	_ = x[Skein1024_200-45945]
	_ = x[Skein1024_208-45946]
	_ = x[Skein1024_216-45947]
	_ = x[Skein1024_224-45948]
	_ = x[Skein1024_232-45949]
	_ = x[Skein1024_240-45950]
	_ = x[Skein1024_248-45951]
	_ = x[Skein1024_256-45952]
	_ = x[Skein1024_264-45953]
	_ = x[Skein1024_272-45954]
	_ = x[Skein1024_280-45955]
	_ = x[Skein1024_288-45956]
	_ = x[Skein1024_296-45957]
	_ = x[Skein1024_304-45958]
	_ = x[Skein1024_312-45959]
	_ = x[Skein1024_320-45960]
	_ = x[Skein1024_328-45961]
	_ = x[Skein1024_336-45962]
	_ = x[Skein1024_344-45963]
	_ = x[Skein1024_352-45964]
	_ = x[Skein1024_360-45965]
	_ = x[Skein1024_368-45966]
	_ = x[Skein1024_376-45967]
	_ = x[Skein1024_384-45968]
	_ = x[Skein1024_392-45969]
	_ = x[Skein1024_400-45970]
	_ = x[Skein1024_408-45971]
	_ = x[Skein1024_416-45972]
	_ = x[Skein1024_424-45973]
	_ = x[Skein1024_432-45974]
	_ = x[Skein1024_440-45975]
	_ = x[Skein1024_448-45976]
	_ = x[Skein1024_456-45977]
	_ = x[Skein1024_464-45978]
	_ = x[Skein1024_472-45979]
	_ = x[Skein1024_480-45980]
	_ = x[Skein1024_488-45981]
	_ = x[Skein1024_496-45982]
	_ = x[Skein1024_504-45983]
	_ = x[Skein1024_512-45984]
	_ = x[Skein1024_520-45985]
	_ = x[Skein1024_528-45986]
	_ = x[Skein1024_536-45987]
	_ = x[Skein1024_544-45988]
	_ = x[Skein1024_552-45989]
	_ = x[Skein1024_560-45990]
	_ = x[Skein1024_568-45991]
	_ = x[Skein1024_576-45992]
	_ = x[Skein1024_584-45993]
	_ = x[Skein1024_592-45994]
	_ = x[Skein1024_600-45995]
	_ = x[Skein1024_608-45996]
	_ = x[Skein1024_616-45997]
	_ = x[Skein1024_624-45998]
	_ = x[Skein1024_632-45999]
	_ = x[Skein1024_640-46000]
	_ = x[Skein1024_648-46001]
	_ = x[Skein1024_656-46002]
	_ = x[Skein1024_664-46003]
	_ = x[Skein1024_672-46004]
	_ = x[Skein1024_680-46005]
	_ = x[Skein1024_688-46006]
	_ = x[Skein1024_696-46007]
	_ = x[Skein1024_704-46008]
	_ = x[Skein1024_712-46009]
	_ = x[Skein1024_720-46010]
	_ = x[Skein1024_728-46011]
	_ = x[Skein1024_736-46012]
	_ = x[Skein1024_744-46013]
	_ = x[Skein1024_752-46014]
	_ = x[Skein1024_760-46015]
	_ = x[Skein1024_768-46016]
	_ = x[Skein1024_776-46017]
	_ = x[Skein1024_784-46018]
	_ = x[Skein1024_792-46019]
	_ = x[Skein1024_800-46020]
	_ = x[Skein1024_808-46021]
	_ = x[Skein1024_816-46022]
	_ = x[Skein1024_824-46023]
	_ = x[Skein1024_832-46024]
	_ = x[Skein1024_840-46025]
	_ = x[Skein1024_848-46026]
	_ = x[Skein1024_856-46027]
	_ = x[Skein1024_864-46028]
	_ = x[Skein1024_872-46029]
	_ = x[Skein1024_880-46030]
	_ = x[Skein1024_888-46031]
	_ = x[Skein1024_896-46032]
	_ = x[Skein1024_904-46033]
	_ = x[Skein1024_912-46034]
	_ = x[Skein1024_920-46035]
	_ = x[Skein1024_928-46036]
	_ = x[Skein1024_936-46037]
	_ = x[Skein1024_944-46038]
	_ = x[Skein1024_952-46039]
	_ = x[Skein1024_960-46040]
	_ = x[Skein1024_968-46041]
	_ = x[Skein1024_976-46042]
	_ = x[Skein1024_984-46043]
	_ = x[Skein1024_992-46044]
	_ = x[Skein1024_1000-46045]
	_ = x[Skein1024_1008-46046]
	_ = x[Skein1024_1016-46047]
	_ = x[Skein1024_1024-46048]
	_ = x[PoseidonBls12_381A2Fc1-46081]
	_ = x[PoseidonBls12_381A2Fc1Sc-46082]
	_ = x[ZeroxcertImprint256-52753]
	_ = x[FilCommitmentUnsealed-61697]
	_ = x[FilCommitmentSealed-61698]
	_ = x[Plaintextv2-7367777]
	_ = x[HolochainAdrV0-8417572]
	_ = x[HolochainAdrV1-8483108]
	_ = x[HolochainKeyV0-9728292]
	_ = x[HolochainKeyV1-9793828]
	_ = x[HolochainSigV0-10645796]
	_ = x[HolochainSigV1-10711332]
	_ = x[SkynetNs-11639056]
	_ = x[ArweaveNs-11704592]
}

const _Code_name = "identitycidv1cidv2cidv3ip4tcpsha1sha2-256sha2-512sha3-512sha3-384sha3-256sha3-224shake-128shake-256keccak-224keccak-256keccak-384keccak-512blake3sha2-384dccpmurmur3-x64-64murmur3-32ip6ip6zonepathmulticodecmultihashmultiaddrmultibasednsdns4dns6dnsaddrprotobufcborrawdbl-sha2-256rlpbencodedag-pbdag-cborlibp2p-keygit-rawtorrent-infotorrent-fileleofcoin-blockleofcoin-txleofcoin-prsctpdag-josedag-coseeth-blocketh-block-listeth-tx-trieeth-txeth-tx-receipt-trieeth-tx-receipteth-state-trieeth-account-snapshoteth-storage-trieeth-receipt-log-trieeth-reciept-logaes-128aes-192aes-256chacha-128chacha-256bitcoin-blockbitcoin-txbitcoin-witness-commitmentzcash-blockzcash-txcaip-50streamidstellar-blockstellar-txmd4md5bmtdecred-blockdecred-txipld-nsipfs-nsswarm-nsipns-nszeronetsecp256k1-pubbls12_381-g1-pubbls12_381-g2-pubx25519-pubed25519-pubbls12_381-g1g2-pubdash-blockdash-txswarm-manifestswarm-feedudpp2p-webrtc-starp2p-webrtc-directp2p-stardustp2p-circuitdag-jsonudtutpunixthreadp2phttpsoniononion3garlic64garlic32tlsnoisequicwswssp2p-websocket-starhttpswhid-1-snpjsonmessagepacklibp2p-peer-recordlibp2p-relay-rsvpcar-index-sortedcar-multihash-index-sortedsha2-256-trunc254-paddedsha2-224sha2-512-224sha2-512-256murmur3-x64-128ripemd-128ripemd-160ripemd-256ripemd-320x11p256-pubp384-pubp521-pubed448-pubx448-pubrsa-pubed25519-privsecp256k1-privx25519-privkangarootwelvesm3-256blake2b-8blake2b-16blake2b-24blake2b-32blake2b-40blake2b-48blake2b-56blake2b-64blake2b-72blake2b-80blake2b-88blake2b-96blake2b-104blake2b-112blake2b-120blake2b-128blake2b-136blake2b-144blake2b-152blake2b-160blake2b-168blake2b-176blake2b-184blake2b-192blake2b-200blake2b-208blake2b-216blake2b-224blake2b-232blake2b-240blake2b-248blake2b-256blake2b-264blake2b-272blake2b-280blake2b-288blake2b-296blake2b-304blake2b-312blake2b-320blake2b-328blake2b-336blake2b-344blake2b-352blake2b-360blake2b-368blake2b-376blake2b-384blake2b-392blake2b-400blake2b-408blake2b-416blake2b-424blake2b-432blake2b-440blake2b-448blake2b-456blake2b-464blake2b-472blake2b-480blake2b-488blake2b-496blake2b-504blake2b-512blake2s-8blake2s-16blake2s-24blake2s-32blake2s-40blake2s-48blake2s-56blake2s-64blake2s-72blake2s-80blake2s-88blake2s-96blake2s-104blake2s-112blake2s-120blake2s-128blake2s-136blake2s-144blake2s-152blake2s-160blake2s-168blake2s-176blake2s-184blake2s-192blake2s-200blake2s-208blake2s-216blake2s-224blake2s-232blake2s-240blake2s-248blake2s-256skein256-8skein256-16skein256-24skein256-32skein256-40skein256-48skein256-56skein256-64skein256-72skein256-80skein256-88skein256-96skein256-104skein256-112skein256-120skein256-128skein256-136skein256-144skein256-152skein256-160skein256-168skein256-176skein256-184skein256-192skein256-200skein256-208skein256-216skein256-224skein256-232skein256-240skein256-248skein256-256skein512-8skein512-16skein512-24skein512-32skein512-40skein512-48skein512-56skein512-64skein512-72skein512-80skein512-88skein512-96skein512-104skein512-112skein512-120skein512-128skein512-136skein512-144skein512-152skein512-160skein512-168skein512-176skein512-184skein512-192skein512-200skein512-208skein512-216skein512-224skein512-232skein512-240skein512-248skein512-256skein512-264skein512-272skein512-280skein512-288skein512-296skein512-304skein512-312skein512-320skein512-328skein512-336skein512-344skein512-352skein512-360skein512-368skein512-376skein512-384skein512-392skein512-400skein512-408skein512-416skein512-424skein512-432skein512-440skein512-448skein512-456skein512-464skein512-472skein512-480skein512-488skein512-496skein512-504skein512-512skein1024-8skein1024-16skein1024-24skein1024-32skein1024-40skein1024-48skein1024-56skein1024-64skein1024-72skein1024-80skein1024-88skein1024-96skein1024-104skein1024-112skein1024-120skein1024-128skein1024-136skein1024-144skein1024-152skein1024-160skein1024-168skein1024-176skein1024-184skein1024-192skein1024-200skein1024-208skein1024-216skein1024-224skein1024-232skein1024-240skein1024-248skein1024-256skein1024-264skein1024-272skein1024-280skein1024-288skein1024-296skein1024-304skein1024-312skein1024-320skein1024-328skein1024-336skein1024-344skein1024-352skein1024-360skein1024-368skein1024-376skein1024-384skein1024-392skein1024-400skein1024-408skein1024-416skein1024-424skein1024-432skein1024-440skein1024-448skein1024-456skein1024-464skein1024-472skein1024-480skein1024-488skein1024-496skein1024-504skein1024-512skein1024-520skein1024-528skein1024-536skein1024-544skein1024-552skein1024-560skein1024-568skein1024-576skein1024-584skein1024-592skein1024-600skein1024-608skein1024-616skein1024-624skein1024-632skein1024-640skein1024-648skein1024-656skein1024-664skein1024-672skein1024-680skein1024-688skein1024-696skein1024-704skein1024-712skein1024-720skein1024-728skein1024-736skein1024-744skein1024-752skein1024-760skein1024-768skein1024-776skein1024-784skein1024-792skein1024-800skein1024-808skein1024-816skein1024-824skein1024-832skein1024-840skein1024-848skein1024-856skein1024-864skein1024-872skein1024-880skein1024-888skein1024-896skein1024-904skein1024-912skein1024-920skein1024-928skein1024-936skein1024-944skein1024-952skein1024-960skein1024-968skein1024-976skein1024-984skein1024-992skein1024-1000skein1024-1008skein1024-1016skein1024-1024poseidon-bls12_381-a2-fc1poseidon-bls12_381-a2-fc1-sczeroxcert-imprint-256fil-commitment-unsealedfil-commitment-sealedplaintextv2holochain-adr-v0holochain-adr-v1holochain-key-v0holochain-key-v1holochain-sig-v0holochain-sig-v1skynet-nsarweave-ns"

var _Code_map = map[Code]string{
	0:        _Code_name[0:8],
	1:        _Code_name[8:13],
	2:        _Code_name[13:18],
	3:        _Code_name[18:23],
	4:        _Code_name[23:26],
	6:        _Code_name[26:29],
	17:       _Code_name[29:33],
	18:       _Code_name[33:41],
	19:       _Code_name[41:49],
	20:       _Code_name[49:57],
	21:       _Code_name[57:65],
	22:       _Code_name[65:73],
	23:       _Code_name[73:81],
	24:       _Code_name[81:90],
	25:       _Code_name[90:99],
	26:       _Code_name[99:109],
	27:       _Code_name[109:119],
	28:       _Code_name[119:129],
	29:       _Code_name[129:139],
	30:       _Code_name[139:145],
	32:       _Code_name[145:153],
	33:       _Code_name[153:157],
	34:       _Code_name[157:171],
	35:       _Code_name[171:181],
	41:       _Code_name[181:184],
	42:       _Code_name[184:191],
	47:       _Code_name[191:195],
	48:       _Code_name[195:205],
	49:       _Code_name[205:214],
	50:       _Code_name[214:223],
	51:       _Code_name[223:232],
	53:       _Code_name[232:235],
	54:       _Code_name[235:239],
	55:       _Code_name[239:243],
	56:       _Code_name[243:250],
	80:       _Code_name[250:258],
	81:       _Code_name[258:262],
	85:       _Code_name[262:265],
	86:       _Code_name[265:277],
	96:       _Code_name[277:280],
	99:       _Code_name[280:287],
	112:      _Code_name[287:293],
	113:      _Code_name[293:301],
	114:      _Code_name[301:311],
	120:      _Code_name[311:318],
	123:      _Code_name[318:330],
	124:      _Code_name[330:342],
	129:      _Code_name[342:356],
	130:      _Code_name[356:367],
	131:      _Code_name[367:378],
	132:      _Code_name[378:382],
	133:      _Code_name[382:390],
	134:      _Code_name[390:398],
	144:      _Code_name[398:407],
	145:      _Code_name[407:421],
	146:      _Code_name[421:432],
	147:      _Code_name[432:438],
	148:      _Code_name[438:457],
	149:      _Code_name[457:471],
	150:      _Code_name[471:485],
	151:      _Code_name[485:505],
	152:      _Code_name[505:521],
	153:      _Code_name[521:541],
	154:      _Code_name[541:556],
	160:      _Code_name[556:563],
	161:      _Code_name[563:570],
	162:      _Code_name[570:577],
	163:      _Code_name[577:587],
	164:      _Code_name[587:597],
	176:      _Code_name[597:610],
	177:      _Code_name[610:620],
	178:      _Code_name[620:646],
	192:      _Code_name[646:657],
	193:      _Code_name[657:665],
	202:      _Code_name[665:672],
	206:      _Code_name[672:680],
	208:      _Code_name[680:693],
	209:      _Code_name[693:703],
	212:      _Code_name[703:706],
	213:      _Code_name[706:709],
	214:      _Code_name[709:712],
	224:      _Code_name[712:724],
	225:      _Code_name[724:733],
	226:      _Code_name[733:740],
	227:      _Code_name[740:747],
	228:      _Code_name[747:755],
	229:      _Code_name[755:762],
	230:      _Code_name[762:769],
	231:      _Code_name[769:782],
	234:      _Code_name[782:798],
	235:      _Code_name[798:814],
	236:      _Code_name[814:824],
	237:      _Code_name[824:835],
	238:      _Code_name[835:853],
	240:      _Code_name[853:863],
	241:      _Code_name[863:870],
	250:      _Code_name[870:884],
	251:      _Code_name[884:894],
	273:      _Code_name[894:897],
	275:      _Code_name[897:912],
	276:      _Code_name[912:929],
	277:      _Code_name[929:941],
	290:      _Code_name[941:952],
	297:      _Code_name[952:960],
	301:      _Code_name[960:963],
	302:      _Code_name[963:966],
	400:      _Code_name[966:970],
	406:      _Code_name[970:976],
	421:      _Code_name[976:979],
	443:      _Code_name[979:984],
	444:      _Code_name[984:989],
	445:      _Code_name[989:995],
	446:      _Code_name[995:1003],
	447:      _Code_name[1003:1011],
	448:      _Code_name[1011:1014],
	454:      _Code_name[1014:1019],
	460:      _Code_name[1019:1023],
	477:      _Code_name[1023:1025],
	478:      _Code_name[1025:1028],
	479:      _Code_name[1028:1046],
	480:      _Code_name[1046:1050],
	496:      _Code_name[1050:1061],
	512:      _Code_name[1061:1065],
	513:      _Code_name[1065:1076],
	769:      _Code_name[1076:1094],
	770:      _Code_name[1094:1111],
	1024:     _Code_name[1111:1127],
	1025:     _Code_name[1127:1153],
	4114:     _Code_name[1153:1177],
	4115:     _Code_name[1177:1185],
	4116:     _Code_name[1185:1197],
	4117:     _Code_name[1197:1209],
	4130:     _Code_name[1209:1224],
	4178:     _Code_name[1224:1234],
	4179:     _Code_name[1234:1244],
	4180:     _Code_name[1244:1254],
	4181:     _Code_name[1254:1264],
	4352:     _Code_name[1264:1267],
	4608:     _Code_name[1267:1275],
	4609:     _Code_name[1275:1283],
	4610:     _Code_name[1283:1291],
	4611:     _Code_name[1291:1300],
	4612:     _Code_name[1300:1308],
	4613:     _Code_name[1308:1315],
	4864:     _Code_name[1315:1327],
	4865:     _Code_name[1327:1341],
	4866:     _Code_name[1341:1352],
	7425:     _Code_name[1352:1366],
	21325:    _Code_name[1366:1373],
	45569:    _Code_name[1373:1382],
	45570:    _Code_name[1382:1392],
	45571:    _Code_name[1392:1402],
	45572:    _Code_name[1402:1412],
	45573:    _Code_name[1412:1422],
	45574:    _Code_name[1422:1432],
	45575:    _Code_name[1432:1442],
	45576:    _Code_name[1442:1452],
	45577:    _Code_name[1452:1462],
	45578:    _Code_name[1462:1472],
	45579:    _Code_name[1472:1482],
	45580:    _Code_name[1482:1492],
	45581:    _Code_name[1492:1503],
	45582:    _Code_name[1503:1514],
	45583:    _Code_name[1514:1525],
	45584:    _Code_name[1525:1536],
	45585:    _Code_name[1536:1547],
	45586:    _Code_name[1547:1558],
	45587:    _Code_name[1558:1569],
	45588:    _Code_name[1569:1580],
	45589:    _Code_name[1580:1591],
	45590:    _Code_name[1591:1602],
	45591:    _Code_name[1602:1613],
	45592:    _Code_name[1613:1624],
	45593:    _Code_name[1624:1635],
	45594:    _Code_name[1635:1646],
	45595:    _Code_name[1646:1657],
	45596:    _Code_name[1657:1668],
	45597:    _Code_name[1668:1679],
	45598:    _Code_name[1679:1690],
	45599:    _Code_name[1690:1701],
	45600:    _Code_name[1701:1712],
	45601:    _Code_name[1712:1723],
	45602:    _Code_name[1723:1734],
	45603:    _Code_name[1734:1745],
	45604:    _Code_name[1745:1756],
	45605:    _Code_name[1756:1767],
	45606:    _Code_name[1767:1778],
	45607:    _Code_name[1778:1789],
	45608:    _Code_name[1789:1800],
	45609:    _Code_name[1800:1811],
	45610:    _Code_name[1811:1822],
	45611:    _Code_name[1822:1833],
	45612:    _Code_name[1833:1844],
	45613:    _Code_name[1844:1855],
	45614:    _Code_name[1855:1866],
	45615:    _Code_name[1866:1877],
	45616:    _Code_name[1877:1888],
	45617:    _Code_name[1888:1899],
	45618:    _Code_name[1899:1910],
	45619:    _Code_name[1910:1921],
	45620:    _Code_name[1921:1932],
	45621:    _Code_name[1932:1943],
	45622:    _Code_name[1943:1954],
	45623:    _Code_name[1954:1965],
	45624:    _Code_name[1965:1976],
	45625:    _Code_name[1976:1987],
	45626:    _Code_name[1987:1998],
	45627:    _Code_name[1998:2009],
	45628:    _Code_name[2009:2020],
	45629:    _Code_name[2020:2031],
	45630:    _Code_name[2031:2042],
	45631:    _Code_name[2042:2053],
	45632:    _Code_name[2053:2064],
	45633:    _Code_name[2064:2073],
	45634:    _Code_name[2073:2083],
	45635:    _Code_name[2083:2093],
	45636:    _Code_name[2093:2103],
	45637:    _Code_name[2103:2113],
	45638:    _Code_name[2113:2123],
	45639:    _Code_name[2123:2133],
	45640:    _Code_name[2133:2143],
	45641:    _Code_name[2143:2153],
	45642:    _Code_name[2153:2163],
	45643:    _Code_name[2163:2173],
	45644:    _Code_name[2173:2183],
	45645:    _Code_name[2183:2194],
	45646:    _Code_name[2194:2205],
	45647:    _Code_name[2205:2216],
	45648:    _Code_name[2216:2227],
	45649:    _Code_name[2227:2238],
	45650:    _Code_name[2238:2249],
	45651:    _Code_name[2249:2260],
	45652:    _Code_name[2260:2271],
	45653:    _Code_name[2271:2282],
	45654:    _Code_name[2282:2293],
	45655:    _Code_name[2293:2304],
	45656:    _Code_name[2304:2315],
	45657:    _Code_name[2315:2326],
	45658:    _Code_name[2326:2337],
	45659:    _Code_name[2337:2348],
	45660:    _Code_name[2348:2359],
	45661:    _Code_name[2359:2370],
	45662:    _Code_name[2370:2381],
	45663:    _Code_name[2381:2392],
	45664:    _Code_name[2392:2403],
	45825:    _Code_name[2403:2413],
	45826:    _Code_name[2413:2424],
	45827:    _Code_name[2424:2435],
	45828:    _Code_name[2435:2446],
	45829:    _Code_name[2446:2457],
	45830:    _Code_name[2457:2468],
	45831:    _Code_name[2468:2479],
	45832:    _Code_name[2479:2490],
	45833:    _Code_name[2490:2501],
	45834:    _Code_name[2501:2512],
	45835:    _Code_name[2512:2523],
	45836:    _Code_name[2523:2534],
	45837:    _Code_name[2534:2546],
	45838:    _Code_name[2546:2558],
	45839:    _Code_name[2558:2570],
	45840:    _Code_name[2570:2582],
	45841:    _Code_name[2582:2594],
	45842:    _Code_name[2594:2606],
	45843:    _Code_name[2606:2618],
	45844:    _Code_name[2618:2630],
	45845:    _Code_name[2630:2642],
	45846:    _Code_name[2642:2654],
	45847:    _Code_name[2654:2666],
	45848:    _Code_name[2666:2678],
	45849:    _Code_name[2678:2690],
	45850:    _Code_name[2690:2702],
	45851:    _Code_name[2702:2714],
	45852:    _Code_name[2714:2726],
	45853:    _Code_name[2726:2738],
	45854:    _Code_name[2738:2750],
	45855:    _Code_name[2750:2762],
	45856:    _Code_name[2762:2774],
	45857:    _Code_name[2774:2784],
	45858:    _Code_name[2784:2795],
	45859:    _Code_name[2795:2806],
	45860:    _Code_name[2806:2817],
	45861:    _Code_name[2817:2828],
	45862:    _Code_name[2828:2839],
	45863:    _Code_name[2839:2850],
	45864:    _Code_name[2850:2861],
	45865:    _Code_name[2861:2872],
	45866:    _Code_name[2872:2883],
	45867:    _Code_name[2883:2894],
	45868:    _Code_name[2894:2905],
	45869:    _Code_name[2905:2917],
	45870:    _Code_name[2917:2929],
	45871:    _Code_name[2929:2941],
	45872:    _Code_name[2941:2953],
	45873:    _Code_name[2953:2965],
	45874:    _Code_name[2965:2977],
	45875:    _Code_name[2977:2989],
	45876:    _Code_name[2989:3001],
	45877:    _Code_name[3001:3013],
	45878:    _Code_name[3013:3025],
	45879:    _Code_name[3025:3037],
	45880:    _Code_name[3037:3049],
	45881:    _Code_name[3049:3061],
	45882:    _Code_name[3061:3073],
	45883:    _Code_name[3073:3085],
	45884:    _Code_name[3085:3097],
	45885:    _Code_name[3097:3109],
	45886:    _Code_name[3109:3121],
	45887:    _Code_name[3121:3133],
	45888:    _Code_name[3133:3145],
	45889:    _Code_name[3145:3157],
	45890:    _Code_name[3157:3169],
	45891:    _Code_name[3169:3181],
	45892:    _Code_name[3181:3193],
	45893:    _Code_name[3193:3205],
	45894:    _Code_name[3205:3217],
	45895:    _Code_name[3217:3229],
	45896:    _Code_name[3229:3241],
	45897:    _Code_name[3241:3253],
	45898:    _Code_name[3253:3265],
	45899:    _Code_name[3265:3277],
	45900:    _Code_name[3277:3289],
	45901:    _Code_name[3289:3301],
	45902:    _Code_name[3301:3313],
	45903:    _Code_name[3313:3325],
	45904:    _Code_name[3325:3337],
	45905:    _Code_name[3337:3349],
	45906:    _Code_name[3349:3361],
	45907:    _Code_name[3361:3373],
	45908:    _Code_name[3373:3385],
	45909:    _Code_name[3385:3397],
	45910:    _Code_name[3397:3409],
	45911:    _Code_name[3409:3421],
	45912:    _Code_name[3421:3433],
	45913:    _Code_name[3433:3445],
	45914:    _Code_name[3445:3457],
	45915:    _Code_name[3457:3469],
	45916:    _Code_name[3469:3481],
	45917:    _Code_name[3481:3493],
	45918:    _Code_name[3493:3505],
	45919:    _Code_name[3505:3517],
	45920:    _Code_name[3517:3529],
	45921:    _Code_name[3529:3540],
	45922:    _Code_name[3540:3552],
	45923:    _Code_name[3552:3564],
	45924:    _Code_name[3564:3576],
	45925:    _Code_name[3576:3588],
	45926:    _Code_name[3588:3600],
	45927:    _Code_name[3600:3612],
	45928:    _Code_name[3612:3624],
	45929:    _Code_name[3624:3636],
	45930:    _Code_name[3636:3648],
	45931:    _Code_name[3648:3660],
	45932:    _Code_name[3660:3672],
	45933:    _Code_name[3672:3685],
	45934:    _Code_name[3685:3698],
	45935:    _Code_name[3698:3711],
	45936:    _Code_name[3711:3724],
	45937:    _Code_name[3724:3737],
	45938:    _Code_name[3737:3750],
	45939:    _Code_name[3750:3763],
	45940:    _Code_name[3763:3776],
	45941:    _Code_name[3776:3789],
	45942:    _Code_name[3789:3802],
	45943:    _Code_name[3802:3815],
	45944:    _Code_name[3815:3828],
	45945:    _Code_name[3828:3841],
	45946:    _Code_name[3841:3854],
	45947:    _Code_name[3854:3867],
	45948:    _Code_name[3867:3880],
	45949:    _Code_name[3880:3893],
	45950:    _Code_name[3893:3906],
	45951:    _Code_name[3906:3919],
	45952:    _Code_name[3919:3932],
	45953:    _Code_name[3932:3945],
	45954:    _Code_name[3945:3958],
	45955:    _Code_name[3958:3971],
	45956:    _Code_name[3971:3984],
	45957:    _Code_name[3984:3997],
	45958:    _Code_name[3997:4010],
	45959:    _Code_name[4010:4023],
	45960:    _Code_name[4023:4036],
	45961:    _Code_name[4036:4049],
	45962:    _Code_name[4049:4062],
	45963:    _Code_name[4062:4075],
	45964:    _Code_name[4075:4088],
	45965:    _Code_name[4088:4101],
	45966:    _Code_name[4101:4114],
	45967:    _Code_name[4114:4127],
	45968:    _Code_name[4127:4140],
	45969:    _Code_name[4140:4153],
	45970:    _Code_name[4153:4166],
	45971:    _Code_name[4166:4179],
	45972:    _Code_name[4179:4192],
	45973:    _Code_name[4192:4205],
	45974:    _Code_name[4205:4218],
	45975:    _Code_name[4218:4231],
	45976:    _Code_name[4231:4244],
	45977:    _Code_name[4244:4257],
	45978:    _Code_name[4257:4270],
	45979:    _Code_name[4270:4283],
	45980:    _Code_name[4283:4296],
	45981:    _Code_name[4296:4309],
	45982:    _Code_name[4309:4322],
	45983:    _Code_name[4322:4335],
	45984:    _Code_name[4335:4348],
	45985:    _Code_name[4348:4361],
	45986:    _Code_name[4361:4374],
	45987:    _Code_name[4374:4387],
	45988:    _Code_name[4387:4400],
	45989:    _Code_name[4400:4413],
	45990:    _Code_name[4413:4426],
	45991:    _Code_name[4426:4439],
	45992:    _Code_name[4439:4452],
	45993:    _Code_name[4452:4465],
	45994:    _Code_name[4465:4478],
	45995:    _Code_name[4478:4491],
	45996:    _Code_name[4491:4504],
	45997:    _Code_name[4504:4517],
	45998:    _Code_name[4517:4530],
	45999:    _Code_name[4530:4543],
	46000:    _Code_name[4543:4556],
	46001:    _Code_name[4556:4569],
	46002:    _Code_name[4569:4582],
	46003:    _Code_name[4582:4595],
	46004:    _Code_name[4595:4608],
	46005:    _Code_name[4608:4621],
	46006:    _Code_name[4621:4634],
	46007:    _Code_name[4634:4647],
	46008:    _Code_name[4647:4660],
	46009:    _Code_name[4660:4673],
	46010:    _Code_name[4673:4686],
	46011:    _Code_name[4686:4699],
	46012:    _Code_name[4699:4712],
	46013:    _Code_name[4712:4725],
	46014:    _Code_name[4725:4738],
	46015:    _Code_name[4738:4751],
	46016:    _Code_name[4751:4764],
	46017:    _Code_name[4764:4777],
	46018:    _Code_name[4777:4790],
	46019:    _Code_name[4790:4803],
	46020:    _Code_name[4803:4816],
	46021:    _Code_name[4816:4829],
	46022:    _Code_name[4829:4842],
	46023:    _Code_name[4842:4855],
	46024:    _Code_name[4855:4868],
	46025:    _Code_name[4868:4881],
	46026:    _Code_name[4881:4894],
	46027:    _Code_name[4894:4907],
	46028:    _Code_name[4907:4920],
	46029:    _Code_name[4920:4933],
	46030:    _Code_name[4933:4946],
	46031:    _Code_name[4946:4959],
	46032:    _Code_name[4959:4972],
	46033:    _Code_name[4972:4985],
	46034:    _Code_name[4985:4998],
	46035:    _Code_name[4998:5011],
	46036:    _Code_name[5011:5024],
	46037:    _Code_name[5024:5037],
	46038:    _Code_name[5037:5050],
	46039:    _Code_name[5050:5063],
	46040:    _Code_name[5063:5076],
	46041:    _Code_name[5076:5089],
	46042:    _Code_name[5089:5102],
	46043:    _Code_name[5102:5115],
	46044:    _Code_name[5115:5128],
	46045:    _Code_name[5128:5142],
	46046:    _Code_name[5142:5156],
	46047:    _Code_name[5156:5170],
	46048:    _Code_name[5170:5184],
	46081:    _Code_name[5184:5209],
	46082:    _Code_name[5209:5237],
	52753:    _Code_name[5237:5258],
	61697:    _Code_name[5258:5281],
	61698:    _Code_name[5281:5302],
	7367777:  _Code_name[5302:5313],
	8417572:  _Code_name[5313:5329],
	8483108:  _Code_name[5329:5345],
	9728292:  _Code_name[5345:5361],
	9793828:  _Code_name[5361:5377],
	10645796: _Code_name[5377:5393],
	10711332: _Code_name[5393:5409],
	11639056: _Code_name[5409:5418],
	11704592: _Code_name[5418:5428],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}
