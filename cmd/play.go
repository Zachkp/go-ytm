/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-ytm/player"
	"strings"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Search and play song with mpv and yt-dlp",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args, " ")
		// TODO: UTILIZE BUBBLETEA AND LIPGLOSS HERE
		fmt.Println("play called")
		player.StreamSong(query)
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
