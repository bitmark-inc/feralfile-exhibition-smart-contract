// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title SeriesRegistry
 * @dev A registry system for managing digital art series and their creators.
 * This contract handles artist-series relationships, collaborative series,
 * delegation permissions, and comprehensive registry operations.
 */
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

contract SeriesRegistry {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableSet for EnumerableSet.UintSet;

    // ------------------------------------------------------------------------
    // DATA STRUCTURES
    // ------------------------------------------------------------------------

    /**
     * @dev The maximum batch size for write operations
     */
    uint256 public constant MAX_BATCH_SIZE = 50;

    /**
     * @dev Each Series stores:
     * - administratorID: The ID of the administrator of the series
     * - metadataURI: URI hosting the series metadata
     * - tokenDataURI: URI hosting the series token data
     * - artistIDs: Set of artist IDs associated with this series
     */
    struct Series {
        uint256 administratorID;
        EnumerableSet.UintSet artistIDs;
        string metadataURI;
        string tokenDataURI;
    }

    /**
     * @dev A counter for the next series ID
     *      This is used to generate a unique ID for each series
     */
    uint256 private nextSeriesID = 1;

    /**
     * @dev A counter for the next artist ID
     *      This is used to generate a unique ID for each artist
     */
    uint256 private nextArtistID = 1;

    /**
     * @dev A mapping of artist addresses to their ID
     */
    mapping(address => uint256) private artistAddressIDMap;

    /**
     * @dev A mapping of artist IDs to their address
     */
    mapping(uint256 => address) private artistIDAddressMap;

    /**
     * @dev A mapping of artist ID to a set of series IDs
     */
    mapping(uint256 => EnumerableSet.UintSet) private artistSeriesMap;

    /**
     * @dev A mapping of delegator address to a set of delegatee addresses
     */
    mapping(address => EnumerableSet.AddressSet) private addressDelegateeMap;

    /**
     * @dev A mapping of series IDs to a set of delegatee addresses
     */
    mapping(uint256 => EnumerableSet.AddressSet) private seriesDelegateeMap;

    /**
     * @dev A mapping of series IDs to series data
     */
    mapping(uint256 => Series) private seriesRegistry;

    /**
     * @dev A set of all series IDs
     */
    EnumerableSet.UintSet private seriesIDsRegistry;

    /**
     * @dev A mapping of series IDs to a mapping of pending collaboration invitee to their inviter
     */
    mapping(uint256 => mapping(address => address))
        private seriesCollaborationInviteeInviterMap;

    /**
     * @dev A mapping of series IDs to a set of collaboration invitees
     */
    mapping(uint256 => EnumerableSet.AddressSet)
        private seriesCollaborationInviteesMap;

    /**
     * @dev A mapping of collaboration invitee address to a list of series IDs
     */
    mapping(address => EnumerableSet.UintSet)
        private collaborationInviteeSeriesMap;

    /**
     * @dev A mapping of inviter address to a set of series IDs they've sent invitations for
     */
    mapping(address => EnumerableSet.UintSet) private inviterSeriesMap;

    // ------------------------------------------------------------------------
    // EVENTS
    // ------------------------------------------------------------------------

    event UpdateArtistAddress(
        address indexed oldAddress,
        address indexed newAddress
    );
    event AssignSeries(
        uint256 indexed seriesID,
        address indexed assignerAddress,
        address indexed assigneeAddress
    );
    event AddDelegatee(address indexed delegator, address indexed delegatee);
    event RemoveDelegatee(address indexed delegator, address indexed delegatee);

    event RegisterSeries(uint256 indexed seriesID);
    event UpdateSeries(uint256 indexed seriesID);
    event DeleteSeries(uint256 indexed seriesID);
    event OptOutSeries(uint256 indexed seriesID, address indexed artistAddress);

    event InviteCollaborator(
        uint256 indexed seriesID,
        address indexed inviterAddress,
        address indexed inviteeAddress
    );
    event OptInCollaboration(
        uint256 indexed seriesID,
        address indexed collaboratorAddress
    );
    event CancelCollaborationInvitation(
        uint256 indexed seriesID,
        address indexed inviteeAddress
    );
    event AssignAdministrator(
        uint256 indexed seriesID,
        address indexed assignerAddress,
        address indexed assigneeAddress
    );

    // ------------------------------------------------------------------------
    // ERRORS
    // ------------------------------------------------------------------------

    error SeriesNotExistsError(uint256 _seriesID);
    error NotSeriesArtistError(uint256 _seriesID, address _address);
    error NotArtistError(address _address);
    error ArtistExistsError(address _address);
    error NotAdministratorError(uint256 _seriesID, address _address);
    error AdministratorCannotSelfAssignError(address _address);
    error NotAuthorizedError(address _address);
    error GenericError(string _message);
    error InvitationNotExistsError(uint256 _seriesID, address _inviteeAddress);
    error NotPendingInvitationError(uint256 _seriesID, address _inviteeAddress);
    error DuplicateInvitationError(uint256 _seriesID, address _inviteeAddress);
    error DuplicateDelegateeError(address _delegator, address _delegatee);
    error NotDelegateeError(address _delegator, address _delegatee);
    error NotCollaborationInviterError(address _address);
    error NotCollaborativeSeriesError(uint256 _seriesID);
    error SameAddressError(address _address);

    // ------------------------------------------------------------------------
    // MODIFIERS
    // ------------------------------------------------------------------------

    /**
     * @dev Modifier to ensure the caller is either a series delegatee or an artist
     * @param _seriesID The ID of the series to check
     */
    modifier onlyArtistOrDelegatee(uint256 _seriesID) {
        if (
            !_isCallerArtist(_seriesID) && !_isCallerSeriesDelegatee(_seriesID)
        ) {
            revert NotAuthorizedError(msg.sender);
        }
        _;
    }

    /**
     * @dev Modifier to ensure the caller is an artist
     * @param _seriesID The ID of the series to check
     */
    modifier onlyArtist(uint256 _seriesID) {
        if (!_isCallerArtist(_seriesID)) {
            revert NotArtistError(msg.sender);
        }
        _;
    }

    /**
     * @dev Modifier to ensure the caller is a series administrator
     * @param _seriesID The ID of the series to check
     */
    modifier onlyAdministrator(uint256 _seriesID) {
        if (!_isCallerAdministrator(_seriesID)) {
            revert NotAdministratorError(_seriesID, msg.sender);
        }
        _;
    }

    /**
     * @dev Modifier to ensure the caller is a series collaboration inviter
     * @param _seriesID The ID of the series to check
     * @param _inviteeAddress The address of the invitee to check
     */
    modifier onlyCollaborationInviter(
        uint256 _seriesID,
        address _inviteeAddress
    ) {
        if (_inviteeAddress == address(0)) {
            revert GenericError("invitee address is zero");
        }
        if (!_isCallerCollaborationInviter(_seriesID, _inviteeAddress)) {
            revert NotCollaborationInviterError(msg.sender);
        }
        _;
    }

    /**
     * @notice Ensures the series exists
     * @param _seriesID The ID of the series to check
     */
    modifier seriesExists(uint256 _seriesID) {
        if (bytes(seriesRegistry[_seriesID].metadataURI).length == 0) {
            revert SeriesNotExistsError(_seriesID);
        }
        _;
    }

    // ------------------------------------------------------------------------
    // REGISTRY FUNCTIONS
    // ------------------------------------------------------------------------

    /**
     * @notice Registers a new series
     * @param _artistAddress The address of the artist associated with the series
     * @param _metadataURI The metadata URI of the series
     * @param _tokenDataURI The token data URI of the series
     * @return The ID of the newly created series
     */
    function registerSeries(
        address _artistAddress,
        string calldata _metadataURI,
        string calldata _tokenDataURI
    ) external returns (uint256) {
        return _registerSeries(_artistAddress, _metadataURI, _tokenDataURI);
    }

    /**
     * @notice Batch adds new series
     * @param _seriesArtists The addresses of the artists associated with the series
     * @param _metadataURIs The metadata URIs of the series
     * @param _tokenDataURIs The token data URIs of the series
     * @return The IDs of the newly created series
     */
    function batchRegisterSeries(
        address[] calldata _seriesArtists,
        string[] calldata _metadataURIs,
        string[] calldata _tokenDataURIs
    ) external returns (uint256[] memory) {
        uint256 seriesLength = _seriesArtists.length;
        _validateSeriesWriteBatch(seriesLength, _metadataURIs, _tokenDataURIs);

        uint256[] memory seriesIDs = new uint256[](seriesLength);
        for (uint256 i = 0; i < seriesLength; ) {
            seriesIDs[i] = _registerSeries(
                _seriesArtists[i],
                _metadataURIs[i],
                _tokenDataURIs[i]
            );
            unchecked {
                i++;
            }
        }
        return seriesIDs;
    }

    /**
     * @notice Updates an existing series' metadata and token data
     * @param _seriesID The ID of the series to update
     * @param _metadataURI The metadata URI of the series
     * @param _tokenDataURI The token data URI of the series
     */
    function updateSeries(
        uint256 _seriesID,
        string calldata _metadataURI,
        string calldata _tokenDataURI
    ) external {
        _updateSeries(_seriesID, _metadataURI, _tokenDataURI);
    }

    /**
     * @notice Batch updates existing series' metadata and token data
     * @param _seriesIDs The IDs of the series to update
     * @param _metadataURIs The metadata URIs of the series
     * @param _tokenDataURIs The token data URIs of the series
     */
    function batchUpdateSeries(
        uint256[] calldata _seriesIDs,
        string[] calldata _metadataURIs,
        string[] calldata _tokenDataURIs
    ) external {
        uint256 seriesLength = _seriesIDs.length;
        _validateSeriesWriteBatch(seriesLength, _metadataURIs, _tokenDataURIs);
        for (uint256 i = 0; i < seriesLength; ) {
            _updateSeries(_seriesIDs[i], _metadataURIs[i], _tokenDataURIs[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Deletes a series
     * @param _seriesID The ID of the series to delete
     */
    function deleteSeries(uint256 _seriesID) external {
        _deleteSeries(_seriesID);
    }

    /**
     * @notice Batch deletes series
     * @param _seriesIDs The IDs of the series to delete
     */
    function batchDeleteSeries(uint256[] calldata _seriesIDs) external {
        uint256 seriesLength = _seriesIDs.length;
        for (uint256 i = 0; i < seriesLength; ) {
            _deleteSeries(_seriesIDs[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Allows an artist to opt out of a series
     * @param _seriesID The ID of the series to opt out of
     */
    function optOutSeries(
        uint256 _seriesID
    ) external seriesExists(_seriesID) onlyArtist(_seriesID) {
        // Validate caller is an artist
        address caller = msg.sender;
        uint256 artistID = artistAddressIDMap[caller];
        if (artistID == 0 || !artistSeriesMap[artistID].contains(_seriesID)) {
            revert NotArtistError(caller);
        }

        // Validate series is collaborative
        if (seriesRegistry[_seriesID].artistIDs.length() <= 1) {
            revert NotCollaborativeSeriesError(_seriesID);
        }

        // Remove artist from series
        _removeArtist(_seriesID, artistID);

        // Emit event
        emit OptOutSeries(_seriesID, caller);
    }

    /**
     * @notice Assigns a series to another artist
     * @param _seriesID The ID of the series to assign
     * @param _assigneeAddress The address of the artist to assign the series to
     */
    function assignSeries(
        uint256 _seriesID,
        address _assigneeAddress
    ) external {
        _assignSeries(_seriesID, _assigneeAddress);
    }

    /**
     * @notice Batch assigns series to other artists
     * @param _seriesIDs The IDs of the series to assign
     * @param _assigneeAddresses The addresses of the artists to assign the series to
     */
    function batchAssignSeries(
        uint256[] calldata _seriesIDs,
        address[] calldata _assigneeAddresses
    ) external {
        uint256 seriesIDsLength = _seriesIDs.length;
        _validateAssignmentWriteBatch(seriesIDsLength, _assigneeAddresses);

        for (uint256 i = 0; i < seriesIDsLength; ) {
            _assignSeries(_seriesIDs[i], _assigneeAddresses[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Invites a collaborator to a series
     * @param _seriesID The ID of the series to invite a collaborator to
     * @param _inviteeAddress The address of the invitee to invite
     */
    function inviteCollaborator(
        uint256 _seriesID,
        address _inviteeAddress
    ) external {
        _inviteCollaborator(_seriesID, msg.sender, _inviteeAddress);
    }

    /**
     * @notice Batch invites collaborators to series
     * @param _seriesIDs The IDs of the series to invite collaborators to
     * @param _collaboratorAddresses The addresses of the collaborators to invite
     */
    function batchInviteCollaborators(
        uint256[] calldata _seriesIDs,
        address[] calldata _collaboratorAddresses
    ) external {
        uint256 seriesIDsLength = _seriesIDs.length;
        _validateCollaborationInvitationWriteBatch(
            seriesIDsLength,
            _collaboratorAddresses
        );

        for (uint256 i = 0; i < seriesIDsLength; ) {
            _inviteCollaborator(
                _seriesIDs[i],
                msg.sender,
                _collaboratorAddresses[i]
            );
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Cancels a collaboration invitation
     * @param _seriesID The ID of the series to cancel the invitation for
     * @param _inviteeAddress The address of the invitee to cancel the invitation for
     */
    function cancelCollaborationInvitation(
        uint256 _seriesID,
        address _inviteeAddress
    ) external {
        _cancelCollaborationInvitation(_seriesID, _inviteeAddress);
    }

    /**
     * @notice Batch cancels collaboration invitations
     * @param _seriesIDs The IDs of the series to cancel the invitations for
     * @param _inviteeAddresses The addresses of the invitees to cancel the invitations for
     */
    function batchCancelCollaborationInvitations(
        uint256[] calldata _seriesIDs,
        address[] calldata _inviteeAddresses
    ) external {
        uint256 seriesIDsLength = _seriesIDs.length;
        _validateCollaborationInvitationWriteBatch(
            seriesIDsLength,
            _inviteeAddresses
        );

        for (uint256 i = 0; i < seriesIDsLength; ) {
            _cancelCollaborationInvitation(_seriesIDs[i], _inviteeAddresses[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Opts in to a series
     * @param _seriesID The ID of the series to opt in to
     */
    function optInCollaboration(uint256 _seriesID) external {
        _optInCollaboration(_seriesID);
    }

    /**
     * @notice Batch opts in to series
     * @param _seriesIDs The IDs of the series to opt in to
     */
    function batchOptInCollaboration(uint256[] calldata _seriesIDs) external {
        uint256 seriesIDsLength = _seriesIDs.length;
        if (seriesIDsLength == 0) {
            revert GenericError("seriesCount is zero");
        }
        if (seriesIDsLength > MAX_BATCH_SIZE) {
            revert GenericError("seriesCount is too large");
        }

        for (uint256 i = 0; i < seriesIDsLength; ) {
            _optInCollaboration(_seriesIDs[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Updates an artist's address
     * @param _newAddress The new address of the artist
     */
    function updateArtistAddress(address _newAddress) external {
        // Check if the new address is already assigned
        // If it is, we need to transfer all series from the old address to the new address
        // Otherwise, it will be a normal address change
        if (artistAddressIDMap[_newAddress] != 0) {
            _assignAllSeries(_newAddress);
        } else {
            _updateArtistAddress(_newAddress);
        }
    }

    /**
     * @notice Adds a delegatee to the caller's delegatee list
     * @param _delegatee The address of the delegatee to add
     */
    function addDelegatee(address _delegatee) external {
        _addDelegatee(_delegatee);
    }

    /**
     * @notice Batch adds delegatees to the caller's delegatee list
     * @param _delegatees The addresses of the delegatees to add
     */
    function batchAddDelegatees(address[] calldata _delegatees) external {
        _validateDelegateeWriteBatch(_delegatees);

        uint256 delegateesLength = _delegatees.length;
        for (uint256 i = 0; i < delegateesLength; ) {
            _addDelegatee(_delegatees[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Removes a delegatee from the caller's delegatee list
     * @param _delegatee The address of the delegatee to remove
     */
    function removeDelegatee(address _delegatee) external {
        _removeDelegatee(_delegatee);
    }

    /**
     * @notice Batch removes delegatees from the caller's delegatee list
     * @param _delegatees The addresses of the delegatees to remove
     */
    function batchRemoveDelegatees(address[] calldata _delegatees) external {
        _validateDelegateeWriteBatch(_delegatees);

        uint256 delegateesLength = _delegatees.length;
        for (uint256 i = 0; i < delegateesLength; ) {
            _removeDelegatee(_delegatees[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Assigns a new administrator to a series
     * @param _seriesID The ID of the series to assign the administrator to
     * @param _assigneeAddress The address of the new administrator
     */
    function assignAdministrator(
        uint256 _seriesID,
        address _assigneeAddress
    ) external seriesExists(_seriesID) onlyAdministrator(_seriesID) {
        // Check if the assignee address is valid
        if (_assigneeAddress == address(0)) {
            revert GenericError("administrator address is zero");
        }

        // Check if the assignee is not the caller
        address caller = msg.sender;
        if (caller == _assigneeAddress) {
            revert AdministratorCannotSelfAssignError(_assigneeAddress);
        }

        // Check if the assignee is an artist of the series
        uint256 assigneeArtistID = artistAddressIDMap[_assigneeAddress];
        if (
            assigneeArtistID == 0 ||
            !seriesRegistry[_seriesID].artistIDs.contains(assigneeArtistID)
        ) {
            revert NotSeriesArtistError(_seriesID, _assigneeAddress);
        }

        // Assign the administrator
        seriesRegistry[_seriesID].administratorID = assigneeArtistID;

        // Emit event
        emit AssignAdministrator(_seriesID, caller, _assigneeAddress);
    }

    // ------------------------------------------------------------------------
    // VIEW FUNCTIONS
    // ------------------------------------------------------------------------

    /**
     * @notice Returns all series IDs
     * @return An array of series IDs
     */
    function getSeriesIDs() external view returns (uint256[] memory) {
        return seriesIDsRegistry.values();
    }

    /**
     * @notice Returns the total number of series
     * @return The total number of series
     */
    function getTotalSeries() external view returns (uint256) {
        return seriesIDsRegistry.length();
    }

    /**
     * @notice Returns the administrator ID for a given series
     * @param _seriesID The ID of the series to get the administrator ID for
     * @return The administrator ID of the series
     */
    function getSeriesAdministratorID(
        uint256 _seriesID
    ) external view returns (uint256) {
        return seriesRegistry[_seriesID].administratorID;
    }

    /**
     * @notice Returns the administrator address for a given series
     * @param _seriesID The ID of the series to get the administrator address for
     * @return The administrator address of the series
     */
    function getSeriesAdministratorAddress(
        uint256 _seriesID
    ) external view returns (address) {
        return artistIDAddressMap[seriesRegistry[_seriesID].administratorID];
    }

    /**
     * @notice Returns all artist IDs for a given series
     * @param _seriesID The ID of the series to get the artist IDs for
     * @return An array of artist IDs associated with the series
     */
    function getSeriesArtistIDs(
        uint256 _seriesID
    ) external view returns (uint256[] memory) {
        return seriesRegistry[_seriesID].artistIDs.values();
    }

    /**
     * @notice Returns all artist addresses for a given series
     * @param _seriesID The ID of the series to get the artist addresses for
     * @return An array of artist addresses associated with the series
     */
    function getSeriesArtistAddresses(
        uint256 _seriesID
    ) external view returns (address[] memory) {
        uint256[] memory artistIDs = seriesRegistry[_seriesID]
            .artistIDs
            .values();
        address[] memory artistAddresses = new address[](artistIDs.length);
        for (uint256 i = 0; i < artistIDs.length; ) {
            artistAddresses[i] = artistIDAddressMap[artistIDs[i]];
            unchecked {
                i++;
            }
        }
        return artistAddresses;
    }

    /**
     * @notice Returns the metadata URI of a given series
     * @param _seriesID The ID of the series to get the metadata URI for
     * @return The metadata URI of the series
     */
    function getSeriesMetadataURI(
        uint256 _seriesID
    ) external view returns (string memory) {
        return seriesRegistry[_seriesID].metadataURI;
    }

    /**
     * @notice Returns the token data URI of a given series
     * @param _seriesID The ID of the series to get the token data URI for
     * @return The token data URI of the series
     */
    function getSeriesTokenDataURI(
        uint256 _seriesID
    ) external view returns (string memory) {
        return seriesRegistry[_seriesID].tokenDataURI;
    }

    /**
     * @notice Returns all series IDs associated with an artist address
     * @param _artistAddress The address of the artist to get the series IDs for
     * @return An array of series IDs associated with the artist
     */
    function getArtistSeriesIDs(
        address _artistAddress
    ) external view returns (uint256[] memory) {
        uint256 artistID = artistAddressIDMap[_artistAddress];
        return artistSeriesMap[artistID].values();
    }

    /**
     * @notice Returns the total number of series associated with an artist
     * @param _artistAddress The address of the artist to get the total number of series for
     * @return The total number of series associated with the artist
     */
    function getTotalArtistSeries(
        address _artistAddress
    ) external view returns (uint256) {
        uint256 artistID = artistAddressIDMap[_artistAddress];
        return artistSeriesMap[artistID].length();
    }

    /**
     * @notice Returns the address associated with a given artist ID
     * @param _artistID The ID of the artist to get the address for
     * @return The address of the artist
     */
    function getArtistAddress(
        uint256 _artistID
    ) external view returns (address) {
        return artistIDAddressMap[_artistID];
    }

    /**
     * @notice Returns the artist ID associated with an address
     * @param _artistAddress The address of the artist to get the ID for
     * @return The ID of the artist
     */
    function getArtistID(
        address _artistAddress
    ) external view returns (uint256) {
        return artistAddressIDMap[_artistAddress];
    }

    /**
     * @notice Returns all collaboration invitees for a given series
     * @param _seriesID The ID of the series to get the collaboration invitees for
     * @return An array of addresses associated with the series' collaboration invitees
     */
    function getCollaborationInvitees(
        uint256 _seriesID
    ) external view returns (address[] memory) {
        return seriesCollaborationInviteesMap[_seriesID].values();
    }

    /**
     * @notice Returns all series IDs associated with an invitee address
     * @param _invitee The address of the invitee to get the series IDs for
     * @return An array of series IDs associated with the invitee
     */
    function getCollaborationInviteeSeriesIDs(
        address _invitee
    ) external view returns (uint256[] memory) {
        return collaborationInviteeSeriesMap[_invitee].values();
    }

    /**
     * @notice Returns all series IDs associated with an inviter address
     * @param _inviter The address of the inviter to get the series IDs for
     * @return An array of series IDs associated with the inviter
     */
    function getCollaborationInviterSeriesIDs(
        address _inviter
    ) external view returns (uint256[] memory) {
        return inviterSeriesMap[_inviter].values();
    }

    /**
     * @notice Returns all delegatees for a given delegator address
     * @param _delegatorAddress The address of the delegator to get the delegatees for
     * @return An array of delegatees associated with the delegator
     */
    function getDelegatees(
        address _delegatorAddress
    ) external view returns (address[] memory) {
        return addressDelegateeMap[_delegatorAddress].values();
    }

    /**
     * @notice Returns all delegatees for a given series
     * @param _seriesID The ID of the series to get the delegatees for
     * @return An array of delegatees associated with the series
     */
    function getSeriesDelegatees(
        uint256 _seriesID
    ) external view returns (address[] memory) {
        return seriesDelegateeMap[_seriesID].values();
    }

    // ------------------------------------------------------------------------
    // INTERNAL FUNCTIONS
    // ------------------------------------------------------------------------

    /**
     * @notice Creates a new series with the specified artists
     * @param _artistAddress The address of the artist associated with the series
     * @param _metadataURI The metadata URI of the series
     * @param _tokenDataURI The token data URI of the series
     * @return The ID of the newly created series
     */
    function _registerSeries(
        address _artistAddress,
        string memory _metadataURI,
        string memory _tokenDataURI
    ) internal returns (uint256) {
        // Validate parameters
        if (_artistAddress == address(0)) {
            revert GenericError("artist address is zero");
        }
        if (bytes(_metadataURI).length == 0) {
            revert GenericError("empty metadataURI");
        }
        if (bytes(_tokenDataURI).length == 0) {
            revert GenericError("empty tokenDataURI");
        }

        // Check if the caller is authorized to register the series
        if (!_isCaller(_artistAddress) && !_isCallerDelegatee(_artistAddress)) {
            revert NotAuthorizedError(msg.sender);
        }

        // Register the series
        uint256 seriesID = nextSeriesID++;
        Series storage series = seriesRegistry[seriesID];
        series.metadataURI = _metadataURI;
        series.tokenDataURI = _tokenDataURI;
        seriesIDsRegistry.add(seriesID);

        // Add artists
        uint256 artistID = _addArtist(seriesID, _artistAddress);

        // Set the series administrator as the artist
        series.administratorID = artistID;

        // Add existing delegatees to series delegatee mapping
        if (addressDelegateeMap[_artistAddress].length() > 0) {
            uint256 delegateeLength = addressDelegateeMap[_artistAddress]
                .length();
            for (uint256 i = 0; i < delegateeLength; ) {
                seriesDelegateeMap[seriesID].add(
                    addressDelegateeMap[_artistAddress].at(i)
                );
                unchecked {
                    i++;
                }
            }
        }

        // Emit event
        emit RegisterSeries(seriesID);

        return seriesID;
    }

    /**
     * @notice Updates an existing series' metadata and token data
     * @param _seriesID The ID of the series to update
     * @param _metadataURI The metadata URI of the series
     * @param _tokenDataURI The token data URI of the series
     */
    function _updateSeries(
        uint256 _seriesID,
        string memory _metadataURI,
        string memory _tokenDataURI
    ) internal seriesExists(_seriesID) onlyArtistOrDelegatee(_seriesID) {
        // Validate parameters
        if (bytes(_metadataURI).length == 0) {
            revert GenericError("empty metadataURI");
        }
        if (bytes(_tokenDataURI).length == 0) {
            revert GenericError("empty tokenDataURI");
        }

        // Update series
        Series storage series = seriesRegistry[_seriesID];
        series.metadataURI = _metadataURI;
        series.tokenDataURI = _tokenDataURI;

        // Emit event
        emit UpdateSeries(_seriesID);
    }

    /**
     * @notice Deletes a series and cleans up related data
     * @param _seriesID The ID of the series to delete
     */
    function _deleteSeries(
        uint256 _seriesID
    ) internal seriesExists(_seriesID) onlyAdministrator(_seriesID) {
        // Clean up pending collaboration invitations
        _removeSeriesCollaborationInvitations(_seriesID);

        // Clean up series delegate
        delete seriesDelegateeMap[_seriesID];

        // Delete series
        uint256[] memory artistIDs = seriesRegistry[_seriesID]
            .artistIDs
            .values();
        _batchRemoveArtists(_seriesID, artistIDs);
        delete seriesRegistry[_seriesID];
        seriesIDsRegistry.remove(_seriesID);

        // Emit event
        emit DeleteSeries(_seriesID);
    }

    /**
     * @notice Validates batch parameters for series write operations
     * @param _seriesCount The number of series
     * @param _metadataURIs The metadata URIs of the series
     * @param _tokenDataURIs The token data URIs of the series
     */
    function _validateSeriesWriteBatch(
        uint256 _seriesCount,
        string[] memory _metadataURIs,
        string[] memory _tokenDataURIs
    ) internal pure {
        if (_seriesCount == 0) {
            revert GenericError("seriesCount is zero");
        }
        if (_seriesCount > MAX_BATCH_SIZE) {
            revert GenericError("seriesCount is too large");
        }
        if (
            _seriesCount != _metadataURIs.length ||
            _seriesCount != _tokenDataURIs.length
        ) {
            revert GenericError("arrayLength mismatch");
        }
    }

    /**
     * @notice Validates batch parameters for collaboration invitation write operations
     * @param _seriesCount The number of series
     * @param _inviteeAddresses The addresses of the invitees
     */
    function _validateCollaborationInvitationWriteBatch(
        uint256 _seriesCount,
        address[] calldata _inviteeAddresses
    ) internal pure {
        if (_seriesCount == 0) {
            revert GenericError("seriesCount is zero");
        }
        if (_seriesCount > MAX_BATCH_SIZE) {
            revert GenericError("seriesCount is too large");
        }
        if (_inviteeAddresses.length != _seriesCount) {
            revert GenericError("arrayLength mismatch");
        }
    }

    /**
     * @notice Validates batch parameters for delegatee write operations
     * @param _delegatees The addresses of the delegatees
     */
    function _validateDelegateeWriteBatch(
        address[] calldata _delegatees
    ) internal pure {
        if (_delegatees.length == 0) {
            revert GenericError("delegateeCount is zero");
        }
        if (_delegatees.length > MAX_BATCH_SIZE) {
            revert GenericError("delegateeCount is too large");
        }
    }

    /**
     * @notice Validates batch parameters for assignment write operations
     * @param _seriesCount The number of series
     * @param _assigneeAddresses The addresses of the artists to assign the series to
     */
    function _validateAssignmentWriteBatch(
        uint256 _seriesCount,
        address[] calldata _assigneeAddresses
    ) internal pure {
        if (_seriesCount == 0) {
            revert GenericError("seriesCount is zero");
        }
        if (_seriesCount > MAX_BATCH_SIZE) {
            revert GenericError("seriesCount is too large");
        }
        if (_assigneeAddresses.length != _seriesCount) {
            revert GenericError("arrayLength mismatch");
        }
    }

    /**
     * @notice Checks if the caller is the same as the address
     * @param addr The address to check
     * @return `true` if the caller is the same as the address, `false` otherwise
     */
    function _isCaller(address addr) internal view returns (bool) {
        return msg.sender == addr;
    }

    /**
     * @notice Checks if the caller is a series artist
     * @param _seriesID The ID of the series to check
     */
    function _isCallerArtist(uint256 _seriesID) internal view returns (bool) {
        uint256 artistID = artistAddressIDMap[msg.sender];
        return artistID != 0 && artistSeriesMap[artistID].contains(_seriesID);
    }

    /**
     * @notice Checks if the caller is a series administrator
     * @param _seriesID The ID of the series to check
     */
    function _isCallerAdministrator(
        uint256 _seriesID
    ) internal view returns (bool) {
        uint256 callerID = artistAddressIDMap[msg.sender];
        if (callerID == 0) {
            return false;
        }

        return seriesRegistry[_seriesID].administratorID == callerID;
    }

    /**
     * @notice Checks if the caller is a series collaboration inviter
     * @param _seriesID The ID of the series to check
     * @param _inviteeAddress The address of the invitee to check
     */
    function _isCallerCollaborationInviter(
        uint256 _seriesID,
        address _inviteeAddress
    ) internal view returns (bool) {
        return
            seriesCollaborationInviteeInviterMap[_seriesID][_inviteeAddress] ==
            msg.sender;
    }

    /**
     * @notice Checks if the caller is a series delegate
     * @param _seriesID The ID of the series to check
     */
    function _isCallerSeriesDelegatee(
        uint256 _seriesID
    ) internal view returns (bool) {
        return seriesDelegateeMap[_seriesID].contains(msg.sender);
    }

    /**
     * @notice Checks if the caller is a delegatee
     * @param _delegatorAddress The address of the delegator to check
     */
    function _isCallerDelegatee(
        address _delegatorAddress
    ) internal view returns (bool) {
        return addressDelegateeMap[_delegatorAddress].contains(msg.sender);
    }

    /**
     * @notice Removes artists from a series
     * @param _seriesID The ID of the series to remove the artists from
     * @param _artistIDs The IDs of the artists to remove from the series
     */
    function _batchRemoveArtists(
        uint256 _seriesID,
        uint256[] memory _artistIDs
    ) internal {
        uint256 artistLength = _artistIDs.length;
        for (uint256 i = 0; i < artistLength; ) {
            _removeArtist(_seriesID, _artistIDs[i]);
            unchecked {
                i++;
            }
        }
    }

    /**
     * @notice Adds an artist to a series
     * @param _seriesID The ID of the series to add the artist to
     * @param _artistAddress The address of the artist to add to the series
     * @return The ID of the artist
     */
    function _addArtist(
        uint256 _seriesID,
        address _artistAddress
    ) internal returns (uint256) {
        uint256 artistID = artistAddressIDMap[_artistAddress];
        if (artistID == 0) {
            artistID = nextArtistID++;
            artistAddressIDMap[_artistAddress] = artistID;
            artistIDAddressMap[artistID] = _artistAddress;
        }

        artistSeriesMap[artistID].add(_seriesID);
        seriesRegistry[_seriesID].artistIDs.add(artistID);
        return artistID;
    }

    /**
     * @notice Removes an artist from a series
     * @param _seriesID The ID of the series to remove the artist from
     * @param _artistID The ID of the artist to remove
     */
    function _removeArtist(uint256 _seriesID, uint256 _artistID) internal {
        if (_artistID == 0) {
            revert GenericError("artist not found");
        }
        address artistAddress = artistIDAddressMap[_artistID];

        EnumerableSet.UintSet storage artistIDs = seriesRegistry[_seriesID]
            .artistIDs;

        // Remove the artist from the series
        artistIDs.remove(_artistID);
        artistSeriesMap[_artistID].remove(_seriesID);

        // Assign the first artist as the series administrator if the removed artist is the administrator
        if (seriesRegistry[_seriesID].administratorID == _artistID) {
            if (artistIDs.length() > 0) {
                seriesRegistry[_seriesID].administratorID = artistIDs.at(0);
            } else {
                delete seriesRegistry[_seriesID].administratorID;
            }
        }

        // To avoid the orphaned artist, we delete the artist address ID map
        // entry if the artist is not associated with any series
        if (artistSeriesMap[_artistID].length() == 0) {
            delete artistAddressIDMap[artistAddress];
            delete artistIDAddressMap[_artistID];
        }

        // Clean up the collaboration invitation sent by the artist for this series
        _removeSeriesCollaborationInvitationsByInviter(
            _seriesID,
            artistAddress
        );
    }

    /**
     * @notice Removes all collaboration invitations from a series
     * @param _seriesID The ID of the series to remove the collaboration invitations from
     */
    function _removeSeriesCollaborationInvitations(uint256 _seriesID) internal {
        EnumerableSet.AddressSet
            storage collaborationInvitees = seriesCollaborationInviteesMap[
                _seriesID
            ];
        uint256 collaborationInviteesLength = collaborationInvitees.length();
        for (uint256 i = 0; i < collaborationInviteesLength; ) {
            address inviteeAddress = collaborationInvitees.at(i);
            address inviterAddress = seriesCollaborationInviteeInviterMap[
                _seriesID
            ][inviteeAddress];

            collaborationInviteeSeriesMap[inviteeAddress].remove(_seriesID);
            delete seriesCollaborationInviteeInviterMap[_seriesID][
                inviteeAddress
            ];
            inviterSeriesMap[inviterAddress].remove(_seriesID);

            unchecked {
                i++;
            }
        }

        delete seriesCollaborationInviteesMap[_seriesID];
    }

    /**
     * @notice Invites a collaborator to a series
     * @param _seriesID The ID of the series to invite the collaborator to
     * @param _inviterAddress The address of the inviter
     * @param _inviteeAddress The address of the invitee to invite
     */
    function _inviteCollaborator(
        uint256 _seriesID,
        address _inviterAddress,
        address _inviteeAddress
    ) internal seriesExists(_seriesID) onlyArtistOrDelegatee(_seriesID) {
        // Validate invitee address is valid
        if (_inviteeAddress == address(0)) {
            revert GenericError("invitee address is zero");
        }

        // Validate the existence of invitation
        if (
            seriesCollaborationInviteesMap[_seriesID].contains(_inviteeAddress)
        ) {
            revert DuplicateInvitationError(_seriesID, _inviteeAddress);
        }

        // Validate series has the same artist address
        uint256 artistID = artistAddressIDMap[_inviteeAddress];
        if (artistID != 0 && artistSeriesMap[artistID].contains(_seriesID)) {
            revert ArtistExistsError(_inviteeAddress);
        }

        // Add invitation
        collaborationInviteeSeriesMap[_inviteeAddress].add(_seriesID);
        seriesCollaborationInviteesMap[_seriesID].add(_inviteeAddress);
        seriesCollaborationInviteeInviterMap[_seriesID][
            _inviteeAddress
        ] = _inviterAddress;
        inviterSeriesMap[_inviterAddress].add(_seriesID);

        // Emit event
        emit InviteCollaborator(_seriesID, _inviterAddress, _inviteeAddress);
    }

    /**
     * @notice Cancels a collaboration invitation to a series
     * @param _seriesID The ID of the series to cancel the collaboration invitation from
     * @param _inviteeAddress The address of the invitee to cancel the invitation for
     */
    function _cancelCollaborationInvitation(
        uint256 _seriesID,
        address _inviteeAddress
    )
        internal
        seriesExists(_seriesID)
        onlyCollaborationInviter(_seriesID, _inviteeAddress)
    {
        // Remove invitation
        _removeSeriesCollaborationInvitationByInvitee(
            _seriesID,
            _inviteeAddress
        );

        // Emit event
        emit CancelCollaborationInvitation(_seriesID, _inviteeAddress);
    }

    /**
     * @notice Opts in to a series
     * @param _seriesID The ID of the series to opt in to
     */
    function _optInCollaboration(
        uint256 _seriesID
    ) internal seriesExists(_seriesID) {
        address sender = msg.sender;

        // Validate the existence of invitation
        _ensureCollaborationInvitationExistence(_seriesID, sender);

        // Add artist to series
        _addArtist(_seriesID, sender);

        // Remove pending collaboration invitations
        _removeSeriesCollaborationInvitationByInvitee(_seriesID, sender);

        // Emit event
        emit OptInCollaboration(_seriesID, sender);
    }

    /**
     * @notice Removes a collaboration invitation from a series
     * @param _seriesID The ID of the series to remove the collaboration invitation from
     * @param _inviteeAddress The address of the invitee to remove the collaboration invitation for
     */
    function _removeSeriesCollaborationInvitationByInvitee(
        uint256 _seriesID,
        address _inviteeAddress
    ) internal {
        address inviterAddress = seriesCollaborationInviteeInviterMap[
            _seriesID
        ][_inviteeAddress];

        collaborationInviteeSeriesMap[_inviteeAddress].remove(_seriesID);
        seriesCollaborationInviteesMap[_seriesID].remove(_inviteeAddress);
        delete seriesCollaborationInviteeInviterMap[_seriesID][_inviteeAddress];

        // If this is the last invitation for the inviter, remove the inviter
        // from the series set
        if (
            seriesCollaborationInviteesMap[_seriesID].length() == 0 ||
            !_hasMoreInvitationsFromInviter(_seriesID, inviterAddress)
        ) {
            inviterSeriesMap[inviterAddress].remove(_seriesID);
        }
    }

    /**
     * @notice Checks if the inviter has more invitations for this series
     * @param _seriesID The ID of the series to check
     * @param _inviterAddress The address of the inviter to check
     * @return `true` if the inviter has more invitations for this series, `false` otherwise
     */
    function _hasMoreInvitationsFromInviter(
        uint256 _seriesID,
        address _inviterAddress
    ) internal view returns (bool) {
        address[] memory invitees = seriesCollaborationInviteesMap[_seriesID]
            .values();
        for (uint256 i = 0; i < invitees.length; ) {
            if (
                seriesCollaborationInviteeInviterMap[_seriesID][invitees[i]] ==
                _inviterAddress
            ) {
                return true;
            }

            unchecked {
                i++;
            }
        }
        return false;
    }

    /**
     * @notice Swaps the inviter of a series
     * @param _seriesID The ID of the series to swap the inviter of
     * @param _oldInviterAddress The address of the old inviter
     * @param _newInviterAddress The address of the new inviter
     */
    function _swapSeriesCollaborationInviter(
        uint256 _seriesID,
        address _oldInviterAddress,
        address _newInviterAddress
    ) internal {
        // Get all invitees for this series
        address[] memory invitees = seriesCollaborationInviteesMap[_seriesID]
            .values();

        // Swap the inviter of each invitee
        for (uint256 i = 0; i < invitees.length; ) {
            address invitee = invitees[i];
            if (
                seriesCollaborationInviteeInviterMap[_seriesID][invitee] ==
                _oldInviterAddress
            ) {
                seriesCollaborationInviteeInviterMap[_seriesID][
                    invitee
                ] = _newInviterAddress;
            }

            unchecked {
                i++;
            }
        }

        // Remove the old inviter from the series set
        inviterSeriesMap[_oldInviterAddress].remove(_seriesID);

        // Add the new inviter to the series set
        inviterSeriesMap[_newInviterAddress].add(_seriesID);
    }

    /**
     * @notice Removes all collaboration invitations for an inviter from a series
     * @param _seriesID The ID of the series to remove the collaboration invitations for
     * @param _inviterAddress The address of the inviter to remove the collaboration invitations for
     */
    function _removeSeriesCollaborationInvitationsByInviter(
        uint256 _seriesID,
        address _inviterAddress
    ) internal {
        // Get all invitees for this series
        address[] memory invitees = seriesCollaborationInviteesMap[_seriesID]
            .values();

        for (uint256 j = invitees.length; j > 0; ) {
            address invitee = invitees[j - 1];

            // If the inviter sent this invitation
            if (
                seriesCollaborationInviteeInviterMap[_seriesID][invitee] ==
                _inviterAddress
            ) {
                // Remove invitation
                collaborationInviteeSeriesMap[invitee].remove(_seriesID);
                seriesCollaborationInviteesMap[_seriesID].remove(invitee);
                delete seriesCollaborationInviteeInviterMap[_seriesID][invitee];
            }

            unchecked {
                j--;
            }
        }

        // Remove the inviter from the series set
        inviterSeriesMap[_inviterAddress].remove(_seriesID);
    }

    /**
     * @notice Ensures the collaboration invitation exists
     * @param _seriesID The ID of the series to check
     * @param _inviteeAddress The address of the invitee to check
     */
    function _ensureCollaborationInvitationExistence(
        uint256 _seriesID,
        address _inviteeAddress
    ) internal view {
        // Validate invitee address is valid
        if (_inviteeAddress == address(0)) {
            revert GenericError("invitee address is zero");
        }

        // Validate the existence of invitation
        if (
            !seriesCollaborationInviteesMap[_seriesID].contains(_inviteeAddress)
        ) {
            revert NotPendingInvitationError(_seriesID, _inviteeAddress);
        }
    }

    /**
     * @notice Adds a delegatee to the caller's delegatee list
     * @param _delegatee The address of the delegatee to add
     */
    function _addDelegatee(address _delegatee) internal {
        // Validate delegatee address is valid
        if (_delegatee == address(0)) {
            revert GenericError("delegatee address is zero");
        }

        // Add delegatee
        address delegator = msg.sender;
        bool added = addressDelegateeMap[delegator].add(_delegatee);
        if (!added) {
            revert DuplicateDelegateeError(delegator, _delegatee);
        }

        // Consider to add delegatee to series delegatee mappings
        uint256 artistID = artistAddressIDMap[delegator];
        if (artistID != 0) {
            uint256 seriesLength = artistSeriesMap[artistID].length();
            for (uint256 i = 0; i < seriesLength; ) {
                uint256 seriesID = artistSeriesMap[artistID].at(i);
                seriesDelegateeMap[seriesID].add(_delegatee);
                unchecked {
                    i++;
                }
            }
        }

        // Emit event
        emit AddDelegatee(delegator, _delegatee);
    }

    /**
     * @notice Removes a delegatee from the caller's delegatee list
     * @param _delegatee The address of the delegatee to remove
     */
    function _removeDelegatee(address _delegatee) internal {
        // Validate delegatee address is valid
        if (_delegatee == address(0)) {
            revert GenericError("delegatee address is zero");
        }

        // Remove delegatee
        address delegator = msg.sender;
        bool removed = addressDelegateeMap[delegator].remove(_delegatee);
        if (!removed) {
            revert NotDelegateeError(delegator, _delegatee);
        }

        // Consider to remove delegatee from series delegatee mappings
        uint256 artistID = artistAddressIDMap[delegator];
        if (artistID != 0) {
            uint256 seriesLength = artistSeriesMap[artistID].length();
            for (uint256 i = 0; i < seriesLength; ) {
                uint256 seriesID = artistSeriesMap[artistID].at(i);

                // Remove delegatee from series delegatee mapping
                seriesDelegateeMap[seriesID].remove(_delegatee);

                // Remove series collaboration invitations for the delegatee
                _removeSeriesCollaborationInvitationsByInviter(
                    seriesID,
                    _delegatee
                );

                unchecked {
                    i++;
                }
            }
        }

        // Emit event
        emit RemoveDelegatee(delegator, _delegatee);
    }

    /**
     * @notice Updates the artist address
     * @param _newAddress The new address of the artist
     */
    function _updateArtistAddress(address _newAddress) internal {
        // Validate the existence of the artist
        address artistAddress = msg.sender;
        uint256 artistID = artistAddressIDMap[artistAddress];
        if (artistID == 0) {
            revert NotArtistError(artistAddress);
        }

        // Validate the new address is not the same as the current address
        if (artistAddress == _newAddress) {
            revert SameAddressError(artistAddress);
        }

        // Validate the new address is not zero
        if (_newAddress == address(0)) {
            revert GenericError("new address is zero");
        }

        // Update artist address ID mapping
        artistAddressIDMap[artistAddress] = 0;
        artistAddressIDMap[_newAddress] = artistID;

        // Update artist ID address mapping
        artistIDAddressMap[artistID] = _newAddress;

        // Swap delegatees and clean up pending collaboration invitations
        uint256[] memory seriesIDs = artistSeriesMap[artistID].values();
        for (uint256 i = 0; i < seriesIDs.length; ) {
            uint256 seriesID = seriesIDs[i];

            // Swap delegatees
            _swapSeriesDelegatees(seriesID, artistAddress, _newAddress);

            // Swap collaboration inviter
            _swapSeriesCollaborationInviter(
                seriesID,
                artistAddress,
                _newAddress
            );

            // Remove the invitation sent by delegatees
            address[] memory delegatees = addressDelegateeMap[artistAddress]
                .values();

            for (uint256 j = 0; j < delegatees.length; ) {
                address delegatee = delegatees[j];

                // If the delegatee is not an inviter of the series, skip
                if (!inviterSeriesMap[delegatee].contains(seriesID)) {
                    unchecked {
                        j++;
                    }
                    continue;
                }

                // Remove the invitation sent by the delegatee
                _removeSeriesCollaborationInvitationsByInviter(
                    seriesID,
                    delegatee
                );

                unchecked {
                    j++;
                }
            }

            unchecked {
                i++;
            }
        }

        // Emit event
        emit UpdateArtistAddress(artistAddress, _newAddress);
    }

    /**
     * @notice Assigns all series to another address
     * @param _assigneeAddress The address of the artist to assign the series to
     */
    function _assignAllSeries(address _assigneeAddress) internal {
        // Validate the correctness of the assignee address
        if (_assigneeAddress == address(0)) {
            revert GenericError("assignee address is zero");
        }

        // Validate the existence of the assigner
        address assignerAddress = msg.sender;
        uint256 assignerID = artistAddressIDMap[assignerAddress];
        if (assignerID == 0) {
            revert NotArtistError(assignerAddress);
        }

        // Get all series IDs from the assigner
        EnumerableSet.UintSet storage assignerSeries = artistSeriesMap[
            assignerID
        ];
        for (uint256 i = assignerSeries.length(); i > 0; i--) {
            uint256 seriesID = assignerSeries.at(i - 1);
            _assignSeries(seriesID, _assigneeAddress);
        }
    }

    /**
     * @notice Assigns a series to another artist
     * @param _seriesID The ID of the series to assign
     * @param _assigneeAddress The address of the artist to assign the series to
     */
    function _assignSeries(
        uint256 _seriesID,
        address _assigneeAddress
    ) internal onlyArtist(_seriesID) {
        // Validate the correctness of the assignee address
        if (_assigneeAddress == address(0)) {
            revert GenericError("assignee address is zero");
        }

        // Validate the existence of the assigner
        address assignerAddress = msg.sender;
        uint256 assignerID = artistAddressIDMap[assignerAddress];
        if (assignerID == 0) {
            revert NotArtistError(assignerAddress);
        }

        // Validate the assignee is not the same as the assigner
        if (assignerAddress == _assigneeAddress) {
            revert SameAddressError(assignerAddress);
        }

        // Validate the assignee is not a collaborator
        uint256 assigneeID = artistAddressIDMap[_assigneeAddress];
        if (
            assigneeID != 0 &&
            seriesRegistry[_seriesID].artistIDs.contains(assigneeID)
        ) {
            revert ArtistExistsError(_assigneeAddress);
        }

        // 1. Add assignee to series
        assigneeID = _addArtist(_seriesID, _assigneeAddress);

        // 2. Assign series administrator
        if (seriesRegistry[_seriesID].administratorID == assignerID) {
            seriesRegistry[_seriesID].administratorID = assigneeID;
        }

        // 3. Remove assigner from series
        _removeArtist(_seriesID, assignerID);

        // 4. Swap delegatees
        _swapSeriesDelegatees(_seriesID, assignerAddress, _assigneeAddress);

        // Emit event
        emit AssignSeries(_seriesID, assignerAddress, _assigneeAddress);
    }

    /**
     * @notice Swaps the delegatees of two delegators
     * @param _seriesID The ID of the series to swap the delegatees of
     * @param _oldDelegator The address of the old delegator
     * @param _newDelegator The address of the new delegator
     *
     * @dev This function assumes that all validation has been done.
     * It basically swaps the series delegatees of two delegators.
     */
    function _swapSeriesDelegatees(
        uint256 _seriesID,
        address _oldDelegator,
        address _newDelegator
    ) internal {
        // 1. Remove old delegator's delegatees from series delegatee mapping
        EnumerableSet.AddressSet storage seriesDelegatees = seriesDelegateeMap[
            _seriesID
        ];
        address[] memory oldDelegatees = addressDelegateeMap[_oldDelegator]
            .values();
        for (uint256 i = 0; i < oldDelegatees.length; ) {
            seriesDelegatees.remove(oldDelegatees[i]);
            unchecked {
                i++;
            }
        }

        // 2. Add new delegator's delegatees to series delegatee mapping
        address[] memory newDelegatees = addressDelegateeMap[_newDelegator]
            .values();
        for (uint256 i = 0; i < newDelegatees.length; ) {
            seriesDelegatees.add(newDelegatees[i]);
            unchecked {
                i++;
            }
        }
    }
}
