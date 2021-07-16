FROM scratch
EXPOSE 8080
ENTRYPOINT ["/case-study"]
COPY ./bin/ /
COPY ./conf/ /conf