# Cloudflare Lite API

```shell
Create:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -content=12.13.14.15   -proxy=true  -akt=create
Create:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=CNAME  -content=example.com   -proxy=true  -akt=create
Create:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test             -type=CNAME  -content=@             -proxy=true  -akt=create
Create MX:    cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=mx.example.com   -type=MX     -content=mx.server.com -proxy=10    -akt=create
Update:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -content=12.13.14.16   -proxy=false -akt=update
Delete:       cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -akt=delete
List:         cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -akt=list
Purge Cache:  cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -akt=purge
```

if email not given, xkey used as Bearer

```shell
Create:       cflite -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -content=12.13.14.15   -proxy=true  -akt=create
Create:       cflite -xkey=XAuthKey -domain=example.com -name=test.example.com -type=CNAME  -content=example.com   -proxy=true  -akt=create
Create:       cflite -xkey=XAuthKey -domain=example.com -name=test             -type=CNAME  -content=@             -proxy=true  -akt=create
Create MX:    cflite -xkey=XAuthKey -domain=example.com -name=mx.example.com   -type=MX     -content=mx.server.com -proxy=10    -akt=create
Update:       cflite -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -content=12.13.14.16   -proxy=false -akt=update
Delete:       cflite -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A      -akt=delete
List:         cflite -xkey=XAuthKey -domain=example.com -akt=list
Purge Cache:  cflite -xkey=XAuthKey -domain=example.com -akt=purge
```
