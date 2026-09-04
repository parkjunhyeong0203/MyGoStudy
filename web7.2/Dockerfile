# 1단계: 빌드 환경 (Go 설치된 이미지)
FROM golang:1.26-alpine3.22 AS builder

WORKDIR /app

# 의존성 파일 복사 및 다운로드
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사 및 실행 파일 빌드
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 2단계: 실제 실행 환경 (최소화된 이미지)
FROM scratch

WORKDIR /app

USER 65532:65532

# 빌드 단계에서 생성된 실행 파일만 복사
COPY --from=builder /app/main .

# 포트 개방 (예: 3000)
EXPOSE 3000

# 서버 실행
CMD ["./main"]