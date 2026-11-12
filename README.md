# esb.dance

Static site for [esb.dance](https://esb.dance), hosted on GitHub Pages.

## Current status

The repository contains a static HTML migration of six exported pages with local images, videos, stylesheets, and the downloadable CV. The contact page is temporarily removed. The current design is preserved as the starting point for a design pass. GitHub Actions builds and deploys it to GitHub Pages on every push to `main`.

## Local development

Install Node.js 24 or newer, pnpm 12.4.1, and [Git LFS](https://git-lfs.com/). The MP4 videos are stored in Git LFS. On macOS, install it with `brew install git-lfs`. After cloning:

```sh
git lfs install --local
git lfs pull
pnpm install --frozen-lockfile
pnpm dev
```

Eleventy builds the site and serves it at `http://localhost:8080` with automatic browser reload. It watches pages, shared templates, and assets. Press Control+C to stop it. `make dev` runs the same command.

## Repository contents

- `www/`: HTML page templates and static assets.
- `templates/`: Shared Nunjucks layout, header, and footer in HTML files.
- `eleventy.config.js`: Template engine, input and output directories, and asset copying.
- `package.json` and `pnpm-lock.yaml`: Commands and pinned dependencies.
- `.oxfmtrc.json`: Formatter configuration.
- `dist/`: Generated, ignored deployment output.
- `Makefile`: Shortcuts for pnpm commands.
- `AGENTS.md`: Instructions for contributors and coding agents.
- `PROJECT.md`: Project direction, decisions, and open questions.

## Migration direction

1. Review the static migration against the original site.
2. Make the design and content changes we agree on.
3. Decide on the contact implementation and settle font hosting.
4. Move the custom domain after review.

## Editing the site

Edit `www/index.html` for the homepage, or `www/<page>/index.html` for another page. Pages include `bio`, `teaching`, `choreography`, `collaborators` and `program`.

Use `www/assets/styles.css` for all styling and `www/assets/site.js` for navigation and video behavior. Edit `templates/header.html` for shared navigation, `templates/footer.html` for the shared footer, and `templates/layout.html` for document metadata and asset references. The footer contains the Mailchimp subscription form on every rendered page except the 404 page. The contact page is temporarily removed while its implementation is decided. The legacy homepage alias remains a standalone redirect. Navigation automatically marks the current route, and the choreography page retains its hero treatment.

The stylesheet contains shared color and typography tokens, reusable components, named page layouts, and navigation states. Body links share `--body-link-color` and one underline rule. Use `.prose` for rich text, `.photo-block` for photographs, `.card-grid` for repeated cards, and `.composition` for positioned grids. Page layouts use readable names such as `.home-news-layout` and `.biography-photos-layout`; search for the page's layout comment when changing its spacing. Per-image crop values remain inline in HTML.

Keep new styles in this single file and reuse component classes before adding page-specific rules. Check both desktop and mobile when changing layout rules.

Pages use Nunjucks through Eleventy's HTML template engine. YAML front matter selects `layout: layout.html` and sets `title`. The choreography page also sets `bodyClass: has-hero`. To add a page, copy an existing page to `www/<route>/index.html` and update its metadata and content. Add its navigation link in the shared header if needed. The redirect remains standalone HTML; the 404 page explicitly sets `permalink: /404.html`.

## Formatting

```sh
pnpm format
pnpm format:check
```

Oxfmt formats CSS, HTML, and JavaScript source files, excluding generated output and dependencies. The equivalent Make commands are `make format` and `make format-check`. Text elements with `white-space: pre-wrap` have formatter ignore directives to preserve their visible whitespace. Keep those directives when editing these elements.

## Building for deployment

```sh
pnpm build
```

`make build` runs the same command. Eleventy renders pages and copies assets into `dist/`, preserving routes, the legacy homepage redirect, `404.html`, and `.nojekyll`. Publish **`dist/`** through GitHub Pages; `www/` contains template source and must not be published directly.

Build and development commands clean `dist/` before starting so deleted sources do not remain in a fresh build. After deleting or renaming a page during development, restart the server to clear stale output. Eleventy reports template errors; a failed build must not be deployed.

The workflow in `.github/workflows/deploy.yml` installs the pinned pnpm version and Node.js 24, installs dependencies with the frozen lockfile, checks formatting, builds, tests the generated output, and publishes `dist/` on pushes to `main`. Its checkout downloads LFS objects so the artifact contains actual videos rather than LFS pointers. Run `pnpm test` locally to build and check the generated site.

Do not edit `dist/`: it is generated output.

Record decisions and changes in direction in [PROJECT.md](PROJECT.md).
