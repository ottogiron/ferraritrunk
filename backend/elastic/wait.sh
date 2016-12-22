until curl 127.0.0.1:9200/_cluster/health?pretty 2>&1 | grep status | egrep "(green|yellow)"; do
    printf '.'
    sleep 1
done