module.exports = {
  publicPath: process.env.NODE_ENV === 'production'
      ? ''
      : '/',
  outputDir: '../resources/app/',
  "transpileDependencies": [
    "vuetify"
  ],
  "css": {
    "loaderOptions": {
      "scss": {
        "prependData": "@import \"~@/styles/_global.scss\";"
      }
    }
  },
}