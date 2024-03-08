#!/bin/bash

# Fetch target branch
TARGET_BRANCH=$1

# Fetch changes
git fetch origin $TARGET_BRANCH
CHANGED_FILES=$(git diff --name-only FETCH_HEAD $(git merge-base FETCH_HEAD $TARGET_BRANCH) | grep '.sol$')
echo "Changed files: $CHANGED_FILES"

# Read the mapping file
declare -A MAPPING

# Use jq to iterate over each key-value pair in the JSON object.
while IFS= read -r line; do
  contract=$(jq -r '.contract' <<< "$line")
  IFS=$'\n' read -r -d '' -a tests < <(jq -r '.tests[]' <<< "$line" && printf '\0')
  
  # Add each test to the mapping for the contract.
  for test in "${tests[@]}"; do
    MAPPING["$contract"]+="$test "
  done
done < <(jq -c 'to_entries[] | {contract: .key, tests: .value}' .github/scripts/contract-test-mapping.json)


# Initialize a variable to hold test files to run
TEST_FILES=""

# Loop through changed files and append relevant test files
for FILE in $CHANGED_FILES; do
  if [[ -n "${MAPPING[$FILE]}" ]]; then
    TEST_FILES+="${MAPPING[$FILE]}"
  fi
done

# Remove duplicate test files
TEST_FILES=$(echo "$TEST_FILES" | tr ' ' '\n' | sort -u | tr '\n' ' ')

# Run Truffle tests if there are any test files to run
if [ -n "$TEST_FILES" ]; then
  echo "Running tests for changed contracts:"
  echo $TEST_FILES
  truffle test $TEST_FILES | tee test-output.txt
else
  echo "No contract changes detected."
fi