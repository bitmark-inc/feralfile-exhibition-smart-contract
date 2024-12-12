// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract SeriesIndexer is Ownable {
    // Counter
    uint256 private nextSeriesID = 1;
    uint256 private nextArtistID = 1;

    struct Series {
        string metadata;
        string contractTokenData;
        uint256[] artistIDs;
    }

    // Artist Management
    mapping(uint256 => address) private artistIDToAddress;
    mapping(address => uint256) private addressToArtistID;
    mapping(uint256 => bool) private ownerRightsRevokedForArtistID;

    // Mappings for series and artists
    mapping(uint256 => Series) private seriesDetails;
    mapping(uint256 => uint256[]) private artistIDSeriesIDs;
    mapping(uint256 => mapping(uint256 => bool)) private isArtistIDInSeries;// SeriesID => ArtistID => Bool
    mapping(uint256 => mapping(uint256 => bool)) private seriesPendingCoArtist;// SeriesID => ArtistID => Bool

    // Pending Requests Management
    mapping(uint256 => uint256[]) private artistPendingCoArtistRequests;// artistID => SeriesID[]
    mapping(uint256 => mapping(uint256 => uint256)) private artistPendingCoArtistRequestIndex;// artistID => SeriesID => Index

    // Events
    event ArtistAddressUpdated(uint256 indexed artistID, address oldAddress, address newAddress);
    event SeriesIndexed(
        uint256 indexed seriesID,
        uint256[] artistIDs,
        string metadata,
        string seriesContractTokenCID
    );
    event SeriesUpdated(
        uint256 indexed seriesID,
        uint256[] artistIDs,
        string metadata,
        string seriesContractTokenCID
    );
    event SeriesDeleted(uint256 indexed seriesID);
    event CoArtistProposed(uint256 indexed seriesID, uint256 indexed proposedArtistID);
    event CoArtistConfirmed(uint256 indexed seriesID, uint256 indexed confirmedArtistID);
    event CoArtistProposalCancelled(
        uint256 indexed seriesID, 
        uint256 indexed proposerArtistID, 
        uint256 indexed cancelledArtistID
    );

    // Modifiers
    modifier onlyOwnerOrArtist(uint256 seriesID) {
        _checkOwnerOrArtist(seriesID);
        _;
    }

    modifier onlyArtist(uint256 seriesID) {
        _checkArtist(seriesID);
        _;
    }

    modifier seriesExists(uint256 seriesID) {
        require(bytes(seriesDetails[seriesID].metadata).length > 0, "Series does not exist");
        _;
    }
    // ------------------------
    // Internal Helper Functions
    // ------------------------

    function _checkOwnerOrArtist(uint256 seriesID) internal view {
        if (msg.sender == owner()) {
            require(ownerCanModifySeries(seriesID), "Owner rights revoked by all artists in series");
        } else {
            _checkArtist(seriesID);
        }
    }

    function _checkArtist(uint256 seriesID) internal view {
        uint256 artistID = addressToArtistID[msg.sender];
        require(artistID != 0 && isArtistIDInSeries[seriesID][artistID], "Caller not a series artist");
    }

    function _checkMetadataAndTokenCID(string memory metadata, string memory tokenIDsMapCID) internal pure {
        require(bytes(metadata).length > 0, "Empty metadata CID");
        require(bytes(tokenIDsMapCID).length > 0, "Empty token map CID");
    }

    function _processArtistAddress(address artistAddr) internal {
        require(artistAddr != address(0), "Zero address not allowed");
        uint256 artistID = addressToArtistID[artistAddr];
        if (artistID == 0) {
            // Assign a new artistID
            artistID = nextArtistID++;
            artistIDToAddress[artistID] = artistAddr;
            addressToArtistID[artistAddr] = artistID;
        }
    }

    function _processArtistAddresses(address[] memory artists) internal {
        for (uint256 i = 0; i < artists.length; i++) {
            _processArtistAddress(artists[i]);
        }
    }

    function _addArtistsToSeries(uint256 seriesID, address[] memory artistAddrs) internal returns (uint256[] memory) {
        // Ensure each address has an artistID
        _processArtistAddresses(artistAddrs);

        // Convert addresses to artistIDs
        uint256[] memory artistIDs = new uint256[](artistAddrs.length);
        for (uint256 i = 0; i < artistAddrs.length; i++) {
            uint256 artistID = addressToArtistID[artistAddrs[i]];
            artistIDs[i] = artistID;

            seriesDetails[seriesID].artistIDs.push(artistID);
            isArtistIDInSeries[seriesID][artistID] = true;
            artistIDSeriesIDs[artistID].push(seriesID);
        }

        return artistIDs;
    }

    function _createSeries(
        address[] memory artistAddrs,
        string memory metadata,
        string memory tokenIDsMapCID
    ) internal returns (uint256) {
        uint256 seriesID = nextSeriesID++;
        _checkMetadataAndTokenCID(metadata, tokenIDsMapCID);

        // Mark series as existing
        seriesDetails[seriesID].metadata = metadata;
        seriesDetails[seriesID].contractTokenData = tokenIDsMapCID;

        uint256[] memory artistIDs = _addArtistsToSeries(seriesID, artistAddrs);

        emit SeriesIndexed(seriesID, artistIDs, metadata, tokenIDsMapCID);
        return seriesID;
    }

    function _removeSeriesFromArtist(uint256 artistID, uint256 seriesID) internal {
        uint256[] storage seriesIDs = artistIDSeriesIDs[artistID];
        uint256 length = seriesIDs.length;
        for (uint256 i = 0; i < length; i++) {
            if (seriesIDs[i] == seriesID) {
                seriesIDs[i] = seriesIDs[length - 1];
                seriesIDs.pop();
                break;
            }
        }
        isArtistIDInSeries[seriesID][artistID] = false;
    }

    function _removeArtistFromSeries(uint256 seriesID, uint256 artistID) internal {
        uint256[] storage artistIDs = seriesDetails[seriesID].artistIDs;
        uint256 length = artistIDs.length;
        // Remove artistID from series' artist list
        for (uint256 i = 0; i < length; i++) {
            if (artistIDs[i] == artistID) {
                artistIDs[i] = artistIDs[length - 1];
                artistIDs.pop();
                break;
            }
        }

        // Remove series from artist's series list
        _removeSeriesFromArtist(artistID, seriesID);

        emit SeriesUpdated(seriesID, artistIDs, seriesDetails[seriesID].metadata, seriesDetails[seriesID].contractTokenData);
    }

    // ------------------------
    // External and Public Functions
    // ------------------------

    function addSeries(
        string memory metadata,
        string memory tokenIDsMapCID
    ) external returns (uint256) {
        address[] memory artistAddrs = new address[](1);
        artistAddrs[0] = msg.sender;
        return _createSeries(artistAddrs, metadata, tokenIDsMapCID);
    }

    function ownerAddSeries(
        address[] memory artistAddrs,
        string memory metadata,
        string memory tokenIDsMapCID
    ) external onlyOwner returns (uint256) {
        require(artistAddrs.length > 0, "No artists");
        _checkArtistsNotRevoked(artistAddrs);
        return _createSeries(artistAddrs, metadata, tokenIDsMapCID);
    }

    function ownerBatchAddSeries(
        address[][] memory artistsArray,
        string[] memory metadatas,
        string[] memory tokenIDsMapCIDs
    ) external onlyOwner returns (uint256[] memory) {
        uint256 length = artistsArray.length;
        require(
            length == metadatas.length &&
            length == tokenIDsMapCIDs.length,
            "Array length mismatch"
        );
        require(length > 0, "No series data");
        require(length <= 50, "Batch size too large"); // Prevent DOS attacks

        uint256[] memory createdSeriesIDs = new uint256[](length);
        for (uint256 i = 0; i < length; i++) {
            require(artistsArray[i].length > 0, "No artists for a series");
            _checkArtistsNotRevoked(artistsArray[i]);
            createdSeriesIDs[i] = _createSeries(artistsArray[i], metadatas[i], tokenIDsMapCIDs[i]);
        }
        return createdSeriesIDs;
    }

    function updateSeries(
        uint256 seriesID,
        string memory metadata,
        string memory tokenIDsMapCID
    ) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        _checkMetadataAndTokenCID(metadata, tokenIDsMapCID);

        seriesDetails[seriesID].metadata = metadata;
        seriesDetails[seriesID].contractTokenData = tokenIDsMapCID;

        emit SeriesUpdated(seriesID, seriesDetails[seriesID].artistIDs, metadata, tokenIDsMapCID);
    }

    function deleteSeries(uint256 seriesID) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        // Remove series from all artist's series lists
        uint256[] storage artistIDs = seriesDetails[seriesID].artistIDs;
        uint256 length = artistIDs.length;
        for (uint256 i = 0; i < length; i++) {
            _removeSeriesFromArtist(artistIDs[i], seriesID);
        }

        // Delete series details
        delete seriesDetails[seriesID];

        // Emit event
        emit SeriesDeleted(seriesID);
    }

    function removeSelfFromSeries(uint256 seriesID) external onlyArtist(seriesID) {
        uint256 artistID = addressToArtistID[msg.sender];
        _removeArtistFromSeries(seriesID, artistID);
    }

    function ownerUpdateSeriesArtists(
        uint256 seriesID,
        address[] memory artistAddrs
    ) external seriesExists(seriesID) onlyOwner {
        require(ownerCanModifySeries(seriesID), "Owner rights revoked by all artists");
        _checkArtistsNotRevoked(artistAddrs);

        // Clean up current artists
        uint256[] storage currentArtistIDs = seriesDetails[seriesID].artistIDs;
        for (uint256 i = 0; i < currentArtistIDs.length; i++) {
            uint256 artistID = currentArtistIDs[i];
            isArtistIDInSeries[seriesID][artistID] = false;
            _removeSeriesFromArtist(artistID, seriesID);
        }
        delete seriesDetails[seriesID].artistIDs;

        // Add new artists
        _addArtistsToSeries(seriesID, artistAddrs);

        // Emit event using storage values for consistency
        emit SeriesUpdated(
            seriesID,
            seriesDetails[seriesID].artistIDs,
            seriesDetails[seriesID].metadata,
            seriesDetails[seriesID].contractTokenData
        );
    }

    // Co-Artist Management Functions
    function proposeCoArtist(uint256 seriesID, address proposedArtistAddr) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        require(proposedArtistAddr != address(0), "Invalid address");
        
        // Ensure proposed artist has an ID
        _processArtistAddress(proposedArtistAddr);

        uint256 proposedArtistID = addressToArtistID[proposedArtistAddr];

        require(!seriesPendingCoArtist[seriesID][proposedArtistID], "Already proposed");

        seriesPendingCoArtist[seriesID][proposedArtistID] = true;
        _addPendingRequest(proposedArtistID, seriesID);

        emit CoArtistProposed(seriesID, proposedArtistID);
    }

    function cancelProposeCoArtist(uint256 seriesID, address proposedArtistAddr) external seriesExists(seriesID) onlyOwnerOrArtist(seriesID) {
        uint256 proposedArtistID = addressToArtistID[proposedArtistAddr];
        require(seriesPendingCoArtist[seriesID][proposedArtistID], "No proposal exists");

        seriesPendingCoArtist[seriesID][proposedArtistID] = false;
        _removePendingRequest(proposedArtistID, seriesID);

        uint256 artistID = addressToArtistID[msg.sender];
        emit CoArtistProposalCancelled(seriesID, artistID, proposedArtistID);
    }

    function confirmAsCoArtist(uint256 seriesID) external seriesExists(seriesID) {
        uint256 artistID = addressToArtistID[msg.sender];
        require(artistID != 0, "Not an artist");
        require(seriesPendingCoArtist[seriesID][artistID], "No pending proposal");

        seriesDetails[seriesID].artistIDs.push(artistID);
        isArtistIDInSeries[seriesID][artistID] = true;
        artistIDSeriesIDs[artistID].push(seriesID);

        seriesPendingCoArtist[seriesID][artistID] = false;
        _removePendingRequest(artistID, seriesID);

        emit CoArtistConfirmed(seriesID, artistID);
    }

    // ------------------------
    // Pending Requests Management
    // ------------------------
    function _addPendingRequest(uint256 artistID, uint256 seriesID) internal {
        artistPendingCoArtistRequests[artistID].push(seriesID);
        artistPendingCoArtistRequestIndex[artistID][seriesID] = artistPendingCoArtistRequests[artistID].length - 1;
    }

    function _removePendingRequest(uint256 artistID, uint256 seriesID) internal {
        uint256 index = artistPendingCoArtistRequestIndex[artistID][seriesID];
        uint256 last = artistPendingCoArtistRequests[artistID][artistPendingCoArtistRequests[artistID].length - 1];

        artistPendingCoArtistRequests[artistID][index] = last;
        artistPendingCoArtistRequestIndex[artistID][last] = index;

        artistPendingCoArtistRequests[artistID].pop();
        delete artistPendingCoArtistRequestIndex[artistID][seriesID];
    }

    // Owner Rights Management
    function revokeOwnerRightsForArtist() external {
        uint256 artistID = addressToArtistID[msg.sender];
        require(artistID != 0, "Not an artist");
        require(!ownerRightsRevokedForArtistID[artistID], "Already revoked");
        ownerRightsRevokedForArtistID[artistID] = true;
    }

    function approveOwnerRightsForArtist() external {
        uint256 artistID = addressToArtistID[msg.sender];
        require(artistID != 0, "Not an artist");
        require(ownerRightsRevokedForArtistID[artistID], "Not revoked");
        ownerRightsRevokedForArtistID[artistID] = false;
    }

    function updateArtistAddress(uint256 artistID, address newAddress) external {
        require(newAddress != address(0), "Invalid new address");
        require(addressToArtistID[newAddress] == 0, "Address already assigned to an artist");

        address oldAddress = artistIDToAddress[artistID];
        require(oldAddress != address(0), "Invalid artistID");

        uint256 callerID = addressToArtistID[msg.sender];
        bool isCallerArtist = (callerID == artistID && callerID != 0);
        bool isOwnerWithRights = (msg.sender == owner() && !ownerRightsRevokedForArtistID[artistID]);

        require(isCallerArtist || isOwnerWithRights, "Not authorized to update address");

        addressToArtistID[oldAddress] = 0;
        artistIDToAddress[artistID] = newAddress;
        addressToArtistID[newAddress] = artistID;

        emit ArtistAddressUpdated(artistID, oldAddress, newAddress);
    }

    // ------------------------
    // View Functions
    // ------------------------

    function _checkArtistsNotRevoked(address[] memory artistAddrs) internal view {
        for (uint256 i = 0; i < artistAddrs.length; i++) {
            uint256 artistID = addressToArtistID[artistAddrs[i]];
            // If artistID is 0, this address has never been assigned an artistID and thus cannot have revoked.
            // If artistID != 0, check if revoked.
            if (artistID != 0 && ownerRightsRevokedForArtistID[artistID]) {
                revert("One of the artists revoked owner rights");
            }
        }
    }

    function ownerCanModifySeries(uint256 seriesID) public view returns (bool) {
        uint256[] memory artistIDs = seriesDetails[seriesID].artistIDs;
        for (uint256 i = 0; i < artistIDs.length; i++) {
            if (!ownerRightsRevokedForArtistID[artistIDs[i]]) {
                return true;
            }
        }
        return false;
    }

    function ownerRightsRevoked(address artistAddr) external view returns (bool) {
        uint256 artistID = addressToArtistID[artistAddr];
        return ownerRightsRevokedForArtistID[artistID];
    }

    function getSeriesArtistIDs(uint256 seriesID) external view returns (uint256[] memory) {
        return seriesDetails[seriesID].artistIDs;
    }

    function getSeriesArtistAddresses(uint256 seriesID) external view returns (address[] memory) {
        uint256[] memory artistIDs = seriesDetails[seriesID].artistIDs;
        address[] memory artistAddrs = new address[](artistIDs.length);
        for (uint256 i = 0; i < artistIDs.length; i++) {
            artistAddrs[i] = artistIDToAddress[artistIDs[i]];
        }
        return artistAddrs;
    }

    function getArtistSeriesIDs(address artistAddr) external view returns (uint256[] memory) {
        uint256 artistID = addressToArtistID[artistAddr];
        return artistIDSeriesIDs[artistID];
    }

    function getSeriesMetadata(uint256 seriesID) external view returns (string memory) {
        return seriesDetails[seriesID].metadata;
    }

    function getSeriesContractTokenData(uint256 seriesID) external view returns (string memory) {
        return seriesDetails[seriesID].contractTokenData;
    }

    function getPendingRequests(address artistAddr) external view returns (uint256[] memory) {
        uint256 artistID = addressToArtistID[artistAddr];
        return artistPendingCoArtistRequests[artistID];
    }

    function getArtistAddress(uint256 artistID) external view returns (address) {
        return artistIDToAddress[artistID];
    }

    function getAddressArtistID(address artistAddr) external view returns (uint256) {
        return addressToArtistID[artistAddr];
    }
}