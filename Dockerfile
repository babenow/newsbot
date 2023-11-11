FROM golang
LABEL authors="Maks B. <babenoff.code@outlook.com>"

ENTRYPOINT ["top", "-b"]