# Contributing

xaligo welcomes contributors, sponsors, bug reports, examples, documentation
improvements, and real-world diagram feedback.

There are no formal contribution rules. The basic flow is :)

1. Open an issue.
2. Assign yourself to the issue.
3. Send a pull request with the change.

Small fixes are valuable. A corrected example, a clearer error message, a
better route around overlapping labels, or a missing documentation note can be
more useful than a large rewrite.

## What To Contribute

Good contribution areas include:

- `.xal` examples for real architecture patterns.
- Rendering bugs in Excalidraw, SVG, PPTX, PDF, Excel, XYFlow, or Isoflow output.
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
make security-setup # first run, or after scanner version updates
make security-check # required before every commit
go test ./...
GOOS=js GOARCH=wasm go build -o /tmp/xaligo-browser.wasm ./cmd/wasm
npm --prefix external/pptx-exporter test
git diff --check
cargo install mdbook-tabs --version 1.0.4 --locked
mdbook build docs
```

For documentation-only changes, the security check is still required;
`mdbook build docs` and `git diff --check` cover the remaining verification.

Keep changes focused. Prefer a small fix with a regression test over a broad
refactor. When changing output behavior, update the relevant docs and examples
so future users can find the new behavior.

## Sponsor 💖

Sponsorship starts at 1 USD per unit. It helps motivate ongoing development.

Sponsor logos, icons, and names can be listed in this documentation.

Donations are used to maintain the development environment, help cover
community operation costs, and support developer living expenses.

If the donor's identity is clear, in-kind gifts are also welcome 🎁. Please
contact us by X or email before sending anything.

Following Xaligo on X and starring the GitHub repository are also encouraging ⭐.
Please consider following and starring the project. (^^)

Donate via PayPal: [Support from 1 USD](https://www.paypal.com/ncp/payment/6PVX83Y9DMSQJ)

## Contact 📬

X: [@XaligoOrg](https://x.com/XaligoOrg)

Email: `xaligo@outlook.com`.

## License

xaligo itself is MIT licensed, so the project is intentionally easy to use,
modify, redistribute, and build on.

Some bundled icons and generated assets have their own attribution or license
requirements. Preserve the bundled license and attribution files when changing
assets or packaging.
