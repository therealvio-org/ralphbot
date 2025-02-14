import { FlatCompat } from "@eslint/eslintrc"
import { fileURLToPath } from "node:url"
import { fixupConfigRules, fixupPluginRules } from "@eslint/compat"
import _import from "eslint-plugin-import"
import js from "@eslint/js"
import path from "node:path"
import sortImportsEs6Autofix from "eslint-plugin-sort-imports-es6-autofix"
import tsParser from "@typescript-eslint/parser"
import typescriptEslint from "@typescript-eslint/eslint-plugin"

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const compat = new FlatCompat({
  baseDirectory: __dirname,
  recommendedConfig: js.configs.recommended,
  allConfig: js.configs.all,
})

export default [
  {
    files: ["**/*.ts", "**/*.tsx"],
  },
  {
    ignores: ["**/*.js", "**/node_modules", "**/*.mjs"],
  },
  ...fixupConfigRules(compat.extends("prettier", "plugin:import/typescript")),
  {
    plugins: {
      "@typescript-eslint": typescriptEslint,
      import: fixupPluginRules(_import),
      "sort-imports-es6-autofix": sortImportsEs6Autofix,
    },

    languageOptions: {
      parser: tsParser,
      ecmaVersion: 5,
      sourceType: "module",

      parserOptions: {
        project: "./tsconfig.json",
      },
    },

    rules: {
      "@typescript-eslint/no-unused-vars": [
        "error",
        {
          argsIgnorePattern: "^_",
        },
      ],

      "import/no-extraneous-dependencies": ["warn"],
      "sort-imports-es6-autofix/sort-imports-es6": ["error"],
    },
  },
]
