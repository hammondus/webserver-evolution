#!/bin/bash

diff_versions() {
  local old="$1"
  local new="$2"

# The various version use the same testdata
rsync -a ./testdata $old
rsync -a ./testdata $new

  git diff --no-index "$old" "$new" \
    ':(exclude)**/.DS_Store' \
    ':(exclude)README.md' > "${old%%-*}-${new%%-*}.diff"
}

diff_versions "v1-super-simple"    "v2-prevent-dot-dir"
diff_versions "v2-prevent-dot-dir" "v3-no-escape"
diff_versions "v3-no-escape"       "v4-testing"