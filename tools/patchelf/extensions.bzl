"""Patchelf binaries."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _deps_impl(_ctx):
    """https://github.com/NixOS/patchelf"""
    build_file_content = """exports_files(["bin/patchelf"])
    """
    http_archive(
        name = "patchelf_amd64_linux",
        url = "https://github.com/NixOS/patchelf/releases/download/0.18.0/patchelf-0.18.0-x86_64.tar.gz",
        build_file_content = build_file_content,
        sha256 = "ce84f2447fb7a8679e58bc54a20dc2b01b37b5802e12c57eece772a6f14bf3f0",
    )
    http_archive(
        name = "patchelf_arm64_linux",
        url = "https://github.com/NixOS/patchelf/releases/download/0.18.0/patchelf-0.18.0-aarch64.tar.gz",
        build_file_content = build_file_content,
        sha256 = "ae13e2effe077e829be759182396b931d8f85cfb9cfe9d49385516ea367ef7b2",
    )
    return _ctx.extension_metadata(
        root_module_direct_deps = ["patchelf_amd64_linux", "patchelf_arm64_linux"],
        root_module_direct_dev_deps = [],
        reproducible = True,
    )

patchelf = module_extension(
    implementation = _deps_impl,
)
