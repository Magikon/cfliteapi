# cfliteapi
cloudflare lite api (list, create, update, delete)

Create:    cflite -domain=example.com -name=test.example.com -type=A  -content=12.13.14.15   -proxy=true  -akt=create"
Create MX: cflite -domain=example.com -name=mx.example.com   -type=MX -content=mx.server.com -proxy=10    -akt=create"
Update:    cflite -domain=example.com -name=test.example.com -type=A  -content=12.13.14.16   -proxy=false -akt=update"
Delete:    cflite -domain=example.com -name=test.example.com -type=A  -akt=delete"
List:      cflite -domain=example.com -akt=list"
