FROM scratch
MAINTAINER cwc cuiwenchang@k2data.com.cn
#REPO gitee.com/k2tf/authx
ADD bi-proxy /
ADD config.yaml /
EXPOSE 8080
#VOLUME /data/conf
ENTRYPOINT ["/bi-proxy", "/config.yaml"]
