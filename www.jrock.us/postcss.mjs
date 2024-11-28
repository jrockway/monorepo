import { default as postCssPlugin } from "@deanc/esbuild-plugin-postcss";
import { default as postcssImport } from "postcss-import";
import { default as postcssMixins } from "postcss-mixins";
import { default as postcssNested } from "postcss-nested";
import { default as postcssPresetEnv } from "postcss-preset-env";
import { default as postcssColor } from "postcss-color-mod-function";
import { default as postcssUrl } from "postcss-url";
import { default as cssnano } from "cssnano";

export default {
  assetNames: "[name]",
  plugins: [
    postCssPlugin({
      plugins: [
        postcssImport,
        postcssMixins,
        postcssNested,
        postcssPresetEnv({
          stage: 1,
          preserve: true,
          features: { "custom-properties": true },
        }),
        postcssColor,
        postcssUrl,
        cssnano,
      ],
    }),
  ],
  loader: {
    ".png": "file",
    ".woff": "file",
    ".woff2": "file",
    ".eot": "file",
    ".ttf": "file",
    ".svg": "file",
  },
};
