# Migrate db

docker exec -it dd-api /bin/sh -c "cd /go/src/github.com/koheishinpuku/dice-dout-api && goose up"

# Rollback

docker exec -it dd-api /bin/sh -c "cd /go/src/github.com/koheishinpuku/dice-dout-api && go run script/db/db.go rollback"
