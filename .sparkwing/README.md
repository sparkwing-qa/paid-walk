# .sparkwing/

This directory holds this repo's [sparkwing](https://sparkwing.dev) pipeline
definitions. Pipelines are Go programs registered in `sparkwing.yaml` and run
via `sparkwing run <name>`.

Add a pipeline:

```
sparkwing pipeline new --pipeline <name>
```

## Layout

```
.sparkwing/
  .gitignore          ignores the cached pipeline binary
  sparkwing.yaml      registry of every pipeline (name -> entrypoint)
  jobs/               Go package holding pipeline definitions; scaffold lands one .go file per pipeline
  main.go             thin entrypoint; delegates to sparkwing.Main
  go.mod / go.sum     module + pinned SDK version
  sparkwing-pipeline  cached compiled binary (gitignored, regenerated)
```

## Agents

Run `sparkwing info --for-agent` for current, one-wake discovery context.
Do not copy runtime command catalogs into durable instruction files.
