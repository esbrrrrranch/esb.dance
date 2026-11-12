export default (eleventyConfig) => {
  eleventyConfig.addPassthroughCopy("www/assets");
  eleventyConfig.addPassthroughCopy("www/.nojekyll");
  eleventyConfig.addWatchTarget("templates/");

  return {
    dir: {
      includes: "../templates",
      input: "www",
      output: "dist",
    },
    htmlTemplateEngine: "njk",
    templateFormats: ["html"],
  };
};
