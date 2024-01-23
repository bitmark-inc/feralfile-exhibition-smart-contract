# SPDX-License-Identifier: ISC
# Copyright (c) 2019-2023 Bitmark Inc.
# Use of this source code is governed by an ISC
# license that can be found in the LICENSE file.

.PHONY:

truffle-check: ; @which truffle > /dev/null
abigen-check: ; @which abigen > /dev/null
jq-check: ; @which jq >/dev/null
npm-check: ;@which npm >/dev/null

check: truffle-check abigen-check jq-check npm-check

clean:
	rm -rf **/abi.go

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
	jq -r ".abi" build/contracts/FeralFileAirdropV1.json > ./build/FeralFileAirdropV1.abi

build: build-contract
	mkdir -p ./binding/feralfile-exhibition-v2 && \
	mkdir -p ./binding/feralfile-exhibition-v3 && \
	mkdir -p ./binding/feralfile-exhibition-v4 && \
	mkdir -p ./binding/feralfile-english-auction && \
	mkdir -p ./binding/feralfile-airdrop-v1 && \
	abigen --abi ./build/FeralfileExhibitionV2.abi --bin ./build/FeralfileExhibitionV2.bin --pkg feralfilev2 -type FeralfileExhibitionV2 --out ./binding/feralfile-exhibition-v2/abi.go
	abigen --abi ./build/FeralfileExhibitionV3.abi --bin ./build/FeralfileExhibitionV3.bin --pkg feralfilev3 -type FeralfileExhibitionV3 --out ./binding/feralfile-exhibition-v3/abi.go
	abigen --abi ./build/FeralfileExhibitionV4.abi --bin ./build/FeralfileExhibitionV4.bin --pkg feralfilev4 -type FeralfileExhibitionV4 --out ./binding/feralfile-exhibition-v4/abi.go
	abigen --abi ./build/FeralfileEnglishAuction.abi --bin ./build/FeralfileEnglishAuction.bin --pkg english_auction -type FeralfileEnglishAuction --out ./binding/feralfile-english-auction/abi.go
	abigen --abi ./build/FeralFileAirdropV1.abi --bin ./build/FeralFileAirdropV1.bin --pkg airdropv1 -type FeralFileAirdropV1 --out ./binding/feralfile-airdrop-v1/abi.go
