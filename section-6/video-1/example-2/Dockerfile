FROM golang:1.10-alpine
ENV APPLICATION_PORT=7951
RUN mkdir /app
ADD main /app/main
EXPOSE 7951
CMD ["/app/main"]
