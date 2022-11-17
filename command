Basic
- curl localhost:{port}/{path}/.... -i

Query
- curl "localhost:{port}/{path}?{name=value}" 
QueryParser
- curl "localhost:{port}/{path}?{name=value}&{name=value}&..." 

Wildcard
- curl localhost:{port}/{path}/.../.../... -i

Group
- curl localhost"{port}/{group}/{path}/... -i

ENV
- curl localhost"{port}/{path}/... -i | jq

POST
- curl localhost:8000/body -H "content-type:application/json" -d '{"id": 6, "name": "pkk"}'

Authorization
- curl localhost:8000/user -i -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM2MDAsImlzcyI6IjEifQ.0J4Msd92uYQ3SAakUryZ4qWUI5oTbxrebsHDs0oad3Q"