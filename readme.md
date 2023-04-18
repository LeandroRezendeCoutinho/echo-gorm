### Echo + GORM

Simple implementation REST API for Echo with GORM

## Explaining test
- Use benchmarck tool in route /dogs
    - `autocannon http://localhost:5000/dogs`
- Populate database with 10 records
```curl
curl --location --request GET 'localhost:5000/dogs' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Ralph",
    "breed":"Pastor",
    "age": 12,
    "isGoodBoy":true
}'
```

#### Benchmark
![bench](https://github.com/LeandroRezendeCoutinho/echo-gorm/blob/main/img/echo_gorm_bench.png)

#### Memory and CPU
![bench](https://github.com/LeandroRezendeCoutinho/echo-gorm/blob/main/img/echo_mem.png)

References:
- https://gin-gonic.com/
- https://gorm.io/
