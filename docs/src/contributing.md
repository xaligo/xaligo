# Contributing and Sponsorship

xaligo welcomes contributors, sponsors, bug reports, examples, documentation
improvements, and real-world diagram feedback.

The project has a simple working spirit:

1. Open an issue.
2. Report what is wrong or what could be better.
3. If you can, fix it yourself and send a pull request.
4. Reflect improvements back into this repository whenever possible.

Small fixes are valuable. A corrected example, a clearer error message, a
better route around overlapping labels, or a missing documentation note can be
more useful than a large rewrite.

## What To Contribute

Good contribution areas include:

- `.xal` examples for real architecture patterns.
- Rendering bugs in Excalidraw, SVG, PPTX, XYFlow, or Isoflow output.
- Connector routing improvements.
- AWS, Tabler, or Yamaha catalog corrections.
- Documentation improvements.
- Diagnostics and validation messages.
- Tests that capture confusing edge cases.
- Packaging, CI, and release workflow improvements.

When reporting rendering issues, include:

- The smallest `.xal` file that reproduces the problem.
- The command you ran.
- The output format.
- A screenshot or generated artifact when it helps.
- What you expected and what actually happened.

## Pull Request Guidelines

Before opening a pull request:

```bash
go test ./...
git diff --check
mdbook build docs
```

For documentation-only changes, `mdbook build docs` and `git diff --check` are
usually enough.

Keep changes focused. Prefer a small fix with a regression test over a broad
refactor. When changing output behavior, update the relevant docs and examples
so future users can find the new behavior.

## Sponsorship

Sponsorship helps keep xaligo moving: more examples, better routing, stronger
exports, clearer documentation, and faster issue response.

Sponsors are welcome to support:

- Maintenance of the CLI and rendering pipeline.
- Better SVG/PPTX/Excalidraw output quality.
- Documentation and sample diagrams.
- Catalog updates and asset import workflows.
- Triage and fixes for real-world diagram reports.

If you want to sponsor work, open an issue describing the area you care about.
The best sponsorships are attached to concrete improvements that can land back
in this repository for everyone.

For sponsorship discussions or direct project contact, email
`xaligo@outlook.com`. Issues are still preferred for concrete bug reports and
feature requests so the discussion and resulting improvements remain visible to
the community.

## License

xaligo itself is MIT licensed, so the project is intentionally easy to use,
modify, redistribute, and build on.

Some bundled icons and generated assets have their own attribution or license
requirements. Preserve the bundled license and attribution files when changing
assets or packaging.
