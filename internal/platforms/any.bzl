"""any.bzl contains a rule for making a rule's output platform-independent."""

def _any_transition_impl(settings, attr):
    return {
        "//command_line_option:platforms": ["//:any"],
    }

_any_transition = transition(
    implementation = _any_transition_impl,
    inputs = [],
    outputs = [
        "//command_line_option:platforms",
    ],
)

def _platform_independent_output_impl(ctx):
    return DefaultInfo(files = depset(ctx.files.src))

platform_independent_output = rule(
    implementation = _platform_independent_output_impl,
    attrs = {
        "src": attr.label(cfg = _any_transition),
        "_allowlist_function_transition": attr.label(
            default = "@bazel_tools//tools/allowlists/function_transition_allowlist",
        ),
    },
)
