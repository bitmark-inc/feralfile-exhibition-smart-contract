
/**
 * Generic helper: expects a revert with a specific custom error name.
 * It automatically finds that error in the ABI, checks the name, and
 * optionally decodes + checks parameter values.
 */
async function expectCustomError(
    promise,
    contractABI,
    expectedErrorName,
    expectedArgs = []
) {
    try {
        await promise;
        expect.fail("Expected revert not received");
    } catch (error) {
        // If we cannot find 'error.data', we can't decode any further
        if (!error.data) {
            expect.fail(`No error.data found. Got error:\n${error}`);
        }

        // The first 4 bytes (8 hex chars after '0x') is the selector
        const selector = error.data.result.slice(2, 10);

        // Try to find which custom error in the ABI matches this selector
        const errorEntry = findErrorBySelector(contractABI, selector);
        if (!errorEntry) {
            expect.fail(
                `Unrecognized custom error (selector = 0x${selector}). Expected "${expectedErrorName}".`
            );
        }

        // Confirm the error name
        const actualErrorName = errorEntry.name;
        expect(
            actualErrorName,
            `Expected custom error "${expectedErrorName}" but got "${actualErrorName}"`
        ).to.equal(expectedErrorName);

        // Decode the parameters
        const encodedParams = error.data.result.slice(10); // skip the selector
        const paramTypes = errorEntry.inputs.map((input) => input.type);
        const decoded = web3.eth.abi.decodeParameters(
            paramTypes,
            encodedParams
        );

        // Compare the decoded parameters if user provided expectedArgs
        if (expectedArgs.length > 0) {
            expect(
                errorEntry.inputs.length,
                "Mismatch in number of expected args vs. the error definition"
            ).to.equal(expectedArgs.length);

            for (let i = 0; i < expectedArgs.length; i++) {
                // Param name from the ABI
                const paramName = errorEntry.inputs[i].name;
                const actualVal = decoded[i];
                const expectedVal = expectedArgs[i];
                expect(
                    actualVal.toString(),
                    `Parameter mismatch for "${paramName}" at index [${i}]`
                ).to.equal(expectedVal.toString());
            }
        }
    }
}

/**
 * Searches the contract ABI for a custom error whose selector matches `selectorHex`.
 */
function findErrorBySelector(contractABI, selectorHex) {
    for (let entry of contractABI) {
        if (entry.type === "error") {
            // Construct the signature: e.g. "MyError(uint256,address)"
            const signature = `${entry.name}(${entry.inputs
                .map((i) => i.type)
                .join(",")})`;
            // Take keccak256(signature), compare first 4 bytes
            const hash = web3.utils.sha3(signature);
            const abiSelector = hash.slice(2, 10); // skip "0x" and take 8 chars
            if (abiSelector === selectorHex) {
                return entry;
            }
        }
    }
    return null;
}

module.exports = {
    expectCustomError
};