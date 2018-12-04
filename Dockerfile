FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD ./gopdf /gopdf
ENTRYPOINT ["/gopdf"]