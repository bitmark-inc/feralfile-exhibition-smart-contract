# SPDX-License-Identifier: ISC
# Copyright (c) 2019-2023 Bitmark Inc.
# Use of this source code is governed by an ISC
# license that can be found in the LICENSE file.

.PHONY:

truffle-check: ; @which truffle > /dev/null
abigen-check: ; @which abigen > /dev/null
jq-check: ; @which jq >/dev/null
npm-check: ;@which npm >/dev/null
sol-merger-check: ;@which sol-merger >/dev/null

check: truffle-check abigen-check jq-check npm-check

clean:
	rm -rf **/abi.go

merge: sol-merger-check
	rm -rf /tmp/sol-merger && rm -rf ./contracts/sol-merger && \
	mkdir /tmp/sol-merger && mkdir ./contracts/sol-merger && \
	sol-merger --export-plugin SPDXLicenseRemovePlugin ./contracts/FeralfileArtworkV2.sol /tmp/sol-merger && \
    sol-merger --export-plugin SPDXLicenseRemovePlugin ./contracts/FeralfileArtworkV3.sol /tmp/sol-merger && \
    sol-merger --export-plugin SPDXLicenseRemovePlugin ./contracts/FeralfileArtworkV4.sol /tmp/sol-merger && \
	sol-merger --export-plugin SPDXLicenseRemovePlugin ./contracts/FeralfileEnglishAuction.sol /tmp/sol-merger && \
	sol-merger --export-plugin SPDXLicenseRemovePlugin ./contracts/FeralFileAirdropV1.sol /tmp/sol-merger && \
	sol-merger --export-plugin SPDXLicenseRemovePlugin ./contracts/OwnerData.sol /tmp/sol-merger && \
	mv /tmp/sol-merger/FeralfileArtworkV2.sol ./contracts/sol-merger/FeralfileExhibitionV2.sol && \
    mv /tmp/sol-merger/FeralfileArtworkV3.sol ./contracts/sol-merger/FeralfileExhibitionV3.sol && \
    mv /tmp/sol-merger/FeralfileArtworkV4.sol ./contracts/sol-merger/FeralfileExhibitionV4.sol && \
	mv /tmp/sol-merger/FeralfileEnglishAuction.sol ./contracts/sol-merger/FeralfileEnglishAuction.sol && \
	mv /tmp/sol-merger/FeralFileAirdropV1.sol ./contracts/sol-merger/FeralFileAirdropV1.sol && \
	mv /tmp/sol-merger/OwnerData.sol ./contracts/sol-merger/OwnerData.sol

build-contract: check
	npm install && \
	truffle compile && \
	jq -r ".bytecode" build/contracts/FeralfileExhibitionV2.json > ./build/FeralfileExhibitionV2.bin && \
	jq -r ".abi" build/contracts/FeralfileExhibitionV2.json > ./build/FeralfileExhibitionV2.abi && \
	jq -r ".bytecode" build/contracts/FeralfileExhibitionV3.json > ./build/FeralfileExhibitionV3.bin && \
	jq -r ".abi" build/contracts/FeralfileExhibitionV3.json > ./build/FeralfileExhibitionV3.abi && \
	jq -r ".bytecode" build/contracts/FeralfileExhibitionV4.json > ./build/FeralfileExhibitionV4.bin && \
	jq -r ".abi" build/contracts/FeralfileExhibitionV4.json > ./build/FeralfileExhibitionV4.abi && \
	jq -r ".bytecode" build/contracts/FeralfileEnglishAuction.json > ./build/FeralfileEnglishAuction.bin && \
	jq -r ".abi" build/contracts/FeralfileEnglishAuction.json > ./build/FeralfileEnglishAuction.abi && \
	jq -r ".bytecode" build/contracts/FeralFileAirdropV1.json > ./build/FeralFileAirdropV1.bin && \
	jq -r ".abi" build/contracts/FeralFileAirdropV1.json > ./build/FeralFileAirdropV1.abi && \
	jq -r ".bytecode" build/contracts/OwnerData.json > ./build/OwnerData.bin && \
	jq -r ".abi" build/contracts/OwnerData.json > ./build/OwnerData.abi

build: build-contract
	mkdir -p ./go-binding/feralfile-exhibition-v2 && \
	mkdir -p ./go-binding/feralfile-exhibition-v3 && \
	mkdir -p ./go-binding/feralfile-exhibition-v4 && \
	mkdir -p ./go-binding/feralfile-english-auction && \
	mkdir -p ./go-binding/feralfile-airdrop-v1 && \
	mkdir -p ./go-binding/owner-data && \
	abigen --abi ./build/FeralfileExhibitionV2.abi --bin ./build/FeralfileExhibitionV2.bin --pkg feralfilev2 -type FeralfileExhibitionV2 --out ./go-binding/feralfile-exhibition-v2/abi.go
	abigen --abi ./build/FeralfileExhibitionV3.abi --bin ./build/FeralfileExhibitionV3.bin --pkg feralfilev3 -type FeralfileExhibitionV3 --out ./go-binding/feralfile-exhibition-v3/abi.go
	abigen --abi ./build/FeralfileExhibitionV4.abi --bin ./build/FeralfileExhibitionV4.bin --pkg feralfilev4 -type FeralfileExhibitionV4 --out ./go-binding/feralfile-exhibition-v4/abi.go
	abigen --abi ./build/FeralfileEnglishAuction.abi --bin ./build/FeralfileEnglishAuction.bin --pkg english_auction -type FeralfileEnglishAuction --out ./go-binding/feralfile-english-auction/abi.go
	abigen --abi ./build/FeralFileAirdropV1.abi --bin ./build/FeralFileAirdropV1.bin --pkg airdropv1 -type FeralFileAirdropV1 --out ./go-binding/feralfile-airdrop-v1/abi.go && \
	abigen --abi ./build/OwnerData.abi --bin ./build/OwnerData.bin --pkg ownerdata -type OwnerData --out ./go-binding/owner-data/abi.go
