FROM centos:7.6.1810
RUN mkdir -p /opt/tweek
COPY tweek2021-pod-d /opt/tweek

WORKDIR /opt/tweek
ENTRYPOINT ["./tweek2021-pod-d"]

EXPOSE 8090
