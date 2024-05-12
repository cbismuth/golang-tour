GOLANG_VERSION = "1.22.2"

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# https://github.com/bazelbuild/rules_go/tags
RULES_GO_VERSION = "v0.47.0"

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "af47f30e9cbd70ae34e49866e201b3f77069abb111183f2c0297e7e74ba6bbc0",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/{0}/rules_go-{0}.zip".format(RULES_GO_VERSION),
    ],
)

# https://github.com/bazelbuild/bazel-gazelle/tags
GAZELLE_VERSION = "v0.36.0"

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-dd8ojEsxyB61D1Hi4U9HY8t1SNquEmgXJHBkY3/Z6mI=",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{0}/bazel-gazelle-{0}.tar.gz".format(GAZELLE_VERSION),
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = GOLANG_VERSION)

gazelle_dependencies()

# https://github.com/protocolbuffers/protobuf/tags
PROTOBUF_VERSION = "25.3"

http_archive(
    name = "com_google_protobuf",
    sha256 = "d19643d265b978383352b3143f04c0641eea75a75235c111cc01a1350173180e",
    strip_prefix = "protobuf-{0}".format(PROTOBUF_VERSION),
    urls = [
        "https://github.com/protocolbuffers/protobuf/releases/download/v{0}/protobuf-{0}.tar.gz".format(PROTOBUF_VERSION),
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# https://github.com/bazelbuild/buildtools/tags
BUILDTOOLS_VERSION = "v7.1.1"

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "ae34c344514e08c23e90da0e2d6cb700fcd28e80c02e23e4d5715dddcb42f7b3",
    urls = [
        "https://github.com/bazelbuild/buildtools/releases/download/{0}/buildtools-{0}.tar.gz".format(BUILDTOOLS_VERSION),
    ],
)
