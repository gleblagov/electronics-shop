FROM gomicro/goose

COPY *.sql /migrations/
COPY entry.sh /migrations/

ENTRYPOINT ["/migrations/entry.sh"]
