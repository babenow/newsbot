FROM golang
LABEL authors="bbnf"

ENTRYPOINT ["top", "-b"]