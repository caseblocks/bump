# bump

Increases a counter in VERSION files. Output is to STDOUT and the original VERSION file is left unmodified.

# Usage

Assuming a VERSION file containting `1.2.3.` in the current folder:

    bump major          # Output 2.0.0

    bump minor          # Output 1.3.0

    bump patch          # Output 1.2.4

    bump somethingelse  # Output 1.2.3

To modify a VERSION file:

    bump major > VERSION  # overwrites current file
