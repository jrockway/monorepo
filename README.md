# jrockway's monorepo

I'm tired of maintaining a separate repo for each of my projects. So I'm gradually moving them into
one.

## Building

Install [Bazelisk](https://github.com/bazelbuild/bazelisk), then:

    bazel test ...

I put this in my `.bazelrc`:

    test --test_output=errors
    startup --max_idle_secs=604800
    build --disk_cache=/tmp/bazel
