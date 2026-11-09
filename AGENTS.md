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

## Git LFS

- Track `www/assets/videos/*.mp4` with Git LFS through `.gitattributes`. Keep LFS pointers in Git and full video files in the working tree.
- Install Git LFS before working with videos (`brew install git-lfs` on macOS). After cloning, run `git lfs install --local` and `git lfs pull` before building.
- Stage new or changed videos normally with `git add`. Check `git lfs ls-files` and run `git lfs fsck` after video changes. Commit `.gitattributes` changes with the affected assets.
- Keep `lfs: true` in the deployment workflow's checkout step. Build and deploy actual video files in `dist/`, never LFS pointers.
- Adding a tracking rule does not remove existing large blobs from history. Rewrite history only when requested, preserve original author and committer metadata, and retain a local recovery reference.
- Upload LFS objects only when pushing is explicitly authorized. For an assets-only upload, use `git lfs push origin main`; this does not push the Git branch. If objects are skipped, use `git lfs push --object-id origin <oid>...` for the specific authorized assets.
- Verify remote availability after uploading. Local LFS files and a successful command alone are not proof. Check GitHub's LFS download API or fetch into an empty temporary LFS storage directory with `git -c lfs.storage=<directory> lfs fetch origin <commit>`.
- Confirm that required LFS objects are available remotely before pushing commits that reference them. Report LFS upload and Git branch push results separately.

## Working conventions

- Use Standard Technical English in documentation and code comments.
- Write concise comments as complete sentences with proper capitalization and a final period.
- Defer formatting, builds, and linting during iteration unless needed for the task. 
- To deploy changes, push directly to main after local verification, no pull requests neede.
