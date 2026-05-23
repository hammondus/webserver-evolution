#!/bin/bash

# old="v1-super-simple"
# new="v2-prevent-dot-dir"

# git diff --no-index "$old" "$new" \
#   ':(exclude)**/.DS_Store' \
#   ':(exclude)README.md' > $new/v1-v2.diff

# old="v2-prevent-dot-dir"
# new="v3-no-escape"

# git diff --no-index "$old" "$new" \
#   ':(exclude)**/.DS_Store' \
#   ':(exclude)README.md' > $new/v2-v3.diff


#!/bin/bash

diff_versions() {
  local old="$1"
  local new="$2"
  git diff --no-index "$old" "$new" \
    ':(exclude)**/.DS_Store' \
    ':(exclude)README.md' > "${old%%-*}-${new%%-*}.diff"
}

diff_versions "v1-super-simple"    "v2-prevent-dot-dir"
diff_versions "v2-prevent-dot-dir" "v3-no-escape"
diff_versions "v3-no-escape" "v4-testing"