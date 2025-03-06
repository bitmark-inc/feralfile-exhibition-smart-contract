const { expect } = require("chai");

function expectBigNumberArraysEqual(actual, expected) {
    expect(actual.length).to.equal(
        expected.length,
        `Expected array length is ${expected.length} but actual is ${actual.length}`
    );

    for (let i = 0; i < actual.length; i++) {
        expect(actual[i]).to.be.a.bignumber.that.equals(
            expected[i],
            `Expected value at index ${i} is ${expected[i]} but actual is ${actual[i]}`
        );
    }
}

function expectAddressArraysEqual(actual, expected) {
    expect(actual.length).to.equal(
        expected.length,
        `Expected array length is ${expected.length} but actual is ${actual.length}`
    );

    for (let i = 0; i < actual.length; i++) {
        const normalizedActual = web3.utils.toChecksumAddress(actual[i]);
        const normalizedExpected = web3.utils.toChecksumAddress(expected[i]);
        expect(normalizedActual).to.equal(
            normalizedExpected,
            `Expected value at index ${i} is ${normalizedExpected} but actual is ${normalizedActual}`
        );
    }
}

module.exports = {
    expectBigNumberArraysEqual,
    expectAddressArraysEqual,
};
