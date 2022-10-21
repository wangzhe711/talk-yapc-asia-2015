FROM golang:1.19

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.3/zsh-in-docker.sh)"
RUN apt install -y graphviz
RUN apt install -y net-tools
RUN go install golang.org/x/perf/cmd/benchstat@latest

WORKDIR /demo

COPY go.mod ./

RUN go mod download && go mod verify

RUN echo "alias bench='go test -v -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=profile.cpu -memprofile=profile.mem'" >> /root/.zshrc
RUN echo "alias cpu='go tool pprof demo.test progfile.cpu'" >> /root/.zshrc

# note: 0.0.0.0 is needed becase we're using a web browser outside of the container
RUN echo "alias web='go tool pprof -http=0.0.0.0:8082 -no_browser demo.test prof.block'" >> /root/.zshrc

COPY . .

WORKDIR ./demo
# cpu & memory profiling
#RUN go test -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=profile.cpu_1 -memprofile=profile.mem_1 | tee res1
#RUN git checkout fprintf
#RUN go test -bench=^BenchmarkHi$ -benchtime=2s -cpuprofile=profile.cpu_2 -memprofile=profile.mem_2 | tee res2

# contention profiling
RUN go test -bench=. -blockprofile=prof.block | tee res
WORKDIR ./..