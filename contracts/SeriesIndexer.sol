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

    // Co-Artist Management
    mapping(uint256 => BitMaps.BitMap) private seriesPendingCoArtist;
    mapping(uint256 => uint256[]) private seriesPendingCoArtists;   // List of series's pending co-artists
    mapping(uint256 => uint256[]) private artistPendingCoArtistSeries;   // List of artist's pending co-artist series

    // ============ Events ============

    event ArtistAddressUpdated(
        uint256 indexed artistID, 
        address indexed oldAddress, 
        address indexed newAddress
    );
    
    event SeriesIndexed(uint256 indexed seriesID);
    event SeriesUpdated(uint256 indexed seriesID);
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
     * @dev Creates a new series
     */
    function addSeries(
        address[] calldata artistAddresses,
        string calldata metadataURI,
        string calldata tokenIDsMapURI
    ) external returns (uint256) {
        return _createSeries(artistAddresses, metadataURI, tokenIDsMapURI);
    }

    /**
     * @dev Batch creation of series
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
        _validateSeriesWriteBatch(seriesIDs.length, metadataURIs, tokenIDsMapURIs);
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
    function resignFromSeries(uint256 seriesID) external seriesExists(seriesID) onlyArtist(seriesID) {
        uint256 artistID = artistAddressToID[msg.sender];

        _unlinkArtistFromSeries(seriesID, artistID);
        _unlinkSeriesFromArtist(artistID, seriesID);

        emit SeriesUpdated(seriesID);
    }

    /**
     * @dev Updates artist list for a series (owner only)
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

    // ============ Public/External Functions: Co-Artist Management ============

    /**
     * @dev Proposes a new co-artist for a series
     */
    function proposeCoArtist(
        uint256 seriesID,
        address proposedArtistAddress
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        _ensureArtistHasID(proposedArtistAddress);
        uint256 proposedArtistID = artistAddressToID[proposedArtistAddress];

        if (seriesArtistExist[seriesID].get(proposedArtistID)) {
            revert ArtistAlreadyInSeriesError(seriesID, proposedArtistAddress);
        }
        if (seriesPendingCoArtist[seriesID].get(proposedArtistID)) {
            revert AlreadyProposedError(seriesID, proposedArtistAddress);
        }

        seriesPendingCoArtist[seriesID].set(proposedArtistID);
        _addCoArtistProposal(proposedArtistID, seriesID);

        emit CoArtistProposed(seriesID, proposedArtistID);
    }

    /**
     * @dev Cancels a co-artist proposal
     */
    function cancelProposeCoArtist(
        uint256 seriesID,
        address proposedArtistAddress
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        uint256 proposedArtistID = artistAddressToID[proposedArtistAddress];
        if (!seriesPendingCoArtist[seriesID].get(proposedArtistID)) {
            revert NoProposalExistsError(seriesID, proposedArtistAddress);
        }

        seriesPendingCoArtist[seriesID].unset(proposedArtistID);
        _removeCoArtistProposal(proposedArtistID, seriesID);

        uint256 artistID = artistAddressToID[msg.sender];
        emit CoArtistProposalCancelled(seriesID, artistID, proposedArtistID);
    }

    /**
     * @dev Confirms participation as a co-artist
     */
    function confirmAsCoArtist(uint256 seriesID) external seriesExists(seriesID) {
        uint256 artistID = artistAddressToID[msg.sender];
        if (artistID == 0) {
            revert NotAnArtistError(msg.sender);
        }
        if (!seriesPendingCoArtist[seriesID].get(artistID)) {
            revert NotAPendingProposalError(seriesID, msg.sender);
        }

        address[] memory artistAddresses = new address[](1);
        artistAddresses[0] = msg.sender;

        seriesPendingCoArtist[seriesID].unset(artistID);
        _removeCoArtistProposal(artistID, seriesID);
        _addArtistsToSeries(seriesID, artistAddresses);

        emit CoArtistConfirmed(seriesID, artistID);
    }

    // ============ Public/External Functions: Artist Management ============

    /**
     * @dev Revokes owner rights for the calling artist
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
     * @dev Approves owner rights for the calling artist
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
     * @dev Updates an artist's address
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
     * @notice Returns `true` if at least one artist in the series
     *         has not revoked the owner's right to modify the series.
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
     * @notice Checks if the given `artistAddress` has revoked owner rights
     */
    function hasArtistRevokedOwnerRights(address artistAddress) external view returns (bool) {
        uint256 artistID = artistAddressToID[artistAddress];
        return artistOwnerRightsRevokedBitMap.get(artistID);
    }

    /**
     * @notice Returns all artist IDs for a given series
     */
    function getSeriesArtistIDs(uint256 seriesID) external view returns (uint256[] memory) {
        return seriesRegistry[seriesID].artistIDs;
    }

    /**
     * @notice Returns the artist addresses for a given series
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
     * @notice Returns all series IDs associated with a given artist address
     */
    function getArtistSeriesIDs(address artistAddress) external view returns (uint256[] memory) {
        uint256 artistID = artistAddressToID[artistAddress];
        return artists[artistID].seriesIDs;
    }

    /**
     * @notice Returns metadataURI of a given series
     */
    function getSeriesMetadataURI(uint256 seriesID) external view returns (string memory) {
        return seriesRegistry[seriesID].metadataURI;
    }

    /**
     * @notice Returns contract-level token data (e.g., IPFS CID) of a given series
     */
    function getSeriesContractTokenDataURI(uint256 seriesID) external view returns (string memory) {
        return seriesRegistry[seriesID].contractTokenDataURI;
    }

    /**
     * @notice Returns the pending co-artist requests for a given artist address
     */
    function getArtistPendingCoArtistSeries(address artistAddress) external view returns (uint256[] memory) {
        uint256 artistID = artistAddressToID[artistAddress];
        return artistPendingCoArtistSeries[artistID];
    }

    /**
     * @notice Returns the pending co-artist list for a given series
     */
    function getSeriesPendingCoArtists(uint256 seriesID) external view returns (address[] memory) {
        uint256[] memory artistIDs = seriesPendingCoArtists[seriesID];
        address[] memory artistAddresses = new address[](artistIDs.length);
        for (uint256 i = 0; i < artistIDs.length; i++) {
            artistAddresses[i] = artists[artistIDs[i]].artistAddress;
        }
        return artistAddresses;
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
    function getAddressArtistID(address artistAddress) external view returns (uint256) {
        return artistAddressToID[artistAddress];
    }

    // ============ Internal Functions ============

    /**
     * @dev Creates a new series with the specified artists
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

        emit SeriesIndexed(seriesID);
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

        Series storage series = seriesRegistry[seriesID];
        series.metadataURI = metadataURI;
        series.contractTokenDataURI = tokenIDsMapURI;

        emit SeriesUpdated(seriesID);
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
        uint256[] memory pendingCoArtists = seriesPendingCoArtists[seriesID];
        for (uint256 i = 0; i < pendingCoArtists.length; i++) {
            uint256 artistID = pendingCoArtists[i];
            seriesPendingCoArtist[seriesID].unset(artistID);
            _removeArtistPendingCoArtistSeries(artistID, seriesID);
            _removeSeriesPendingCoArtist(seriesID, artistID);
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
     * @dev Removes a series from an artist's list
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
     * @dev Removes an artist from a series
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
     * @dev Validates that artists can only add series that only have themselves as artist
     */
    function _ensureCallerIsOnlyArtist(
        address[] memory artistAddresses
    ) internal view {
        if (artistAddresses.length != 1 || artistAddresses[0] != msg.sender) {
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
     * @dev Checks if the caller is contract owner
     */
    function _isCallerOwner() internal view returns (bool) {
        return msg.sender == owner();
    }

    /**
     * @dev Checks if the caller is a series artist
     */
    function _requireCallerIsSeriesArtist(uint256 seriesID) internal view {
        uint256 artistID = artistAddressToID[msg.sender];
        if (artistID == 0 || !seriesArtistExist[seriesID].get(artistID)) {
            revert CallerNotASeriesArtistError(seriesID, msg.sender);
        }
    }

    /**
     * @dev Ensures all provided artist addresses have not revoked owner rights
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
     * @dev Ensures an address is assigned an artist ID. If not, creates one.
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
     * @dev Adds a pending co-artist request
     */
    function _addCoArtistProposal(uint256 artistID, uint256 seriesID) internal {
        // Add to artist tracking
        artistPendingCoArtistSeries[artistID].push(seriesID);
        
        // Add to series tracking
        seriesPendingCoArtists[seriesID].push(artistID);
    }

    /**
     * @dev Removes a pending co-artist request
     */
    function _removeCoArtistProposal(uint256 artistID, uint256 seriesID) internal {
        _removeArtistPendingCoArtistSeries(artistID, seriesID);
        _removeSeriesPendingCoArtist(seriesID, artistID);
    }

    /**
     * @dev Removes a pending co-artist request from artist
     */
    function _removeArtistPendingCoArtistSeries(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage pendingSeries = artistPendingCoArtistSeries[artistID];
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
     * @dev Removes a pending co-artist request from series
     */
    function _removeSeriesPendingCoArtist(uint256 seriesID, uint256 artistID) internal {
        uint256[] storage pendingArtists = seriesPendingCoArtists[seriesID];
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
     * @dev Adds multiple artists to a series
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