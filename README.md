```bash
cd URL-Shortener/api
go mod tidy

cd ..
docker compose up -d

curl -X POST http://localhost:3000/api/v1 \
  -H "Content-Type: application/json" \
  -d '{"url":"https://youtu.be/SR5tSMlmmWU?si=kukXc5PS1NoCEjfd"}'

# possible output
# {"url":"https://youtu.be/SR5tSMlmmWU?si=kukXc5PS1NoCEjfd","short":"localhost:3000/a3b5b3","expiry":24,"rate_limit":8,"rate_limit_reset":21}%
```
