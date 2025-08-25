#!/usr/bin/env python3
#
# Tags, pushes, and publishes.

import os
import subprocess
import sys


def git_tag(tag: str) -> None:
    try:
        _: subprocess.CompletedProcess[bytes] = subprocess.run(["git", "tag", tag])
    except subprocess.CalledProcessError as error:
        raise error
    except Exception as error:
        print("Unexpected error:")
        raise error
    return


def git_push(tag: str) -> None:
    try:
        _: subprocess.CompletedProcess[bytes] = subprocess.run(
            ["git", "push", "origin", tag]
        )
    except subprocess.CalledProcessError as error:
        raise error
    except Exception as error:
        print("Unexpected error:")
        raise error
    return


def publish(tag: str) -> None:
    os.environ["GOPROXY"] = "proxy.golang.org"
    try:
        _: subprocess.CompletedProcess[bytes] = subprocess.run(
            ["go", "list", "-m", f"github.com/rewgs/go-pathlib@{tag}"]
        )
    except subprocess.CalledProcessError as error:
        raise error
    except Exception as error:
        print("Unexpected error:")
        raise error
    return


def main(tag: str):
    git_tag(tag)
    git_push(tag)
    publish(tag)
    return


if __name__ == "__main__":
    try:
        main(sys.argv[1])
    except IndexError:
        raise Exception("Required: git tag")
