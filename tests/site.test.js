import assert from "node:assert/strict";
import { readFile, readdir } from "node:fs/promises";
import test from "node:test";

const readOutput = (path) => readFile(`dist/${path}`, "utf8");

test("pages preserve routes, canonical URLs, navigation, and subscriptions", async () => {
  const routes = [
    "",
    "bio",
    "choreography",
    "collaborators",
    "program",
    "teaching",
  ];
  for (const route of routes) {
    const url = route ? `/${route}/` : "/";
    const html = await readOutput(route ? `${route}/index.html` : "index.html");
    assert.ok(html.includes(`href="https://esb.dance${url}"`));
    assert.ok(html.includes('class="subscribe-form"'));
    assert.equal(html.includes('class="has-hero"'), route === "choreography");
    assert.equal(
      (html.match(/aria-current="page"/g) || []).length,
      route === "program" ? 0 : 1,
    );
    if (route !== "program") {
      assert.match(html, new RegExp(`href="${url}"\\s+aria-current="page"`));
    }
    assert.doesNotMatch(html, /\{\{|\{%/);
  }
});

test("404 and legacy redirect retain their special behavior", async () => {
  const notFound = await readOutput("404.html");
  assert.doesNotMatch(notFound, /rel="canonical"|class="subscribe-form"/);
  const redirect = await readOutput("emily-schoen-branch/index.html");
  assert.match(redirect, /content="0; url=\/" http-equiv="refresh"/);
  assert.doesNotMatch(redirect, /class="site-header"/);
  assert.equal(await readOutput(".nojekyll"), "");
});

test("static assets are copied unchanged and videos are not LFS pointers", async () => {
  const paths = await readdir("www/assets", {
    recursive: true,
    withFileTypes: true,
  });
  for (const entry of paths.filter(
    (entry) => entry.isFile() && !entry.name.startsWith("."),
  )) {
    const source = `${entry.parentPath}/${entry.name}`;
    const original = await readFile(source);
    assert.deepEqual(
      await readFile(source.replace(/^www\//, "dist/")),
      original,
    );
    if (entry.name.endsWith(".mp4")) {
      assert.ok(original.length > 1024, `${source} must contain a full video`);
      assert.notEqual(original.subarray(0, 7).toString(), "version");
    }
  }
});
