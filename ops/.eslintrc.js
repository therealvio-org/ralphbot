module.exports = {
  extends: ["prettier"],
  parser: "@typescript-eslint/parser",
  plugins: ["@typescript-eslint", "import", "sort-imports-es6-autofix"],
  rules: {
    "@typescript-eslint/no-unused-vars": ["error", { argsIgnorePattern: "^_" }],
    "import/no-extraneous-dependencies": ["warn"],
    "sort-imports-es6-autofix/sort-imports-es6": ["error"],
  },
  parserOptions: {
    sourceType: "module",
  },
  root: true,
}
