# chmtime

chmtime changes the mtime of a file. It's like touch -t, but with RFC3339 time
format with support for nanosecond precision.

## Usage example

```
$ touch foo
$ chmtime 2023-01-01T10:00:00.123456789Z foo
old mtime: 2023-09-24T10:22:00.502255655+02:00
new mtime: 2023-01-01T11:00:00.123456789+01:00
```
