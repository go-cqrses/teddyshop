#!/usr/bin/env sh

bazel run //:gazelle -- update-repos -from_file=./go.mod -to_macro "go_vendor.bzl%go_repositories"
