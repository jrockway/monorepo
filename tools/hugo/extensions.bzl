"""Hugo binaries."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _deps_impl(_ctx):
    """https://github.com/gohugoio/hugo/releases"""
    build_file_content = """exports_files(["hugo"])
    """
    print(build_file_content)
    http_archive(
        name = "hugo_amd64_linux",
        url = "https://github.com/gohugoio/hugo/releases/download/v0.139.2/hugo_0.139.2_linux-amd64.tar.gz",
        build_file_content = build_file_content,
        sha256 = "6b22df5394f3e52726e173f6cdc97a3b45a7ab3c82912fb85a2eb635495f6af1",
    )
    http_archive(
        name = "hugo_arm64_linux",
        url = "https://github.com/gohugoio/hugo/releases/download/v0.139.2/hugo_0.139.2_linux-arm64.tar.gz",
        build_file_content = build_file_content,
        sha256 = "b1f865f8aa34131cc786916109d24f835a2f526a4b5e1e348bb3172a8c3f2828",
    )
    http_archive(
        name = "hugo_macos",
        url = "https://github.com/gohugoio/hugo/releases/download/v0.139.2/hugo_0.139.2_darwin-universal.tar.gz",
        build_file_content = build_file_content,
        sha256 = "4d1a465d21acf1e284dd8f460a7109351082c231bd80e3b272033c82f2379642",
    )
    return _ctx.extension_metadata(
        root_module_direct_deps = ["hugo_amd64_linux", "hugo_arm64_linux", "hugo_macos"],
        root_module_direct_dev_deps = [],
    )

hugo = module_extension(
    implementation = _deps_impl,
)
