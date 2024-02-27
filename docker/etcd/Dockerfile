FROM alpine:latest
LABEL version="1.0.0" \
      autho="jun_w@709944968"

ADD etcd-v3.5.12-linux-amd64/etcd /usr/local/bin/
ADD etcd-v3.5.12-linux-amd64/etcdctl /usr/local/bin/
RUN mkdir -p /data
COPY etcd.conf /data
VOLUME [ "/data" ]
# Alpine Linux doesn't use pam, which means that there is no /etc/nsswitch.conf,
# but Golang relies on /etc/nsswitch.conf to check the order of DNS resolving
# (see https://github.com/golang/go/commit/9dee7771f561cf6aee081c0af6658cc81fac3918)
# To fix this we just create /etc/nsswitch.conf and add the following line:
EXPOSE 2379 2380
# Define default command.
CMD ["/usr/local/bin/etcd","--config-file","/data/etcd.conf"]
