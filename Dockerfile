# FROM sonarsource/sonar-scanner-cli:latest AS sonarqube_scan
# WORKDIR /app
# COPY . .
# RUN ls -list    
# # sonar.projectName property used for providing human-friendly project name in addition 
# # for projectKey
# RUN sonar-scanner \
#     -Dsonar.host.url="http://localhost:9090" \
#     -Dsonar.projectKey="AltaStore" \
#     -Dsonar.sources="." \
#     -Dsonar.go.coverage.reportPaths="coverage.out" \
#     -Dsonar.login="ff9f3a19daf80a937e25559f51f3d6049b0525a6" \ 
#     -Dsonar.exclusions="business/**/service_test.go, business/errors.go, api/**, app/**, config/**, modules/**, util/**"\
#     -Dsonar.test.exclusions="business/**/service_test.go"\  

# stage I - khusus build dengan envinroment yang sama
FROM golang:1.16-alpine AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main 
# EXPOSE 8080
# CMD ["/app/main"]

# stage 2
FROM alpine:3.14
WORKDIR /root/
# COPY --from=builder /app/config.json .
# COPY --from=builder /app/.env /config/
COPY /config/.env /config/.env
COPY --from=builder /app/main .
EXPOSE 9000
CMD ["./main"]