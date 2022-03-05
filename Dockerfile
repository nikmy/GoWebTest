FROM golang:latest
ADD . .
EXPOSE 80
CMD [ "go", "run", "main.go" ]