FROM redis:5.0.4
LABEL version="1.0.0" \
      autho="jun_w@709944968"
RUN mkdir -p /data      
COPY redis.conf /data/redis.conf
#暴露短口
EXPOSE 6379

#macos查看挂载目录 docker run -it --privileged --pid=host debian nsenter -t 1 -m -u -n -i sh
VOLUME [ "/data" ]

# 启动命令
CMD [ "redis-server","/data/redis.conf"]
  