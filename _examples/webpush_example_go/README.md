## Webpush example

Using python is the easiest way to test this example program. First, `pip install pywebpush`, then

```ipython
from pywebpush import webpush

webpush({"endpoint": "https://myendpointdomain.com/randomtoken", "keys": {"p256dh": "BPKufsdbDfPs_W7edekFlEyPESyUhmZkkn8WRohe6gdYUvLIdmZ9oTXdMmnOgxY5mcbwBXXPAQjutnLe9pxib7A=", "auth":"AQIDBAUGBwgJCgsMDQ4PEA"}},
"This is my message")
```
where `p256dh` and `auth` are values output by the example program.

Run the example program with `go run .`

Here, Python is the (pretend) application server, this example program is the end-user application, and they can be used with any distributor that supports the latest D-Bus UnifiedPush spec.
