# go-pathlib

An implementation of Python's [`pathlib`](https://docs.python.org/3/library/pathlib.html#) for Go.

This package is in extremely early development. Contributions are welcome!

## Similarities to `pathlib`

Like Python's `pathlib`, `go-pathlib` contains to base types: `Path` and `PurePath`.

`PurePath` is the base type upon which all else rests. `PurePath`s provide purely computational operations _without I/O_.

`Path` types, also known as "concrete paths," contain embedded `PurePath` structs and provide I/O operations.

Additionally, both `PurePath` and `Path` contain OS-specific implementations: `PurePosixPath`, `PureWindowsPath`, `PosixPath`, and `WindowsPath`. Their relationships can be diagrammed as follows:

```mermaid
graph TD;
    PurePath <-- Path;
```

## Deviations from `pathlib`

`Path` and `PurePath` are interfaces, so properties in `pathlib` are member functions in `go-pathlib`, e.g. `Path.name` becomes `Path.Name()`.

Some quality-of-life improvements have been added as well, such as `Path.AsString()`.

## Parity with `pathlib`

| `pathlib` function                                                                    | `go-pathlib` function | status      |
| :------------------------------------------------------------------------------------ | :-------------------- | :---------- |
| `Path.absolute()`                                                                     | `Absolute()`          | todo        |
| `Path.chmod()`                                                                        | `Chmod()`             | todo        |
| `Path.cwd()`                                                                          | `Cwd()`               | todo        |
| [`Path.exists()`](https://docs.python.org/3/library/pathlib.html#pathlib.Path.exists) | `Exists()`            | in progress |
| `Path.expanduser()`                                                                   | `ExpandUser()`        | todo        |
| `Path.hardlink_to()`                                                                  | `HardlinkTo()`        | todo        |
| `Path.is_dir()`                                                                       | `IsDir()`             | todo        |
| `Path.is_file()`                                                                      | `IsFile()`            | todo        |
| `Path.is_junction()`                                                                  | `IsJunction()`        | todo        |
| `Path.is_mount()`                                                                     | `IsMount()`           | todo        |
| `Path.is_symlink()`                                                                   | `IsSymlink()`         | todo        |
| `Path.iterdir()`                                                                      | `Iterdir()`           | todo        |
| `Path.lchmod()`                                                                       | `Lchmod()`            | todo        |
| `Path.lstat()`                                                                        | `Lstat()`             | todo        |
| `Path.mkdir()`                                                                        | `Mkdir()`             | todo        |
| `Path.resolve()`                                                                      | `Resolve()`           | todo        |
| `Path.readlink()`                                                                     | `Readlink()`          | todo        |
| `Path.rename()`                                                                       | `Rename()`            | todo        |
| `Path.replace()`                                                                      | `Replace()`           | todo        |
| `Path.samefile()`                                                                     | `Samefile()`          | todo        |
| `Path.stat()`                                                                         | `Stat()`              | todo        |
| `Path.symlink_to()`                                                                   | `SymlinkTo()`         | todo        |
| `Path.unlink()`                                                                       | `Unlink()`            | todo        |
| `Path.rmdir()`                                                                        | `Rmdir()`             | todo        |
| `Path.walk()`                                                                         | `Walk()`              | todo        |
