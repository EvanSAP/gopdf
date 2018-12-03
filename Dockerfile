FROM scratch
ADD ./hello-world /hello-world
ENTRYPOINT ["/hello-world"]