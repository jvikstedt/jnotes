FROM golang:1.7.3-wheezy

run curl https://glide.sh/get | sh
run go get -v github.com/codegangsta/gin

run mkdir -p /go/src/github.com/jvikstedt/jnotes
WORKDIR /go/src/github.com/jvikstedt/jnotes

COPY glide.lock glide.yaml /go/src/github.com/jvikstedt/jnotes/
run glide install

copy . /go/src/github.com/jvikstedt/jnotes

run go install -v

EXPOSE 3000
