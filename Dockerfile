FROM scratch
ADD ./gopdf /gopdf
ENTRYPOINT ["/gopdf"]