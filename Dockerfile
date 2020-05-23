FROM golang:1.13

WORKDIR /go/src/github.com/Maximilan4/connor

COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN make

ENTRYPOINT ["./build/connor"]
