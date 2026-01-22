module.exports = {
  root: true,
  env: { browser: true, node: true, es2022: true },
  extends: ['eslint:recommended', 'plugin:vue/vue3-recommended', 'prettier'],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
};
