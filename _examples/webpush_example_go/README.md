## Webpush example

Using python is the easiest way to test this example program. First, `pip install pywebpush`, then

```ipython
from pywebpush import webpush

webpush({"endpoint": "https://nc.malhotra.cc/index.php/apps/uppush/push/a1df0031-b602-4198-a3
c8-01ddc577ffcd", "keys": {"p256dh": "BPKufsdbDfPs_W7edekFlEyPESyUhmZkkn8WRohe6gdYUvLIdmZ9oTXdMmnOgxY5mcbwBXXPAQjutnLe9pxib7A=", "auth":"AQIDBAUGBwgJCgsMDQ4PEA"}},
"This is my message")
```
where `p256dh` and `auth` are values output by the example program.

Run the example program with `go run .`
