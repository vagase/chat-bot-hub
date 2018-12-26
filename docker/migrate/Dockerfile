FROM migrate/migrate

ARG mirror=mirrors.ustc.edu.cn
ADD modifysource.sh /home/scripts/modifysource.sh
RUN /bin/sh /home/scripts/modifysource.sh $mirror
RUN cat /etc/apk/repositories

RUN apk add --no-cache --update 'su-exec>=0.2'

# commet temporarily, because it will cause compatibility issues on Mac docker.

#COPY docker-entrypoint.sh /usr/local/bin/
#ENTRYPOINT ["docker-entrypoint.sh"]
#CMD ["--help"]