package edit

import (
	"log/slog"
	"os"
	"os/exec"
)

// EditScreenshot opens a simple image editor as a subprocess to edit the newly-created screenshot
// By default this will be pinta but can be overridden with the GSHOOT_EDITOR env var
func EditScreenshot(filename string) error {
	binary := "pinta"

	if s := os.Getenv("GSHOOT_EDITOR"); s != "" {
		binary = s
	}

	cmd := exec.Command(binary, filename)

	if err := cmd.Run(); err != nil {
		slog.Error("Failed to run image editor", slog.Any("editor", binary), slog.Any("err", err))
		return err
	}

	return nil
}
