# koios 🔱

### Incremental recommender engine built with Golang & Redis

> Koios (Coeus) was a Titan god of intelligence and farsight, meaning that, due to his inquisitive mind and desire to learn, he was with gained knowledge and understanding able to see beyond the obvious. He was also identified as a god of wisdom and heavenly oracles.

<img src="https://live.staticflickr.com/8335/8360108363_1fb50122bb_b.jpg" width="300" height="300" align="right"/>

### Quickstart

Following the guidelines of the Twelve Factor App (https://12factor.net/), environment variables are kept in a dotenv file at the root directory.

````zsh
docker run -d -p 6379:6379 --name recommender redis

go run server/*.go

curl -H "Authorization: Bearer API_SECRET" "http://localhost:1323/api/rate?user=u1&item=i1&rating=0.88"
````

### REST API Endpoints

- /api/rate?user=[uid]&item=[itemid]&rating=[amount]
- /api/recommend?user=[uid]
- /api/update


### Copyright Notice

The code was based on:
- https://redislabs.com/docs/quick-guide-recommendations-using-redis/ 
- https://github.com/RedisLabs/redis-recommend.
