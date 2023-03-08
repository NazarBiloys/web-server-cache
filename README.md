# web-server-cache
Cache images on server side (nginx) with ability to drop cache by endpoint

- run `docker-compose up -d`
- go twice to endpoint `http://localhost:90/image/tyger.jpg`
- after this nginx save to cache image
- if you want to remove some image from cache, run endpint `http://localhost:90/purge/tyger.jpg`
