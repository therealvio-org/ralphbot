module.exports = {
  extends: ["prettier", "plugin:import/typescript"],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    sourceType: "module",
    project: "./tsconfig.json",
  },
  plugins: ["@typescript-eslint", "import", "sort-imports-es6-autofix"],
  rules: {
    "@typescript-eslint/no-unused-vars": ["error", { argsIgnorePattern: "^_" }],
    "import/no-extraneous-dependencies": ["warn"],
    "sort-imports-es6-autofix/sort-imports-es6": ["error"],
  },
  root: true,
}
