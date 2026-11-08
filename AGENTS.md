# Project guidance

## Direction and decisions

- Build esb.dance as a static site hosted on GitHub Pages.
- Read `README.md` and `PROJECT.md` before making project changes.
- Keep `PROJECT.md` updated with substantive decisions, their rationale, overall project direction, and unresolved questions.
- Date decision entries and distinguish agreed decisions from proposals and open questions.
- Update existing notes when decisions change so the current direction remains clear.
- Keep notes concise and avoid recording routine activity or sensitive information.
- Keep shared document markup in `templates/` and page content in `www/`. Build with `make build` and deploy `dist/`; do not edit generated output.
- Keep styling in `www/assets/styles.css`. Use readable component and page classes, and check desktop and mobile layouts after layout changes.

## Working conventions

- Use Standard Technical English in documentation and code comments.
- Write concise comments as complete sentences with proper capitalization and a final period.
- Defer formatting, builds, and linting during iteration unless needed for the task. 
- To deploy changes, push directly to main after local verification, no pull requests neede.
