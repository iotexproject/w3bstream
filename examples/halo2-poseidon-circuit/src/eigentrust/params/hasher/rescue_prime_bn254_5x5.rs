use super::*;
use halo2_curves::bn256::Fr;

#[derive(Clone, Debug)]
/// Configures a structure.
pub struct Params;

impl Sbox for Params {
    fn sbox_expr<F: FieldExt>(exp: Expression<F>) -> Expression<F> {
        let exp2 = exp.clone() * exp.clone();
        let exp4 = exp2.clone() * exp2;
        exp4 * exp
    }

    fn sbox_f<F: FieldExt>(f: F) -> F {
        let f2 = f * f;
        let f4 = f2 * f2;
        f4 * f
    }

    fn sbox_inv_f<F: FieldExt>(f: F) -> F {
        // Pow by inverse of 5
        f.pow([
            14981214993055009997,
            6006880321387387405,
            10624953561019755799,
            2789598613442376532,
        ])
    }
}

impl RoundParams<Fr, 5> for Params {
    fn partial_rounds() -> usize {
        0
    }

    fn full_rounds() -> usize {
        8
    }

    fn round_constants_raw() -> Vec<&'static str> {
        [
            "0x2ca4a8f397401b89877423695357441aa6f56488841db4b55ffb8d911b947e44",
            "0x245a7b16a5b3b10cf19fbec864eb6fc2559d88e25ccc9b27b1c6b7dede3c9190",
            "0x19e20da5a280c6b52b228c4cc14bec5ac4e39f3fd464dbe6a0203de8b5d7d795",
            "0x0f8263b25f9b8a46cfd68979b111feecf4b734ef44093df98aeed58c36803238",
            "0x11d3d70e0ce66186372e7fe386ad3f5ccf2c811231720e3f0192a8b8995eedfe",
            "0x01cfc193c968b4c6d20ec2df4f110906a8360bc62150ca0400a53620699932bf",
            "0x2d69bbded4695ac9a93ff4a23a17c103376d0ea8c7c5663ad283ec8428b51e25",
            "0x071237ec067c09821a6045191b2281cbd82186de8ba1374cbc96b87dc1b61c07",
            "0x24161b3910f19b028ffadca01be8fa665dfccab8c044cdc3397cdaac1666a08d",
            "0x22e6c619a6ecf9a856dbe0e5306d0b99eb4d7ad1b1a6d33ce3ae4ef635dedfc6",
            "0x12fe8da5193f1d8310d56f9666c2818cb773104fb5e16daf89956ac6e0ec4cec",
            "0x103dc1169a48ca6394a5f58f327b5ec3b34f35f5be85be08def563c534614fbb",
            "0x0b40a5360e10b322fb67174e2df67079dd7abd217d125c0251482a15ae6e14eb",
            "0x13f6c28b697c88a387a69d712e32fbea0350e1a6233ef755df28168d434531b4",
            "0x10aa4f4aafc4ac840abf70f0efdd37a9e5b103152ff754a45e998243444177d4",
            "0x123fdba3acb01a1af93ab0beeb66dc7e3b6c455ede611a5eddd1c63da1306f32",
            "0x14b11a1f018a3684f810508727a472808f620d9b9cdd28708735267a3969fd93",
            "0x14e6815fc9838e55a9e8681bd8781d63eaa51f494ab1168b38fee42f33451297",
            "0x20d3e0623f8878b6ea3cc84c227edb0d7f92602593f0b19bceea9bbdfb044ce8",
            "0x2e866c6a7fbbbfddfeb0bda9f375ab1f7c406c21b165d46f1728149544ac7b2e",
            "0x0af23ac272eb2c056d6e0f10df465150330663a97e3ce6c90ab1224d32dddda7",
            "0x00e5e60f14b9d190aaa3e7944b90c1bbd231692e1dbf8e5869dad2d9f708e9c6",
            "0x17b44d163738b1ea585c1ec820ee91fe89cf92f9221bd46dc29c6bdf9542b90b",
            "0x28ef1c64cfa2b272bec8850b67f7c340cb146d9c084d59119d4befc9e4a683a4",
            "0x1c36102272038b2518169af737850e34d421e5b8e96c4f6d6496ee880adc544b",
            "0x19aa5606180777d0437629fb73fa97338875d2332d8d478df1de9061269092fa",
            "0x1aec324236183446d97fa67c33bcb6128b8a7dc85f1686f9d5bf2c34ce3663e1",
            "0x1411983e589c744567fe19e737b9d2c942923d466406de9642f401f72f1d3ec0",
            "0x0900bd14426240fc675ca4a3c84a2c85fe6e958ccf92ffa49a72d182ce67ff30",
            "0x0edfdabb93d0650b2d0da4aa90f83dd46a3d6d868534b611f01c6efde225570c",
            "0x05e1e0c7ce6bf5e75308fd960c058d0cf47da1bee7cf6c5f730a76aab750fa3e",
            "0x27902233323d1871fb83f97c40480264391deedb2fcee12864adf503d442f144",
            "0x034927884e7a8042701e54e5c08d97cb80a6875c45390c6c5dd25c642f3cf233",
            "0x168bec5a3792b86e6d5d946fe3697653401501b080e03fb4a5d5b643752ab01a",
            "0x24f01f489c63d0812decff6cbdebadbb23b3a469fe7b1ba29c82dd5cb15a0642",
            "0x0f629e3f29346352ae58ff95d2a2faaafe60abb5d0129c212024e1d299d69670",
            "0x0ad63006a8c4dc5183f961e485d1271865f7e217e2165f391f62a30cc0f31013",
            "0x0360d6a52b4059c7b6f1a9784f2f60e953cb143e265cd90c9edeb5db3c5f13d9",
            "0x2f17067d317d1ccc87726275d791e7c0e613a344b703171443a772cd745f4f68",
            "0x0784da44edbf903e6f5aa79f7b4591d39bfdf1f56c2ac7b5f2b0520fcdf4a4cc",
        ]
        .to_vec()
    }

    fn mds_raw() -> [[&'static str; 5]; 5] {
        [
            [
                "0x0d5365c702a1d156ecb069c1a5b6f3fefa94d552d01f19e0cea4ba91e4537c0b",
                "0x0e0ec9e239372ffcc147499c18127fcd14df23c28afea070d9a5ca5a8e2edad1",
                "0x139b19872a35491f40fb3ae9e5d85fc5af24bd86e334c4dd256da7318de3d3b0",
                "0x0b9abb8cac6d528f3cd7583d153402fb66d1bdd7fad8bfc993ffbd462c26dcb9",
                "0x25ef5785f15274631fcb54ee87fe105c042eaf901054d5cb2ea1f140c7795eda",
            ],
            [
                "0x116b9e8295f0086e020f035edc4ec2b5e65187597edb3d37de19589968b8d95a",
                "0x17f2bb6779aa19aa40b9903101f34f3079b64947538c68eec839700e4d3f4dc6",
                "0x2748eafccb493093f5b6b8eed565ed7c5d6a196cfd09afaadcc1e4ef41723ff2",
                "0x177e934ef325737ab20d869f5570b6d442938a9ba31a69931fdce3b5681f3dab",
                "0x20be0041859b35e0ebfa48cb6b55623cdb72cca5866939f0854e2ec3271c9d28",
            ],
            [
                "0x1c794b13d4b66d9883849539f556b020117666918dae1210cc4c713c4137c980",
                "0x196241af44073291925c16017a812df6b64258f4b0e570fb12381945105fd50f",
                "0x0820fbc1ecd8b272543eb1fdcfb2a04a8897b42d545ad050fa4743ba248df73b",
                "0x1d630dd3a25b7aaa9684f4f641b9d763ebd202d0dc44245aaac81e37379c822b",
                "0x0d79e1f5cf020c1d584cb948c22d55a9d5f98c9b2b6eae028d346fa5cb0ecebc",
            ],
            [
                "0x1f9743a6cceebff6284adf5b93119ffa18ac050da856193e19f9f0d8b5673ea4",
                "0x023d9a2dbf719235ecd411ae6234fb6fe55a1ff2be57568089c198aa54888aef",
                "0x11a2ed471ebf4a4fb2471cf834a842b85c09ebadd54cb6decd9d42a1a026b7c4",
                "0x2b32d181fb89f8afa92d9139bee4d4eabfb1b2e3831499dcf1b261913c3b978c",
                "0x1894483c1c129daae2eae33b791ab6709b2d387453f531f5def0fa4c89d07947",
            ],
            [
                "0x14de4fd37b770cabbb10d2ca60e89e300d24cc1c04e936d6d3bb66ac2953769b",
                "0x2caaef6f5c7fdc4c9b039d78d96c05eacf06e2e72b108d420001fcb9ba1883ae",
                "0x144dd8dffcabe999775811f306b50ccb43bcb1d5c6a661e89d095bd2aed983c5",
                "0x2b535d04594ad253957b83629c6899a05f0332764aa5ab5708ddb230a9f7e722",
                "0x0819a28061b1b476f1c62584e24040eea0c8a6afd5e052b8823eb3bee7ecf144",
            ],
        ]
    }
}
