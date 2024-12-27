// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/BitMaps.sol";

/**
 * @title SeriesIndexer
 * @dev A contract for managing series and artists in a collaborative art platform.
 *      This contract allows artists to create series, manage co-artists, and control owner rights.
 */
contract SeriesIndexer is Ownable {
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
    error ArtistRevokedOwnerRightsError(address artistAddr);
    error NoProposalExistsError(uint256 seriesID, address proposedArtistAddr);
    error NotAPendingProposalError(uint256 seriesID, address caller);
    error ZeroAddressNotAllowedError();
    error NoSeriesDataError();
    error BatchSizeTooLargeError(uint256 size);
    error ArrayLengthMismatchError(uint256 lenArtists, uint256 lenMetas, uint256 lenTokens);
    error AlreadyProposedError(uint256 seriesID, address proposedArtistAddr);
    error ArtistAlreadyInSeriesError(uint256 seriesID, address proposedArtistAddr);
    error NotAnArtistError(address caller);
    error EmptyMetadataURIError();
    error EmptyTokenMapURIError();

    // ============ Structs ============

    struct Series {
        string metadataURI;              // IPFS hash or similar identifier for series metadata
        string contractTokenDataURI;     // Token-related data for the series
        uint256[] artistIDs;          // List of artists associated with this series
        uint256[] pendingCoArtists;   // List of pending co-artists
    }

    struct Artist {
        address artistAddress;           // Artist's wallet address
        uint256[] seriesIDs;            // List of series associated with this artist
        uint256[] pendingCoArtistSeries; // List of pending co-artist series
    }

    // ============ State Variables ============

    // ID Counters
    uint256 private nextSeriesID = 1;
    uint256 private nextArtistID = 1;

    // Artist Management
    mapping(uint256 => Artist) private artists;
    mapping(address => uint256) private addressToArtistID;
    mapping(uint256 => BitMaps.BitMap) private artistOwnerRightsRevokedBitMap;

    // Series Management
    mapping(uint256 => Series) private seriesDetails;

    // Series - Artist Relationship
    mapping(uint256 => mapping(uint256 => bool)) private seriesArtistExist;

    // Co-Artist Management
    mapping(uint256 => mapping(uint256 => bool)) private seriesPendingCoArtist;
    mapping(uint256 => mapping(uint256 => uint256)) private seriesPendingCoArtistIndex;
    mapping(uint256 => mapping(uint256 => uint256)) private artistPendingCoArtistSeriesIndex;

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
     * @dev Creates a new series
     */
    function addSeries(
        address[] calldata artistAddrs,
        string calldata metadataURI,
        string calldata tokenIDsMapURI
    ) external returns (uint256) {
        return _createSeries(artistAddrs, metadataURI, tokenIDsMapURI);
    }

    /**
     * @dev Batch creation of series
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
     * @dev Batch deletes series and cleans up related data
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
    function removeSelfFromSeries(uint256 seriesID) external seriesExists(seriesID) onlyArtist(seriesID) {
        uint256 artistID = addressToArtistID[msg.sender];

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
     * @dev Updates artist list for a series (owner only)
     */
    function ownerUpdateSeriesArtists(
        uint256 seriesID,
        address[] calldata artistAddrs
    ) external seriesExists(seriesID) onlyOwner {
        if (!ownerCanModifySeries(seriesID)) {
            revert OwnerRightsRevokedForThisSeries(seriesID);
        }
        _validateArtistsNotRevoked(artistAddrs);

        // Remove current artists
        Series storage series = seriesDetails[seriesID];
        for (uint256 i = 0; i < series.artistIDs.length; i++) {
            uint256 artistID = series.artistIDs[i];
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
        _ensureArtistHasID(proposedArtistAddr);
        uint256 proposedArtistID = addressToArtistID[proposedArtistAddr];

        if (seriesArtistExist[seriesID][proposedArtistID]) {
            revert ArtistAlreadyInSeriesError(seriesID, proposedArtistAddr);
        }
        if (seriesPendingCoArtist[seriesID][proposedArtistID]) {
            revert AlreadyProposedError(seriesID, proposedArtistAddr);
        }

        seriesPendingCoArtist[seriesID][proposedArtistID] = true;
        _addPendingCoArtistRequest(proposedArtistID, seriesID);

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
        _removePendingCoArtistRequest(proposedArtistID, seriesID);

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

        address[] memory artistAddrs = new address[](1);
        artistAddrs[0] = msg.sender;

        seriesPendingCoArtist[seriesID][artistID] = false;
        _removePendingCoArtistRequest(artistID, seriesID);
        _addArtistsToSeries(seriesID, artistAddrs);

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
        if (_isArtistOwnerRightRevoked(artistID)) {
            return;
        }
        _setArtistOwnerRightRevokedBool(artistID, true);
    }

    /**
     * @dev Approves owner rights for the calling artist
     */
    function approveOwnerRightsForArtist() external {
        uint256 artistID = addressToArtistID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (!_isArtistOwnerRightRevoked(artistID)) {
            return;
        }
        _setArtistOwnerRightRevokedBool(artistID, false);
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
        if (artists[artistID].artistAddress == address(0)) {
            revert InvalidArtistIDError(artistID);
        }

        uint256 callerID = addressToArtistID[msg.sender];
        bool isCallerArtist = (callerID == artistID && callerID != 0);
        bool isOwnerWithRights = (msg.sender == owner() && !_isArtistOwnerRightRevoked(artistID));

        if (!isCallerArtist && !isOwnerWithRights) {
            revert NotAuthorizedError(msg.sender);
        }

        address oldAddress = artists[artistID].artistAddress;
        addressToArtistID[oldAddress] = 0;
        artists[artistID].artistAddress = newAddress;
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
        if (ids.length == 0) {
            // If no artists, owner can modify the series
            return true;
        }
        for (uint256 i = 0; i < ids.length; i++) {
            // If any artist has not revoked, return true
            if (!_isArtistOwnerRightRevoked(ids[i])) {
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
        return _isArtistOwnerRightRevoked(artistID);
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
            artistAddrs[i] = artists[ids[i]].artistAddress;
        }
        return artistAddrs;
    }

    /**
     * @notice Returns all series IDs associated with a given artist address
     */
    function getArtistSeriesIDs(address artistAddr) external view returns (uint256[] memory) {
        uint256 artistID = addressToArtistID[artistAddr];
        return artists[artistID].seriesIDs;
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
    function getArtistPendingCoArtistSeries(address artistAddr) external view returns (uint256[] memory) {
        uint256 artistID = addressToArtistID[artistAddr];
        return artists[artistID].pendingCoArtistSeries;
    }

    /**
     * @notice Returns the pending co-artist list for a given series
     */
    function getSeriesPendingCoArtists(uint256 seriesID) external view returns (address[] memory) {
        uint256[] memory artistIDs = seriesDetails[seriesID].pendingCoArtists;
        address[] memory artistAddrs = new address[](artistIDs.length);
        for (uint256 i = 0; i < artistIDs.length; i++) {
            artistAddrs[i] = artists[artistIDs[i]].artistAddress;
        }
        return artistAddrs;
    }

    /**
     * @notice Returns the address associated with a given artist ID
     */
    function getArtistAddress(uint256 artistID) external view returns (address) {
        return artists[artistID].artistAddress;
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
     * @dev Deletes a series and cleans up related data
     */
    function _deleteSeries(uint256 seriesID) 
        internal 
        seriesExists(seriesID) 
        onlyOwnerOrArtist(seriesID) 
    {
        // Clean up all pending co-artist requests for this series in one loop
        uint256[] memory pendingCoArtists = seriesDetails[seriesID].pendingCoArtists;
        for (uint256 i = 0; i < pendingCoArtists.length; i++) {
            uint256 artistID = pendingCoArtists[i];
            seriesPendingCoArtist[seriesID][artistID] = false;
            _removeArtistPendingCoArtistSeries(artistID, seriesID);
        }

        Series memory series = seriesDetails[seriesID];
        
        // Remove series from all artists
        for (uint256 i = 0; i < series.artistIDs.length; i++) {
            _removeSeriesFromArtist(series.artistIDs[i], seriesID);
        }

        delete seriesDetails[seriesID];

        emit SeriesDeleted(seriesID);
    }

    /**
     * @dev Removes a series from an artist's list
     */
    function _removeSeriesFromArtist(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage seriesIDs = artists[artistID].seriesIDs;
        
        for (uint256 i = 0; i < seriesIDs.length; i++) {
            if (seriesIDs[i] == seriesID) {
                seriesIDs[i] = seriesIDs[seriesIDs.length - 1];
                seriesIDs.pop();
                break;
            }
        }
        
        seriesArtistExist[seriesID][artistID] = false;
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
        if (artistID == 0 || !seriesArtistExist[seriesID][artistID]) {
            revert CallerNotASeriesArtistError(seriesID, msg.sender);
        }
    }

    /**
     * @dev Ensures all provided artist addresses have not revoked owner rights
     */
    function _validateArtistsNotRevoked(address[] memory artistAddrs) internal view {
        for (uint256 i = 0; i < artistAddrs.length; i++) {
            uint256 artistID = addressToArtistID[artistAddrs[i]];
            if (artistID != 0 && _isArtistOwnerRightRevoked(artistID)) {
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
            artists[artistID].artistAddress = artistAddr;
            addressToArtistID[artistAddr] = artistID;
        }
    }

    /**
     * @dev Adds a pending co-artist request
     */
    function _addPendingCoArtistRequest(uint256 artistID, uint256 seriesID) internal {
        // Add to artist tracking
        artists[artistID].pendingCoArtistSeries.push(seriesID);
        artistPendingCoArtistSeriesIndex[artistID][seriesID] =
            artists[artistID].pendingCoArtistSeries.length - 1;
        
        // Add to series tracking
        seriesDetails[seriesID].pendingCoArtists.push(artistID);
        seriesPendingCoArtistIndex[seriesID][artistID] =
            seriesDetails[seriesID].pendingCoArtists.length - 1;
    }

    /**
     * @dev Removes a pending co-artist request
     */
    function _removePendingCoArtistRequest(uint256 artistID, uint256 seriesID) internal {
        _removeArtistPendingCoArtistSeries(artistID, seriesID);
        _removeSeriesPendingCoArtist(seriesID, artistID);
    }

    /**
     * @dev Removes a pending co-artist request from artist
     */
    function _removeArtistPendingCoArtistSeries(uint256 artistID, uint256 seriesID) private {
        uint256[] storage artistCoArtistSeries = artists[artistID].pendingCoArtistSeries;
        uint256 artistCoArtistSeriesIndex = artistPendingCoArtistSeriesIndex[artistID][seriesID];
        uint256 lastArtistCoArtistSeriesIndex = artistCoArtistSeries.length - 1;
        
        if (artistCoArtistSeriesIndex != lastArtistCoArtistSeriesIndex) {
            uint256 lastValue = artistCoArtistSeries[lastArtistCoArtistSeriesIndex];
            artistCoArtistSeries[artistCoArtistSeriesIndex] = lastValue;
            artistPendingCoArtistSeriesIndex[artistID][lastValue] = artistCoArtistSeriesIndex;
        }
        
        artistCoArtistSeries.pop();
        delete artistPendingCoArtistSeriesIndex[artistID][seriesID];
    }

    /**
     * @dev Removes a pending co-artist request from series
     */
    function _removeSeriesPendingCoArtist(uint256 seriesID, uint256 artistID) private {
        uint256[] storage seriesPendingCoArtists = seriesDetails[seriesID].pendingCoArtists;
        uint256 seriesPendingArtistIndex = seriesPendingCoArtistIndex[seriesID][artistID];
        uint256 lastSeriesPendingArtistIndex = seriesPendingCoArtists.length - 1;
        
        if (seriesPendingArtistIndex != lastSeriesPendingArtistIndex) {
            uint256 lastValue = seriesPendingCoArtists[lastSeriesPendingArtistIndex];
            seriesPendingCoArtists[seriesPendingArtistIndex] = lastValue;
            seriesPendingCoArtistIndex[seriesID][lastValue] = seriesPendingArtistIndex;
        }
        
        seriesPendingCoArtists.pop();
        delete seriesPendingCoArtistIndex[seriesID][artistID];
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
            seriesArtistExist[seriesID][artistID] = true;
            artists[artistID].seriesIDs.push(seriesID);
        }

        return artistIDs;
    }

    // ============ BitMap Management Functions ============
    
    // Helper function to get OwnerRightRevoked bitmap position of artistID.
    function _getArtistOwnerRightRevokedBitMapPosition(uint256 artistID)
        private
        pure
        returns (uint256 chunkKey, uint8 bitIndex)
    {
        chunkKey = artistID / 256;
        bitIndex = uint8(artistID % 256);
    }
    
    function _isArtistOwnerRightRevoked(uint256 artistID) internal view returns (bool) {
        (uint256 chunk, uint8 bitIndex) = _getArtistOwnerRightRevokedBitMapPosition(artistID);
        return artistOwnerRightsRevokedBitMap[chunk].get(bitIndex);
    }

    // Artist in series functions
    function _setArtistOwnerRightRevokedBool(uint256 artistID, bool boolean) internal {
        (uint256 chunk, uint8 bitIndex) = _getArtistOwnerRightRevokedBitMapPosition(artistID);
        artistOwnerRightsRevokedBitMap[chunk].setTo(bitIndex, boolean);
    }
}