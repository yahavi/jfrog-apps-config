package jfrogappsconfig

const version = "1.0"

type JFrogAppsConfig struct {
	Version      string      `json:"version"`
	Applications Application `json:"applications,omitempty"`
}

type Application struct {
	Name            string    `json:"name"`
	SourceRoot      string    `json:"source_root,omitempty"`
	ExcludePatterns []string  `json:"exclude_patterns,omitempty"`
	ExcludeScanners []string  `json:"exclude_scanners,omitempty"`
	Scanners        []Scanner `json:"scanners,omitempty"`
}

type Scanner struct {
	WorkingDirs     []string `json:"working_dirs,omitempty"`
	ExcludePatterns []string `json:"exclude_patterns,omitempty"`
}

type SastScanner struct {
	Scanner
	Language string `json:"language,omitempty"`
}
