#!/bin/bash

set -euo pipefail

commit_sha=$(git rev-parse HEAD)
echo "COMMIT_SHA $commit_sha"

git_branch=$(git rev-parse --abbrev-ref HEAD)
echo "GIT_BRANCH $git_branch"

git_tree_status=$(git diff-index --quiet HEAD -- && echo 'Clean' || echo 'Modified')
echo "GIT_TREE_STATUS $git_tree_status"

SHA256SUM="sha256sum"
if command -v shasum &> /dev/null; then
    # Mac OS ships shasum instead of sha256sum.
    SHA256SUM="shasum -a 256"
fi

ekglue_version="v0.0.70-pre.g$(git describe --match=ekglue-* --long --always --dirty=".$(git diff HEAD | $SHA256SUM | cut -c 1-10)")"
echo "EKGLUE_VERSION ${ekglue_version}"
