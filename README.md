# OfflinePush

Offline Push Using GeTui Rest API v2.

send push via proxy, which is an `adaptor`.

```txt
|other code| ----> (proxy pkg) ---->|push.svr|
```

maybe using a proxy svr (auth/stats/user/push api, all in one) would be better!

stats/user/push svr are base on auth svr (using cache pkg, of course).

```txt
|stats.svr| & |user.svr| & |push.svr| ----> (cache pkg) ----> |auth.svr|
```
