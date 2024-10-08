Команда для запуска `centrifugo`:
```bash
docker run --name centrifugo -v $(pwd)/config.json:/centrifugo/config.json -p 10000:10000 -p 8000:8000 --ulimit nofile=65535:65535 centrifugo/centrifugo:v5.4 centrifugo -c config.json
```

Админка по адресу `127.0.0.1:8000`  
GRPC-апи по адресу `127.0.0.1:10000`