FROM golang:1.16-alpine
ENV PORT 8080

WORKDIR /app
# RUN go install

# install go hot reroad tool
RUN go get -u github.com/cosmtrek/air
EXPOSE 8080

CMD ["air"]