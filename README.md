# 🕒 timestamp-cli

`timestamp` is a command to generate timestamps (in milliseconds) with a human readable interface.

Supported periods: `minute` `hour` `day` `month` `year`

### Examples:
```
$ timestamp 50 minutes ago
> 1579317748880
```
```
$ timestamp 5 hours from now
> 1579338789212
```

## Build
**In order to build the timestamp command you need go, I used go v1.13 (but this should probably work for older version since I'm not using anything fancy)**

```
$ make
> ...
$ ./build/timestamp
> 1579338789212
```

## Install

* `$ make install`
* And you're good to go, this script install timestamp in /usr/bin
