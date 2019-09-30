load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/go-cqrses/teddyshop
# gazelle:exclude hack
# gazelle:exclude third_party

gazelle(
    name = "gazelle",
)

## Buildifier to test WORKSPACE and BUILD(.bazel) files.

load("@com_github_bazelbuild_buildtools//buildifier:def.bzl", "buildifier")

buildifier(
    name = "buildifier",
    diff_command = "diff -u",
    mode = "diff",
    verbose = True,
)

buildifier(
    name = "buildifier.fix",
    mode = "fix",
)
