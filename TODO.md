# TODO

## post v0.0.1-alpha.1

- Rename `path` and `purepath`. The former conflicts with the `path` package. Perhaps `goPath` and `goPurePath`?
- Returning too many errors. Let `New()` and `Exists()` just return a single value -- I'll never use the errors they're currently returning outside of the function, and it makes using this module less expresive.
