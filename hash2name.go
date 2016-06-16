package ikaring

var weapons = map[string]string{
	"3359908b8a0baf5b2c28402ad74d1d754c5adc3f9971832fac75fd4344f5c428": ".52ガロン",
	"b73d45cca8d831e9b032f0642e4df95fcf1be1843df99e289b315adcab3c450d": ".52ガロンデコ",
	"24216c6e64d6d4a97cb76c642233b678a7d663a266a6501714c58094c3727766": ".96ガロン",
	"d1f55d0d3a8a5bfd209e87b75b67567c98038873eef190bc25ef8ed1027d1d2c": ".96ガロンデコ",
	"8a5e4de6e62b281659e0573fd95576debfd96fb03d08ee6b0e9a81401a719dc9": "14式竹筒銃・丙",
	"27151f9470ea6763ecf45269efae872ac1b8bb84d6efaae95a1cad01cf12791e": "14式竹筒銃・乙",
	"f9a34915ba879bea89ffb3293fc91ddb19d2fc4acf8a481f09b1222bc98ab080": "14式竹筒銃・甲",
	"7ad0f51b984c6144e99cbda675e0234c513733c9ba4b9a100cf4487b2999f618": "3Kスコープ",
	"6c76f1afc2956b053e35e649800a21aa2efe44518f7208d2ccb8d179e4f02247": "3Kスコープカスタム",
	"5fb8ad7830c398b738bc972e25a966b6cf6e010b848c2f0dee2f8d936898158b": "H3リールガン",
	"bd2f37f3bea997e4c6752a5e30657159c40d4e45d46b7e4f73e927b3d15ec261": "H3リールガンD",
	"ddd6a002bf375420d00e996de139475a977f17898dfa33fa3a4661d77976c38a": "H3リールガンチェリー",
	"44e2035a09dfc82a2484c48bb69ded59354e45cd4a2062f60a1fa2d0871773e0": "L3リールガン",
	"aab295e6db118bf0100e55885e1647df69d0f37ed7f3d1deae38a06082d175a6": "L3リールガンD",
	"45eab6318ef9aa736036106a96d09ec4effafb118860ce6fd10036f301359b78": "N-ZAP83",
	"ec3a13c7f0aef55dad5bb7bd3db5be6f21c943811d7cad14c6a74b18ab199605": "N-ZAP85",
	"aeb8e1fe59cf32f8755a44a704267405322b17b6926f1692e86aa9c9924dd795": "N-ZAP89",
	"96dcf4f0b4224c5eb9b07850af74912358b8eb521ac3668ad72ffc78af3861ca": "Rブラスターエリート",
	"9420d3bda417cf06d492fe857c291467bee2890dbdfc9063f06ab5b606d80442": "Rブラスターエリートデコ",
	"eb49639a24f7763d18c3d3375ec660395ae31bdf4227166f3c255d1f0b142717": "オクタシューター レプリカ",
	"e088d953b16e68118a3036a54b0eab02cdc00ea61ba4abfc85f1f98e83d49f4c": "カーボンローラー",
	"5fb8898f8ce4c15bb16ec7530531f90743475ea2824561e7ac7b386b330a4364": "カーボンローラーデコ",
	"be4595427ab39aec0bd5186280bc2003fe6003120e3a887124a8ebb1f65c9a18": "シャープマーカー",
	"421271a78563e434a4b79579be720a33fd5cee461cf517862858b6fed6a0e058": "シャープマーカーネオ",
	"d97d42d7ac13478131330fbf3de80fc85c78a4d50fcc11cea7d9776385615396": "ジェットスイーパー",
	"6c38c99840cb78ad525e1e53e1c686397b1360e83ef71b9dd4cb62ad2ce38d92": "ジェットスイーパーカスタム",
	"686933ec4f43c67c153fd62984601c37d2a136c42878405724bb65020700937a": "スクイックリンα",
	"fe8da69077740b1db12a2e1cc2375231271fee3d05ee0f5ba54d9bde83a4baab": "スクイックリンβ",
	"017d74dcd5f231dd4c1eeb3b49e850cf253926b99fde425e1ad8b617707245a4": "スクイックリンγ",
	"a63ba19229be34cc366f6fab6e2a79d3fb4d4901cbeb4ea23170dab7e9beb7cd": "スクリュースロッシャー",
	"58105e03607ff7ee7f8837c42ac46d4500a65b3481845a07abea5d7a0c7dc9ea": "スクリュースロッシャーネオ",
	"2fee2bac1e22b2ef19107e3eb58aabc4fd54e9ced476d1d2ac0b7d176d11d2ba": "スプラシューター",
	"e67b1303cfe0d3335e7e078244c84c17792e2fb54dbd789433fecaf102bd3362": "スプラシューターコラボ",
	"94ba5a26a250d54de3510d2c252d40a5f489eee19563b8e153f73128879e19de": "スプラシューターワサビ",
	"a1c5e419c8b284aa684726d27b2472b69d1a7215d570579d85ad657da7180f9a": "スプラスコープ",
	"0005d0c1f017e6fe6601282187125619b2c036fdb26ec8b948c8c40d3449159c": "スプラスコープベントー",
	"b4c444616489c4e38cdb3ce282d6f69ba94f3e330b7d76d7208e8a672378ea8b": "スプラスコープワカメ",
	"612e90b58d01f14c1f6858f933f55656d09d83283f125cdd272b58e1555c1186": "スプラスピナー",
	"2e3b37440632af8895518ad0d97b7867e1823a50e1dd5039912d07e4adc5cc7a": "スプラスピナーコラボ",
	"73369ac2737480e9e992bc60d9c72e5f0a52fa2c6a9efeff88de6ce0d376b25a": "スプラスピナーリペア",
	"234817cbabd95f27d60fb213122c1b98ef41a13457aa6b57a2e7dbf5f2fb8c71": "スプラチャージャー",
	"1a78b7b9c89c10a908b718f21a2efe4aa49f7b0e1b8e67f1a135ee89720b0a8a": "スプラチャージャーベントー",
	"1e971e6aecaa03575042e8050a0c8431edbec68b72f41527c1ca6ac9fd024741": "スプラチャージャーワカメ",
	"062629ff15ff75912af49ca5efc53e30e0fa183543774d0effa5d06939dddd74": "スプラローラー",
	"ff801fb34b6acf0ccb763f2f98c98553ce5d298a144855b36ffed89730e358d2": "スプラローラーコラボ",
	"fc4118d94dfbff16ce4283298a86d27530e02bffbecdf8fbd4928ed0457fc02b": "スプラローラーコロコロ",
	"cd4b9d0fab7239d83e0d728826e565ba6bf64a8fa63eb81489345e1276c6600b": "ダイナモローラー",
	"6009c0a233e080a8ab70d0fd291c23098e12bbf877cee983f5da31b36be4997a": "ダイナモローラーテスラ",
	"57b4fbb7dce52c9cc99734964109f85b67b7e6c8e8edc6eda2820e8f7f7f5f82": "ダイナモローラーバーンド",
	"c624ffec5c28352485cb210ee152756cb6ca486b44e3b10dc4b9bf04ab953927": "デュアルスイーパー",
	"87293e3e06bd09356e0f7515d77ac7bd00be416492c6e8c21879a627cd3babb9": "デュアルスイーパーカスタム",
	"36fe4791f65cc1db834a4d040b22a2ea637e8ab75ad37ce3db4caea3247e30af": "ノヴァブラスター",
	"6b34ccae7b426b59c68e97fa699f0f5a1f799f524bfbdeb52989e0fce9a53039": "ノヴァブラスターネオ",
	"dd1711429bdfe5bd809762b23811e648e47f1d7bf51f48de6f7759cb68ec15aa": "ハイドラント",
	"b55b9b85d532918a14c0c986a6962aadb619944dac31fd11b28208ef6836b90c": "ハイドラントカスタム",
	"251c94c0f851f5227d7795735af968513ac8f008cd024043e86fdb87b1f089fa": "バケットスロッシャー",
	"cd409da57775c7ce7b987195c426d1c075669350538ad8f8eeb756e43468fbe1": "バケットスロッシャーソーダ",
	"f4a3cedfdd3d4119ddcc5c3ea2d327e63e1159250f2fccef4cc3692bfd3de973": "バケットスロッシャーデコ",
	"08ac41b5dfeb33c35fa2d03269f78ea4dad63b88cc6be3221b733897c7e82135": "バレルスピナー",
	"6371a40e3e9e739422f28a44cf30c3e613ee39ddcc021aba9c5c7909250539a0": "バレルスピナーデコ",
	"91fb83a0943ba211e2e13281758b5f4de5f1989647f9b8810e4c5823722b1e0c": "バレルスピナーリミックス",
	"ad5ce637d338080d1ce2c2703370f8616b4499eca84ecca24136b848de967331": "パブロ",
	"de94a1dabe2d403274cc65d08516ae25db2c13fe9b9447f114612747a6151767": "パブロ・ヒュー",
	"739840486a61b8252795a6668d02ce21f4b08415c95f7f737dbd77aba49cfa32": "パーマネント・パブロ",
	"79d770d4a8c5d062756ee1626ca635dc20acbe0c23ca20c0cf5c3891848cb965": "ヒッセン",
	"3f309d25639d80baa19485f7ae6099c3d82ae3ef8ffd1f69786ecf7f868bc124": "ヒッセン・ヒュー",
	"37ea525217902df9a399b2ddac6b5443daa1e16cbb65f829c8427246ba12b13f": "ヒーローシューター レプリカ",
	"c5b48a8e03781a6d34f2978ac3a1a12f90b301be8f1f4d622a1796697ba603ce": "ヒーローチャージャーレプリカ",
	"b5fe8e89fe8c36b356f2265eb92096571528fbd40776391e853b262c37a0f8a0": "ヒーローローラーレプリカ",
	"b44b45fcc1cfc3862796f01f5a99ba119929bbead5b2a89f316e90f481abcd4c": "プライムシューター",
	"43a94c610c74e6d92b75424133d47d01b97b91f7fdb7ea904e385e2ed20c6149": "プライムシューターコラボ",
	"0b1c45c1490a76251a4a4e4e7d51521ffec6ef07098bba811cb4dbd24694980d": "プライムシューターベリー",
	"7bef265d011060ef6cf88033ec2ecd6a1c73790caefe7bcbd7be323cc3584cbd": "プロモデラーMG",
	"1610a5bfcddbd258fed1bd2262aa7eafb867d4c4b8dd7bfb3493c79d51ebf750": "プロモデラーPG",
	"de63b9c7c5f9f5ab76456328b11f2f6ac8f17c7c91ff402727625fc46d0a3055": "プロモデラーRG",
	"efe0cba0dfeb40bbea197c9ead9241bb719b8ed6bcd2131b339c45df051379b2": "ホクサイ",
	"8707d8047c8dbd35d6cefa2a08d55be9f939dcaea4db44b7e0e4aa599736cb01": "ホクサイ・ヒュー",
	"282fd792a220632b6c34bfbb564629a67b9f540d1a1ab212f9e2f30c9174a820": "ホットブラスター",
	"d78618b66b7562fbbb1a2455fa959216204890da362aca1b0726e09aad7b9408": "ホットブラスターカスタム",
	"56f4e14ac7cb57af50030071da35700af501985ceca37351330eee6df0c2c2da": "ボールドマーカー",
	"67980ff4205ba0792b96f6f19ccb6438c4b1a10d87d194f8d4261be33bb48805": "ボールドマーカー7",
	"4c30e234da557cb93ea4e7282005edcf4112dc1b6a43124ad21f6c3b3366cfad": "ボールドマーカーネオ",
	"259b88ef89bc14c74559a6d05a14efa9482557e89190a39748a79031524b60c8": "もみじシューター",
	"142657a7ff4bb94523b1dd730721a2020828e83e9dba2fe2dfe395c740989ae1": "ラピッドブラスター",
	"eeece3664b74f630f68f923b2a423e0ba3a1859d4abf36a1ee96056025de1f1e": "ラピッドブラスターデコ",
	"9f8bc13b5a7c929e0bb4b5e46f93a197b8407e55da3c421c3ef59979fb4ea8f2": "リッター3K",
	"230a8eee09130d6cf9a1d32d29808fd5f0110a58d4998130cec89c4eaf58a977": "リッター3Kカスタム",
	"cfb91dda41634286360f37d61172e8c3a276ae2f89f7719c572f68576ea3ad7e": "ロングブラスター",
	"9f39d37e7eb03c0d975d72710602452c6e939c5faf8d44c1cc5aa3766dc670ec": "ロングブラスターカスタム",
	"456535821d077524a06b46347ac177be89a164c6333932a0ff17ad2fd3a2059c": "ロングブラスターネクロ",
	"c9f7cead9ee5ada35437d7c2ea8ddae6ca1dacfc6c9b01d5939cfc0ff59fe0ea": "わかばシューター",
}
