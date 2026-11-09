# esb.dance

Static site for [esb.dance](https://esb.dance), hosted on GitHub Pages.

## Current status

The repository contains a static HTML migration of six exported pages with local images, videos, stylesheets, and the downloadable CV. The contact page is temporarily removed. The current design is preserved as the starting point for a design pass. GitHub Actions builds and deploys it to GitHub Pages on every push to `main`.

## Local development

Install [Go](https://go.dev/dl/) 1.25 or newer and [Git LFS](https://git-lfs.com/). The MP4 videos are stored in Git LFS. On macOS, install it with `brew install git-lfs`. After cloning, initialize LFS and download the videos:

```sh
git lfs install --local
git lfs pull
```

Then run this command from the repository directory:

```sh
make dev
```

`modd` builds the site, starts `devd`, and watches `www/`, `templates/`, and the Go build code. The server opens the site in your browser and reloads it when generated files change. Edit an HTML template or CSS file and save it to see the change. Press Control+C in the terminal to stop the server.

If Make is not installed, run the equivalent command directly:

```sh
go tool modd
```

Go downloads and compiles the pinned [modd](https://github.com/cortesi/modd) and [devd](https://github.com/cortesi/devd) tools on first use, so the first launch requires internet access and may take longer. Neither tool requires a separate installation. Building modd from source requires CGO and a C compiler; on macOS, install Xcode Command Line Tools if they are not already available. Tool versions and dependency checksums are tracked in `go.mod` and `go.sum`.

## Repository contents

- `www/`: Page templates and static assets.
- `templates/`: Shared layout, header, and footer templates.
- `cmd/site/`: Go site builder and its tests.
- `dist/`: Generated, ignored deployment output.
- `modd.conf`: File watching, rebuilding, and preview server configuration.
- `Makefile`: Local development commands.
- `go.mod` and `go.sum`: Pinned Go tools and their dependencies.
- `AGENTS.md`: Instructions for contributors and coding agents.
- `PROJECT.md`: Project direction, decisions, and open questions.

## Migration direction

1. Review the static migration against the original site.
2. Make the design and content changes we agree on.
3. Decide on the contact implementation and settle font hosting.
4. Move the custom domain after review.

## Editing the site

Edit `www/index.html` for the homepage, or `www/<page>/index.html` for another page. Pages include `bio`, `teaching`, `choreography`, `collaborators` and `program`.

Use `www/assets/styles.css` for all styling and `www/assets/site.js` for navigation and video behavior. Edit `templates/header.html` for shared navigation, `templates/footer.html` for the shared footer, and `templates/layout.html` for document metadata and asset references. The footer contains the Mailchimp subscription form on every rendered page except the 404 page. The contact page is temporarily removed while its implementation is decided. The homepage redirect remains a standalone redirect. Navigation automatically marks the current route, and the choreography page retains its hero treatment.

The stylesheet contains shared color and typography tokens, reusable components, named page layouts, and navigation states. Body links share `--body-link-color` and one underline rule. Use `.prose` for rich text, `.photo-block` for photographs, `.card-grid` for repeated cards, and `.composition` for positioned grids. Page layouts use readable names such as `.home-news-layout` and `.biography-photos-layout`; search for the page's layout comment when changing its spacing. Per-image crop values remain inline in HTML.

Keep new styles in this single file and reuse component classes before adding page-specific rules. Check both desktop and mobile when changing layout rules.

Pages use Go's `html/template`: each page defines its `title` and `content`, then includes the shared layout. The choreography page also defines `body-class`. To add a page, copy an existing page template to `www/<route>/index.html` and edit those definitions. Add its navigation link in the shared header if needed. The redirect page remains standalone HTML.

## Building for deployment

```sh
make build
```

Without Make, run `go run ./cmd/site`. This renders HTML and copies assets into `dist/`, preserving routes, the homepage redirect, `404.html`, and `.nojekyll`. Publish **`dist/`** through GitHub Pages; `www/` now contains template source and must not be published directly. The workflow in `.github/workflows/deploy.yml` runs the tests, builds, and publishes `dist/` on every push to `main`.

The deployment workflow downloads LFS objects before building, so the published artifact contains the actual videos rather than LFS pointers.

Do not edit `dist/`: it is generated output, and rebuilding removes files whose sources were deleted. Template errors fail the build before updating output. During development, the last successful output remains available until the error is fixed. Run `make test` or `go test ./cmd/site` to check the builder.

Record decisions and changes in direction in [PROJECT.md](PROJECT.md).
