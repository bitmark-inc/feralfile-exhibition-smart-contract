// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/Base64.sol";

import "./Authorizable.sol";
import "./Decentraland.sol";
import "./FeralfileArtworkV3.sol";

contract FeralfileExhibitionV31 is FeralfileExhibitionV3 {
    using Strings for uint256;

    mapping(uint256 => string) public externalArtworkIPFSCID; // artworkID => string (IPFS CID)

    constructor(
        string memory name_,
        string memory symbol_,
        uint256 secondarySaleRoyaltyBPS_,
        address royaltyPayoutAddress_,
        string memory contractURI_,
        string memory tokenBaseURI_,
        bool isBurnable_,
        bool isBridgeable_
    )
        FeralfileExhibitionV3(
            name_,
            symbol_,
            secondarySaleRoyaltyBPS_,
            royaltyPayoutAddress_,
            contractURI_,
            tokenBaseURI_,
            isBurnable_,
            isBridgeable_
        )
    {}

    address public decentralandAddress =
        0xF87E31492Faf9A91B02Ee0dEAAd50d51d56D5d4d;
    uint256 public decentralandTokenID =
        115792089237316195423570985008687907826047395311965486962387615413371672723439;

    function updateDecentralandInfo(address contractAddress, uint256 tokenId)
        external
        onlyOwner
    {
        decentralandAddress = contractAddress;
        decentralandTokenID = tokenId;
    }

    /// @notice createArtworkWithIPFSCID use for create list of artworks in a transaction
    /// @param artwork - the array of artwork
    function createArtworkWithIPFSCID(
        Artwork memory artwork,
        string memory artworkIPFSCID
    ) external onlyAuthorized {
        require(
            bytes(artworkIPFSCID).length != 0,
            "artwor IPFS CID can not be empty"
        );

        uint256 artworkID = _createArtwork(
            artwork.fingerprint,
            artwork.title,
            artwork.artistName,
            artwork.editionSize,
            artwork.AEAmount,
            artwork.PPAmount
        );

        externalArtworkIPFSCID[artworkID] = artworkIPFSCID;
    }

    /// @notice buildArtworkData returns an object of artwork which would push to the actually artwork
    /// @param artworkID - the artwork ID for building artwork data
    function buildArtworkData(uint256 artworkID)
        private
        view
        returns (string memory)
    {
        Decentraland dc = Decentraland(decentralandAddress);
        uint256[] memory editionIDs = allArtworkEditions[artworkID];
        bytes memory ownersArray = bytes("[");

        for (uint256 i = 0; i < editionIDs.length; i++) {
            ownersArray = abi.encodePacked(
                ownersArray,
                '"',
                Strings.toHexString(ownerOf(editionIDs[i])),
                '",'
            );
        }

        ownersArray = abi.encodePacked(ownersArray, "]");

        return
            string(
                abi.encodePacked(
                    "{"
                    'landOwner:"',
                    Strings.toHexString(dc.ownerOf(decentralandTokenID)),
                    '", ownerArray:',
                    string(ownersArray),
                    "}"
                )
            );
    }

    /// @notice buildIframe returns a base64 encoded data for ff-frame
    /// @param artworkData - the artwork data which would bring into the artwork
    /// @param iframeURI - the artwork URL to loaded into iframe
    function buildIframe(string memory artworkData, string memory iframeURI)
        private
        pure
        returns (string memory)
    {
        return
            Base64.encode(
                abi.encodePacked(
                    '<!DOCTYPE html><html lang="en"><head><script> var defaultArtworkData= ',
                    artworkData,
                    "</script><script>"
                    'let allowOrigins={"https://feralfile.com":!0};function resizeIframe(t){let e=document.getElementById("mainframe");t&&(e.style.minHeight=t+"px",e.style.height=t+"px")}'
                    'function initData(){pushArtworkDataToIframe(defaultArtworkData)}function pushArtworkDataToIframe(t){t&&document.getElementById("mainframe").contentWindow.postMessage(t,"*")}'
                    'function updateArtowrkData(t){document.getElementById("mainframe").contentWindow.postMessage(t,"*")}window.addEventListener("message",function(t){allowOrigins[t.origin]?'
                    '"update-artwork-data"===t.data.type&&updateArtowrkData(t.data.artworkData):"object"==typeof t.data&&"resize-iframe"===t.data.type&&resizeIframe(t.data.newHeight)});</script>'
                    '</head><body style="overflow-x: hidden; padding: 0; margin: 0; width: 100%;" onload="initData()">'
                    '<iframe id="mainframe" style="display:block; padding: 0; margin: 0; border:none; width: 100%;" src="',
                    iframeURI,
                    '"></iframe> </body></html>'
                )
            );
    }

    /// @notice buildDataURL returns a base64 encoded data for ff-frame
    /// @param artworkID - the artwork ID for building artwork data
    /// @param ipfsCID - the artwork ipfs CID for this specific artwork
    function buildDataURL(uint256 artworkID, string memory ipfsCID)
        private
        view
        returns (string memory)
    {
        Artwork memory artwork = artworks[artworkID];
        return
            string(
                abi.encodePacked(
                    "data:application/json;base64,",
                    Base64.encode(
                        abi.encodePacked(
                            '{"name":"',
                            artwork.title,
                            '", "artwork_id":"',
                            artworkID.toString(),
                            '", "image":"',
                            '", "animation_url":"data:text/html;base64,',
                            buildIframe(
                                buildArtworkData(artworkID),
                                buildIPFSURL(ipfsCID)
                            ),
                            '"}'
                        )
                    )
                )
            );
    }

    /// @notice buildIPFSURL returns a formatted IPFS link based on the _tokenBaseURI
    /// @param ipfsCID - thei IPFS Cid
    function buildIPFSURL(string memory ipfsCID)
        private
        view
        returns (string memory)
    {
        string memory baseURI = _tokenBaseURI;
        if (bytes(baseURI).length == 0) {
            baseURI = "ipfs://";
        }

        return string(abi.encodePacked(baseURI, ipfsCID));
    }

    /// @notice A distinct Uniform Resource Identifier (URI) for a given asset.
    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override
        returns (string memory)
    {
        require(
            _exists(tokenId),
            "ERC721Metadata: URI query for nonexistent token"
        );

        ArtworkEditionIndex
            memory artworkEditionIndex = allArtworkEditionsIndex[tokenId];

        uint256 artworkID = artworkEditionIndex.artworkID;
        string memory artworkIPFSCID = externalArtworkIPFSCID[artworkID];

        if (bytes(artworkIPFSCID).length > 0) {
            return buildDataURL(artworkID, artworkIPFSCID);
        } else {
            return buildIPFSURL(artworkEditions[tokenId].ipfsCID);
        }
    }
}
