// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {MerkleProof} from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

/**
 * @title MerkleRegistry
 * @notice A registry contract to store multiple Merkle roots along with their metadata.
 *         - Allows both enumeration (via allEntries array)
 *         - Direct lookups using rootToIndex mapping
 *         - On-chain verification of Merkle proofs
 */
contract MerkleRegistry {
    using MerkleProof for bytes32[];

    // ------------------------------------------------------------------------
    // ERRORS
    // ------------------------------------------------------------------------

    error MerkleRootAlreadyExists();
    error MerkleRootIsZero();
    error ProofURIIsEmpty();
    error IndexOutOfRange();
    error RootNotFound();
    error ArraysLengthMismatch();
    error NotSubmitter();
    error MetadataURIIsEmpty();

    // ------------------------------------------------------------------------
    // DATA STRUCTURES
    // ------------------------------------------------------------------------

    /**
     * @dev Each MerkleEntry stores:
     * - merkleRoot: the root hash of a Merkle tree
     * - proofURI:   IPFS (or other) URI referencing the proof file
     * - metadataURI: IPFS (or other) URI referencing the metadata file
     * - submitter:  address who submitted this root
     * - timestamp:  block timestamp when it was submitted
     */
    struct MerkleEntry {
        bytes32 merkleRoot;
        string proofURI;
        string metadataURI;
        address submitter;
        uint256 timestamp;
    }

    /**
     * @dev Array of all Merkle entries for enumeration (on-chain or off-chain).
     *      Users can index into this array to see each submission.
     */
    MerkleEntry[] public allEntries;

    /**
     * @dev Maps a Merkle root to its index in the `allEntries` array.
     *      If a root is not found, index will be 0, but it might also
     *      overlap with the real first entry.
     */
    mapping(bytes32 => uint256) public rootToIndex;

    // ------------------------------------------------------------------------
    // EVENTS
    // ------------------------------------------------------------------------

    event MerkleRootRegistered(
        uint256 indexed index,
        bytes32 indexed merkleRoot,
        address indexed submitter,
        string proofURI,
        string metadataURI
    );

    event MerkleRootProofURIUpdated(
        uint256 indexed index,
        bytes32 indexed merkleRoot,
        string proofURI
    );

    event MerkleRootMetadataURIUpdated(
        uint256 indexed index,
        bytes32 indexed merkleRoot,
        string metadataURI
    );

    // ------------------------------------------------------------------------
    // MODIFIERS
    // ------------------------------------------------------------------------

    modifier onlySubmitter(bytes32 _merkleRoot) {
        uint256 index = _ensureRootExists(_merkleRoot);
        if (msg.sender != allEntries[index].submitter) {
            revert NotSubmitter();
        }
        _;
    }

    // ------------------------------------------------------------------------
    // CONSTRUCTOR
    // ------------------------------------------------------------------------

    /**
     * @dev Optionally store a single IPFS URI for the *standard* Merkle generation code.
     */
    string public sourceCodeURI;

    constructor(string memory _sourceCodeURI) {
        sourceCodeURI = _sourceCodeURI;
    }

    // ------------------------------------------------------------------------
    // REGISTRY FUNCTIONS
    // ------------------------------------------------------------------------

    /**
     * @notice Register multiple Merkle roots in a single transaction.
     * @param _merkleRoots Array of Merkle roots to register
     * @param _proofURIs Array of proof URIs
     * @param _metadataURIs Array of metadata URIs
     */
    function batchRegisterMerkleRoots(
        bytes32[] calldata _merkleRoots,
        string[] calldata _proofURIs,
        string[] calldata _metadataURIs
    ) external {
        // Check if the arrays are of the same length
        if (
            _merkleRoots.length != _proofURIs.length ||
            _merkleRoots.length != _metadataURIs.length
        ) {
            revert ArraysLengthMismatch();
        }

        for (uint256 i = 0; i < _merkleRoots.length; i++) {
            _registerMerkleRoot(
                _merkleRoots[i],
                _proofURIs[i],
                _metadataURIs[i]
            );
        }
    }

    /**
     * @notice Register a new Merkle root into the array + mapping.
     * @param _merkleRoot  The root of the Merkle tree
     * @param _proofURI    IPFS URI of the proof file
     * @param _metadataURI IPFS URI of the metadata file (optional)
     */
    function registerMerkleRoot(
        bytes32 _merkleRoot,
        string calldata _proofURI,
        string calldata _metadataURI
    ) external {
        _registerMerkleRoot(_merkleRoot, _proofURI, _metadataURI);
    }

    /**
     * @notice Register a new Merkle root into the array + mapping.
     * @param _merkleRoot  The root of the Merkle tree
     * @param _proofURI    IPFS URI of the proof file
     * @param _metadataURI IPFS URI of the metadata file
     */
    function _registerMerkleRoot(
        bytes32 _merkleRoot,
        string memory _proofURI,
        string memory _metadataURI
    ) internal {
        // Check if the root is zero
        if (_merkleRoot == bytes32(0)) {
            revert MerkleRootIsZero();
        }

        // Check if the _proofURI is empty
        if (bytes(_proofURI).length == 0) {
            revert ProofURIIsEmpty();
        }

        // If the root is already registered, revert or handle differently
        uint256 existingIndex = rootToIndex[_merkleRoot];
        if (existingIndex != 0) {
            revert MerkleRootAlreadyExists();
        }

        if (allEntries.length > 0 && _merkleRoot == allEntries[0].merkleRoot) {
            // Check if it's truly the same root or if default 0 is overlapping
            revert MerkleRootAlreadyExists();
        }

        // Create the new entry
        MerkleEntry memory entry = MerkleEntry({
            merkleRoot: _merkleRoot,
            proofURI: _proofURI,
            submitter: msg.sender,
            timestamp: block.timestamp,
            metadataURI: _metadataURI
        });

        // Push to array
        allEntries.push(entry);

        // Map the Merkle root to the new index
        uint256 newIndex = allEntries.length - 1;
        rootToIndex[_merkleRoot] = newIndex;

        // Emit event
        emit MerkleRootRegistered(
            newIndex,
            _merkleRoot,
            msg.sender,
            _proofURI,
            _metadataURI
        );
    }

    // ------------------------------------------------------------------------
    // UPDATE FUNCTIONS
    // ------------------------------------------------------------------------

    /**
     * @notice Update the proof URI for a specific Merkle root.
     * @param _merkleRoot The root of the Merkle tree
     * @param _proofURI   The new IPFS URI for the proof file
     */
    function updateMerkleRootProofURI(
        bytes32 _merkleRoot,
        string calldata _proofURI
    ) external onlySubmitter(_merkleRoot) {
        // Check if the _proofURI is empty
        if (bytes(_proofURI).length == 0) {
            revert ProofURIIsEmpty();
        }

        // Create a storage pointer to the entry
        MerkleEntry storage entry = allEntries[rootToIndex[_merkleRoot]];
        entry.proofURI = _proofURI;

        emit MerkleRootProofURIUpdated(
            rootToIndex[_merkleRoot],
            _merkleRoot,
            _proofURI
        );
    }

    /**
     * @notice Update the metadata URI for a specific Merkle root.
     * @param _merkleRoot The root of the Merkle tree
     * @param _metadataURI The new IPFS URI for the metadata file
     */
    function updateMerkleRootMetadataURI(
        bytes32 _merkleRoot,
        string calldata _metadataURI
    ) external onlySubmitter(_merkleRoot) {
        // Check if the _metadataURI is empty
        if (bytes(_metadataURI).length == 0) {
            revert MetadataURIIsEmpty();
        }

        // Create a storage pointer to the entry
        MerkleEntry storage entry = allEntries[rootToIndex[_merkleRoot]];
        entry.metadataURI = _metadataURI;

        emit MerkleRootMetadataURIUpdated(
            rootToIndex[_merkleRoot],
            _merkleRoot,
            _metadataURI
        );
    }

    /**
     * @notice Returns the total number of Merkle entries.
     *         This can help with enumerating in a front-end.
     */
    function getEntryCount() external view returns (uint256) {
        return allEntries.length;
    }

    /**
     * @notice Retrieve a particular entry by its array index.
     * @param index The index in allEntries array
     * @return The MerkleEntry struct
     */
    function getMerkleEntryByIndex(
        uint256 index
    ) external view returns (MerkleEntry memory) {
        if (index >= allEntries.length) {
            revert IndexOutOfRange();
        }
        return allEntries[index];
    }

    /**
     * @notice Retrieve a Merkle entry by its root directly,
     *         using the rootToIndex mapping.
     * @param root The merkle root
     * @return The MerkleEntry struct
     */
    function getMerkleEntryByRoot(
        bytes32 root
    ) external view returns (MerkleEntry memory) {
        uint256 index = _ensureRootExists(root);
        return allEntries[index];
    }

    // ------------------------------------------------------------------------
    // ON-CHAIN VERIFICATION
    // ------------------------------------------------------------------------

    /**
     * @notice Verify that a given leaf is part of the Merkle tree identified by `root`.
     * @param root  The merkle root to check against
     * @param leaf  The leaf (hash) being verified
     * @param proof The array of sibling hashes for the merkle proof
     * @return true if the leaf is part of the tree, false otherwise
     */
    function verifyLeafByRoot(
        bytes32 root,
        bytes32 leaf,
        bytes32[] calldata proof
    ) external view returns (bool) {
        _ensureRootExists(root);

        // Using OpenZeppelin's MerkleProof library:
        // Reconstruct the path from leaf up using the proof, compare with root
        return proof.verify(root, leaf);
    }

    /**
     * @notice Verify that a given leaf is part of the Merkle tree at a specific array index.
     * @param index Index in the allEntries array
     * @param leaf  The leaf (hash) being verified
     * @param proof The array of sibling hashes for the merkle proof
     * @return true if the leaf is part of the tree, false otherwise
     */
    function verifyLeafByIndex(
        uint256 index,
        bytes32 leaf,
        bytes32[] calldata proof
    ) external view returns (bool) {
        if (index >= allEntries.length) {
            revert IndexOutOfRange();
        }
        bytes32 root = allEntries[index].merkleRoot;
        if (root == bytes32(0)) {
            revert RootNotFound();
        }

        return proof.verify(root, leaf);
    }

    /**
     * @notice Ensure that a root exists in the registry.
     * @param root The merkle root to check
     */
    function _ensureRootExists(
        bytes32 root
    ) private view returns (uint256 index) {
        index = rootToIndex[root];
        if (
            index >= allEntries.length || allEntries[index].merkleRoot != root
        ) {
            revert RootNotFound();
        }
    }
}
