
#!/bin/bash
curl -X POST http://127.0.0.1:8080/must-login -H 'content-type: application/json' -d '{"user":"test"}'
curl -X POST http://127.0.0.1:8080/must-login -H 'content-type: application/json' -d '{"user":"test","password":"123"}'
curl -X POST http://127.0.0.1:8080/must-login -H 'content-type: application/json' -d '{"user":"test","password":"123456"}'

curl -X POST http://127.0.0.1:8080/login -H 'content-type: application/json' -d '{"user":"test"}'
curl -X POST http://127.0.0.1:8080/login -H 'content-type: application/json' -d '{"user":"test","password":"123"}'
curl -X POST http://127.0.0.1:8080/login -H 'content-type: application/json' -d '{"user":"test","password":"123456"}'

curl "http://127.0.0.1:8080/bookable?check_in=2021-01-01&check_out=2021-01-02"
curl "http://127.0.0.1:8080/bookable?check_in=2021-01-02&check_out=2021-01-03"