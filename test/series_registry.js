const { expect } = require("chai");
const { BN, expectEvent } = require("@openzeppelin/test-helpers");
const SeriesRegistry = artifacts.require("SeriesRegistry");
const SeriesRegistryArtifact = require("../build/contracts/SeriesRegistry.json");
const { expectCustomError } = require("./helper/expectCustomError.js");
const {
    expectBigNumberArraysEqual,
    expectAddressArraysEqual,
} = require("./helper/array.js");
const abi = SeriesRegistryArtifact.abi;

contract("SeriesRegistry", (accounts) => {
    let instance;
    const metadataURI = "ipfs://QmS4ghgMgfFvqPjB4WKXHaN15ZyT4K4JYZxUY5C21U123";
    const tokenDataURI = "ipfs://QmS4ghgMgfFvqPjB4WKXHaN15ZyT4K4JYZxUY5C21U321";
    const zeroAddress = "0x0000000000000000000000000000000000000000";

    beforeEach(async () => {
        instance = await SeriesRegistry.new();
    });

    // ------------------------------------------------------------------------
    // registerSeries
    // ------------------------------------------------------------------------

    describe("registerSeries", () => {
        it("should be able to register a new series by artist that has no delegatees", async () => {
            const artistAddress = accounts[0];
            const expectedArtistID = new BN(1);
            const expectedSeriesCount = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check series details
            // 2.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2 check series administrator ID should be the artist ID
            const administratorID =
                await instance.getSeriesAdministratorID(seriesID);
            expect(administratorID).to.be.a.bignumber.that.equals(
                expectedArtistID
            );

            // 2.3 check series artist addresses
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress]
            );

            // 2.4 check series artist IDs
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 2.5 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                metadataURI
            );

            // 2.6 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                tokenDataURI
            );

            // 3. check total series
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(expectedSeriesCount);

            // 4. check artist addresses
            expect(await instance.getArtistAddress(expectedArtistID)).to.equal(
                artistAddress
            );

            // 5. check artist ID
            expect(
                await instance.getArtistID(artistAddress)
            ).to.be.a.bignumber.that.equals(expectedArtistID);

            // 6. check artist series
            expectBigNumberArraysEqual(
                await instance.getArtistSeriesIDs(artistAddress),
                [seriesID]
            );

            // 7. check event emitted
            expectEvent(tx, "RegisterSeries", {
                seriesID: seriesID,
            });
        });

        it("should be able to register a new series by artist that has delegatees", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedSeriesCount = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 2. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. check series details
            // 3.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 3.2 check series administrator ID should be the artist ID
            const administratorID =
                await instance.getSeriesAdministratorID(seriesID);
            expect(administratorID).to.be.a.bignumber.that.equals(
                expectedArtistID
            );

            // 3.3 check series artist addresses
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress]
            );

            // 3.4 check series artist IDs
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 3.5 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                metadataURI
            );

            // 3.6 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                tokenDataURI
            );

            // 4. check total series
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(expectedSeriesCount);

            // 5. check artist addresses
            expect(await instance.getArtistAddress(expectedArtistID)).to.equal(
                artistAddress
            );

            // 6. check artist ID
            expect(
                await instance.getArtistID(artistAddress)
            ).to.be.a.bignumber.that.equals(expectedArtistID);

            // 7. check artist series ID
            expectBigNumberArraysEqual(
                await instance.getArtistSeriesIDs(artistAddress),
                [seriesID]
            );

            // 8. check series delegatees should include the delegatee address
            expectAddressArraysEqual(
                await instance.getSeriesDelegatees(seriesID),
                [delegateeAddress]
            );

            // 9. check artist delegatees should include the delegatee address
            expectAddressArraysEqual(
                await instance.getDelegatees(artistAddress),
                [delegateeAddress]
            );

            // 10. check event emitted
            expectEvent(tx, "RegisterSeries", {
                seriesID: seriesID,
            });
        });

        it("should be able to register a new series by delegatee", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedSeriesCount = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 2. delegatee register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: delegateeAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. check series details
            // 3.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 3.2 check series administrator ID should be the artist ID
            const administratorID =
                await instance.getSeriesAdministratorID(seriesID);
            expect(administratorID).to.be.a.bignumber.that.equals(
                expectedArtistID
            );

            // 3.3 check series artist addresses
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress]
            );

            // 3.4 check series artist IDs
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 3.5 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                metadataURI
            );

            // 3.6 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                tokenDataURI
            );

            // 4. check total series
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(expectedSeriesCount);

            // 5. check artist addresses
            expect(await instance.getArtistAddress(expectedArtistID)).to.equal(
                artistAddress
            );

            // 6. check artist ID
            expect(
                await instance.getArtistID(artistAddress)
            ).to.be.a.bignumber.that.equals(expectedArtistID);

            // 7. check artist series ID
            expectBigNumberArraysEqual(
                await instance.getArtistSeriesIDs(artistAddress),
                [seriesID]
            );

            // 8. check series delegatees should include the delegatee address
            expectAddressArraysEqual(
                await instance.getSeriesDelegatees(seriesID),
                [delegateeAddress]
            );

            // 9. check artist delegatees should include the delegatee address
            expectAddressArraysEqual(
                await instance.getDelegatees(artistAddress),
                [delegateeAddress]
            );

            // 10. check event emitted
            expectEvent(tx, "RegisterSeries", {
                seriesID: seriesID,
            });
        });

        it("should NOT be able to register a new series if the caller is not an artist or delegatee", async () => {
            const artistAddress = accounts[0];
            const someoneAddress = accounts[1];

            await expectCustomError(
                instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    {
                        from: someoneAddress,
                    }
                ),
                abi,
                "NotAuthorizedError"
            );
        });

        it("should NOT be able to register a new series if the artist address is zero", async () => {
            await expectCustomError(
                instance.registerSeries(zeroAddress, metadataURI, tokenDataURI),
                abi,
                "GenericError",
                ["artist address is zero"]
            );
        });

        it("should NOT be able to register a new series if the metadataURI is empty", async () => {
            const artistAddress = accounts[0];
            await expectCustomError(
                instance.registerSeries(artistAddress, "", tokenDataURI),
                abi,
                "GenericError",
                ["empty metadataURI"]
            );
        });

        it("should NOT be able to register a new series if the tokenDataURI is empty", async () => {
            const artistAddress = accounts[0];
            await expectCustomError(
                instance.registerSeries(artistAddress, metadataURI, ""),
                abi,
                "GenericError",
                ["empty tokenDataURI"]
            );
        });
    });

    // ------------------------------------------------------------------------
    // updateSeries
    // ------------------------------------------------------------------------

    describe("updateSeries", () => {
        it("should be able to update a series by artist", async () => {
            const artistAddress = accounts[0];
            const expectedSeriesID = new BN(1);
            const newMetadataURI = "new metadataURI";
            const newTokenDataURI = "new tokenDataURI";

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check the series details
            // 2.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                metadataURI
            );

            // 2.3 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                tokenDataURI
            );

            // 2.4 check series artist addresses
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress]
            );

            // 3. artist update series
            const updateTx = await instance.updateSeries(
                seriesID,
                newMetadataURI,
                newTokenDataURI,
                { from: artistAddress }
            );

            // 4. check series details
            // 4.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 4.2 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                newMetadataURI
            );

            // 4.3 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                newTokenDataURI
            );

            // 4.4 check series artist addresses
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress]
            );

            // 5. check event emitted
            expectEvent(updateTx, "UpdateSeries", {
                seriesID: seriesID,
            });
        });

        it("should be able to update a series by delegatee", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const expectedSeriesID = new BN(1);
            const newMetadataURI = "new metadataURI";
            const newTokenDataURI = "new tokenDataURI";

            // 1. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 2. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. check series details
            // 3.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 3.2 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                metadataURI
            );

            // 3.3 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                tokenDataURI
            );

            // 4. Delegatee update series
            const updateTx = await instance.updateSeries(
                seriesID,
                newMetadataURI,
                newTokenDataURI,
                { from: delegateeAddress }
            );

            // 5. check series details
            // 5.1 check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 5.2 check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                newMetadataURI
            );

            // 5.3 check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                newTokenDataURI
            );

            // 5. check event emitted
            expectEvent(updateTx, "UpdateSeries", {
                seriesID: seriesID,
            });
        });

        it("should NOT be able to update a series if the caller is not an artist or delegatee", async () => {
            const artistAddress = accounts[0];
            const someoneAddress = accounts[1];
            const newMetadataURI = "new metadataURI";
            const newTokenDataURI = "new tokenDataURI";

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. update series by someone is not artist or delegatee
            await expectCustomError(
                instance.updateSeries(
                    seriesID,
                    newMetadataURI,
                    newTokenDataURI,
                    {
                        from: someoneAddress,
                    }
                ),
                abi,
                "NotAuthorizedError"
            );
        });

        it("should NOT be able to update a series if the metadataURI is empty", async () => {
            const artistAddress = accounts[0];
            const newMetadataURI = "";
            const newTokenDataURI = "new tokenDataURI";

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. update series with metadataURI is empty
            await expectCustomError(
                instance.updateSeries(
                    seriesID,
                    newMetadataURI,
                    newTokenDataURI,
                    {
                        from: artistAddress,
                    }
                ),
                abi,
                "GenericError",
                ["empty metadataURI"]
            );
        });

        it("should NOT be able to update a series if the tokenDataURI is empty", async () => {
            const artistAddress = accounts[0];
            const newMetadataURI = "new metadataURI";
            const newTokenDataURI = "";

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. update series with tokenDataURI is empty
            await expectCustomError(
                instance.updateSeries(
                    seriesID,
                    newMetadataURI,
                    newTokenDataURI,
                    {
                        from: artistAddress,
                    }
                ),
                abi,
                "GenericError",
                ["empty tokenDataURI"]
            );
        });

        it("should NOT be able to update a series if the series does not exist", async () => {
            const artistAddress = accounts[0];
            const seriesID = new BN(1);
            const newMetadataURI = "new metadataURI";
            const newTokenDataURI = "new tokenDataURI";

            // 1. update series that does not exist
            await expectCustomError(
                instance.updateSeries(
                    seriesID,
                    newMetadataURI,
                    newTokenDataURI,
                    {
                        from: artistAddress,
                    }
                ),
                abi,
                "SeriesNotExistsError"
            );
        });
    });

    // ------------------------------------------------------------------------
    // deleteSeries
    // ------------------------------------------------------------------------

    describe("deleteSeries", () => {
        it("should be able to delete the series by artist has administrator role and has only one series", async () => {
            const artistAddress = accounts[0];
            const expectedArtistID = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check series registered successfully
            // 2.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2. check series metadata URI
            expect(await instance.getSeriesMetadataURI(seriesID)).to.equal(
                metadataURI
            );

            // 2.3. check series token data URI
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.equal(
                tokenDataURI
            );

            // 2.4. check series artist addresses
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress]
            );

            // 2.5. check series administrator
            expect(
                await instance.getSeriesAdministratorID(seriesID)
            ).to.be.a.bignumber.that.equals(expectedArtistID);

            // 3. artist delete series
            const deleteTx = await instance.deleteSeries(seriesID, {
                from: artistAddress,
            });

            // 4. check total series should be 0
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 5. check total artist series should be 0
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 6. check artist existence globally should be false
            expect(await instance.getArtistAddress(expectedArtistID)).to.equal(
                zeroAddress
            );

            // 7. check artist existence in series should be false
            expect(await instance.getSeriesArtistIDs(seriesID)).to.be.empty;

            // 8. check event emitted
            expectEvent(deleteTx, "DeleteSeries", {
                seriesID: seriesID,
            });
        });

        it("should be able to delete the series by artist has administrator role and has delegatees", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 2. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. check series registered successfully
            // 3.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 3.2. check administrator
            expect(
                await instance.getSeriesAdministratorID(seriesID)
            ).to.be.a.bignumber.that.equals(expectedArtistID);

            // 3.3. check artist IDs
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 4. artist delete series
            const deleteTx = await instance.deleteSeries(seriesID, {
                from: artistAddress,
            });

            // 5. check total series should be 0
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 6. check total artist series should be 0
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 7. check artist existence globally should be false
            expect(await instance.getArtistAddress(expectedArtistID)).to.equal(
                zeroAddress
            );

            // 8. check artist existence in series should be false
            expect(await instance.getSeriesArtistIDs(seriesID)).to.be.empty;

            // 9. check event emitted
            expectEvent(deleteTx, "DeleteSeries", {
                seriesID: seriesID,
            });
        });

        it("should be able to delete a series by artist has administrator role and has multiple series", async () => {
            const artistAddress = accounts[0];
            const expectedArtistID = new BN(1);
            const registeredSeriesIDs = [];

            // 1. artist register multiple series (2 series)
            for (let i = 0; i < 2; i++) {
                const tx = await instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    { from: artistAddress }
                );

                // 1.1. check series ID
                const seriesID = tx.logs[0].args.seriesID;
                expect(seriesID).to.be.a.bignumber.that.equals(new BN(i + 1));

                // 1.2. check artist ID
                expectBigNumberArraysEqual(
                    await instance.getSeriesArtistIDs(seriesID),
                    [expectedArtistID]
                );

                // 1.3. add series ID to registeredSeriesIDs
                registeredSeriesIDs.push(seriesID);
            }

            // 2. artist delete one series
            const deletingSeriesID = registeredSeriesIDs[0];
            const deleteTx = await instance.deleteSeries(deletingSeriesID, {
                from: artistAddress,
            });

            // 3. check total series count should be 1
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(new BN(1));

            // 4. check total artist series should be 1
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equals(new BN(1));

            // 5. check artist existence globally should be true
            expect(await instance.getArtistAddress(expectedArtistID)).to.equal(
                artistAddress
            );

            // 6. check artist existence in series should be false
            expect(await instance.getSeriesArtistIDs(deletingSeriesID)).to.be
                .empty;

            // 7. check event emitted
            expectEvent(deleteTx, "DeleteSeries", {
                seriesID: deletingSeriesID,
            });
        });

        it("should be able to delete a series by collaborator after the series creator has opted out", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedCollaboratorID = new BN(2);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check series registered successfully
            // 2.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2. check series artist ID
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 3. artist invite collaborator to series
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 4. collaborator opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 5. check collaborator should be in series
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID, expectedCollaboratorID]
            );

            // 6. artist opt out of series
            await instance.optOutSeries(seriesID, {
                from: artistAddress,
            });

            // 7. check artist should not be in series
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedCollaboratorID]
            );

            // 8. collaborator delete series
            const deleteTx = await instance.deleteSeries(seriesID, {
                from: collaboratorAddress,
            });

            // 9. check total series should be 0
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 10. check total artist series should be 0
            expect(
                await instance.getTotalArtistSeries(collaboratorAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 11. check event emitted
            expectEvent(deleteTx, "DeleteSeries", {
                seriesID: seriesID,
            });
        });

        it("should NOT be able to delete a series if the caller isn't series artist", async () => {
            const artistAddress = accounts[0];
            const someoneAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check series registered successfully
            // 2.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2. check series artist ID
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 3. delete series by someone and check error is `NotAdministratorError`
            await expectCustomError(
                instance.deleteSeries(seriesID, {
                    from: someoneAddress,
                }),
                abi,
                "NotAdministratorError"
            );
        });

        it("should NOT be able to delete a series if the caller is collaborator but not administrator", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedCollaboratorID = new BN(2);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check series registered successfully
            // 2.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2. check series artist ID
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 3. artist invite collaborator to series
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 4. collaborator opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 5. check collaborator should be in series
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID, expectedCollaboratorID]
            );

            // 6. delete series by collaborator and check error is `NotAdministratorError`
            await expectCustomError(
                instance.deleteSeries(seriesID, {
                    from: collaboratorAddress,
                }),
                abi,
                "NotAdministratorError"
            );
        });

        it("should NOT be able to delete a series if the series does not exist", async () => {
            const artistAddress = accounts[0];
            const seriesID = new BN(1);

            // 1. delete series that does not exist
            await expectCustomError(
                instance.deleteSeries(seriesID, {
                    from: artistAddress,
                }),
                abi,
                "SeriesNotExistsError"
            );
        });
    });

    // ------------------------------------------------------------------------
    // optOutSeries
    // ------------------------------------------------------------------------

    describe("optOutSeries", () => {
        it("should be able to opt out of a series if artist has only one series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedCollaboratorID = new BN(2);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check series registered successfully
            // 2.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 2.2. check series artist ID
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 3. artist invite collaborator to series
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 4. collaborator opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 5. check series artist IDs should include the both artist and collaborator
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID, expectedCollaboratorID]
            );

            // 6. artist opt out of series
            const optOutTx = await instance.optOutSeries(seriesID, {
                from: artistAddress,
            });

            // 7. check artist should not be in series
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedCollaboratorID]
            );

            // 8. check collaborator existence globally should be true
            expect(
                await instance.getArtistAddress(expectedCollaboratorID)
            ).to.equal(collaboratorAddress);

            // 9. check collaborator existence in series should be true
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedCollaboratorID]
            );

            // 10. check artist existence in series should be false
            expect(await instance.getArtistSeriesIDs(artistAddress)).to.be
                .empty;

            // 11. check artist existence globally should be false
            expect(
                await instance.getArtistAddress(expectedArtistID)
            ).to.be.equal(zeroAddress);

            // 12. check series administrator should be the collaborator address
            expect(
                await instance.getSeriesAdministratorID(seriesID)
            ).to.be.a.bignumber.that.equals(expectedCollaboratorID);

            // 13. check event emitted
            expectEvent(optOutTx, "OptOutSeries", {
                seriesID: seriesID,
                artistAddress: artistAddress,
            });
        });

        it("should be able to opt out of a series if artist has multiple series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedCollaboratorID = new BN(2);
            const expectedSeriesIDs = [];

            // 1. register multiple series (2 series)
            for (let i = 0; i < 2; i++) {
                const tx = await instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    { from: artistAddress }
                );
                const seriesID = tx.logs[0].args.seriesID;
                expectedSeriesIDs.push(seriesID);

                // 1.1. check series ID
                expect(seriesID).to.be.a.bignumber.that.equals(
                    expectedSeriesIDs[i]
                );

                // 1.2. check series artist ID
                expectBigNumberArraysEqual(
                    await instance.getSeriesArtistIDs(seriesID),
                    [expectedArtistID]
                );
            }

            // 2. artist invite collaborator to a series
            const seriesID = expectedSeriesIDs[0];
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. collaborator opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 4. artist opt out series
            const optOutTx = await instance.optOutSeries(seriesID, {
                from: artistAddress,
            });

            // 5. check series existence should be true
            expect(await instance.getSeriesMetadataURI(seriesID)).to.be.equal(
                metadataURI
            );
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.be.equal(
                tokenDataURI
            );

            // 6. check total series count should be 2
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(new BN(2));

            // 7. check total artist series should be 1
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equals(new BN(1));

            // 8. check artist existence globally should be true
            expect(
                await instance.getArtistAddress(expectedArtistID)
            ).to.be.equal(artistAddress);

            // 9. check artist existence in series should be false
            expectBigNumberArraysEqual(
                await instance.getArtistSeriesIDs(artistAddress),
                [expectedSeriesIDs[1]]
            );

            // 10. check collaborator existence in series should be true
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedCollaboratorID]
            );

            // 11. check series administrator should be the collaborator address
            expect(
                await instance.getSeriesAdministratorID(seriesID)
            ).to.be.a.bignumber.that.equals(expectedCollaboratorID);

            // 12. check event emitted
            expectEvent(optOutTx, "OptOutSeries", {
                seriesID: seriesID,
                artistAddress: artistAddress,
            });
        });

        it("should be able to opt out of a series if the artist has some pending collaboration invitations", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress1 = accounts[1];
            const collaboratorAddress2 = accounts[2];
            const expectedArtistID = new BN(1);
            const expectedCollaboratorID1 = new BN(2);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 1.1. check series ID
            expect(seriesID).to.be.a.bignumber.that.equals(expectedSeriesID);

            // 1.2. check series artist ID
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID]
            );

            // 2. artist invite collaborators to series (2 collaborators)
            await instance.inviteCollaborator(seriesID, collaboratorAddress1, {
                from: artistAddress,
            });
            await instance.inviteCollaborator(seriesID, collaboratorAddress2, {
                from: artistAddress,
            });

            // 3. one collaborator opt in to series, another collaborator doesn't opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress1,
            });

            // 3.1. check collaborator 1 should be in series and collaborator 2 should not be in series
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID, expectedCollaboratorID1]
            );

            // 4. artist opt out series
            const optOutTx = await instance.optOutSeries(seriesID, {
                from: artistAddress,
            });

            // 5. check series existence should be true
            expect(await instance.getSeriesMetadataURI(seriesID)).to.be.equal(
                metadataURI
            );
            expect(await instance.getSeriesTokenDataURI(seriesID)).to.be.equal(
                tokenDataURI
            );

            // 6. check total series count should be 1
            expect(
                await instance.getTotalSeries()
            ).to.be.a.bignumber.that.equals(new BN(1));

            // 7. check total artist series should be 0
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 8. check total collaborator 1 series should be 1
            expect(
                await instance.getTotalArtistSeries(collaboratorAddress1)
            ).to.be.a.bignumber.that.equals(new BN(1));

            // 9. check total collaborator 2 series should be 0
            expect(
                await instance.getTotalArtistSeries(collaboratorAddress2)
            ).to.be.a.bignumber.that.equals(new BN(0));

            // 10. check artist existence globally should be false
            expect(
                await instance.getArtistAddress(expectedArtistID)
            ).to.be.equal(zeroAddress);

            // 11. check pending collaboration invitations should be empty
            expect(await instance.getCollaborationInvitees(seriesID)).to.be
                .empty;
            expect(
                await instance.getCollaborationInviterSeriesIDs(artistAddress)
            ).to.be.empty;

            // 12. check event emitted
            expectEvent(optOutTx, "OptOutSeries", {
                seriesID: seriesID,
                artistAddress: artistAddress,
            });
        });

        it("should NOT be able to opt out of a series if the caller is not an artist", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const someoneAddress = accounts[2];

            // 1. register series by artist
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist invite collaborator to series
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. collaborator opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 4. someone else opt out of series
            expectCustomError(
                instance.optOutSeries(seriesID, {
                    from: someoneAddress,
                }),
                abi,
                "NotArtistError"
            );
        });

        it("should NOT be able to opt out of a series if the series does not exist", async () => {
            expectCustomError(
                instance.optOutSeries(new BN(1), {
                    from: accounts[0],
                }),
                abi,
                "SeriesNotExistsError"
            );
        });

        it("should NOT be able to opt out of a series if the series is not collaborative", async () => {
            const artistAddress = accounts[0];

            // 1. register series by artist
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. opt out series by artist
            expectCustomError(
                instance.optOutSeries(seriesID, {
                    from: artistAddress,
                }),
                abi,
                "NotCollaborativeSeriesError"
            );
        });
    });

    // ------------------------------------------------------------------------
    // inviteCollaborator
    // ------------------------------------------------------------------------

    describe("inviteCollaborator", () => {
        it("should be able to create collaboration invitation by artist", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedSeriesID = new BN(1);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist create collaboration invitation
            const inviteTx = await instance.inviteCollaborator(
                seriesID,
                collaboratorAddress,
                {
                    from: artistAddress,
                }
            );

            // 3. check pending collaboration invitation
            expectAddressArraysEqual(
                await instance.getCollaborationInvitees(seriesID),
                [collaboratorAddress]
            );
            expectBigNumberArraysEqual(
                await instance.getCollaborationInviterSeriesIDs(artistAddress),
                [seriesID]
            );
            expectBigNumberArraysEqual(
                await instance.getCollaborationInviteeSeriesIDs(
                    collaboratorAddress
                ),
                [seriesID]
            );

            // 4. check event emitted
            expectEvent(inviteTx, "InviteCollaborator", {
                seriesID: seriesID,
                inviterAddress: artistAddress,
                inviteeAddress: collaboratorAddress,
            });
        });

        it("should be able to create collaboration invitation by delegatee", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const delegateeAddress = accounts[2];

            // 1. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 2. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. delegatee create collaboration invitation
            const inviteTx = await instance.inviteCollaborator(
                seriesID,
                collaboratorAddress,
                {
                    from: delegateeAddress,
                }
            );

            // 4. check pending collaboration invitation
            expectAddressArraysEqual(
                await instance.getCollaborationInvitees(seriesID),
                [collaboratorAddress]
            );
            expectBigNumberArraysEqual(
                await instance.getCollaborationInviterSeriesIDs(
                    delegateeAddress
                ),
                [seriesID]
            );
            expectBigNumberArraysEqual(
                await instance.getCollaborationInviteeSeriesIDs(
                    collaboratorAddress
                ),
                [seriesID]
            );

            // 5. check event emitted
            expectEvent(inviteTx, "InviteCollaborator", {
                seriesID: seriesID,
                inviterAddress: delegateeAddress,
                inviteeAddress: collaboratorAddress,
            });
        });

        it("should NOT be able to create collaboration invitation if the caller is not an artist or delegatee", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const someoneAddress = accounts[2];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. someone else create collaboration invitation
            expectCustomError(
                instance.inviteCollaborator(seriesID, collaboratorAddress, {
                    from: someoneAddress,
                }),
                abi,
                "NotAuthorizedError"
            );
        });

        it("should NOT be able to create collaboration invitation if the series does not exist", async () => {
            expectCustomError(
                instance.inviteCollaborator(new BN(1), accounts[1], {
                    from: accounts[0],
                }),
                abi,
                "SeriesNotExistsError"
            );
        });

        it("should NOT be able to create collaboration invitation if the collaborator is already in the series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. invitee opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 4. check error is `ArtistExistsError`
            expectCustomError(
                instance.inviteCollaborator(seriesID, collaboratorAddress, {
                    from: artistAddress,
                }),
                abi,
                "ArtistExistsError"
            );
        });

        it("should NOT be able to create collaboration invitation if the collaborator is already invited to the series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. check error is `DuplicateInvitationError`
            expectCustomError(
                instance.inviteCollaborator(seriesID, collaboratorAddress, {
                    from: artistAddress,
                }),
                abi,
                "DuplicateInvitationError"
            );
        });

        it("should NOT be able to create collaboration invitation if the collaborator address is zero", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. check error is `GenericError` with message `invitee address is zero`
            expectCustomError(
                instance.inviteCollaborator(seriesID, zeroAddress, {
                    from: artistAddress,
                }),
                abi,
                "GenericError",
                ["invitee address is zero"]
            );
        });
    });

    // ------------------------------------------------------------------------
    // cancelCollaborationInvitation
    // ------------------------------------------------------------------------

    describe("cancelCollaborationInvitation", () => {
        it("should be able to cancel a collaboration invitation invited by artist", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. artist cancel collaboration invitation
            const cancelTx = await instance.cancelCollaborationInvitation(
                seriesID,
                collaboratorAddress,
                { from: artistAddress }
            );

            // 4. check pending collaboration invitation should be empty
            expect(await instance.getCollaborationInvitees(seriesID)).to.be
                .empty;

            // 5. check pending collaborative series should be empty
            expect(
                await instance.getCollaborationInviterSeriesIDs(artistAddress)
            ).to.be.empty;

            // 6. check event emitted
            expectEvent(cancelTx, "CancelCollaborationInvitation", {
                seriesID: seriesID,
                inviteeAddress: collaboratorAddress,
            });
        });

        it("should be able to cancel a collaboration invitation invited by delegatee", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const delegateeAddress = accounts[2];

            // 1. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 2. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. delegatee create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: delegateeAddress,
            });

            // 4. delegatee cancel collaboration invitation
            const cancelTx = await instance.cancelCollaborationInvitation(
                seriesID,
                collaboratorAddress,
                { from: delegateeAddress }
            );

            // 5. check pending collaboration invitation should be empty
            expect(await instance.getCollaborationInvitees(seriesID)).to.be
                .empty;

            // 6. check pending collaborative series should be empty
            expect(
                await instance.getCollaborationInviterSeriesIDs(artistAddress)
            ).to.be.empty;

            // 7. check event emitted
            expectEvent(cancelTx, "CancelCollaborationInvitation", {
                seriesID: seriesID,
                inviteeAddress: collaboratorAddress,
            });
        });

        it("should NOT be able to cancel a collaboration invitation if the series does not exist", async () => {
            expectCustomError(
                instance.cancelCollaborationInvitation(new BN(1), accounts[1], {
                    from: accounts[0],
                }),
                abi,
                "SeriesNotExistsError"
            );
        });

        it("should NOT be able to cancel a collaboration invitation if the collaborator is not invited to the series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist cancel collaboration invitation for collaborator that is not invited to the series
            expectCustomError(
                instance.cancelCollaborationInvitation(
                    seriesID,
                    collaboratorAddress,
                    {
                        from: artistAddress,
                    }
                ),
                abi,
                "NotCollaborationInviterError",
                [artistAddress]
            );
        });

        it("should NOT be able to cancel a collaboration invitation if it's not invited by the caller", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const someoneAddress = accounts[2];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. someone else cancel collaboration invitation
            expectCustomError(
                instance.cancelCollaborationInvitation(
                    seriesID,
                    collaboratorAddress,
                    {
                        from: someoneAddress,
                    }
                ),
                abi,
                "NotCollaborationInviterError",
                [someoneAddress]
            );
        });
    });

    // ------------------------------------------------------------------------
    // optInCollaboration
    // ------------------------------------------------------------------------

    describe("optInCollaboration", () => {
        it("should be able to opt in to a series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];
            const expectedArtistID = new BN(1);
            const expectedCollaboratorID = new BN(2);

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. collaborator opt in to series
            const optInTx = await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 4. check pending collaboration invitation should be empty
            expect(await instance.getCollaborationInvitees(seriesID)).to.be
                .empty;

            // 5. check pending collaborative series should be empty
            expect(
                await instance.getCollaborationInviterSeriesIDs(artistAddress)
            ).to.be.empty;

            // 6. check total artist in series should be 2
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistAddress, collaboratorAddress]
            );
            expectBigNumberArraysEqual(
                await instance.getSeriesArtistIDs(seriesID),
                [expectedArtistID, expectedCollaboratorID]
            );

            // 7. check event emitted
            expectEvent(optInTx, "OptInCollaboration", {
                seriesID: seriesID,
                collaboratorAddress: collaboratorAddress,
            });
        });

        it("should NOT be able to opt in to a series if the series does not exist", async () => {
            expectCustomError(
                instance.optInCollaboration(new BN(1), {
                    from: accounts[1],
                }),
                abi,
                "SeriesNotExistsError"
            );
        });

        it("should NOT be able to opt in to a series if the collaborator is not invited to the series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. collaborator opt in to series
            expectCustomError(
                instance.optInCollaboration(seriesID, {
                    from: collaboratorAddress,
                }),
                abi,
                "NotPendingInvitationError",
                [seriesID, collaboratorAddress]
            );
        });

        it("should NOT be able to opt in to a series if the collaborator is already in the series", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist register series
            const tx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 2. artist send collaboration invitation to collaborator
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. collaborator opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 4. collaborator opt in to series again
            expectCustomError(
                instance.optInCollaboration(seriesID, {
                    from: collaboratorAddress,
                }),
                abi,
                "NotPendingInvitationError",
                [seriesID, collaboratorAddress]
            );
        });
    });

    // ------------------------------------------------------------------------
    // updateArtistAddress
    // ------------------------------------------------------------------------

    describe("updateArtistAddress", () => {
        it("should be able to update an artist's address to a new address hasn't been assigned to any series", async () => {
            const artistAddress = accounts[0];
            const newArtistAddress = accounts[1];
            const seriesIDs = [];

            // 1. artist register series
            for (let i = 0; i < 2; i++) {
                const tx = await instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    { from: artistAddress }
                );
                seriesIDs.push(tx.logs[0].args.seriesID);
            }

            // 2. artist update artist address to a new address hasn't been assigned to any series
            const tx = await instance.updateArtistAddress(newArtistAddress, {
                from: artistAddress,
            });

            // 3. check series artist addresses should include only the new address
            for (let i = 0; i < seriesIDs.length; i++) {
                expectAddressArraysEqual(
                    await instance.getSeriesArtistAddresses(seriesIDs[i]),
                    [newArtistAddress]
                );
            }

            // 4. check total artist series
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));
            expect(
                await instance.getTotalArtistSeries(newArtistAddress)
            ).to.be.a.bignumber.that.equals(new BN(2));

            // 5. check event emitted
            expectEvent(tx, "UpdateArtistAddress", {
                oldAddress: artistAddress,
                newAddress: newArtistAddress,
            });
        });

        it("should be able to update an artist's address to a new address that has been assigned to another series", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];

            // 1. artist A register series
            const txA = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesIDA = txA.logs[0].args.seriesID;

            // 2. artist B register series
            const txB = await instance.registerSeries(
                artistBAddress,
                metadataURI,
                tokenDataURI,
                { from: artistBAddress }
            );
            const seriesIDB = txB.logs[0].args.seriesID;

            // 3. artist A update artist address to artist B's address
            const tx = await instance.updateArtistAddress(artistBAddress, {
                from: artistAAddress,
            });

            // 4. check series artist addresses should include the new address
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesIDA),
                [artistBAddress]
            );
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesIDB),
                [artistBAddress]
            );

            // 5. check total artist series
            expect(
                await instance.getTotalArtistSeries(artistAAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));
            expect(
                await instance.getTotalArtistSeries(artistBAddress)
            ).to.be.a.bignumber.that.equals(new BN(2));

            // 6. check event emitted
            expectEvent(tx, "AssignSeries", {
                seriesID: seriesIDA,
                assignerAddress: artistAAddress,
                assigneeAddress: artistBAddress,
            });
        });

        it("should be able to update an artist's address to a new address that both of them has delegatees", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];
            const delegateeAAddress = accounts[2];
            const delegateeBAddress = accounts[3];

            // 1. artist A add delegatee
            await instance.addDelegatee(delegateeAAddress, {
                from: artistAAddress,
            });

            // 2. artist B add delegatee
            await instance.addDelegatee(delegateeBAddress, {
                from: artistBAddress,
            });

            // 3. artist A register series
            const tx = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 4. artist A update artist address to person B's address
            const updateTx = await instance.updateArtistAddress(
                artistBAddress,
                {
                    from: artistAAddress,
                }
            );

            // 5. check series artist addresses should include the new address
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistBAddress]
            );

            // 6. check series delegatees should include the new address's delegatee address
            expectAddressArraysEqual(
                await instance.getSeriesDelegatees(seriesID),
                [delegateeBAddress]
            );

            // 7. check the delegatee address should be the same as before
            expectAddressArraysEqual(
                await instance.getDelegatees(artistAAddress),
                [delegateeAAddress]
            );
            expectAddressArraysEqual(
                await instance.getDelegatees(artistBAddress),
                [delegateeBAddress]
            );

            // 8. check event emitted
            expectEvent(updateTx, "UpdateArtistAddress", {
                oldAddress: artistAAddress,
                newAddress: artistBAddress,
            });
        });

        it("should be able to update an artist's address to a new address that the old address has pending collaboration invitations", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];
            const collaboratorAddress1 = accounts[2];
            const collaboratorAddress2 = accounts[3];
            const delegateeAddress = accounts[4];

            // 1. artist A add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAAddress,
            });

            // 2. artist A register series
            const tx = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID = tx.logs[0].args.seriesID;

            // 3. artist A create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress1, {
                from: artistAAddress,
            });

            // 4. delegatee create collaboration invitation
            await instance.inviteCollaborator(seriesID, collaboratorAddress2, {
                from: delegateeAddress,
            });

            // 5. artist A update artist address to artist B's address
            const updateTx = await instance.updateArtistAddress(
                artistBAddress,
                {
                    from: artistAAddress,
                }
            );

            // 6. check series artist addresses should include only the new address
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [artistBAddress]
            );

            // 7. check total artist series
            expect(
                await instance.getTotalArtistSeries(artistAAddress)
            ).to.be.a.bignumber.that.equals(new BN(0));
            expect(
                await instance.getTotalArtistSeries(artistBAddress)
            ).to.be.a.bignumber.that.equals(new BN(1));

            // 8. check pending collaboration invitations.
            // since the collaborator is invited by the artist A, so the collaboration invitation
            // should be transferred to the artist B.
            // All the invitation created by artist A's delegatee should be terminated.

            // 8.1. the collaboration invitation should be still.
            expectAddressArraysEqual(
                await instance.getCollaborationInvitees(seriesID),
                [collaboratorAddress1]
            );

            // 8.2. the artist A's collaboration invitations should be terminated.
            expect(
                await instance.getCollaborationInviterSeriesIDs(artistAAddress)
            ).to.be.empty;

            // 8.3. the artist A's collaboration invitations should be transferred to the artist B.
            expectBigNumberArraysEqual(
                await instance.getCollaborationInviterSeriesIDs(artistBAddress),
                [seriesID]
            );

            // 8.4. the artist A's delegatee's collaboration invitations should be terminated.
            expect(
                await instance.getCollaborationInviterSeriesIDs(
                    delegateeAddress
                )
            ).to.be.empty;

            // 9. check event emitted
            expectEvent(updateTx, "UpdateArtistAddress", {
                oldAddress: artistAAddress,
                newAddress: artistBAddress,
            });
        });

        it("should NOT be able to update an artist's address to a new address that is the same as the current address", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );

            // 2. update artist address to the same address
            expectCustomError(
                instance.updateArtistAddress(artistAddress, {
                    from: artistAddress,
                }),
                abi,
                "SameAddressError",
                [artistAddress]
            );
        });

        it("should NOT be able to update an artist's address to a new address that is zero", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );

            // 2. update artist address to zero address
            expectCustomError(
                instance.updateArtistAddress(zeroAddress, {
                    from: artistAddress,
                }),
                abi,
                "GenericError",
                ["new address is zero"]
            );
        });

        it("should NOT be able to update an artist's address if the caller is not an artist", async () => {
            expectCustomError(
                instance.updateArtistAddress(accounts[1], {
                    from: accounts[0],
                }),
                abi,
                "NotArtistError",
                [accounts[0]]
            );
        });
    });

    // ------------------------------------------------------------------------
    // addDelegatee
    // ------------------------------------------------------------------------

    describe("addDelegatee", () => {
        it("should be able to add a delegatee by artist", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const seriesIDs = [];

            // 1. artist register multiple series (2 series)
            for (let i = 0; i < 2; i++) {
                const tx = await instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    { from: artistAddress }
                );
                seriesIDs.push(tx.logs[0].args.seriesID);
            }

            // 2. artist add delegatee
            const tx = await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 3. check all series delegatees should include the delegatee address
            for (let i = 0; i < seriesIDs.length; i++) {
                expectAddressArraysEqual(
                    await instance.getSeriesDelegatees(seriesIDs[i]),
                    [delegateeAddress]
                );
            }

            // 4. check artist delegatees should include the delegatee address
            expectAddressArraysEqual(
                await instance.getDelegatees(artistAddress),
                [delegateeAddress]
            );

            // 5. check event emitted
            expectEvent(tx, "AddDelegatee", {
                delegator: artistAddress,
                delegatee: delegateeAddress,
            });
        });

        it("should be able to add a delegatee by a person who is not an artist", async () => {
            const someoneAddress = accounts[1];
            const delegateeAddress = accounts[2];

            // 1. someone add delegatee
            const tx = await instance.addDelegatee(delegateeAddress, {
                from: someoneAddress,
            });

            // 2. check delegatee should be the same as the delegatee address
            expectAddressArraysEqual(
                await instance.getDelegatees(someoneAddress),
                [delegateeAddress]
            );

            // 3. check event emitted
            expectEvent(tx, "AddDelegatee", {
                delegator: someoneAddress,
                delegatee: delegateeAddress,
            });
        });

        it("should NOT be able to add a delegatee if the delegatee address is zero", async () => {
            expectCustomError(
                instance.addDelegatee(zeroAddress, {
                    from: accounts[0],
                }),
                abi,
                "GenericError",
                ["delegatee address is zero"]
            );
        });

        it("should NOT be able to add a delegatee if the delegatee is already a delegatee", async () => {
            const delegatorAddress = accounts[0];
            const delegateeAddress = accounts[1];

            // 1. add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: delegatorAddress,
            });

            // 2. add delegatee again
            expectCustomError(
                instance.addDelegatee(delegateeAddress, {
                    from: delegatorAddress,
                }),
                abi,
                "DuplicateDelegateeError",
                [delegatorAddress, delegateeAddress]
            );
        });
    });

    // ------------------------------------------------------------------------
    // removeDelegatee
    // ------------------------------------------------------------------------

    describe("removeDelegatee", () => {
        it("should be able to remove a delegatee by artist", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const seriesIDs = [];

            // 1. artist register multiple series (2 series)
            for (let i = 0; i < 2; i++) {
                const tx = await instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    { from: artistAddress }
                );
                seriesIDs.push(tx.logs[0].args.seriesID);
            }

            // 2. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 3. check all series delegatees should include the delegatee address
            for (let i = 0; i < seriesIDs.length; i++) {
                expectAddressArraysEqual(
                    await instance.getSeriesDelegatees(seriesIDs[i]),
                    [delegateeAddress]
                );
            }

            // 4. check artist delegatees should include the delegatee address
            expectAddressArraysEqual(
                await instance.getDelegatees(artistAddress),
                [delegateeAddress]
            );

            // 5. artist remove delegatee
            const tx = await instance.removeDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 6. check all series delegatees should NOT include the delegatee address
            for (let i = 0; i < seriesIDs.length; i++) {
                expect(await instance.getSeriesDelegatees(seriesIDs[i])).to.be
                    .empty;
            }

            // 7. check artist delegatees should NOT include the delegatee address
            expect(await instance.getDelegatees(artistAddress)).to.be.empty;

            // 8. check event emitted
            expectEvent(tx, "RemoveDelegatee", {
                delegator: artistAddress,
                delegatee: delegateeAddress,
            });
        });

        it("should be able to remove a delegatee by someone else", async () => {
            const delegatorAddress = accounts[0];
            const delegateeAddress = accounts[1];

            // 1. delegator add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: delegatorAddress,
            });

            // 2. check delegatee address related to the delegator address
            expectAddressArraysEqual(
                await instance.getDelegatees(delegatorAddress),
                [delegateeAddress]
            );

            // 3. delegator remove delegatee
            const tx = await instance.removeDelegatee(delegateeAddress, {
                from: delegatorAddress,
            });

            // 4. check his delegatees should NOT include the delegatee address
            expect(await instance.getDelegatees(delegatorAddress)).to.be.empty;

            // 5. check event emitted
            expectEvent(tx, "RemoveDelegatee", {
                delegator: delegatorAddress,
                delegatee: delegateeAddress,
            });
        });

        it("should be able to remove a delegatee even there are some pending collaboration invitations sent by the delegatee", async () => {
            const artistAddress = accounts[0];
            const delegateeAddress = accounts[1];
            const collaboratorAddress = accounts[2];
            const seriesIDs = [];

            // 1. artist register series
            for (let i = 0; i < 2; i++) {
                const tx = await instance.registerSeries(
                    artistAddress,
                    metadataURI,
                    tokenDataURI,
                    { from: artistAddress }
                );
                seriesIDs.push(tx.logs[0].args.seriesID);
            }

            // 2. artist add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 3. delegatee create collaboration invitation
            for (let i = 0; i < seriesIDs.length; i++) {
                await instance.inviteCollaborator(
                    seriesIDs[i],
                    collaboratorAddress,
                    { from: delegateeAddress }
                );
            }

            // 4. artist remove delegatee
            const tx = await instance.removeDelegatee(delegateeAddress, {
                from: artistAddress,
            });

            // 5. check artist delegatees should NOT include the delegatee address
            expect(await instance.getDelegatees(artistAddress)).to.be.empty;

            // 6. check series delegatees should NOT include the delegatee address
            for (let i = 0; i < seriesIDs.length; i++) {
                expect(await instance.getSeriesDelegatees(seriesIDs[i])).to.be
                    .empty;
            }

            // 7. check pending collaboration invitations should be empty
            for (let i = 0; i < seriesIDs.length; i++) {
                expect(await instance.getCollaborationInvitees(seriesIDs[i])).to
                    .be.empty;
            }

            // 8. check event emitted
            expectEvent(tx, "RemoveDelegatee", {
                delegator: artistAddress,
                delegatee: delegateeAddress,
            });
        });

        it("should NOT be able to remove a delegatee if the delegatee address is zero", async () => {
            expectCustomError(
                instance.removeDelegatee(zeroAddress, {
                    from: accounts[0],
                }),
                abi,
                "GenericError",
                ["delegatee address is zero"]
            );
        });

        it("should NOT be able to remove a delegatee if the delegatee is not a delegatee", async () => {
            expectCustomError(
                instance.removeDelegatee(accounts[1], {
                    from: accounts[0],
                }),
                abi,
                "NotDelegateeError",
                [accounts[0], accounts[1]]
            );
        });
    });

    // ------------------------------------------------------------------------
    // assignSeries
    // ------------------------------------------------------------------------

    describe("assignSeries", () => {
        it("should be able to assign a series to an artist", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];
            const expectedArtistBID = new BN(2);

            // 1. artist A register series
            const registerTxA = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID1 = registerTxA.logs[0].args.seriesID;

            // 2. artist B register series
            const registerTxB = await instance.registerSeries(
                artistBAddress,
                metadataURI,
                tokenDataURI,
                { from: artistBAddress }
            );
            const seriesID2 = registerTxB.logs[0].args.seriesID;

            // 3. artist A assign series to artist B
            const assignTx = await instance.assignSeries(
                seriesID1,
                artistBAddress,
                { from: artistAAddress }
            );

            // 4. check artist A series count should be 0
            expect(
                await instance.getTotalArtistSeries(artistAAddress)
            ).to.be.a.bignumber.that.equal(new BN(0));

            // 5. check artist B series count should be 2
            expect(
                await instance.getTotalArtistSeries(artistBAddress)
            ).to.be.a.bignumber.that.equal(new BN(2));

            // 6. check series artists
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID1),
                [artistBAddress]
            );
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID2),
                [artistBAddress]
            );

            // 7. check series administrator should be artist B
            expect(
                await instance.getSeriesAdministratorID(seriesID1)
            ).to.be.a.bignumber.that.equal(expectedArtistBID);
            expect(
                await instance.getSeriesAdministratorID(seriesID2)
            ).to.be.a.bignumber.that.equal(expectedArtistBID);

            // 9. check event emitted
            expectEvent(assignTx, "AssignSeries", {
                seriesID: seriesID1,
                assignerAddress: artistAAddress,
                assigneeAddress: artistBAddress,
            });
        });

        it("should be able to assign a series to someone else that is not an artist", async () => {
            const artistAddress = accounts[0];
            const assigneeAddress = accounts[1];
            const expectedAssigneeID = new BN(2);

            // 1. artist A register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist A assign series to person B
            const assignTx = await instance.assignSeries(
                seriesID,
                assigneeAddress,
                { from: artistAddress }
            );

            // 3. check artist A series count should be 0
            expect(
                await instance.getTotalArtistSeries(artistAddress)
            ).to.be.a.bignumber.that.equal(new BN(0));

            // 4. check series artists should include person B
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID),
                [assigneeAddress]
            );

            // 5. check series administrator should be person B
            expect(
                await instance.getSeriesAdministratorID(seriesID)
            ).to.be.a.bignumber.that.equal(expectedAssigneeID);

            // 6. check event emitted
            expectEvent(assignTx, "AssignSeries", {
                seriesID: seriesID,
                assignerAddress: artistAddress,
                assigneeAddress: assigneeAddress,
            });
        });

        it("should be able to assign a series to an artist that has delegatees", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];
            const delegateeAddress = accounts[2];

            // 1. artist A register series
            const registerTxA = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID1 = registerTxA.logs[0].args.seriesID;

            // 2. artist B add delegatee
            await instance.addDelegatee(delegateeAddress, {
                from: artistBAddress,
            });

            // 3. artist B register series
            const registerTxB = await instance.registerSeries(
                artistBAddress,
                metadataURI,
                tokenDataURI,
                { from: artistBAddress }
            );
            const seriesID2 = registerTxB.logs[0].args.seriesID;

            // 4. artist A assign series to artist B
            const assignTx = await instance.assignSeries(
                seriesID1,
                artistBAddress,
                { from: artistAAddress }
            );

            // 5. check artist A series count should be 0
            expect(
                await instance.getTotalArtistSeries(artistAAddress)
            ).to.be.a.bignumber.that.equal(new BN(0));

            // 6. check artist B series count should be 2
            expect(
                await instance.getTotalArtistSeries(artistBAddress)
            ).to.be.a.bignumber.that.equal(new BN(2));

            // 7. check series artists should NOT include artist A
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID1),
                [artistBAddress]
            );
            expectAddressArraysEqual(
                await instance.getSeriesArtistAddresses(seriesID2),
                [artistBAddress]
            );

            // 8. check series delegatees should include the artist B's delegatee address
            expectAddressArraysEqual(
                await instance.getSeriesDelegatees(seriesID1),
                [delegateeAddress]
            );
            expectAddressArraysEqual(
                await instance.getSeriesDelegatees(seriesID2),
                [delegateeAddress]
            );

            // 9. check artist delegatees should be the same as before
            expect(await instance.getDelegatees(artistAAddress)).to.be.empty;
            expectAddressArraysEqual(
                await instance.getDelegatees(artistBAddress),
                [delegateeAddress]
            );

            // 10. check event emitted
            expectEvent(assignTx, "AssignSeries", {
                seriesID: seriesID1,
                assignerAddress: artistAAddress,
                assigneeAddress: artistBAddress,
            });
        });

        it("should NOT be able to assign a series to an artist that is already a collaborator", async () => {
            const artistAddress = accounts[0];
            const collaboratorAddress = accounts[1];

            // 1. artist A register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist A invite artist B to series
            await instance.inviteCollaborator(seriesID, collaboratorAddress, {
                from: artistAddress,
            });

            // 3. artist B opt in to series
            await instance.optInCollaboration(seriesID, {
                from: collaboratorAddress,
            });

            // 4. artist A assign series to artist B
            expectCustomError(
                instance.assignSeries(seriesID, collaboratorAddress, {
                    from: artistAddress,
                }),
                abi,
                "ArtistExistsError",
                [collaboratorAddress]
            );
        });

        it("should NOT be able to assign a series to the same address", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist assign series to himself
            expectCustomError(
                instance.assignSeries(seriesID, artistAddress, {
                    from: artistAddress,
                }),
                abi,
                "SameAddressError",
                [artistAddress]
            );
        });

        it("should NOT be able to assign a series by an artist that does not exist", async () => {
            const artistAddress = accounts[0];
            const assigneeAddress = accounts[1];
            const assignerAddress = accounts[2];

            // 1. artist A register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist A assign series to himself
            expectCustomError(
                instance.assignSeries(seriesID, assigneeAddress, {
                    from: assignerAddress,
                }),
                abi,
                "NotArtistError",
                [assignerAddress]
            );
        });

        it("should NOT be able to assign a series to a zero address", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist assign series to zero address
            expectCustomError(
                instance.assignSeries(seriesID, zeroAddress, {
                    from: artistAddress,
                }),
                abi,
                "GenericError",
                ["assignee address is zero"]
            );
        });
    });

    // ------------------------------------------------------------------------
    // assignAdministrator
    // ------------------------------------------------------------------------

    describe("assignAdministrator", () => {
        it("should be able to assign an administrator to a series", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];

            // 1. artist A register series
            const registerTxA = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID = registerTxA.logs[0].args.seriesID;

            // 2. artist A invite artist B to series
            await instance.inviteCollaborator(seriesID, artistBAddress, {
                from: artistAAddress,
            });

            // 3. artist B opt in to series
            await instance.optInCollaboration(seriesID, {
                from: artistBAddress,
            });

            // 4. artist A assign administrator to artist B
            const assignTx = await instance.assignAdministrator(
                seriesID,
                artistBAddress,
                { from: artistAAddress }
            );

            // 5. check series administrator should be artist B
            expect(
                await instance.getSeriesAdministratorAddress(seriesID)
            ).to.equal(artistBAddress);

            // 6. check event emitted
            expectEvent(assignTx, "AssignAdministrator", {
                seriesID: seriesID,
                assignerAddress: artistAAddress,
                assigneeAddress: artistBAddress,
            });
        });

        it("should NOT be able to assign an administrator to zero address", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist assign administrator to zero address
            expectCustomError(
                instance.assignAdministrator(seriesID, zeroAddress, {
                    from: artistAddress,
                }),
                abi,
                "GenericError",
                ["administrator address is zero"]
            );
        });
        it("should NOT be able to assign an administrator to himself", async () => {
            const artistAddress = accounts[0];

            // 1. artist register series
            const registerTx = await instance.registerSeries(
                artistAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAddress }
            );
            const seriesID = registerTx.logs[0].args.seriesID;

            // 2. artist assign administrator to himself
            expectCustomError(
                instance.assignAdministrator(seriesID, artistAddress, {
                    from: artistAddress,
                }),
                abi,
                "AdministratorCannotSelfAssignError",
                [artistAddress]
            );
        });

        it("should NOT be able to assign an administrator to an artist that is not in the series", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];

            // 1. artist A register series
            const registerTxA = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID1 = registerTxA.logs[0].args.seriesID;

            // 2. artist B register series
            const registerTxB = await instance.registerSeries(
                artistBAddress,
                metadataURI,
                tokenDataURI,
                { from: artistBAddress }
            );
            const seriesID2 = registerTxB.logs[0].args.seriesID;

            // 3. artist A assign administrator to artist B
            expectCustomError(
                instance.assignAdministrator(seriesID1, artistBAddress, {
                    from: artistAAddress,
                }),
                abi,
                "NotSeriesArtistError",
                [seriesID1, artistBAddress]
            );
        });

        it("should NOT be able to assign an administrator if caller is not the administrator", async () => {
            const artistAAddress = accounts[0];
            const artistBAddress = accounts[1];
            const assignerAddress = accounts[2];

            // 1. artist A register series
            const registerTxA = await instance.registerSeries(
                artistAAddress,
                metadataURI,
                tokenDataURI,
                { from: artistAAddress }
            );
            const seriesID = registerTxA.logs[0].args.seriesID;

            // 2. artist A assign administrator to artist B
            expectCustomError(
                instance.assignAdministrator(seriesID, artistBAddress, {
                    from: assignerAddress,
                }),
                abi,
                "NotAdministratorError",
                [seriesID, assignerAddress]
            );
        });

        it("should NOT be able to assign an administrator for unknown series", async () => {
            const seriesID = new BN(1);
            expectCustomError(
                instance.assignAdministrator(seriesID, accounts[1], {
                    from: accounts[0],
                }),
                abi,
                "SeriesNotExistsError",
                [seriesID]
            );
        });
    });
});
