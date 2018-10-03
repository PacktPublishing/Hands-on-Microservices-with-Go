FROM golang:1.10
RUN mkdir /app
ADD server.key /app
ADD server.pem /app
ADD main /app/main
EXPOSE 8443
CMD ["/app/main"]
