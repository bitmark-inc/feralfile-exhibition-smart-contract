// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/BitMaps.sol";

/**
 * @title SeriesRegistry
 * @dev A contract for managing series and artists in a collaborative art platform.
 *      This contract allows artists to create series, manage collaborators, and control owner rights.
 */
contract SeriesRegistry is Ownable {
    using BitMaps for BitMaps.BitMap;

    // ============ Custom Errors ============

    error OwnerRightsRevokedForThisSeries(uint256 seriesID);
    error SeriesDoesNotExistError(uint256 seriesID);
    error CallerNotASeriesArtistError(uint256 seriesID, address caller);
    error NoArtistsProvidedError();
    error NoArtistsForSeriesError();
    error InvalidNewAddressError(address newAddress);
    error AddressAlreadyAssignedError(address newAddress);
    error InvalidArtistIDError(uint256 artistID);
    error NotAuthorizedError(address caller);
    error ArtistRevokedOwnerRightsError(address artistAddress);
    error NoProposalExistsError(uint256 seriesID, address proposedArtistAddress);
    error NotAPendingProposalError(uint256 seriesID, address caller);
    error ZeroAddressNotAllowedError();
    error NoSeriesDataError();
    error BatchSizeTooLargeError(uint256 size);
    error ArrayLengthMismatchError(uint256 lenArtists, uint256 lenMetas, uint256 lenTokens);
    error AlreadyProposedError(uint256 seriesID, address proposedArtistAddress);
    error ArtistAlreadyInSeriesError(uint256 seriesID, address proposedArtistAddress);
    error NotAnArtistError(address caller);
    error EmptyMetadataURIError();
    error EmptyTokenMapURIError();

    // ============ Structs ============

    struct Series {
        string metadataURI;              // IPFS hash or similar identifier for series metadata
        string contractTokenDataURI;     // Token-related data for the series
        uint256[] artistIDs;          // List of artists associated with this series
    }

    struct Artist {
        address artistAddress;           // Artist's wallet address
        uint256[] seriesIDs;            // List of series associated with this artist
    }

    // ============ State Variables ============

    // ID Counters
    uint256 private nextSeriesID = 1;
    uint256 private nextArtistID = 1;

    // Artist Management
    mapping(uint256 => Artist) private artists;
    mapping(address => uint256) private artistAddressToID;
    BitMaps.BitMap private artistOwnerRightsRevokedBitMap;

    // Series Management
    mapping(uint256 => Series) private seriesRegistry;

    // Series - Artist Relationship
    mapping(uint256 => BitMaps.BitMap) private seriesArtistExist;

    // Collaborator Management
    mapping(uint256 => BitMaps.BitMap) private seriesPendingCollaborator;
    mapping(uint256 => uint256[]) private seriesPendingCollaborators;   // List of series's pending collaborators
    mapping(uint256 => uint256[]) private artistPendingCollaboratorSeries;   // List of artist's pending collaborator series

    // ============ Events ============

    event ArtistAddressUpdated(
        uint256 indexed artistID, 
        address indexed oldAddress, 
        address indexed newAddress
    );
    
    event SeriesRegistered(uint256 indexed seriesID);
    event SeriesUpdated(uint256 indexed seriesID);
    event SeriesDeleted(uint256 indexed seriesID);
    
    event CollaboratorProposed(
        uint256 indexed seriesID, 
        uint256 indexed proposedArtistID
    );
    
    event CollaboratorConfirmed(
        uint256 indexed seriesID, 
        uint256 indexed confirmedArtistID
    );
    
    event CollaboratorProposalCancelled(
        uint256 indexed seriesID,
        uint256 indexed proposerArtistID,
        uint256 indexed cancelledArtistID
    );

    // ============ Modifiers ============

    modifier onlyOwnerOrArtist(uint256 seriesID) {
        if (_isCallerOwner()) {
            // If owner is calling, check if at least one artist in this series has NOT revoked rights
            if (!hasUnrevokedArtist(seriesID)) {
                revert OwnerRightsRevokedForThisSeries(seriesID);
            }
        } else {
            // Otherwise, must be a valid artist in this series
            _requireCallerIsSeriesArtist(seriesID);
        }
        _;
    }

    modifier onlyArtist(uint256 seriesID) {
        _requireCallerIsSeriesArtist(seriesID);
        _;
    }

    modifier seriesExists(uint256 seriesID) {
        // If metadata is empty, we consider the series nonexistent
        if (bytes(seriesRegistry[seriesID].metadataURI).length == 0) {
            revert SeriesDoesNotExistError(seriesID);
        }
        _;
    }

    // ============ Public/External Functions: Series Management ============

    /**
     * @notice Adds a new series
     * @param artistAddresses The addresses of the artists associated with the series
     * @param metadataURI The metadata URI of the series
     * @param tokenIDsMapURI The token IDs map URI of the series
     * @return The ID of the newly created series
     */
    function addSeries(
        address[] calldata artistAddresses,
        string calldata metadataURI,
        string calldata tokenIDsMapURI
    ) external returns (uint256) {
        return _createSeries(artistAddresses, metadataURI, tokenIDsMapURI);
    }

    /**
     * @notice Batch adds new series
     * @param seriesArtists The addresses of the artists associated with the series
     * @param metadataURIs The metadata URIs of the series
     * @param tokenIDsMapURIs The token IDs map URIs of the series
     * @return The IDs of the newly created series
     */
    function batchAddSeries(
        address[][] calldata seriesArtists,
        string[] calldata metadataURIs,
        string[] calldata tokenIDsMapURIs
    ) external returns (uint256[] memory) {
        _validateSeriesWriteBatch(seriesArtists.length, metadataURIs, tokenIDsMapURIs);

        uint256[] memory seriesIDs = new uint256[](seriesArtists.length);
        for (uint256 i = 0; i < seriesArtists.length; i++) {
            seriesIDs[i] = _createSeries(
                seriesArtists[i],
                metadataURIs[i],
                tokenIDsMapURIs[i]
            );
        }
        return seriesIDs;
    }

    /**
     * @notice Updates an existing series' metadata and token data
     * @param seriesID The ID of the series to update
     * @param metadataURI The metadata URI of the series
     * @param tokenIDsMapURI The token IDs map URI of the series
     */
    function updateSeries(
        uint256 seriesID,
        string calldata metadataURI,
        string calldata tokenIDsMapURI
    ) external {
        _updateSeries(seriesID, metadataURI, tokenIDsMapURI);
    }

    /**
     * @notice Batch updates existing series' metadata and token data
     * @param seriesIDs The IDs of the series to update
     * @param metadataURIs The metadata URIs of the series
     * @param tokenIDsMapURIs The token IDs map URIs of the series
     */
    function batchUpdateSeries(
        uint256[] calldata seriesIDs,
        string[] calldata metadataURIs,
        string[] calldata tokenIDsMapURIs
    ) external {
        _validateSeriesWriteBatch(seriesIDs.length, metadataURIs, tokenIDsMapURIs);
        for (uint256 i = 0; i < seriesIDs.length; i++) {
            _updateSeries(seriesIDs[i], metadataURIs[i], tokenIDsMapURIs[i]);
        }
    }

    /**
     * @notice Deletes a series
     * @param seriesID The ID of the series to delete
     */
    function deleteSeries(uint256 seriesID) external {
        _deleteSeries(seriesID);
    }

    /**
     * @notice Batch deletes series
     * @param seriesIDs The IDs of the series to delete
     */
    function batchDeleteSeries(uint256[] calldata seriesIDs) external {
        for (uint256 i = 0; i < seriesIDs.length; i++) {
            _deleteSeries(seriesIDs[i]);
        }
    }

    // ============ Public/External Functions: Series Artist Management ============

    /**
     * @notice Allows an artist to remove themselves from a series
     * @param seriesID The ID of the series to resign from
     */
    function resignFromSeries(uint256 seriesID) external seriesExists(seriesID) onlyArtist(seriesID) {
        uint256 artistID = artistAddressToID[msg.sender];

        _unlinkArtistFromSeries(seriesID, artistID);
        _unlinkSeriesFromArtist(artistID, seriesID);

        emit SeriesUpdated(seriesID);
    }

    /**
     * @notice Updates the artist list for a series (owner only)
     * @param seriesID The ID of the series to update
     * @param artistAddresses The addresses of the artists to add to the series
     */
    function ownerUpdateSeriesArtists(
        uint256 seriesID,
        address[] calldata artistAddresses
    ) external seriesExists(seriesID) onlyOwner {
        if (!hasUnrevokedArtist(seriesID)) {
            revert OwnerRightsRevokedForThisSeries(seriesID);
        }
        _ensureNoArtistsRevokedOwnerRights(artistAddresses);

        // Remove current artists
        Series storage series = seriesRegistry[seriesID];
        for (uint256 i = 0; i < series.artistIDs.length; i++) {
            uint256 artistID = series.artistIDs[i];
            _unlinkSeriesFromArtist(artistID, seriesID);
        }
        delete series.artistIDs;

        // Add new artists
        _addArtistsToSeries(seriesID, artistAddresses);

        emit SeriesUpdated(seriesID);
    }

    // ============ Public/External Functions: Collaborator Management ============

    /**
     * @notice Proposes a new collaborator for a series
     * @param seriesID The ID of the series to propose a collaborator for
     * @param proposedArtistAddress The address of the proposed collaborator
     */
    function proposeCollaborator(
        uint256 seriesID,
        address proposedArtistAddress
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        _ensureArtistHasID(proposedArtistAddress);
        uint256 proposedArtistID = artistAddressToID[proposedArtistAddress];

        if (seriesArtistExist[seriesID].get(proposedArtistID)) {
            revert ArtistAlreadyInSeriesError(seriesID, proposedArtistAddress);
        }
        if (seriesPendingCollaborator[seriesID].get(proposedArtistID)) {
            revert AlreadyProposedError(seriesID, proposedArtistAddress);
        }

        seriesPendingCollaborator[seriesID].set(proposedArtistID);
        _addCollaboratorProposal(proposedArtistID, seriesID);

        emit CollaboratorProposed(seriesID, proposedArtistID);
    }

    /**
     * @notice Cancels a collaborator proposal
     * @param seriesID The ID of the series to cancel the proposal for
     * @param proposedArtistAddress The address of the proposed collaborator
     */
    function cancelProposeCollaborator(
        uint256 seriesID,
        address proposedArtistAddress
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        uint256 proposedArtistID = artistAddressToID[proposedArtistAddress];
        if (!seriesPendingCollaborator[seriesID].get(proposedArtistID)) {
            revert NoProposalExistsError(seriesID, proposedArtistAddress);
        }

        seriesPendingCollaborator[seriesID].unset(proposedArtistID);
        _removeCollaboratorProposal(proposedArtistID, seriesID);

        uint256 artistID = artistAddressToID[msg.sender];
        emit CollaboratorProposalCancelled(seriesID, artistID, proposedArtistID);
    }

    /**
     * @notice Confirms participation as a collaborator
     * @param seriesID The ID of the series to confirm participation for
     */
    function confirmAsCollaborator(uint256 seriesID) external seriesExists(seriesID) {
        uint256 artistID = artistAddressToID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (!seriesPendingCollaborator[seriesID].get(artistID)) {
            revert NotAPendingProposalError(seriesID, msg.sender);
        }

        address[] memory artistAddresses = new address[](1);
        artistAddresses[0] = msg.sender;

        seriesPendingCollaborator[seriesID].unset(artistID);
        _removeCollaboratorProposal(artistID, seriesID);
        _addArtistsToSeries(seriesID, artistAddresses);

        emit CollaboratorConfirmed(seriesID, artistID);
    }

    // ============ Public/External Functions: Artist Management ============

    /**
     * @notice Allow artist to revoke their contract owner rights
     */
    function revokeContractOwnerRights() external {
        uint256 artistID = artistAddressToID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (artistOwnerRightsRevokedBitMap.get(artistID)) {
            return;
        }
        artistOwnerRightsRevokedBitMap.set(artistID);
    }

    /**
     * @notice Allow artist to approve their contract owner rights
     */
    function approveContractOwnerRights() external {
        uint256 artistID = artistAddressToID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (!artistOwnerRightsRevokedBitMap.get(artistID)) {
            return;
        }
        artistOwnerRightsRevokedBitMap.unset(artistID);
    }

    /**
     * @notice Updates an artist's address
     * @param artistID The ID of the artist to update
     * @param newAddress The new address of the artist
     */
    function updateArtistAddress(uint256 artistID, address newAddress) external {
        if (newAddress == address(0)) {
            revert InvalidNewAddressError(newAddress);
        }
        if (artistAddressToID[newAddress] != 0) {
            revert AddressAlreadyAssignedError(newAddress);
        }
        if (artists[artistID].artistAddress == address(0)) {
            revert InvalidArtistIDError(artistID);
        }

        uint256 callerID = artistAddressToID[msg.sender];
        bool isCallerArtist = (callerID == artistID && callerID != 0);
        bool isOwnerWithRights = (_isCallerOwner() && !artistOwnerRightsRevokedBitMap.get(artistID));

        if (!isCallerArtist && !isOwnerWithRights) {
            revert NotAuthorizedError(msg.sender);
        }

        address oldAddress = artists[artistID].artistAddress;
        artistAddressToID[oldAddress] = 0;
        artists[artistID].artistAddress = newAddress;
        artistAddressToID[newAddress] = artistID;

        emit ArtistAddressUpdated(artistID, oldAddress, newAddress);
    }

    // ============ Public/External View Functions ============

    /**
     * @notice Checks if at least one artist in the series has not revoked the owner's right to modify the series
     * @param seriesID The ID of the series to check
     * @return `true` if at least one artist has not revoked the owner's right to modify the series, `false` otherwise
     */
    function hasUnrevokedArtist(uint256 seriesID) public view returns (bool) {
        uint256[] memory ids = seriesRegistry[seriesID].artistIDs;
        if (ids.length == 0) {
            // If no artists, owner can modify the series
            return true;
        }
        for (uint256 i = 0; i < ids.length; i++) {
            // If any artist has not revoked, return true
            if (!artistOwnerRightsRevokedBitMap.get(ids[i])) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice Checks if an artist has revoked their contract owner rights
     * @param artistAddress The address of the artist to check
     * @return `true` if the artist has revoked their contract owner rights, `false` otherwise
     */
    function hasArtistRevokedOwnerRights(address artistAddress) external view returns (bool) {
        uint256 artistID = artistAddressToID[artistAddress];
        return artistOwnerRightsRevokedBitMap.get(artistID);
    }

    /**
     * @notice Returns all artist IDs for a given series
     * @param seriesID The ID of the series to get the artist IDs for
     * @return An array of artist IDs associated with the series
     */
    function getSeriesArtistIDs(uint256 seriesID) external view returns (uint256[] memory) {
        return seriesRegistry[seriesID].artistIDs;
    }

    /**
     * @notice Returns all artist addresses for a given series
     * @param seriesID The ID of the series to get the artist addresses for
     * @return An array of artist addresses associated with the series
     */
    function getSeriesArtistAddresses(uint256 seriesID) external view returns (address[] memory) {
        uint256[] memory ids = seriesRegistry[seriesID].artistIDs;
        address[] memory artistAddresses = new address[](ids.length);
        for (uint256 i = 0; i < ids.length; i++) {
            artistAddresses[i] = artists[ids[i]].artistAddress;
        }
        return artistAddresses;
    }

    /**
     * @notice Returns all series IDs associated with an artist address
     * @param artistAddress The address of the artist to get the series IDs for
     * @return An array of series IDs associated with the artist
     */
    function getArtistSeriesIDs(address artistAddress) external view returns (uint256[] memory) {
        uint256 artistID = artistAddressToID[artistAddress];
        return artists[artistID].seriesIDs;
    }

    /**
     * @notice Returns the metadata URI of a given series
     * @param seriesID The ID of the series to get the metadata URI for
     * @return The metadata URI of the series
     */
    function getSeriesMetadataURI(uint256 seriesID) external view returns (string memory) {
        return seriesRegistry[seriesID].metadataURI;
    }

    /**
     * @notice Returns the contract-level token data (e.g., IPFS CID) of a given series
     * @param seriesID The ID of the series to get the contract-level token data for
     * @return The contract-level token data of the series
     */
    function getSeriesContractTokenDataURI(uint256 seriesID) external view returns (string memory) {
        return seriesRegistry[seriesID].contractTokenDataURI;
    }

    /**
     * @notice Returns the pending collaborator requests for an artist address
     * @param artistAddress The address of the artist to get the pending collaborator requests for
     * @return An array of series IDs associated with the artist's pending collaborator requests
     */
    function getArtistPendingCollaboratorSeries(address artistAddress) external view returns (uint256[] memory) {
        uint256 artistID = artistAddressToID[artistAddress];
        return artistPendingCollaboratorSeries[artistID];
    }

    /**
     * @notice Returns the pending collaborator list for a given series
     * @param seriesID The ID of the series to get the pending collaborator list for
     * @return An array of addresses associated with the series' pending collaborators
     */
    function getSeriesPendingCollaborators(uint256 seriesID) external view returns (address[] memory) {
        uint256[] memory artistIDs = seriesPendingCollaborators[seriesID];
        address[] memory artistAddresses = new address[](artistIDs.length);
        for (uint256 i = 0; i < artistIDs.length; i++) {
            artistAddresses[i] = artists[artistIDs[i]].artistAddress;
        }
        return artistAddresses;
    }

    /**
     * @notice Returns the address associated with a given artist ID
     * @param artistID The ID of the artist to get the address for
     * @return The address of the artist
     */
    function getArtistAddress(uint256 artistID) external view returns (address) {
        return artists[artistID].artistAddress;
    }

    /**
     * @notice Returns the artist ID associated with an address
     * @param artistAddress The address of the artist to get the ID for
     * @return The ID of the artist
     */
    function getAddressArtistID(address artistAddress) external view returns (uint256) {
        return artistAddressToID[artistAddress];
    }

    // ============ Internal Functions ============

    /**
     * @notice Creates a new series with the specified artists
     * @param artistAddresses The addresses of the artists associated with the series
     * @param metadataURI The metadata URI of the series
     * @param tokenIDsMapURI The token IDs map URI of the series
     * @return The ID of the newly created series
     */
    function _createSeries(
        address[] memory artistAddresses,
        string memory metadataURI,
        string memory tokenIDsMapURI
    ) internal returns (uint256) {
        if (artistAddresses.length == 0) {
            revert NoArtistsForSeriesError();
        }
        if (_isCallerOwner()) {
            _ensureNoArtistsRevokedOwnerRights(artistAddresses);
        } else {
            _ensureCallerIsOnlyArtist(artistAddresses);
        }
        _validateMetadataAndTokenURI(metadataURI, tokenIDsMapURI);
        uint256 seriesID = nextSeriesID++;

        Series storage series = seriesRegistry[seriesID];
        series.metadataURI = metadataURI;
        series.contractTokenDataURI = tokenIDsMapURI;

        _addArtistsToSeries(seriesID, artistAddresses);

        emit SeriesRegistered(seriesID);
        return seriesID;
    }

    /**
     * @notice Updates an existing series' metadata and token data
     * @param seriesID The ID of the series to update
     * @param metadataURI The metadata URI of the series
     * @param tokenIDsMapURI The token IDs map URI of the series
     */
    function _updateSeries(
        uint256 seriesID,
        string memory metadataURI,
        string memory tokenIDsMapURI
    ) internal seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        _validateMetadataAndTokenURI(metadataURI, tokenIDsMapURI);

        Series storage series = seriesRegistry[seriesID];
        series.metadataURI = metadataURI;
        series.contractTokenDataURI = tokenIDsMapURI;

        emit SeriesUpdated(seriesID);
    }

    /**
     * @notice Deletes a series and cleans up related data
     * @param seriesID The ID of the series to delete
     */
    function _deleteSeries(uint256 seriesID) 
        internal 
        seriesExists(seriesID) 
        onlyOwnerOrArtist(seriesID) 
    {
        // Clean up all pending collaborator requests for this series in one loop
        uint256[] memory pendingCollaborators = seriesPendingCollaborators[seriesID];
        for (uint256 i = 0; i < pendingCollaborators.length; i++) {
            uint256 artistID = pendingCollaborators[i];
            seriesPendingCollaborator[seriesID].unset(artistID);
            _removeArtistPendingCollaboratorSeries(artistID, seriesID);
            _removeSeriesPendingCollaborator(seriesID, artistID);
        }

        Series memory series = seriesRegistry[seriesID];
        
        // Remove series from all artists
        for (uint256 i = 0; i < series.artistIDs.length; i++) {
            _unlinkSeriesFromArtist(series.artistIDs[i], seriesID);
        }

        delete seriesRegistry[seriesID];

        emit SeriesDeleted(seriesID);
    }

    /**
     * @notice Removes a series from an artist's list
     * @param artistID The ID of the artist to remove the series from
     * @param seriesID The ID of the series to remove
     */
    function _unlinkSeriesFromArtist(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage seriesIDs = artists[artistID].seriesIDs;
        uint256 seriesIDsLength = seriesIDs.length;
        for (uint256 i = 0; i < seriesIDsLength; i++) {
            if (seriesIDs[i] == seriesID) {
                seriesIDs[i] = seriesIDs[seriesIDsLength - 1];
                seriesIDs.pop();
                break;
            }
        }
        
        seriesArtistExist[seriesID].unset(artistID);
    }

    /**
     * @notice Removes an artist from a series
     * @param seriesID The ID of the series to remove the artist from
     * @param artistID The ID of the artist to remove
     */
    function _unlinkArtistFromSeries(uint256 seriesID, uint256 artistID) internal {
        uint256[] storage artistIDs = seriesRegistry[seriesID].artistIDs;
        uint256 artistIDsLength = artistIDs.length;
        // Find and remove artistID from series' artist list
        for (uint256 i = 0; i < artistIDsLength; i++) {
            if (artistIDs[i] == artistID) {
                artistIDs[i] = artistIDs[artistIDsLength - 1];
                artistIDs.pop();
                break;
            }
        }
    }

    /**
     * @notice Validates that artists can only add series that only have themselves as artist
     * @param artistAddresses The addresses of the artists
     */
    function _ensureCallerIsOnlyArtist(
        address[] memory artistAddresses
    ) internal view {
        if (artistAddresses.length != 1 || artistAddresses[0] != msg.sender) {
            revert NotAuthorizedError(msg.sender);
        }
    }

    /**
     * @notice Validates that metadataURI and token URI are not empty
     * @param metadataURI The metadata URI of the series
     * @param tokenIDsMapURI The token IDs map URI of the series
     */
    function _validateMetadataAndTokenURI(
        string memory metadataURI,
        string memory tokenIDsMapURI
    ) internal pure {
        if (bytes(metadataURI).length == 0) {
            revert EmptyMetadataURIError();
        }
        if (bytes(tokenIDsMapURI).length == 0) {
            revert EmptyTokenMapURIError();
        }
    }

    /**
     * @notice Validates batch parameters for series creation
     * @param seriesCount The number of series
     * @param metadataURIs The metadata URIs of the series
     * @param tokenIDsMapURIs The token IDs map URIs of the series
     */
    function _validateSeriesWriteBatch(
        uint256 seriesCount,
        string[] memory metadataURIs,
        string[] memory tokenIDsMapURIs
    ) internal pure {
        if (seriesCount == 0) {
            revert NoSeriesDataError();
        }
        if (seriesCount > 50) {
            revert BatchSizeTooLargeError(seriesCount);
        }
        if (seriesCount != metadataURIs.length || seriesCount != tokenIDsMapURIs.length) {
            revert ArrayLengthMismatchError(
                seriesCount,
                metadataURIs.length,
                tokenIDsMapURIs.length
            );
        }
    }

    /**
     * @notice Checks if the caller is contract owner
     * @return `true` if the caller is the contract owner, `false` otherwise
     */
    function _isCallerOwner() internal view returns (bool) {
        return msg.sender == owner();
    }

    /**
     * @notice Checks if the caller is a series artist
     * @param seriesID The ID of the series to check
     */
    function _requireCallerIsSeriesArtist(uint256 seriesID) internal view {
        uint256 artistID = artistAddressToID[msg.sender];
        if (artistID == 0 || !seriesArtistExist[seriesID].get(artistID)) {
            revert CallerNotASeriesArtistError(seriesID, msg.sender);
        }
    }

    /**
     * @notice Ensures all provided artist addresses have not revoked contract owner rights
     * @param artistAddresses The addresses of the artists
     */
    function _ensureNoArtistsRevokedOwnerRights(address[] memory artistAddresses) internal view {
        for (uint256 i = 0; i < artistAddresses.length; i++) {
            uint256 artistID = artistAddressToID[artistAddresses[i]];
            if (artistID != 0 && artistOwnerRightsRevokedBitMap.get(artistID)) {
                revert ArtistRevokedOwnerRightsError(artistAddresses[i]);
            }
        }
    }

    /**
     * @notice Ensures an address is assigned an artist ID. If not, creates one.
     * @param artistAddress The address of the artist to ensure has an ID
     */
    function _ensureArtistHasID(address artistAddress) internal {
        if (artistAddress == address(0)) {
            revert ZeroAddressNotAllowedError();
        }
        if (artistAddressToID[artistAddress] == 0) {
            uint256 artistID = nextArtistID++;
            artists[artistID].artistAddress = artistAddress;
            artistAddressToID[artistAddress] = artistID;
        }
    }

    /**
     * @notice Adds a pending collaborator request
     * @param artistID The ID of the artist to add the pending collaborator request for
     * @param seriesID The ID of the series to add the pending collaborator request for
     */
    function _addCollaboratorProposal(uint256 artistID, uint256 seriesID) internal {
        // Add to artist tracking
        artistPendingCollaboratorSeries[artistID].push(seriesID);
        
        // Add to series tracking
        seriesPendingCollaborators[seriesID].push(artistID);
    }

    /**
     * @notice Removes a pending collaborator request
     * @param artistID The ID of the artist to remove the pending collaborator request for
     * @param seriesID The ID of the series to remove the pending collaborator request for
     */
    function _removeCollaboratorProposal(uint256 artistID, uint256 seriesID) internal {
        _removeArtistPendingCollaboratorSeries(artistID, seriesID);
        _removeSeriesPendingCollaborator(seriesID, artistID);
    }

    /**
     * @notice Removes a pending collaborator request from artist
     * @param artistID The ID of the artist to remove the pending collaborator request for
     * @param seriesID The ID of the series to remove the pending collaborator request for
     */
    function _removeArtistPendingCollaboratorSeries(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage pendingSeries = artistPendingCollaboratorSeries[artistID];
        uint256 pendingSeriesLength = pendingSeries.length;
        for (uint256 i = 0; i < pendingSeriesLength; i++) {
            if (pendingSeries[i] == seriesID) {
                if (i != pendingSeriesLength - 1) {
                    pendingSeries[i] = pendingSeries[pendingSeriesLength - 1];
                }
                pendingSeries.pop();
                break;
            }
        }
    }

    /**
     * @notice Removes a pending collaborator request from series
     * @param seriesID The ID of the series to remove the pending collaborator request for
     * @param artistID The ID of the artist to remove the pending collaborator request for
     */
    function _removeSeriesPendingCollaborator(uint256 seriesID, uint256 artistID) internal {
        uint256[] storage pendingArtists = seriesPendingCollaborators[seriesID];
        uint256 pendingArtistLength = pendingArtists.length;
        for (uint256 i = 0; i < pendingArtistLength; i++) {
            if (pendingArtists[i] == artistID) {
                if (i != pendingArtistLength - 1) {
                    pendingArtists[i] = pendingArtists[pendingArtistLength - 1];
                }
                pendingArtists.pop();
                break;
            }
        }
    }

    /**
     * @notice Adds multiple artists to a series
     * @param seriesID The ID of the series to add the artists to
     * @param artistAddresses The addresses of the artists to add to the series
     * @return An array of artist IDs
     */
    function _addArtistsToSeries(
        uint256 seriesID,
        address[] memory artistAddresses
    ) internal returns (uint256[] memory) {
        uint256[] memory artistIDs = new uint256[](artistAddresses.length);
        
        for (uint256 i = 0; i < artistAddresses.length; i++) {
            _ensureArtistHasID(artistAddresses[i]);
            uint256 artistID = artistAddressToID[artistAddresses[i]];
            artistIDs[i] = artistID;

            seriesRegistry[seriesID].artistIDs.push(artistID);
            seriesArtistExist[seriesID].set(artistID);
            artists[artistID].seriesIDs.push(seriesID);
        }

        return artistIDs;
    }
}