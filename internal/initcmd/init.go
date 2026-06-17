package initcmd

import (
	"path/filepath"

	"github.com/dmparrow/agent-md/internal/config"
	"github.com/dmparrow/agent-md/internal/fsutil"
)

func Init(root string) error {
	agentDir := filepath.Join(root, config.DirName)
	runsDir := filepath.Join(agentDir, config.RunsDirName)
	configPath := filepath.Join(agentDir, config.ConfigName)

	if err := fsutil.EnsureDir(runsDir); err != nil {
		return err
	}

	return fsutil.WriteFileIfMissing(configPath, []byte(config.DefaultYAML), 0o644)
}
