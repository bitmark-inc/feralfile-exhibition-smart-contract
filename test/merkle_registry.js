const { expect } = require("chai");
const { BN, expectEvent } = require("@openzeppelin/test-helpers");
const abi = require("../build/contracts/MerkleRegistry.json").abi;
const { expectCustomError } = require("./helper/expectCustomError.js");

const MerkleRegistry = artifacts.require("MerkleRegistry");
const SOURCE_CODE_URI = "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
const PROOF_URI = "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
const METADATA_URI = "ipfs://QmaQegRqExfFx8zuR6yscxzUwQJUc96LuNNQiAMK9BsUQe";
const MERKLE_ROOT =
    "0xd7d0e7f3496001d4600e37fa978ccafd1f68bdcf743311c69c28442103f618a1";
const LEAF =
    "0x5bd4c9770efe2fe95612f5cbc1ad0b6e4c158f2351f1ca136700f21d37ed2b5d";
const PROOF = [
    "0x765556eae309fd86ffd22d312a0ce30d62a6d7e2c19f21f55a59a51222bfad55",
    "0x9ee46f6499d611144bb3cdfedd69378d20958a79d31daa1a64c506bd23b70745",
    "0x02008afae295b32f4303bb511cb527bd87a34d853488f9f882e7056fcffdaafa",
];

contract("MerkleRegistry", async (accounts) => {
    let contract;

    beforeEach(async () => {
        contract = await MerkleRegistry.new(SOURCE_CODE_URI);
    });

    it("test constructor", async () => {
        expect(await contract.sourceCodeURI()).to.equal(SOURCE_CODE_URI);
        expect(await contract.getEntryCount()).to.be.a.bignumber.that.equals(
            new BN(0)
        );
    });

    it("test batchRegisterMerkleRoots", async () => {
        let submitter = accounts[0];

        // Test successful registration
        let tx = await contract.batchRegisterMerkleRoots(
            [MERKLE_ROOT],
            [PROOF_URI],
            [METADATA_URI],
            { from: submitter }
        );

        expectEvent(tx, "MerkleRootRegistered", {
            index: new BN(0),
            merkleRoot: MERKLE_ROOT,
            submitter: submitter,
            proofURI: PROOF_URI,
            metadataURI: METADATA_URI,
        });

        // Get block timestamp from the transaction receipt
        const receipt = await web3.eth.getTransactionReceipt(tx.tx);
        const block = await web3.eth.getBlock(receipt.blockNumber);
        const timestamp = block.timestamp;

        // Verify the registration
        expect(await contract.getEntryCount()).to.be.a.bignumber.that.equals(
            new BN(1)
        );
        expect(
            await contract.rootToIndex(MERKLE_ROOT)
        ).to.be.a.bignumber.that.equals(new BN(0));

        // Verify the entry by root
        let entry = await contract.getMerkleEntryByRoot(MERKLE_ROOT);
        expect(entry.merkleRoot).to.equal(MERKLE_ROOT);
        expect(entry.proofURI).to.equal(PROOF_URI);
        expect(entry.metadataURI).to.equal(METADATA_URI);
        expect(entry.submitter).to.equal(submitter);
        expect(entry.timestamp.toString()).to.equal(timestamp.toString());

        // Verify the entry by index
        entry = await contract.getMerkleEntryByIndex(0);
        expect(entry.merkleRoot).to.equal(MERKLE_ROOT);
        expect(entry.proofURI).to.equal(PROOF_URI);
        expect(entry.metadataURI).to.equal(METADATA_URI);
        expect(entry.submitter).to.equal(submitter);
        expect(entry.timestamp.toString()).to.equal(timestamp.toString());

        // Test revert due to existing root
        await expectCustomError(
            contract.batchRegisterMerkleRoots(
                [MERKLE_ROOT],
                [PROOF_URI],
                [METADATA_URI],
                { from: submitter }
            ),
            abi,
            "MerkleRootAlreadyExists",
            []
        );

        // Test revert due to zero roots
        await expectCustomError(
            contract.batchRegisterMerkleRoots(
                ["0x0"],
                [PROOF_URI],
                [METADATA_URI],
                { from: submitter }
            ),
            abi,
            "MerkleRootIsZero",
            []
        );

        // Test revert due to empty proofURI
        await expectCustomError(
            contract.batchRegisterMerkleRoots(
                [MERKLE_ROOT],
                [""],
                [METADATA_URI],
                { from: submitter }
            ),
            abi,
            "ProofURIIsEmpty",
            []
        );
    });

    it("test updateMerkleRootProofURI", async () => {
        let submitter = accounts[0];

        // Register a new Merkle root
        await contract.batchRegisterMerkleRoots(
            [MERKLE_ROOT],
            [PROOF_URI],
            [METADATA_URI],
            { from: submitter }
        );

        // Verify the registration
        let entry = await contract.getMerkleEntryByRoot(MERKLE_ROOT);
        expect(entry.proofURI).to.equal(PROOF_URI);

        // Test successful update
        const newProofURI = "ipfs://QmnewProofURI";
        let tx = await contract.updateMerkleRootProofURI(
            MERKLE_ROOT,
            newProofURI,
            { from: submitter }
        );

        expectEvent(tx, "MerkleRootProofURIUpdated", {
            index: new BN(0),
            merkleRoot: MERKLE_ROOT,
            proofURI: newProofURI,
        });

        // Verify the update
        entry = await contract.getMerkleEntryByRoot(MERKLE_ROOT);
        expect(entry.proofURI).to.equal(newProofURI);

        // Test revert due to empty proofURI
        await expectCustomError(
            contract.updateMerkleRootProofURI(MERKLE_ROOT, "", {
                from: submitter,
            }),
            abi,
            "ProofURIIsEmpty",
            []
        );

        // Test revert due to non-submitter
        const nonSubmitter = accounts[1];
        await expectCustomError(
            contract.updateMerkleRootProofURI(MERKLE_ROOT, newProofURI, {
                from: nonSubmitter,
            }),
            abi,
            "NotSubmitter",
            []
        );

        // Test revert due to non-existent root
        await expectCustomError(
            contract.updateMerkleRootProofURI("0x0", newProofURI, {
                from: submitter,
            }),
            abi,
            "RootNotFound",
            []
        );
    });

    it("test updateMerkleRootMetadataURI", async () => {
        let submitter = accounts[0];

        // Register a new Merkle root
        await contract.batchRegisterMerkleRoots(
            [MERKLE_ROOT],
            [PROOF_URI],
            [METADATA_URI],
            { from: submitter }
        );

        // Verify the registration
        let entry = await contract.getMerkleEntryByRoot(MERKLE_ROOT);
        expect(entry.metadataURI).to.equal(METADATA_URI);

        // Test successful update
        const newMetadataURI = "ipfs://QmnewMetadataURI";
        let tx = await contract.updateMerkleRootMetadataURI(
            MERKLE_ROOT,
            newMetadataURI,
            { from: submitter }
        );

        expectEvent(tx, "MerkleRootMetadataURIUpdated", {
            index: new BN(0),
            merkleRoot: MERKLE_ROOT,
            metadataURI: newMetadataURI,
        });

        // Verify the update
        entry = await contract.getMerkleEntryByRoot(MERKLE_ROOT);
        expect(entry.metadataURI).to.equal(newMetadataURI);

        // Test revert due to empty metadataURI
        await expectCustomError(
            contract.updateMerkleRootMetadataURI(MERKLE_ROOT, "", {
                from: submitter,
            }),
            abi,
            "MetadataURIIsEmpty",
            []
        );

        // Test revert due to non-submitter
        const nonSubmitter = accounts[1];
        await expectCustomError(
            contract.updateMerkleRootMetadataURI(MERKLE_ROOT, newMetadataURI, {
                from: nonSubmitter,
            }),
            abi,
            "NotSubmitter",
            []
        );

        // Test revert due to non-existent root
        await expectCustomError(
            contract.updateMerkleRootMetadataURI("0x0", newMetadataURI, {
                from: submitter,
            }),
            abi,
            "RootNotFound",
            []
        );
    });

    it("test verifyLeafByIndex", async () => {
        let submitter = accounts[0];

        // Register a new Merkle root
        await contract.batchRegisterMerkleRoots(
            [MERKLE_ROOT],
            [PROOF_URI],
            [METADATA_URI],
            { from: submitter }
        );

        // Verify the leaf by index
        let result = await contract.verifyLeafByIndex(0, LEAF, PROOF);
        expect(result).to.be.true;

        // Verify the failure case
        result = await contract.verifyLeafByIndex(0, LEAF, [
            "0x0",
            "0x1",
            "0x2",
        ]);
        expect(result).to.be.false;

        // Test revert due to index out of range
        await expectCustomError(
            contract.verifyLeafByIndex(1, LEAF, PROOF),
            abi,
            "IndexOutOfRange",
            []
        );
    });

    it("test verifyLeafByRoot", async () => {
        let submitter = accounts[0];

        // Register a new Merkle root
        await contract.batchRegisterMerkleRoots(
            [MERKLE_ROOT],
            [PROOF_URI],
            [METADATA_URI],
            { from: submitter }
        );

        // Verify the leaf by root
        let result = await contract.verifyLeafByRoot(MERKLE_ROOT, LEAF, PROOF);
        expect(result).to.be.true;

        // Verify the failure case
        result = await contract.verifyLeafByRoot(MERKLE_ROOT, LEAF, [
            "0x0",
            "0x1",
            "0x2",
        ]);
        expect(result).to.be.false;

        // Test revert due to non-existent root
        await expectCustomError(
            contract.verifyLeafByRoot("0x0", LEAF, PROOF),
            abi,
            "RootNotFound",
            []
        );
    });
});
