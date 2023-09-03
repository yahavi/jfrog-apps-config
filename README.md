# JFrog Applications Config

## Schema:

```yaml
# [Required] JFrog Applications Config version
version: "1.0"

applications:
  # [Required] Application Name
  - name: FrogLeapApp
    # [Optional, default: "."] Application's Root Directory
    source_root: "src"
    # [Optional] Directories to exclude from Scanning Across All Scanners
    exclude_patterns:
      - "docs/"
    # [Optional] Scanners to exclude from JFrog Advanced Security (Options: "secrets", "sast", "iac")
    exclude_scanners:
      - secrets
    # [Optional] Scanner Configurations
    scanners:
      # [Optional] Configuration for Static Analysis Scanner (SAST)
      sast:
        # [Optional] Specify the Programming Language for SAST
        language: java
        # [Optional] Working Directories Specific to SAST (Relative to source_root)
        working_dirs:
          - "src/module1"
          - "src/module2"
        # [Optional] Additional exclude Patterns for this Scanner
        exclude_patterns:
          - "src/module1/test"
        # [Optional] List of specific scan rules to exclude from the scan
        excluded_rules:
          - xss-injection
```
