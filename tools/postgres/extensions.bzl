"""Postgres binaries.

This only handles MacOS.  Linux binaries are handled via rules_distroless.
"""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

def _deps_impl(module_ctx):
    """Postgres binaries for Mac OS (linux binaries are from rules_distroless + apt)."""
    http_archive(
        name = "com_enterprisedb_get_postgresql_macos",
        url = "https://get.enterprisedb.com/postgresql/postgresql-17.2-1-osx-binaries.zip",
        integrity = "sha256-RcNIQv9TlB52C0PuJMMUfhY5NuuVWFXvBSGHRaMIX40=",
        build_file_content = """exports_files(glob(["**"]))""",
        strip_prefix = "pgsql",
    )
    return module_ctx.extension_metadata(
        root_module_direct_deps = ["com_enterprisedb_get_postgresql_macos"],
        root_module_direct_dev_deps = [],
        reproducible = True,
    )

postgres = module_extension(
    implementation = _deps_impl,
)
