package player

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func StreamSong(query string) {
	// TODO: USE BUBBLETEA AND LIPGLOSS HERE
	fmt.Printf("Streaming %s\n", query)

	ytSearch := fmt.Sprintf("ytsearch1:%s", query)

	urlCmd := exec.Command("yt-dlp", "-f", "bestaudio", "-g", ytSearch)
	output, err := urlCmd.Output()
	if err != nil {
		fmt.Println("yd-dlp error:", err)
		return
	}

	streamURL := strings.TrimSpace(string(output))

	mpvCmd := exec.Command("mpv", "--no-video", streamURL)
	mpvCmd.Stdout = os.Stdout
	mpvCmd.Stderr = os.Stderr
	_ = mpvCmd.Run()
}
