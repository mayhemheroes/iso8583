FROM fuzzers/go-fuzz:1.2.0 as builder

ARG project='iso8583'
ARG fuzzTargetPath='test/fuzz-reader'
ARG fuzzTarget='reader-fuzz'

ADD . ${project}
WORKDIR ${project}
RUN GO_ENABLED=0 go-fuzz-build -libfuzzer -o /${fuzzTarget}.a ./${fuzzTargetPath}
RUN clang -fsanitize=fuzzer /${fuzzTarget}.a -o /${fuzzTarget}.libfuzzer

FROM ubuntu:20.04
COPY --from=builder /reader-fuzz.libfuzzer /
