#! /usr/bin/env bash

COMMIT_RANGE=${COMMIT_RANGE:-$(git merge-base origin/master HEAD)".."}

cd "$(git rev-parse --show-toplevel)"

files=()
for file in $(git diff --name-only ${COMMIT_RANGE} ); do
  IFS=$'\n' read -r -a files <<< "$(bazel query $file)"
  bazel query $file
done

buildables=$(bazel query --keep_going --noshow_progress "kind(.*_binary, rdeps(//..., set(${files[*]})))")

if [[ ! -z $buildables ]]; then
  echo "Building binaries"
  bazel build $buildables
fi

tests=$(bazel query --keep_going --noshow_progress "kind(test, rdeps(//..., set(${files[*]}))) except attr('tags', 'manual', //...)")

if [[ ! -z $tests ]]; then
  echo "Running tests"
  bazel test $tests
fi
