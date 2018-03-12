FROM scratch

ADD bin/pwhash /

CMD ["/pwhash"]