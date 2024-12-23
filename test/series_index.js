const { expect } = require("chai");
const { BN, expectRevert, expectEvent } = require("@openzeppelin/test-helpers");
const SeriesIndexer = artifacts.require("SeriesIndexer");
// NOTE: We have to build the abi before running the test
const SeriesIndexerArtifact = require("../build/contracts/SeriesIndexer.json");
const { expectCustomError } = require("./helper/expectCustomError.js");
const seriesIndexerABI = SeriesIndexerArtifact.abi;

contract("SeriesIndexer", (accounts) => {
    const owner = accounts[0];
    const artistA = accounts[1];
    const artistB = accounts[2];
    const artistC = accounts[3]; // will be proposed as co-artist
    const artistD = accounts[5]; // for multiple proposals scenario
    const nonArtist = accounts[4]; // never added as an artist
    const firstSeriesID = new BN("1");
    const secondSeriesID = new BN("2");
    const metadataURI = "QmMetadataURI";
    const tokenMapURI = "QmTokenMapURI";
    const newMetadataURI = "QmMetadataURINew";
    const newTokenMapURI = "QmTokenMapURINew";

    let instance;

    beforeEach(async () => {
        instance = await SeriesIndexer.new({ from: owner });
    });

    describe("Initial State", () => {
        it("Owner should be the contract deployer", async () => {
            const contractOwner = await instance.owner();
            expect(contractOwner).to.equal(owner);
        });
    });

    describe("Adding Series", () => {
        it("Non-owner (artist) can add a series with themselves as the first artist", async () => {
            const tx = await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            expectEvent(tx, "SeriesIndexed", {
                seriesID: firstSeriesID,
            });

            const seriesMeta = await instance.getSeriesMetadataURI(firstSeriesID);
            const seriesTokenMap = await instance.getSeriesContractTokenDataURI(
                firstSeriesID
            );

            expect(seriesMeta).to.equal(metadataURI);
            expect(seriesTokenMap).to.equal(tokenMapURI);

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(1);

            const artistAID = await instance.getAddressArtistID(artistA);
            expect(artistAID.toNumber()).to.not.equal(
                0,
                "ArtistA should have a valid artistID"
            );
        });

        it("Non-owner (artist) can't add a series with others as the artist", async () => {
            await expectCustomError(
                instance.addSeries([artistB], metadataURI, tokenMapURI, {
                    from: artistA,
                }),
                seriesIndexerABI,
                "NotAuthorizedError",
                [artistA]
            );

            await expectCustomError(
                instance.addSeries([artistA, artistB], metadataURI, tokenMapURI, {
                    from: artistA,
                }),
                seriesIndexerABI,
                "NotAuthorizedError",
                [artistA]
            );
        });

        it("Owner can add a series with multiple artists", async () => {
            const tx = await instance.addSeries(
                [artistA, artistB],
                metadataURI,
                tokenMapURI,
                { from: owner }
            );

            expectEvent(tx, "SeriesIndexed", {
                seriesID: firstSeriesID,
            });

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(2);
        });

        it("Should revert when metadata URI or token map URI is empty", async () => {
            // Reverts with custom error EmptyMetadataURIError()
            await expectCustomError(
                instance.addSeries([artistA], "", tokenMapURI, { from: artistA }),
                seriesIndexerABI,
                "EmptyMetadataURIError"
            );

            // Reverts with custom error EmptyTokenMapURIError()
            await expectCustomError(
                instance.addSeries([artistA], metadataURI, "", { from: artistA }),
                seriesIndexerABI,
                "EmptyTokenMapURIError"
            );
        });
    });

    describe("Batch Adding Series by Owner", () => {
        const batchArtistsArray = [[artistA], [artistB]];
        const batchMetadataURIs = ["meta3001", "meta3002"];
        const batchTokenMapURIs = ["token3001", "token3002"];

        it("Owner can batch add multiple series", async () => {
            const tx = await instance.batchAddSeries(
                batchArtistsArray,
                batchMetadataURIs,
                batchTokenMapURIs,
                { from: owner }
            );

            expectEvent(tx, "SeriesIndexed", { seriesID: firstSeriesID });
            expectEvent(tx, "SeriesIndexed", { seriesID: secondSeriesID });

            const seriesMeta1 = await instance.getSeriesMetadataURI(firstSeriesID);
            const seriesTokenMap1 = await instance.getSeriesContractTokenDataURI(
                firstSeriesID
            );
            expect(seriesMeta1).to.equal("meta3001");
            expect(seriesTokenMap1).to.equal("token3001");

            const seriesMeta2 = await instance.getSeriesMetadataURI(
                secondSeriesID
            );
            const seriesTokenMap2 = await instance.getSeriesContractTokenDataURI(
                secondSeriesID
            );
            expect(seriesMeta2).to.equal("meta3002");
            expect(seriesTokenMap2).to.equal("token3002");
        });

        it("Non-owner (artist) can batch add multiple series of themselves", async () => {
            const tx = await instance.batchAddSeries(
                [[artistA], [artistA]],
                batchMetadataURIs,
                batchTokenMapURIs,
                { from: owner }
            );

            expectEvent(tx, "SeriesIndexed", { seriesID: firstSeriesID });
            expectEvent(tx, "SeriesIndexed", { seriesID: secondSeriesID });

            const seriesMeta1 = await instance.getSeriesMetadataURI(firstSeriesID);
            const seriesTokenMap1 = await instance.getSeriesContractTokenDataURI(
                firstSeriesID
            );
            expect(seriesMeta1).to.equal("meta3001");
            expect(seriesTokenMap1).to.equal("token3001");

            const seriesMeta2 = await instance.getSeriesMetadataURI(
                secondSeriesID
            );
            const seriesTokenMap2 = await instance.getSeriesContractTokenDataURI(
                secondSeriesID
            );
            expect(seriesMeta2).to.equal("meta3002");
            expect(seriesTokenMap2).to.equal("token3002");
        });

        it("Owner batch add should revert if one of the artists revoked owner rights", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.revokeOwnerRightsForArtist({ from: artistA });

            // Reverts with custom error ArtistRevokedOwnerRightsError(address artistAddr)
            await expectCustomError(
                instance.batchAddSeries(
                    [[artistA]],
                    ["meta4001"],
                    ["token4001"],
                    { from: owner }
                ),
                seriesIndexerABI,
                "ArtistRevokedOwnerRightsError",
                [artistA]
            );
        });

        it("Non-owner (artist) can't batch add if adding others as artists", async () => {
            await expectCustomError(
                instance.batchAddSeries(
                    [[artistA], [artistB]],
                    batchMetadataURIs,
                    batchTokenMapURIs,
                    { from: artistA }
                ),
                seriesIndexerABI,
                "NotAuthorizedError",
                [artistA]
            );
        });

        it("Batch add should revert on array length mismatch", async () => {
            // Reverts with custom error ArrayLengthMismatchError(uint256, uint256, uint256)
            // The first array's length = 1, while the others have length = 2
            await expectCustomError(
                instance.batchAddSeries(
                    [[artistA]],
                    ["meta5001", "meta5002"],
                    ["token5001", "token5002"],
                    { from: owner }
                ),
                seriesIndexerABI,
                "ArrayLengthMismatchError",
                ["1", "2", "2"] // must pass as strings or BN if we want to match decode
            );
        });
    });

    describe("Updating Series", () => {
        it("Series artist can update the series metadata", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.updateSeries(
                firstSeriesID,
                newMetadataURI,
                newTokenMapURI,
                { from: artistA }
            );
            expectEvent(tx, "SeriesUpdated", {
                seriesID: firstSeriesID,
            });

            const updatedMeta = await instance.getSeriesMetadataURI(firstSeriesID);
            const updatedTokenMap = await instance.getSeriesContractTokenDataURI(
                firstSeriesID
            );
            expect(updatedMeta).to.equal(newMetadataURI);
            expect(updatedTokenMap).to.equal(newTokenMapURI);
        });

        it("Non-owner (artist) can batch update multiple series of themselves", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });

            const tx = await instance.batchUpdateSeries(
                [firstSeriesID, secondSeriesID],
                ["meta5001", "meta5002"],
                ["token5001", "token5002"],
                { from: artistA }
            );
            expectEvent(tx, "SeriesUpdated", {
                seriesID: firstSeriesID,
            });
            expectEvent(tx, "SeriesUpdated", {
                seriesID: firstSeriesID,
            });

            let updatedMeta = await instance.getSeriesMetadataURI(firstSeriesID);
            let updatedTokenMap = await instance.getSeriesContractTokenDataURI(
                firstSeriesID
            );
            expect(updatedMeta).to.equal("meta5001");
            expect(updatedTokenMap).to.equal("token5001");
            updatedMeta = await instance.getSeriesMetadataURI(secondSeriesID);
            updatedTokenMap = await instance.getSeriesContractTokenDataURI(
                secondSeriesID
            );
            expect(updatedMeta).to.equal("meta5002");
            expect(updatedTokenMap).to.equal("token5002");
        });

        it("Non-artist or non-owner cannot update", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Reverts with custom error: CallerNotASeriesArtistError(uint256 seriesID, address caller)
            await expectCustomError(
                instance.updateSeries(firstSeriesID, metadataURI, tokenMapURI, {
                    from: nonArtist,
                }),
                seriesIndexerABI,
                "CallerNotASeriesArtistError",
                [firstSeriesID.toString(), nonArtist] // (seriesID, caller)
            );
        });

        it("Updating a non-existent series should revert", async () => {
            // Reverts with custom error: SeriesDoesNotExistError(uint256 seriesID)
            await expectCustomError(
                instance.updateSeries(
                    new BN("9999"),
                    metadataURI,
                    tokenMapURI,
                    { from: artistA }
                ),
                seriesIndexerABI,
                "SeriesDoesNotExistError",
                ["9999"] // pass as a string
            );
        });
    });

    describe("Owner Rights Revocation", () => {
        it("Artist can revoke owner rights", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            const revoked = await instance.ownerRightsRevoked(artistA);
            expect(revoked).to.be.true;
        });

        it("Artist can re-approve owner rights", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            await instance.approveOwnerRightsForArtist({ from: artistA });
            const revoked = await instance.ownerRightsRevoked(artistA);
            expect(revoked).to.be.false;
        });

        it("When all artists revoke, owner cannot update series", async () => {
            await instance.addSeries(
                [artistA, artistB],
                metadataURI,
                tokenMapURI,
                { from: owner }
            );
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistB });

            // Reverts with custom error: OwnerRightsRevokedForThisSeries(uint256 seriesID)
            await expectCustomError(
                instance.updateSeries(firstSeriesID, "x", "y", { from: owner }),
                seriesIndexerABI,
                "OwnerRightsRevokedForThisSeries",
                [firstSeriesID.toString()]
            );

            await instance.approveOwnerRightsForArtist({ from: artistA });
            const tx = await instance.updateSeries(
                firstSeriesID,
                "metaNew",
                "tokenNew",
                { from: owner }
            );
            expectEvent(tx, "SeriesUpdated", {
                seriesID: firstSeriesID,
            });
        });

        it("Non-artist tries to revoke/approve owner rights should revert", async () => {
            // Reverts with custom error: NotAnArtistError(address caller)
            await expectCustomError(
                instance.revokeOwnerRightsForArtist({ from: nonArtist }),
                seriesIndexerABI,
                "NotAnArtistError",
                [nonArtist]
            );

            // Reverts with custom error: NotAnArtistError(address caller)
            await expectCustomError(
                instance.approveOwnerRightsForArtist({ from: nonArtist }),
                seriesIndexerABI,
                "NotAnArtistError",
                [nonArtist]
            );
        });
    });

    describe("Propose and Confirm Co-Artist", () => {
        it("Artist in a series can propose a new co-artist", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.proposeCoArtist(firstSeriesID, artistC, {
                from: artistA,
            });
            expectEvent(tx, "CoArtistProposed", {
                seriesID: firstSeriesID,
            });

            const pending = await instance.getPendingRequests(artistC);
            expect(pending.map((e) => e.toString())).to.include(
                firstSeriesID.toString()
            );
        });

        it("Co-artist can confirm themselves", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.proposeCoArtist(firstSeriesID, artistC, {
                from: artistA,
            });
            let artistURI = await instance.getAddressArtistID(artistC);

            const tx = await instance.confirmAsCoArtist(firstSeriesID, {
                from: artistC,
            });
            expectEvent(tx, "CoArtistConfirmed", {
                seriesID: firstSeriesID,
                confirmedArtistID: artistURI,
            });

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.map((id) => id.toString())).to.include(
                artistURI.toString()
            );

            const pendingAfter = await instance.getPendingRequests(artistC);
            expect(pendingAfter.length).to.equal(
                0,
                "Pending requests should be cleared after confirmation"
            );
        });

        it("Cannot confirm if not proposed", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.addSeries([artistC], metadataURI, tokenMapURI, {
                from: artistC,
            });
            // Reverts with custom error: NotAPendingProposalError(uint256 seriesID, address caller)
            await expectCustomError(
                instance.confirmAsCoArtist(firstSeriesID, { from: artistC }),
                seriesIndexerABI,
                "NotAPendingProposalError",
                [firstSeriesID.toString(), artistC]
            );
        });

        it("An artist can cancel a co-artist proposal before confirmation", async () => {
            await instance.addSeries(
                [artistA, artistB],
                metadataURI,
                tokenMapURI,
                { from: owner }
            );
            await instance.proposeCoArtist(firstSeriesID, artistC, {
                from: artistA,
            });
            let artistURI = await instance.getAddressArtistID(artistC);

            let pendingBefore = await instance.getPendingRequests(artistC);
            expect(pendingBefore.map((e) => e.toString())).to.include(
                firstSeriesID.toString()
            );

            const tx = await instance.cancelProposeCoArtist(
                firstSeriesID,
                artistC,
                { from: artistB }
            );
            expectEvent(tx, "CoArtistProposalCancelled", {
                seriesID: firstSeriesID,
                proposerArtistID: await instance.getAddressArtistID(artistB),
                cancelledArtistID: artistURI,
            });

            // Now the proposal is canceled, confirming should revert
            // Reverts with custom error: NotAPendingProposalError(...)
            await expectCustomError(
                instance.confirmAsCoArtist(firstSeriesID, { from: artistC }),
                seriesIndexerABI,
                "NotAPendingProposalError",
                [firstSeriesID.toString(), artistC]
            );

            const pendingAfterCancel = await instance.getPendingRequests(
                artistC
            );
            expect(pendingAfterCancel.length).to.equal(
                0,
                "Pending requests should be cleared after cancellation"
            );
        });

        it("Proposing a co-artist for a non-existent series should revert", async () => {
            // Reverts with custom error: SeriesDoesNotExistError(seriesID)
            await expectCustomError(
                instance.proposeCoArtist(new BN("9999"), artistC, {
                    from: artistA,
                }),
                seriesIndexerABI,
                "SeriesDoesNotExistError",
                ["9999"]
            );
        });

        it("Proposing the same co-artist twice for the same series should revert", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Propose artistB
            await instance.proposeCoArtist(firstSeriesID, artistB, {
                from: artistA,
            });
            // Attempt to propose the same co-artist again
            // Reverts with custom error: AlreadyProposedError(uint256 seriesID, address proposedArtistAddr)
            await expectCustomError(
                instance.proposeCoArtist(firstSeriesID, artistB, {
                    from: artistA,
                }),
                seriesIndexerABI,
                "AlreadyProposedError",
                [firstSeriesID.toString(), artistB]
            );
        });

        it("Non-artist cannot propose co-artist", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Reverts with custom error: CallerNotASeriesArtistError(uint256 seriesID, address caller)
            await expectCustomError(
                instance.proposeCoArtist(firstSeriesID, artistB, {
                    from: nonArtist,
                }),
                seriesIndexerABI,
                "CallerNotASeriesArtistError",
                [firstSeriesID.toString(), nonArtist]
            );
        });

        it("Propose co-artist with invalid (zero) address should revert", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Reverts with custom error: ZeroAddressNotAllowedError()
            await expectCustomError(
                instance.proposeCoArtist(
                    firstSeriesID,
                    "0x0000000000000000000000000000000000000000",
                    { from: artistA }
                ),
                seriesIndexerABI,
                "ZeroAddressNotAllowedError"
            );
        });
    });

    describe("Multiple Co-Artist Proposals", () => {
        it("Can propose multiple co-artists and confirm/cancel them independently", async () => {
            await instance.addSeries(
                [artistA, artistB],
                metadataURI,
                tokenMapURI,
                { from: owner }
            );
            // Propose artistD for firstSeriesID
            await instance.proposeCoArtist(firstSeriesID, artistD, {
                from: artistA,
            });
            let artistDID = await instance.getAddressArtistID(artistD);

            // Add another new series
            await instance.addSeries([artistA], "meta1010", "token1010", {
                from: artistA,
            });

            // Propose artistD for secondSeriesID
            await instance.proposeCoArtist(secondSeriesID, artistD, {
                from: artistA,
            });

            // Now artistD should have 2 pending requests
            let pendingD = await instance.getPendingRequests(artistD);
            expect(pendingD.map((e) => e.toString())).to.have.members([
                firstSeriesID.toString(),
                secondSeriesID.toString(),
            ]);

            // Confirm artistD for secondSeriesID first
            await instance.confirmAsCoArtist(secondSeriesID, { from: artistD });

            // Now pending should only have firstSeriesID
            let pendingDAfterOneConfirm = await instance.getPendingRequests(
                artistD
            );
            expect(pendingDAfterOneConfirm.map((e) => e.toString())).to.include(
                firstSeriesID.toString()
            );
            expect(pendingDAfterOneConfirm.length).to.equal(1);

            // Cancel the remaining proposal on firstSeriesID
            const artistAID = await instance.getAddressArtistID(artistA);
            const tx = await instance.cancelProposeCoArtist(
                firstSeriesID,
                artistD,
                { from: artistA }
            );
            expectEvent(tx, "CoArtistProposalCancelled", {
                seriesID: firstSeriesID,
                proposerArtistID: artistAID,
                cancelledArtistID: artistDID,
            });

            // Attempting to confirm the canceled proposal should fail
            // Reverts with custom error: NotAPendingProposalError(...)
            await expectCustomError(
                instance.confirmAsCoArtist(firstSeriesID, { from: artistD }),
                seriesIndexerABI,
                "NotAPendingProposalError",
                [firstSeriesID.toString(), artistD]
            );

            // Pending should now be empty
            let pendingDAfterCancel = await instance.getPendingRequests(
                artistD
            );
            expect(pendingDAfterCancel.length).to.equal(
                0,
                "All pending requests cleared after confirm/cancel operations"
            );
        });
    });

    describe("Updating Artist's address", () => {
        it("Artist can update their own address", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const artistAID = await instance.getAddressArtistID(artistA);
            const newAddressA = accounts[6];

            const tx = await instance.updateArtistAddress(
                artistAID,
                newAddressA,
                { from: artistA }
            );
            expectEvent(tx, "ArtistAddressUpdated", {
                artistID: artistAID,
                oldAddress: artistA,
                newAddress: newAddressA,
            });

            const updatedAddress = await instance.getArtistAddress(artistAID);
            expect(updatedAddress).to.equal(newAddressA);

            const updatedArtistID = await instance.getAddressArtistID(
                newAddressA
            );
            expect(updatedArtistID.toString()).to.equal(artistAID.toString());
        });

        it("Owner can update artist address if owner rights not revoked", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const artistAID = await instance.getAddressArtistID(artistA);
            const newAddressA2 = accounts[7];

            const tx = await instance.updateArtistAddress(
                artistAID,
                newAddressA2,
                { from: owner }
            );
            expectEvent(tx, "ArtistAddressUpdated", {
                artistID: artistAID,
                oldAddress: artistA,
                newAddress: newAddressA2,
            });

            const updatedAddress = await instance.getArtistAddress(artistAID);
            expect(updatedAddress).to.equal(newAddressA2);
        });

        it("Owner cannot update artist address if the artist revoked owner rights", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const artistAID = await instance.getAddressArtistID(artistA);
            await instance.revokeOwnerRightsForArtist({ from: artistA });

            // Reverts with custom error: NotAuthorizedError(address caller)
            await expectCustomError(
                instance.updateArtistAddress(artistAID, artistB, {
                    from: owner,
                }),
                seriesIndexerABI,
                "NotAuthorizedError",
                [owner]
            );
        });

        it("Random address cannot update artist address", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const artistAID = await instance.getAddressArtistID(artistA);

            // Reverts with custom error: NotAuthorizedError(address caller)
            await expectCustomError(
                instance.updateArtistAddress(artistAID, artistB, {
                    from: nonArtist,
                }),
                seriesIndexerABI,
                "NotAuthorizedError",
                [nonArtist]
            );
        });

        it("Cannot update artist address to one already assigned to another artist", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.addSeries([artistB], metadataURI, tokenMapURI, {
                from: artistB,
            });
            const artistAID = await instance.getAddressArtistID(artistA);

            // Reverts with custom error: AddressAlreadyAssignedError(address newAddress)
            await expectCustomError(
                instance.updateArtistAddress(artistAID, artistB, {
                    from: owner,
                }),
                seriesIndexerABI,
                "AddressAlreadyAssignedError",
                [artistB]
            );
        });
    });

    describe("deleteSeries", () => {
        it("should allow owner to delete series", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.deleteSeries(firstSeriesID, {
                from: owner,
            });

            // Verify event emission
            expectEvent(tx, "SeriesDeleted", {
                seriesID: firstSeriesID,
            });

            // Verify series is deleted
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(0);

            // Verify series is removed from artists' series lists
            const artist1Series = await instance.getArtistSeriesIDs(artistA);
            expect(artist1Series.length).to.equal(0);
        });

        it("should allow owner to batch delete series", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.batchDeleteSeries([firstSeriesID, secondSeriesID], { from: owner });
            expectEvent(tx, "SeriesDeleted", {
                seriesID: firstSeriesID,
            });
            expectEvent(tx, "SeriesDeleted", {
                seriesID: secondSeriesID,
            });
            let artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(0);
            artistIDs = await instance.getSeriesArtistIDs(secondSeriesID);
            expect(artistIDs.length).to.equal(0);
        });

        it("should allow artist to delete series", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.deleteSeries(firstSeriesID, { from: artistA });
            expectEvent(tx, "SeriesDeleted", {
                seriesID: firstSeriesID,
            });
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(0);
        });

        it("should allow artist to batch delete series", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.batchDeleteSeries([firstSeriesID, secondSeriesID], { from: artistA });
            expectEvent(tx, "SeriesDeleted", {
                seriesID: firstSeriesID,
            });
            expectEvent(tx, "SeriesDeleted", {
                seriesID: secondSeriesID,
            });
            let artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(0);
            artistIDs = await instance.getSeriesArtistIDs(secondSeriesID);
            expect(artistIDs.length).to.equal(0);
        });

        it("should not allow non-artist to delete series", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Reverts with custom error: CallerNotASeriesArtistError(...)
            await expectCustomError(
                instance.deleteSeries(firstSeriesID, { from: artistB }),
                seriesIndexerABI,
                "CallerNotASeriesArtistError",
                [firstSeriesID.toString(), artistB]
            );
        });

        it("should not allow deleting non-existent series", async () => {
            const nonExistentSeriesID = new BN("9999");
            // Reverts with custom error: SeriesDoesNotExistError(uint256 seriesID)
            await expectCustomError(
                instance.deleteSeries(nonExistentSeriesID, { from: owner }),
                seriesIndexerABI,
                "SeriesDoesNotExistError",
                [nonExistentSeriesID.toString()]
            );
        });
    });

    describe("removeSelfFromSeries", () => {
        it("should allow artist to remove themselves", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const tx = await instance.removeSelfFromSeries(firstSeriesID, {
                from: artistA,
            });

            // Verify event emission
            expect(tx.logs[0].event).to.equal("SeriesUpdated");

            // Check artist is removed from series
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            const artistAID = await instance.getAddressArtistID(artistA);
            expect(artistIDs.map((id) => id.toNumber())).to.not.include(
                artistAID.toNumber()
            );

            // Check series is removed from artist's list
            const artistASeries = await instance.getArtistSeriesIDs(artistA);
            expect(artistASeries.map((id) => id.toNumber())).to.not.include(
                firstSeriesID.toNumber()
            );
        });

        it("should not allow non-artist to remove themselves", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Reverts with custom error: CallerNotASeriesArtistError(...)
            await expectCustomError(
                instance.removeSelfFromSeries(firstSeriesID, { from: artistB }),
                seriesIndexerABI,
                "CallerNotASeriesArtistError",
                [firstSeriesID.toString(), artistB]
            );
        });

        it("should maintain other artists in series after removal", async () => {
            await instance.addSeries(
                [artistA, artistB],
                metadataURI,
                tokenMapURI,
                { from: owner }
            );
            await instance.removeSelfFromSeries(firstSeriesID, {
                from: artistA,
            });

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            const artistBID = await instance.getAddressArtistID(artistB);
            expect(artistIDs.map((id) => id.toNumber())).to.include(
                artistBID.toNumber()
            );
        });
    });

    describe("ownerUpdateSeriesArtists", () => {
        it("should allow owner to update artists", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            const newArtists = [artistB, artistC]; // remove A, add B & C
            const tx = await instance.ownerUpdateSeriesArtists(
                firstSeriesID,
                newArtists,
                { from: owner }
            );

            // Verify event emission
            expect(tx.logs[0].event).to.equal("SeriesUpdated");

            // Check new artists are in series
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            const artistBID = await instance.getAddressArtistID(artistB);
            const artistURI = await instance.getAddressArtistID(artistC);

            const artistIDNumbers = artistIDs.map((id) => id.toNumber());
            expect(artistIDNumbers).to.include(artistBID.toNumber());
            expect(artistIDNumbers).to.include(artistURI.toNumber());

            // Check old artist (A) is removed
            const artistAID = await instance.getAddressArtistID(artistA);
            expect(artistIDNumbers).to.not.include(artistAID.toNumber());
        });

        it("should not allow non-owner to update artists", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // This fails with "Ownable: caller is not the owner", which is a standard OpenZeppelin revert,
            // not a custom error from our contract. We'll keep the old check:
            await expectRevert(
                instance.ownerUpdateSeriesArtists(
                    firstSeriesID,
                    [artistB, artistC],
                    { from: artistA }
                ),
                "Ownable: caller is not the owner"
            );
        });

        it("should not allow update if all artists revoked owner rights", async () => {
            await instance.addSeries(
                [artistA, artistB],
                metadataURI,
                tokenMapURI,
                { from: owner }
            );
            // Revoke rights for all current artists
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistB });

            // Reverts with custom error: AllArtistsRevokedOwnerRightsError(uint256 seriesID)
            await expectCustomError(
                instance.ownerUpdateSeriesArtists(
                    firstSeriesID,
                    [artistC, artistD],
                    { from: owner }
                ),
                seriesIndexerABI,
                "AllArtistsRevokedOwnerRightsError",
                [firstSeriesID.toString()]
            );
        });

        it("should not allow adding artists who revoked owner rights", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            // Revoke rights for artistA
            await instance.revokeOwnerRightsForArtist({ from: artistA });

            await instance.addSeries([artistB], metadataURI, tokenMapURI, {
                from: artistB,
            });
            // Fails with custom error: ArtistRevokedOwnerRightsError(address artistAddr)
            await expectCustomError(
                instance.ownerUpdateSeriesArtists(
                    secondSeriesID,
                    [artistA, artistB],
                    { from: owner }
                ),
                seriesIndexerABI,
                "ArtistRevokedOwnerRightsError",
                [artistA]
            );
        });

        it("should properly update artists' series lists", async () => {
            await instance.addSeries([artistA], metadataURI, tokenMapURI, {
                from: artistA,
            });
            await instance.ownerUpdateSeriesArtists(
                firstSeriesID,
                [artistB, artistC],
                { from: owner }
            );

            // Check artistA's series list doesn't contain the series
            const artistASeries = await instance.getArtistSeriesIDs(artistA);
            expect(artistASeries.map((id) => id.toString())).to.not.include(
                firstSeriesID.toString()
            );

            // Check artistC's series list contains the series
            const artistCSeries = await instance.getArtistSeriesIDs(artistC);
            expect(artistCSeries.map((id) => id.toString())).to.include(
                firstSeriesID.toString()
            );
        });
    });
});
