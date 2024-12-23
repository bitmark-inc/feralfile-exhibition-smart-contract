// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol" as Ownable;

/**
 * @title SeriesIndexer
 * @dev A contract for managing series and artists in a collaborative art platform.
 *      This contract allows artists to create series, manage co-artists, and control owner rights.
 */
contract SeriesIndexer is Ownable.Ownable {
    // ============ Custom Errors ============

    error OwnerRightsRevokedForThisSeries(uint256 seriesID);
    error AllArtistsRevokedOwnerRightsError(uint256 seriesID);
    error SeriesDoesNotExistError(uint256 seriesID);
    error CallerNotASeriesArtistError(uint256 seriesID, address caller);
    error NoArtistsProvidedError();
    error NoArtistsForSeriesError();
    error AlreadyRevokedError(uint256 artistID);
    error NotRevokedError(uint256 artistID);
    error InvalidNewAddressError(address newAddress);
    error AddressAlreadyAssignedError(address newAddress);
    error InvalidArtistIDError(uint256 artistID);
    error NotAuthorizedError(address caller);
    error ArtistRevokedOwnerRightsError(address artistAddr);
    error NoProposalExistsError(uint256 seriesID, address proposedArtistAddr);
    error NotAPendingProposalError(uint256 seriesID, address caller);
    error ZeroAddressNotAllowedError();
    error NoSeriesDataError();
    error BatchSizeTooLargeError(uint256 size);
    error ArrayLengthMismatchError(uint256 lenArtists, uint256 lenMetas, uint256 lenTokens);
    error AlreadyProposedError(uint256 seriesID, address proposedArtistAddr);
    error NotAnArtistError(address caller);
    error EmptyMetadataURIError();
    error EmptyTokenMapURIError();

    // ============ Structs ============

    struct Series {
        string metadataURI;              // IPFS hash or similar identifier for series metadata
        string contractTokenDataURI;     // Token-related data for the series
        uint256[] artistIDs;          // List of artists associated with this series
    }

    // ============ State Variables ============

    // ID Counters
    uint256 private nextSeriesID = 1;
    uint256 private nextArtistID = 1;

    // Artist Management
    mapping(uint256 => address) private artistIDToAddress;
    mapping(address => uint256) private addressToArtistID;
    mapping(uint256 => bool) private ownerRightsRevokedForArtistID;

    // Series Management
    mapping(uint256 => Series) private seriesDetails;
    mapping(uint256 => uint256[]) private artistIDSeriesIDs;
    mapping(uint256 => mapping(uint256 => bool)) private isArtistIDInSeries;
    
    // Co-Artist Management
    mapping(uint256 => mapping(uint256 => bool)) private seriesPendingCoArtist;
    mapping(uint256 => uint256[]) private artistPendingCoArtistRequests;
    mapping(uint256 => mapping(uint256 => uint256)) private artistPendingCoArtistRequestIndex;

    // ============ Events ============

    event ArtistAddressUpdated(
        uint256 indexed artistID, 
        address indexed oldAddress, 
        address indexed newAddress
    );
    
    event SeriesIndexed(
        uint256 indexed seriesID,
        uint256[] artistIDs,
        string metadataURI,
        string seriesContractTokenURI
    );
    
    event SeriesUpdated(
        uint256 indexed seriesID,
        uint256[] artistIDs,
        string metadataURI,
        string seriesContractTokenURI
    );
    
    event SeriesDeleted(uint256 indexed seriesID);
    
    event CoArtistProposed(
        uint256 indexed seriesID, 
        uint256 indexed proposedArtistID
    );
    
    event CoArtistConfirmed(
        uint256 indexed seriesID, 
        uint256 indexed confirmedArtistID
    );
    
    event CoArtistProposalCancelled(
        uint256 indexed seriesID,
        uint256 indexed proposerArtistID,
        uint256 indexed cancelledArtistID
    );

    // ============ Modifiers ============

    modifier onlyOwnerOrArtist(uint256 seriesID) {
        if (msg.sender == owner()) {
            // If owner is calling, check if at least one artist in this series has NOT revoked rights
            if (!ownerCanModifySeries(seriesID)) {
                revert OwnerRightsRevokedForThisSeries(seriesID);
            }
        } else {
            // Otherwise, must be a valid artist in this series
            _requireIsSeriesArtist(seriesID);
        }
        _;
    }

    modifier onlyArtist(uint256 seriesID) {
        _requireIsSeriesArtist(seriesID);
        _;
    }

    modifier seriesExists(uint256 seriesID) {
        // If metadata is empty, we consider the series nonexistent
        if (bytes(seriesDetails[seriesID].metadataURI).length == 0) {
            revert SeriesDoesNotExistError(seriesID);
        }
        _;
    }

    // ============ Public/External Functions: Series Management ============

    /**
     * @dev Creates a new series with the caller as the artist
     */
    function addSeries(
        address[] calldata artistAddrs,
        string calldata metadataURI,
        string calldata tokenIDsMapURI
    ) external returns (uint256) {
        return _createSeries(artistAddrs, metadataURI, tokenIDsMapURI);
    }

    /**
     * @dev Batch creation of series by owner
     */
    function batchAddSeries(
        address[][] calldata artistsArray,
        string[] calldata metadataURIs,
        string[] calldata tokenIDsMapURIs
    ) external returns (uint256[] memory) {
        _validateBatchParameters(artistsArray.length, metadataURIs, tokenIDsMapURIs);

        uint256[] memory createdSeriesIDs = new uint256[](artistsArray.length);
        for (uint256 i = 0; i < artistsArray.length; i++) {
            createdSeriesIDs[i] = _createSeries(
                artistsArray[i],
                metadataURIs[i],
                tokenIDsMapURIs[i]
            );
        }
        return createdSeriesIDs;
    }

    /**
     * @dev Updates an existing series' metadata and token data
     */
    function updateSeries(
        uint256 seriesID,
        string calldata metadataURI,
        string calldata tokenIDsMapURI
    ) external {
        _updateSeries(seriesID, metadataURI, tokenIDsMapURI);
    }

    /**
     * @dev Updates an existing series' metadata and token data
     */
    function batchUpdateSeries(
        uint256[] calldata seriesIDs,
        string[] calldata metadataURIs,
        string[] calldata tokenIDsMapURIs
    ) external {
        _validateBatchParameters(seriesIDs.length, metadataURIs, tokenIDsMapURIs);
        for (uint256 i = 0; i < seriesIDs.length; i++) {
            _updateSeries(seriesIDs[i], metadataURIs[i], tokenIDsMapURIs[i]);
        }
    }

    /**
     * @dev Deletes series and cleans up related data
     */
    function deleteSeries(uint256 seriesID) external {
        _deleteSeries(seriesID);
    }

    /**
     * @dev Deletes series and cleans up related data
     */
    function batchDeleteSeries(uint256[] calldata seriesIDs) external {
        for (uint256 i = 0; i < seriesIDs.length; i++) {
            _deleteSeries(seriesIDs[i]);
        }
    }

    // ============ Public/External Functions: Series Artist Management ============

    /**
     * @dev Allows an artist to remove themselves from a series
     */
    function removeSelfFromSeries(uint256 seriesID) external onlyArtist(seriesID) {
        uint256 artistID = addressToArtistID[msg.sender];
        _removeArtistFromSeries(seriesID, artistID);
    }

    /**
     * @dev Updates artist list for a series (owner only)
     */
    function ownerUpdateSeriesArtists(
        uint256 seriesID,
        address[] calldata artistAddrs
    ) external seriesExists(seriesID) onlyOwner {
        if (!ownerCanModifySeries(seriesID)) {
            revert AllArtistsRevokedOwnerRightsError(seriesID);
        }
        _validateArtistsNotRevoked(artistAddrs);

        // Remove current artists
        Series storage series = seriesDetails[seriesID];
        for (uint256 i = 0; i < series.artistIDs.length; i++) {
            uint256 artistID = series.artistIDs[i];
            isArtistIDInSeries[seriesID][artistID] = false;
            _removeSeriesFromArtist(artistID, seriesID);
        }
        delete series.artistIDs;

        // Add new artists
        _addArtistsToSeries(seriesID, artistAddrs);

        emit SeriesUpdated(
            seriesID,
            series.artistIDs,
            series.metadataURI,
            series.contractTokenDataURI
        );
    }

    // ============ Public/External Functions: Co-Artist Management ============

    /**
     * @dev Proposes a new co-artist for a series
     */
    function proposeCoArtist(
        uint256 seriesID,
        address proposedArtistAddr
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        if (proposedArtistAddr == address(0)) {
            revert ZeroAddressNotAllowedError();
        }
        
        _ensureArtistHasID(proposedArtistAddr);
        uint256 proposedArtistID = addressToArtistID[proposedArtistAddr];

        if (seriesPendingCoArtist[seriesID][proposedArtistID]) {
            revert AlreadyProposedError(seriesID, proposedArtistAddr);
        }

        seriesPendingCoArtist[seriesID][proposedArtistID] = true;
        _addPendingRequest(proposedArtistID, seriesID);

        emit CoArtistProposed(seriesID, proposedArtistID);
    }

    /**
     * @dev Cancels a co-artist proposal
     */
    function cancelProposeCoArtist(
        uint256 seriesID,
        address proposedArtistAddr
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        uint256 proposedArtistID = addressToArtistID[proposedArtistAddr];
        if (!seriesPendingCoArtist[seriesID][proposedArtistID]) {
            revert NoProposalExistsError(seriesID, proposedArtistAddr);
        }

        seriesPendingCoArtist[seriesID][proposedArtistID] = false;
        _removePendingRequest(proposedArtistID, seriesID);

        uint256 artistID = addressToArtistID[msg.sender];
        emit CoArtistProposalCancelled(seriesID, artistID, proposedArtistID);
    }

    /**
     * @dev Confirms participation as a co-artist
     */
    function confirmAsCoArtist(uint256 seriesID) external seriesExists(seriesID) {
        uint256 artistID = addressToArtistID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (!seriesPendingCoArtist[seriesID][artistID]) {
            revert NotAPendingProposalError(seriesID, msg.sender);
        }

        Series storage series = seriesDetails[seriesID];
        series.artistIDs.push(artistID);
        isArtistIDInSeries[seriesID][artistID] = true;
        artistIDSeriesIDs[artistID].push(seriesID);

        seriesPendingCoArtist[seriesID][artistID] = false;
        _removePendingRequest(artistID, seriesID);

        emit CoArtistConfirmed(seriesID, artistID);
    }

    // ============ Public/External Functions: Artist Management ============

    /**
     * @dev Revokes owner rights for the calling artist
     */
    function revokeOwnerRightsForArtist() external {
        uint256 artistID = addressToArtistID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (ownerRightsRevokedForArtistID[artistID]) {
            revert AlreadyRevokedError(artistID);
        }
        ownerRightsRevokedForArtistID[artistID] = true;
    }

    /**
     * @dev Approves owner rights for the calling artist
     */
    function approveOwnerRightsForArtist() external {
        uint256 artistID = addressToArtistID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (!ownerRightsRevokedForArtistID[artistID]) {
            revert NotRevokedError(artistID);
        }
        ownerRightsRevokedForArtistID[artistID] = false;
    }

    /**
     * @dev Updates an artist's address
     */
    function updateArtistAddress(uint256 artistID, address newAddress) external {
        if (newAddress == address(0)) {
            revert InvalidNewAddressError(newAddress);
        }
        if (addressToArtistID[newAddress] != 0) {
            revert AddressAlreadyAssignedError(newAddress);
        }
        if (artistIDToAddress[artistID] == address(0)) {
            revert InvalidArtistIDError(artistID);
        }

        uint256 callerID = addressToArtistID[msg.sender];
        bool isCallerArtist = (callerID == artistID && callerID != 0);
        bool isOwnerWithRights = (msg.sender == owner() && !ownerRightsRevokedForArtistID[artistID]);

        if (!isCallerArtist && !isOwnerWithRights) {
            revert NotAuthorizedError(msg.sender);
        }

        address oldAddress = artistIDToAddress[artistID];
        addressToArtistID[oldAddress] = 0;
        artistIDToAddress[artistID] = newAddress;
        addressToArtistID[newAddress] = artistID;

        emit ArtistAddressUpdated(artistID, oldAddress, newAddress);
    }

    // ============ Public/External View Functions ============

    /**
     * @notice Returns `true` if at least one artist in the series
     *         has not revoked the owner's right to modify the series.
     */
    function ownerCanModifySeries(uint256 seriesID) public view returns (bool) {
        uint256[] memory ids = seriesDetails[seriesID].artistIDs;
        for (uint256 i = 0; i < ids.length; i++) {
            // If any artist has not revoked, return true
            if (!ownerRightsRevokedForArtistID[ids[i]]) {
                return true;
            }
        }
        return false;
    }

    /**
     * @notice Checks if the given `artistAddr` has revoked owner rights
     */
    function ownerRightsRevoked(address artistAddr) external view returns (bool) {
        uint256 artistID = addressToArtistID[artistAddr];
        return ownerRightsRevokedForArtistID[artistID];
    }

    /**
     * @notice Returns all artist IDs for a given series
     */
    function getSeriesArtistIDs(uint256 seriesID) external view returns (uint256[] memory) {
        return seriesDetails[seriesID].artistIDs;
    }

    /**
     * @notice Returns the artist addresses for a given series
     */
    function getSeriesArtistAddresses(uint256 seriesID) external view returns (address[] memory) {
        uint256[] memory ids = seriesDetails[seriesID].artistIDs;
        address[] memory artistAddrs = new address[](ids.length);
        for (uint256 i = 0; i < ids.length; i++) {
            artistAddrs[i] = artistIDToAddress[ids[i]];
        }
        return artistAddrs;
    }

    /**
     * @notice Returns all series IDs associated with a given artist address
     */
    function getArtistSeriesIDs(address artistAddr) external view returns (uint256[] memory) {
        uint256 artistID = addressToArtistID[artistAddr];
        return artistIDSeriesIDs[artistID];
    }

    /**
     * @notice Returns metadataURI of a given series
     */
    function getSeriesMetadataURI(uint256 seriesID) external view returns (string memory) {
        return seriesDetails[seriesID].metadataURI;
    }

    /**
     * @notice Returns contract-level token data (e.g., IPFS CID) of a given series
     */
    function getSeriesContractTokenDataURI(uint256 seriesID) external view returns (string memory) {
        return seriesDetails[seriesID].contractTokenDataURI;
    }

    /**
     * @notice Returns the pending co-artist requests for a given artist address
     */
    function getPendingRequests(address artistAddr) external view returns (uint256[] memory) {
        uint256 artistID = addressToArtistID[artistAddr];
        return artistPendingCoArtistRequests[artistID];
    }

    /**
     * @notice Returns the address associated with a given artist ID
     */
    function getArtistAddress(uint256 artistID) external view returns (address) {
        return artistIDToAddress[artistID];
    }

    /**
     * @notice Returns the artist ID associated with a given address
     */
    function getAddressArtistID(address artistAddr) external view returns (uint256) {
        return addressToArtistID[artistAddr];
    }

    // ============ Internal Functions ============

    /**
     * @dev Creates a new series with the specified artists
     */
    function _createSeries(
        address[] memory artistAddrs,
        string memory metadataURI,
        string memory tokenIDsMapURI
    ) internal returns (uint256) {
        if (artistAddrs.length == 0) {
            revert NoArtistsForSeriesError();
        }
        if (msg.sender == owner()) {
            _validateArtistsNotRevoked(artistAddrs);
        } else {
            _validateOnlySelfAsArtist(artistAddrs);
        }
        uint256 seriesID = nextSeriesID++;
        _validateMetadataAndTokenURI(metadataURI, tokenIDsMapURI);

        Series storage series = seriesDetails[seriesID];
        series.metadataURI = metadataURI;
        series.contractTokenDataURI = tokenIDsMapURI;

        uint256[] memory artistIDs = _addArtistsToSeries(seriesID, artistAddrs);

        emit SeriesIndexed(seriesID, artistIDs, metadataURI, tokenIDsMapURI);
        return seriesID;
    }

    /**
     * @dev Update existing series' metadata and token data
     */
    function _updateSeries(
        uint256 seriesID,
        string memory metadataURI,
        string memory tokenIDsMapURI
    ) internal seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        _validateMetadataAndTokenURI(metadataURI, tokenIDsMapURI);

        Series storage series = seriesDetails[seriesID];
        series.metadataURI = metadataURI;
        series.contractTokenDataURI = tokenIDsMapURI;

        emit SeriesUpdated(seriesID, series.artistIDs, metadataURI, tokenIDsMapURI);
    }

    /**
     * @dev Removes an artist from a series and updates related mappings
     */
    function _removeArtistFromSeries(uint256 seriesID, uint256 artistID) internal {
        Series storage series = seriesDetails[seriesID];
        uint256[] storage artistIDs = series.artistIDs;
        
        // Find and remove artistID from series' artist list
        for (uint256 i = 0; i < artistIDs.length; i++) {
            if (artistIDs[i] == artistID) {
                artistIDs[i] = artistIDs[artistIDs.length - 1];
                artistIDs.pop();
                break;
            }
        }

        // Remove series from artist's list
        _removeSeriesFromArtist(artistID, seriesID);

        emit SeriesUpdated(
            seriesID, 
            series.artistIDs, 
            series.metadataURI, 
            series.contractTokenDataURI
        );
    }

    /**
     * @dev Removes a series from an artist's list
     */
    function _removeSeriesFromArtist(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage seriesIDs = artistIDSeriesIDs[artistID];
        
        for (uint256 i = 0; i < seriesIDs.length; i++) {
            if (seriesIDs[i] == seriesID) {
                seriesIDs[i] = seriesIDs[seriesIDs.length - 1];
                seriesIDs.pop();
                break;
            }
        }
        
        isArtistIDInSeries[seriesID][artistID] = false;
    }

        /**
     * @dev Deletes a series and cleans up related data
     */
    function _deleteSeries(uint256 seriesID) 
        internal 
        seriesExists(seriesID) 
        onlyOwnerOrArtist(seriesID) 
    {
        Series storage series = seriesDetails[seriesID];
        
        // Remove series from all artists
        for (uint256 i = 0; i < series.artistIDs.length; i++) {
            _removeSeriesFromArtist(series.artistIDs[i], seriesID);
        }

        delete seriesDetails[seriesID];
        emit SeriesDeleted(seriesID);
    }

    /**
     * @dev Validates that artists can only add series that only have themselves as artist
     */
    function _validateOnlySelfAsArtist(
        address[] memory artistAddrs
    ) internal view {
        if (artistAddrs.length != 1 || artistAddrs[0] != msg.sender) {
            revert NotAuthorizedError(msg.sender);
        }
    }

    /**
     * @dev Validates that metadataURI and token URI are not empty
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
     * @dev Validates batch parameters for series creation
     */
    function _validateBatchParameters(
        uint256 seriesCount,
        string[] memory metadataURIs,
        string[] memory tokenIDsMapURIs
    ) internal pure {
        if (seriesCount != metadataURIs.length || seriesCount != tokenIDsMapURIs.length) {
            revert ArrayLengthMismatchError(
                seriesCount,
                metadataURIs.length,
                tokenIDsMapURIs.length
            );
        }
        if (seriesCount == 0) {
            revert NoSeriesDataError();
        }
        if (seriesCount > 50) {
            revert BatchSizeTooLargeError(seriesCount);
        }
    }

    /**
     * @dev Checks if the caller is a series artist
     */
    function _requireIsSeriesArtist(uint256 seriesID) internal view {
        uint256 artistID = addressToArtistID[msg.sender];
        if (artistID == 0 || !isArtistIDInSeries[seriesID][artistID]) {
            revert CallerNotASeriesArtistError(seriesID, msg.sender);
        }
    }

    /**
     * @dev Ensures all provided artist addresses have not revoked owner rights
     */
    function _validateArtistsNotRevoked(address[] memory artistAddrs) internal view {
        for (uint256 i = 0; i < artistAddrs.length; i++) {
            uint256 artistID = addressToArtistID[artistAddrs[i]];
            if (artistID != 0 && ownerRightsRevokedForArtistID[artistID]) {
                revert ArtistRevokedOwnerRightsError(artistAddrs[i]);
            }
        }
    }

    /**
     * @dev Ensures an address is assigned an artist ID. If not, creates one.
     */
    function _ensureArtistHasID(address artistAddr) internal {
        if (artistAddr == address(0)) {
            revert ZeroAddressNotAllowedError();
        }
        if (addressToArtistID[artistAddr] == 0) {
            uint256 artistID = nextArtistID++;
            artistIDToAddress[artistID] = artistAddr;
            addressToArtistID[artistAddr] = artistID;
        }
    }

    /**
     * @dev Adds a pending co-artist request
     */
    function _addPendingRequest(uint256 artistID, uint256 seriesID) internal {
        artistPendingCoArtistRequests[artistID].push(seriesID);
        artistPendingCoArtistRequestIndex[artistID][seriesID] =
            artistPendingCoArtistRequests[artistID].length - 1;
    }

    /**
     * @dev Removes a pending co-artist request
     */
    function _removePendingRequest(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage requests = artistPendingCoArtistRequests[artistID];
        uint256 index = artistPendingCoArtistRequestIndex[artistID][seriesID];
        uint256 lastIndex = requests.length - 1;
        
        if (index != lastIndex) {
            uint256 lastValue = requests[lastIndex];
            requests[index] = lastValue;
            artistPendingCoArtistRequestIndex[artistID][lastValue] = index;
        }
        
        requests.pop();
        delete artistPendingCoArtistRequestIndex[artistID][seriesID];
    }

    /**
     * @dev Adds multiple artists to a series
     */
    function _addArtistsToSeries(
        uint256 seriesID,
        address[] memory artistAddrs
    ) internal returns (uint256[] memory) {
        uint256[] memory artistIDs = new uint256[](artistAddrs.length);
        
        for (uint256 i = 0; i < artistAddrs.length; i++) {
            _ensureArtistHasID(artistAddrs[i]);
            uint256 artistID = addressToArtistID[artistAddrs[i]];
            artistIDs[i] = artistID;

            seriesDetails[seriesID].artistIDs.push(artistID);
            isArtistIDInSeries[seriesID][artistID] = true;
            artistIDSeriesIDs[artistID].push(seriesID);
        }

        return artistIDs;
    }
}