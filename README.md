# cfliteapi
cloudflare lite api (list, create, update, delete)

Create config file:        cflite -config -email=some@mail.com -xkey=XAuthKey

Work without config file:  cflite -email=some@mail.com -xkey=XAuthKey -domain=example.com -name=test.example.com -type=A  -content=12.13.14.15   -proxy=true  -akt=create

If config file exist

Create:    cflite -domain=example.com -name=test.example.com -type=A  -content=12.13.14.15   -proxy=true  -akt=create"

Create MX: cflite -domain=example.com -name=mx.example.com   -type=MX -content=mx.server.com -proxy=10    -akt=create"

Update:    cflite -domain=example.com -name=test.example.com -type=A  -content=12.13.14.16   -proxy=false -akt=update"

Delete:    cflite -domain=example.com -name=test.example.com -type=A  -akt=delete"

List:      cflite -domain=example.com -akt=list"
