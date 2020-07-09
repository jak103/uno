module.exports = {
  "moduleFileExtensions": [
    "js",
    "jsx",
    "json",
    "vue"
  ],
  "transform": {
    "^.+\\.vue$": "vue-jest",
    "^.+\\.jsx?$": "babel-jest"
  },
  "transformIgnorePatterns": [
    "<rootDir>/(node_modules)/"
  ],
  "moduleNameMapper": {
    "^@/(.*)$": "<rootDir>/src/$1"
  }
}
