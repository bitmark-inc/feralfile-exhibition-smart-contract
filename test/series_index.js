const { expect } = require("chai");
const { BN, expectRevert, expectEvent } = require("@openzeppelin/test-helpers");
const SeriesIndexer = artifacts.require("SeriesIndexer");

contract("SeriesIndexer", (accounts) => {
    const owner = accounts[0];
    const artistA = accounts[1];
    const artistB = accounts[2];
    const artistC = accounts[3]; // will be proposed as co-artist
    const artistD = accounts[5]; // for multiple proposals scenario
    const nonArtist = accounts[4]; // never added as an artist
    const firstSeriesID = new BN("1");
    const secondSeriesID = new BN("2");
    const metadataCID = "QmMetadataCID";
    const tokenMapCID = "QmTokenMapCID";
    const newMetadataCID = "QmMetadataCIDNew";
    const newTokenMapCID = "QmTokenMapCIDNew";

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
            const tx = await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            expectEvent(tx, "SeriesIndexed", {
                seriesID: firstSeriesID
            });

            const seriesMeta = await instance.getSeriesMetadata(firstSeriesID);
            const seriesTokenMap = await instance.getSeriesContractTokenData(firstSeriesID);

            expect(seriesMeta).to.equal(metadataCID);
            expect(seriesTokenMap).to.equal(tokenMapCID);

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(1);

            const artistAID = await instance.getAddressArtistID(artistA);
            expect(artistAID.toNumber()).to.not.equal(0, "ArtistA should have a valid artistID");
        });

        it("Owner can add a series with multiple artists", async () => {
            const tx = await instance.ownerAddSeries(
                [artistA, artistB],
                metadataCID,
                tokenMapCID,
                { from: owner }
            );

            expectEvent(tx, "SeriesIndexed", {
                seriesID: firstSeriesID
            });

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(2);
        });

        it("Should revert when metadata CID or token map CID is empty", async () => {
            await expectRevert(
                instance.addSeries("", tokenMapCID, { from: artistA }),
                "Empty metadata CID"
            );

            await expectRevert(
                instance.addSeries(metadataCID, "", { from: artistA }),
                "Empty token map CID"
            );
        });
    });

    describe("Batch Adding Series by Owner", () => {
        const batchArtistsArray = [[artistA], [artistB]];
        const batchMetadataCIDs = ["meta3001", "meta3002"];
        const batchTokenMapCIDs = ["token3001", "token3002"];

        it("Owner can batch add multiple series", async () => {
            const tx = await instance.ownerBatchAddSeries(
                batchArtistsArray,
                batchMetadataCIDs,
                batchTokenMapCIDs,
                { from: owner }
            );

            expectEvent(tx, "SeriesIndexed", { seriesID: firstSeriesID });
            expectEvent(tx, "SeriesIndexed", { seriesID: secondSeriesID });

            const seriesMeta1 = await instance.getSeriesMetadata(firstSeriesID);
            const seriesTokenMap1 = await instance.getSeriesContractTokenData(firstSeriesID);
            expect(seriesMeta1).to.equal("meta3001");
            expect(seriesTokenMap1).to.equal("token3001");

            const seriesMeta2 = await instance.getSeriesMetadata(secondSeriesID);
            const seriesTokenMap2 = await instance.getSeriesContractTokenData(secondSeriesID);
            expect(seriesMeta2).to.equal("meta3002");
            expect(seriesTokenMap2).to.equal("token3002");
        });

        it("Batch add should revert if one of the artists revoked owner rights", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistA });

            await expectRevert(
                instance.ownerBatchAddSeries(
                    [[artistA]],
                    ["meta4001"],
                    ["token4001"],
                    { from: owner }
                ),
                "One of the artists revoked owner rights"
            );
        });

        it("Batch add should revert on array length mismatch", async () => {
            await expectRevert(
                instance.ownerBatchAddSeries(
                    [[artistA]], // Mismatch here
                    ["meta5001","meta5002"],
                    ["token5001","token5002"],
                    { from: owner }
                ),
                "Array length mismatch"
            );
        });
    });

    describe("Updating Series", () => {
        it("Series artist can update the series metadata", async () => {
            await instance.addSeries( metadataCID, tokenMapCID, { from: artistA });
            const tx = await instance.updateSeries(firstSeriesID, newMetadataCID, newTokenMapCID, { from: artistA });
            expectEvent(tx, "SeriesUpdated", {
                seriesID: firstSeriesID
            });

            const updatedMeta = await instance.getSeriesMetadata(firstSeriesID);
            const updatedTokenMap = await instance.getSeriesContractTokenData(firstSeriesID);
            expect(updatedMeta).to.equal(newMetadataCID);
            expect(updatedTokenMap).to.equal(newTokenMapCID);
        });

        it("Non-artist or non-owner cannot update", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await expectRevert(
                instance.updateSeries(firstSeriesID, metadataCID, tokenMapCID, { from: nonArtist }),
                "Caller not a series artist"
            );
        });

        it("Updating a non-existent series should revert", async () => {
            await expectRevert(
                instance.updateSeries(new BN("9999"), metadataCID, tokenMapCID, { from: artistA }),
                "Series does not exist"
            );
        });
    });

    describe("Owner Rights Revocation", () => {
        it("Artist can revoke owner rights", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            const revoked = await instance.ownerRightsRevoked(artistA);
            expect(revoked).to.be.true;
        });

        it("Artist can re-approve owner rights", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            await instance.approveOwnerRightsForArtist({ from: artistA });
            const revoked = await instance.ownerRightsRevoked(artistA);
            expect(revoked).to.be.false;
        });

        it("When all artists revoke, owner cannot update series", async () => {
            await instance.ownerAddSeries(
                [artistA, artistB],
                metadataCID,
                tokenMapCID,
                { from: owner }
            );
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistB });

            await expectRevert(
                instance.updateSeries(firstSeriesID, "x", "y", { from: owner }),
                "Owner rights revoked by all artists in series"
            );

            await instance.approveOwnerRightsForArtist({ from: artistA });
            const tx = await instance.updateSeries(firstSeriesID, "metaNew", "tokenNew", { from: owner });
            expectEvent(tx, "SeriesUpdated", {
                seriesID: firstSeriesID
            });
        });

        it("Non-artist tries to revoke/approve owner rights should revert", async () => {
            await expectRevert(
                instance.revokeOwnerRightsForArtist({ from: nonArtist }),
                "Not an artist"
            );

            await expectRevert(
                instance.approveOwnerRightsForArtist({ from: nonArtist }),
                "Not an artist"
            );
        });
    });

    describe("Propose and Confirm Co-Artist", () => {
        it("Artist in a series can propose a new co-artist", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const tx = await instance.proposeCoArtist(firstSeriesID, artistC, { from: artistA });
            expectEvent(tx, "CoArtistProposed", {
                seriesID: firstSeriesID
            });

            const pending = await instance.getPendingRequests(artistC);
            expect(pending.map((e) => e.toString())).to.include(firstSeriesID.toString());
        });

        it("Co-artist can confirm themselves", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.proposeCoArtist(firstSeriesID, artistC, { from: artistA });
            let artistCID = await instance.getAddressArtistID(artistC);
            const tx = await instance.confirmAsCoArtist(firstSeriesID, { from: artistC });
            expectEvent(tx, "CoArtistConfirmed", {
                seriesID: firstSeriesID,
                confirmedArtistID: artistCID
            });

            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.map(id => id.toString())).to.include(artistCID.toString());

            const pendingAfter = await instance.getPendingRequests(artistC);
            expect(pendingAfter.length).to.equal(0, "Pending requests should be cleared after confirmation");
        });

        it("Cannot confirm if not proposed", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistC });
            await expectRevert(
                instance.confirmAsCoArtist(firstSeriesID, { from: artistC }),
                "No pending proposal"
            );
        });

        it("An artist can cancel a co-artist proposal before confirmation", async () => {
            await instance.ownerAddSeries(
                [artistA, artistB],
                metadataCID,
                tokenMapCID,
                { from: owner }
            );
            await instance.proposeCoArtist(firstSeriesID, artistC, { from: artistA });
            let artistCID = await instance.getAddressArtistID(artistC);
            let pendingBefore = await instance.getPendingRequests(artistC);
            expect(pendingBefore.map((e) => e.toString())).to.include(firstSeriesID.toString());

            const tx = await instance.cancelProposeCoArtist(firstSeriesID, artistC, { from: artistB });
            expectEvent(tx, "CoArtistProposalCancelled", {
                seriesID: firstSeriesID,
                proposerArtistID: await instance.getAddressArtistID(artistB),
                cancelledArtistID: artistCID
            });

            await expectRevert(
                instance.confirmAsCoArtist(firstSeriesID, { from: artistC }),
                "No pending proposal"
            );

            const pendingAfterCancel = await instance.getPendingRequests(artistC);
            expect(pendingAfterCancel.length).to.equal(0, "Pending requests should be cleared after cancellation");
        });

        it("Proposing a co-artist for a non-existent series should revert", async () => {
            await expectRevert(
                instance.proposeCoArtist(new BN("9999"), artistC, { from: artistA }),
                "Series does not exist"
            );
        });

        it("Proposing the same co-artist twice for the same series should revert", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            // Propose artistD
            await instance.proposeCoArtist(firstSeriesID, artistB, { from: artistA });
            // Attempt to propose the same co-artist again
            await expectRevert(
                instance.proposeCoArtist(firstSeriesID, artistB, { from: artistA }),
                "Already proposed"
            );
        });

        it("Non-artist cannot propose co-artist", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await expectRevert(
                instance.proposeCoArtist(firstSeriesID, artistB, { from: nonArtist }),
                "Caller not a series artist"
            );
        });

        it("Propose co-artist with invalid address should revert", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await expectRevert(
                instance.proposeCoArtist(firstSeriesID, "0x0000000000000000000000000000000000000000", { from: artistA }),
                "Invalid address"
            );
        });
    });

    describe("Multiple Co-Artist Proposals", () => {
        it("Can propose multiple co-artists and confirm/cancel them independently", async () => {
            await instance.ownerAddSeries(
                [artistA, artistB],
                metadataCID,
                tokenMapCID,
                { from: owner }
            );
            // Propose artistD firstSeriesID
            await instance.proposeCoArtist(firstSeriesID, artistD, { from: artistA });
            let artistDID = await instance.getAddressArtistID(artistD);

            // Add another new series
            await instance.addSeries("meta1010", "token1010", { from: artistA });
            
            // Propose artistD to secondSeriesID
            await instance.proposeCoArtist(secondSeriesID, artistD, { from: artistA });

            // Now artistD should have 2 pending requests
            let pendingD = await instance.getPendingRequests(artistD);
            expect(pendingD.map(e => e.toString())).to.have.members([firstSeriesID.toString(), secondSeriesID.toString()]);

            // Confirm artistD for secondSeriesID first (out-of-order)
            await instance.confirmAsCoArtist(secondSeriesID, { from: artistD });

            // Now pending should only have firstSeriesID
            let pendingDAfterOneConfirm = await instance.getPendingRequests(artistD);
            expect(pendingDAfterOneConfirm.map(e => e.toString())).to.include(firstSeriesID.toString());
            expect(pendingDAfterOneConfirm.length).to.equal(1);

            // Cancel the remaining proposal (f)
            const artistAID = await instance.getAddressArtistID(artistA);
            const tx = await instance.cancelProposeCoArtist(firstSeriesID, artistD, { from: artistA });
            expectEvent(tx, "CoArtistProposalCancelled", {
                seriesID: firstSeriesID,
                proposerArtistID: artistAID,
                cancelledArtistID: artistDID
            });

            // Pending should now be empty
            let pendingDAfterCancel = await instance.getPendingRequests(artistD);
            expect(pendingDAfterCancel.length).to.equal(0, "All pending requests cleared after confirm/cancel operations");
        });
    });

    describe("Updating Artist's address", () => {
        it("Artist can update their own address", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const artistAID = await instance.getAddressArtistID(artistA);
            const newAddressA = accounts[6];
        
            const tx = await instance.updateArtistAddress(artistAID, newAddressA, { from: artistA });
            expectEvent(tx, "ArtistAddressUpdated", {
                artistID: artistAID,
                oldAddress: artistA,
                newAddress: newAddressA
            });
        
            const updatedAddress = await instance.getArtistAddress(artistAID);
            expect(updatedAddress).to.equal(newAddressA);
        
            const updatedArtistID = await instance.getAddressArtistID(newAddressA);
            expect(updatedArtistID.toString()).to.equal(artistAID.toString());
        });
        
        it("Owner can update artist address if owner rights not revoked", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const artistAID = await instance.getAddressArtistID(artistA);
            const newAddressA2 = accounts[7];
        
            const tx = await instance.updateArtistAddress(artistAID, newAddressA2, { from: owner });
            expectEvent(tx, "ArtistAddressUpdated", {
                artistID: artistAID,
                oldAddress: artistA,
                newAddress: newAddressA2
            });
        
            const updatedAddress = await instance.getArtistAddress(artistAID);
            expect(updatedAddress).to.equal(newAddressA2);
        });
        
        it("Owner cannot update artist address if the artist revoked owner rights", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const artistAID = await instance.getAddressArtistID(artistA);
            await instance.revokeOwnerRightsForArtist({ from: artistA });
        
            await expectRevert(
                instance.updateArtistAddress(artistAID, artistB, { from: owner }),
                "Not authorized to update address"
            );
        });
        
        it("Random address cannot update artist address", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const artistAID = await instance.getAddressArtistID(artistA);
            await expectRevert(
                instance.updateArtistAddress(artistAID, artistB, { from: nonArtist }),
                "Not authorized to update address"
            );
        });
        
        it("Cannot update artist address to one already assigned to another artist", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistB });
            const artistAID = await instance.getAddressArtistID(artistA);
            
            // Attempt to assign artistB's address to artistA
            await expectRevert(
                instance.updateArtistAddress(artistAID, artistB, { from: owner }),
                "Address already assigned to an artist"
            );
        });        
    });

    describe("deleteSeries", () => {
        it("should allow owner to delete series", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const tx = await instance.deleteSeries(firstSeriesID, { from: owner });
            
            // Verify event emission
            expect(tx.logs[0].event).to.equal("SeriesDeleted");
            expect(tx.logs[0].args.seriesID.toString()).to.equal(firstSeriesID.toString());

            // Verify series is deleted
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(0);

            // Verify series is removed from artists' series lists
            const artist1Series = await instance.getArtistSeriesIDs(artistA);
            expect(artist1Series.length).to.equal(0);
        });

        it("should allow artist to delete series", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.deleteSeries(firstSeriesID, { from: artistA });
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            expect(artistIDs.length).to.equal(0);
        });

        it("should not allow non-artist to delete series", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await expectRevert(
                instance.deleteSeries(firstSeriesID, { from: artistB }),
                "Caller not a series artist"
            );
        });

        it("should not allow deleting non-existent series", async () => {
            const nonExistentSeriesID = BN("9999");
            await expectRevert(
                instance.deleteSeries(nonExistentSeriesID, { from: owner }),
                "Series does not exist"
            );
        });
    });

    describe("removeSelfFromSeries", () => {
        it("should allow artist to remove themselves", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const tx = await instance.removeSelfFromSeries(firstSeriesID, { from: artistA });
            
            // Verify event emission
            expect(tx.logs[0].event).to.equal("SeriesUpdated");
            
            // Check artist is removed from series
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            const artistAID = await instance.getAddressArtistID(artistA);
            expect(artistIDs.map(id => id.toNumber())).to.not.include(artistAID.toNumber());
            
            // Check series is removed from artist's list
            const artistASeries = await instance.getArtistSeriesIDs(artistA);
            expect(artistASeries.map(id => id.toNumber())).to.not.include(firstSeriesID);
        });

        it("should not allow non-artist to remove themselves", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await expectRevert(
                instance.removeSelfFromSeries(firstSeriesID, { from: artistB }),
                "Caller not a series artist"
            );
        });

        it("should maintain other artists in series after removal", async () => {
            await instance.ownerAddSeries(
                [artistA, artistB],
                metadataCID,
                tokenMapCID,
                { from: owner }
            );
            await instance.removeSelfFromSeries(firstSeriesID, { from: artistA });
            
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            const artistBID = await instance.getAddressArtistID(artistB);
            expect(artistIDs.map(id => id.toNumber())).to.include(artistBID.toNumber());
        });
    });

    describe("ownerUpdateSeriesArtists", () => {
        it("should allow owner to update artists", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            const newArtists = [artistB, artistC]; // Remove artist1, add artist3
            const tx = await instance.ownerUpdateSeriesArtists(firstSeriesID, newArtists, { from: owner });
            
            // Verify event emission
            expect(tx.logs[0].event).to.equal("SeriesUpdated");
            
            // Check new artists are in series
            const artistIDs = await instance.getSeriesArtistIDs(firstSeriesID);
            const artistBID = await instance.getAddressArtistID(artistB);
            const artistCID = await instance.getAddressArtistID(artistC);
            
            const artistIDNumbers = artistIDs.map(id => id.toNumber());
            expect(artistIDNumbers).to.include(artistBID.toNumber());
            expect(artistIDNumbers).to.include(artistCID.toNumber());
            
            // Check old artist is removed
            const artistAID = await instance.getAddressArtistID(artistA);
            expect(artistIDNumbers).to.not.include(artistAID.toNumber());
        });

        it("should not allow non-owner to update artists", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await expectRevert(
                instance.ownerUpdateSeriesArtists(firstSeriesID, [artistB, artistC], { from: artistA }),
                "Ownable: caller is not the owner"
            );
        });

        it("should not allow update if all artists revoked owner rights", async () => {
            await instance.ownerAddSeries(
                [artistA, artistB],
                metadataCID,
                tokenMapCID,
                { from: owner }
            );
            // Revoke rights for all current artists
            await instance.revokeOwnerRightsForArtist({ from: artistA });
            await instance.revokeOwnerRightsForArtist({ from: artistB });

            await expectRevert(
                instance.ownerUpdateSeriesArtists(firstSeriesID, [artistC, artistD], { from: owner }),
                "Owner rights revoked by all artists"
            );
        });

        it("should not allow adding artists who revoked owner rights", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            // Revoke rights for new artist
            await instance.revokeOwnerRightsForArtist({ from: artistA });

            await instance.addSeries(metadataCID, tokenMapCID, { from: artistB });
            await expectRevert(
                instance.ownerUpdateSeriesArtists(secondSeriesID, [artistA, artistB], { from: owner }),
                "One of the artists revoked owner rights"
            );
        });

        it("should properly update artists' series lists", async () => {
            await instance.addSeries(metadataCID, tokenMapCID, { from: artistA });
            await instance.ownerUpdateSeriesArtists(firstSeriesID, [artistB, artistC], { from: owner });
            
            // Check artistA's series list doesn't contain the series
            const artistASeries = await instance.getArtistSeriesIDs(artistA);
            expect(artistASeries.map(id => id.toString())).to.not.include(firstSeriesID.toString());
            
            // Check artistC's series list contains the series
            const artistCSeries = await instance.getArtistSeriesIDs(artistC);
            expect(artistCSeries.map(id => id.toString())).to.include(firstSeriesID.toString());
        });
    });
});
