package jobs

import (
	"context"

	sw "github.com/sparkwing-dev/sparkwing/sparkwing"
)

// Hello is a sparkwing pipeline. See `sparkwing docs read --topic sdk` for SDK helpers.
type Hello struct{ sw.Base }

func (p Hello) ShortHelp() string { return "one-line description of hello" }

// Help is the long-form description; defaults to ShortHelp until you have more to say.
func (p Hello) Help() string { return p.ShortHelp() }

func (Hello) Examples() []sw.Example {
	return []sw.Example{
		{Comment: "Run locally", Command: "sparkwing run hello"},
	}
}

// Plan registers the pipeline's DAG on the passed-in *Plan. run
// carries run-time environment: run.Git (repo state), run.Trigger
// (what started the run, such as local or github_app), run.Pipeline
// (registered name). Typed inputs arrive as the third parameter.
func (Hello) Plan(ctx context.Context, plan *sw.Plan, _ sw.NoInputs, run sw.RunContext) error {
	sw.Job(plan, run.Pipeline, &HelloJob{})
	return nil
}

type HelloJob struct{ sw.Base }

func (j *HelloJob) Work(w *sw.Work) (*sw.WorkStep, error) {
	sw.Step(w, "run", j.run)
	return nil, nil
}

// Paths in Exec / Bash / ReadFile are relative to the repo root, not
// .sparkwing/. See WorkDir().
func (HelloJob) run(ctx context.Context) error {
	sw.Info(ctx, "replace this stub with your logic")
	// Shell out and propagate failure:
	//   if _, err := sw.Bash(ctx, "go test ./...").Run(); err != nil {
	//           return err
	//   }
	return nil
}

func init() {
	sw.RegisterEntrypoint[sw.NoInputs]("Hello", func() sw.Pipeline[sw.NoInputs] { return &Hello{} })
}
