FROM alpine:3.4
ADD build/minesweeper /bin/minesweeper
EXPOSE 3000
ENTRYPOINT ["/bin/minesweeper"]
