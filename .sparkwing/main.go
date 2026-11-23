// Command repo-pipelines serves this repo's pipeline definitions to the Sparkwing host.
package main

import (
	"github.com/sparkwing-dev/sparkwing/sparkwing"

	_ "repo-pipelines/jobs"
)

func main() { sparkwing.Main() }
