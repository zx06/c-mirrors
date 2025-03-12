package parsemd

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// readTestData 从指定路径读取文件内容并返回字符串，如果读取过程中出现错误，会使用 require.NoError 输出错误信息并终止测试。
func readTestData(t *testing.T, path string) string {
	// 打开文件
	file, err := os.Open(path)
	require.NoError(t, err, "Failed to open file")
	// 确保文件在函数结束时关闭
	defer file.Close()

	// 读取文件内容
	byteValue, err := io.ReadAll(file)
	require.NoError(t, err, "Failed to read file")

	// 将字节切片转换为字符串并返回
	return string(byteValue)
}

func TestParseMd(t *testing.T) {
	testCases := []struct {
		filePath string
		expected *MDInfo
	}{
		{
			filePath: "testdata/deploytool.md",
			expected: &MDInfo{
				From: "",
				To:   "deploytool:latest",
				Dockerfile: `# [to]: deploytool
FROM alpine:3.18

# Install dependencies
RUN apk add --no-cache \
    bash \
    curl \
    git \
    jq \
    openssh-client \
    gettext \
    docker-cli \    
    && rm -rf /var/cache/apk/*

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" \
    && install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl \
    && rm kubectl
`,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.filePath, func(t *testing.T) {
			md := readTestData(t, tc.filePath)
			result := ParseMd(md)
			require.Equal(t, tc.expected, result)
		})
	}

}
