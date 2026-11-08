# Project notes

## Overall direction

Build esb.dance as a static site deployed through GitHub Pages.

## Decisions

### 2026-11-07: Establish the migration repository

- Use GitHub Pages as the target host, as requested by Michael.
- Keep project direction and decision rationale in this file so future work has a shared reference.

### 2026-11-07: Add local development tooling

- Use `www/` as the static site root, beginning with a plain HTML test page.
- Use devd for local serving and automatic browser reload through `make dev`.
- Track devd with a Go module tool directive and commit dependency checksums for reproducible setup without a separate tool installation.
- Require Go 1.25 or newer to satisfy the resolved tool dependencies. Provide a direct `go tool` command for contributors without Make.
- Keep the development setup independent of Node.js to simplify setup for non-engineers.

### 2026-11-07: Preserve the current site in static HTML

- Initially use plain HTML, CSS, and small browser scripts without a build step. The shared-template decision below supersedes the original build-free workflow.
- Preserve the exported pages and their text, except for the temporary contact removal agreed below. Use public live-page markup and styles for layout fidelity, supplementing the export with local media and the CV PDF.
- Preserve the main routes with directory index files. Map `/emily-schoen-branch/` to the homepage and restore `/program/` from the export because the live route returned 404.
- Store images, three MP4 videos, the CV, stylesheets, and stylesheet dependencies locally.
- Replace legacy runtime scripts with native video elements and a small accessible menu script. Retain YouTube and Vimeo embeds.
- Contact integration remains deferred. The temporary contact page and links to the old contact form were removed on 2026-11-08, as agreed below.
- Retain remote Adobe font references for visual review. Resolve independent font hosting before cancelling the previous hosting.
- Use the initial migration commit as the imported CSS fidelity baseline. The shared stylesheet consolidation below supersedes retaining vendor CSS in the working site.

### 2026-11-07: Restore responsive presentation

- Restore the original site wrapper and CSS selector attributes. The initial migration removed attributes used for paragraph spacing, list alignment, responsive font sizes, and grid configuration.
- Match the original mobile logo size, hamburger control, full-screen navigation, and navigation breakpoints. Keep keyboard dismissal, focus containment, and background scroll locking.
- Reserve header space inside the first section to avoid double-counting it in mobile banners.
- Compare the live site and static pages at the same viewport size. Mobile checks must include content below the first screen, not just the header and first image.

### 2026-11-07: Consolidate site styling

- Use `www/assets/styles.css` as the only stylesheet, as requested by Michael. Remove imported vendor stylesheets and unused editor font and image assets.
- Replace legacy component classes and generated layout identifiers with readable component and page names. Keep responsive layout values to preserve the current presentation.
- Move embedded page styles into the shared stylesheet. Keep individual image crops and content-specific configuration inline with their HTML.
- Give body links a shared dark gray color and underline. Preserve separate navigation, button, and hero treatments, including the choreography navigation at 75% white and hero title at 50% white.
- Compare all seven pages against the pre-consolidation version at desktop and mobile widths. Keep the stylesheet directly editable without a CSS build pipeline.

### 2026-11-08: Share HTML templates and build static output

- Use Go's standard `html/template` package, as requested by Michael. Keep editable page templates and assets in `www/`, with shared layout, header, and footer definitions in `templates/`.
- Generate deployable files into ignored `dist/` through `make build`. GitHub Pages must publish `dist/`, replacing the previous plan to publish `www/` directly.
- 2026-11-08, agreed: Deploy with a GitHub Actions workflow (`.github/workflows/deploy.yml`) that tests, runs `make build`, and publishes `dist/` on pushes to `main`. The repository Pages source must be set to GitHub Actions.
- Define navigation once and derive its active state from each page's route. Use the shared footer for the subscription form described below.
- Use `modd` for source watching and rebuilding, as requested by Michael, and `devd` for serving generated files and browser reloads. Pin both tools through `go.mod`; `make dev` runs `go tool modd`.
- Preserve routes and page content except for the temporary contact removal agreed below. Render every page before updating output so template errors leave the last successful build available. No Node.js toolchain is required.

### 2026-11-08: Simplify the choreography banner

- Remove the visible video pause control at Michael's request. Preserve muted autoplay and the reduced-motion poster fallback.

### 2026-11-08: Add Mailchimp subscriptions

- Use Michael's supplied Mailchimp form action, email and name fields and bot-trap field in the shared footer. Show it on all remaining rendered pages; the legacy homepage alias still redirects.
- Style the form in `www/assets/styles.css` using the site's typography and colors. Use native email validation and a standard POST to Mailchimp in a new tab, where Mailchimp handles confirmation and errors. Do not load the embed's unrelated SMS scripts or external stylesheet.
- Omit the Mailchimp logo and attribution link at Michael's request.
- Keep the separate contact integration deferred. No live subscription was submitted during implementation.

### 2026-11-08: Store videos in Git LFS

- Agreed: Track `www/assets/videos/*.mp4` with Git LFS and migrate the existing local history to remove the video blobs from ordinary Git transfers. Preserve original author and committer metadata during migration.
- Download LFS objects in the deployment checkout before building and uploading `dist/`, so Pages serves ordinary video files from the deployment artifact.
- Contributors need Git LFS and must download the video objects before building locally.

### 2026-11-08: Temporarily remove the contact page

- Agreed: Remove the contact page and navigation entry while Michael decides how to implement contact. Remove the homepage link to the old contact form and retain its social messaging links.
- Keep the shared Mailchimp subscription form available. Restore a contact page once its implementation is settled.

## Open questions

- What Mailchimp integration and contact flow should be configured when contact work resumes?
- Should the project retain Orpheus Pro and Adobe Garamond Pro through an Adobe Fonts web project, or switch fonts during the design pass?
- What domain configuration will be used for GitHub Pages? The repository and deployment workflow are settled (2026-11-08).

