FROM ubuntu:18.04
ADD ./HiData.Manager /HiData.Manager
ENTRYPOINT ["/HiData.Manager"]
