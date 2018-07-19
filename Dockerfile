FROM scratch
ADD app /
ADD db.conf /
CMD ["/app"]

