FROM scratch
EXPOSE 8080
ENTRYPOINT ["/rest-api-go-gin"]
COPY ./bin/ /