---

# QR 코드 생성기 (Go)

## 프로젝트 개요
이 프로젝트는 사용자가 입력한 URL을 QR 코드로 변환하여 보여주는 간단한 웹 애플리케이션입니다. Go 언어와 Gin 프레임워크를 사용하여 서버를 구현하고, 클라이언트가 입력한 URL을 QR 코드로 생성하여 이미지를 표시합니다.

## 주요 기능
- 사용자가 입력한 URL을 QR 코드로 변환
- 생성된 QR 코드를 웹 페이지에 바로 표시
- QR 코드 이미지를 PNG 형식으로 제공

## 사용 기술
- Go (Gin 프레임워크)
- HTML, CSS
- `skip2/go-qrcode` 라이브러리

## 설치 방법
1. 이 프로젝트를 클론합니다:
   ```bash
   git clone https://github.com/yourusername/qrcode-generator.git
   ```

2. 프로젝트 디렉토리로 이동합니다:
   ```bash
   cd qrcode-generator
   ```

3. 필요한 패키지를 설치합니다:
   ```bash
   go get -u github.com/gin-gonic/gin
   go get -u github.com/skip2/go-qrcode
   ```

## 실행 방법
1. Go 서버를 실행합니다:
   ```bash
   go run main.go
   ```

2. 브라우저에서 `http://localhost:8080`으로 접속합니다.

3. URL을 입력하고 "QR 코드 생성" 버튼을 클릭하여 QR 코드를 생성합니다.

## 디렉토리 구조
```
├── main.go          # Go 서버 코드
├── static
│   └── style.css    # CSS 파일
└── templates
    └── index.html   # HTML 템플릿
```

## 참고 자료
- [Gin 프레임워크 문서](https://gin-gonic.com/docs/)
- [skip2/go-qrcode 문서](https://github.com/skip2/go-qrcode)

---
