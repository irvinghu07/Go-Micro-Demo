FROM alpine
ADD fibo /fibo
ENTRYPOINT [ "/fibo" ]
