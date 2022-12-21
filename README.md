# recdrive

中科大睿客网云盘 rec.ustc.edu.cn/recdrive (unofficial) SDK

## Maintenance

I still have my recdrive account so if the library can not work, I will debug and fix it.
However, further development is not likely since now I do not use it.
If you are looking for a (somehow, more than this project) actively maintained recdrive cli, see the _Related_ section.

## Features

- List a folder
- Upload / Download a file

## Get Started

```bash
# List a folder
recdrive -token <token> ls <dir>
# Upload a file
recdrive -token <token> cp <file> :<dir>
# Download a file
recdrive -token <token> cp :<file> <file>
# Other than `-token <token>`, use env `TOKEN` is also OK
```

## Related

- [taoky/reccli](https://github.com/taoky/reccli)

## License

MIT or Apache 2.0 at your option

SPDX-License-Identifier: MIT OR Apache-2.0
