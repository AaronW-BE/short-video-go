package utils

import (
	"bytes"
	"os/exec"
)

/**
从视频保存一帧图片
*/
func ParseFrame(filepath string, pos int) (bytes.Buffer, error) {
	cmd := exec.Command("ffmpeg", "-i", filepath, "-ss", string(pos), "-vframes", "1", "-f", "singlejpeg", "-")
	var buffer bytes.Buffer
	cmd.Stdout = &buffer
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	return buffer, err
}

/**
保存视频第一帧
*/
func ParseFirstFrame(filepath string) (bytes.Buffer, error) {
	return ParseFrame(filepath, 0)
}

func init() {
	// 检查环境变量是否有ffmpeg
	path, err := exec.LookPath("ffmpeg")
	if err != nil || path == "" {
		panic("ffmpeg module is not exist for current system")
	}
}
