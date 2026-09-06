# Mini Governata Backend

This is the cumulative Go backend for Mini Governata. Phase 0 adds the first domain, application, transport, infrastructure, and context packages. Later phases will extend these same packages instead of creating separate applications.

## Run it

```bash
cd /Users/hcsingh/Documents/Min-Governata/backend
go run ./cmd/trace-lab
```

Run the tests:

```bash
go test ./...
```

## Folder purposes

```text
backend/
├── go.mod
├── cmd/
│   └── trace-lab/main.go
└── internal/
    ├── domain/task/
    ├── application/workflow/
    ├── transport/handler/
    ├── infrastructure/memory/
    └── platform/tracectx/
```

### `backend/`

Contains the entire Go backend. The React application will live separately under `frontend/` beginning in Phase 1.

### `go.mod`

Defines the Go module name, Go version, and external dependencies. It is comparable to part of `package.json`, but for Go. Commands such as `go test ./...` use it to locate the module root.

### `cmd/`

Contains executable entrypoints. Each child directory builds one executable:

```text
cmd/trace-lab  → Phase 0 executable
cmd/api        → HTTP server added later
cmd/worker     → background worker added later
```

Files in `cmd` should mainly create concrete dependencies, connect them, and start the process. Business rules do not belong here.

### `internal/`

Contains application code that should not be imported by unrelated external Go modules. Go enforces this special directory rule. It protects implementation details while allowing Mini Governata packages to use them.

### `internal/domain/`

Contains business concepts and rules that do not depend on Gin, GORM, NATS, or other infrastructure.

### `internal/domain/task/`

Contains the `Task` entity, task statuses, and transition rules. `task` is one domain package inside `domain`; later we can add packages such as `asset`, `policy`, or `audit`.

### `internal/application/`

Coordinates use cases using domain objects and small interfaces. It answers: “What steps must happen to complete this user action?”

### `internal/application/workflow/`

Contains `Engine` and the consumer-owned `TaskStore` interface. It finds a task, applies the domain completion rule, and asks storage to save it.

### `internal/transport/`

Translates an outside request into an application call and translates the result back. The current handler is a simulation; Gin HTTP handlers will be placed here in Phase 3.

### `internal/infrastructure/`

Contains technical implementations of interfaces: memory today, PostgreSQL/GORM later, followed by NATS and object storage adapters.

### `internal/infrastructure/memory/`

Provides the current concrete `TaskStore`. It lets us run and test the workflow without a database.

### `internal/platform/`

Contains small technical capabilities used across boundaries, such as trace context, configuration, logging, and eventually observability setup.

### `internal/platform/tracectx/`

Adds and retrieves a Trace ID from `context.Context`, allowing one operation to be followed through every layer.

## Read the Phase 0 code in this order

1. `cmd/trace-lab/main.go` — the composition root creates and connects every concrete object.
2. `internal/transport/handler/task_handler.go` — simulates the future HTTP boundary and directly calls the engine.
3. `internal/application/workflow/engine.go` — declares `TaskStore`, applies the use case, and calls storage through the interface.
4. `internal/domain/task/task.go` — contains the business rule for completing a task.
5. `internal/infrastructure/memory/task_store.go` — the concrete implementation selected at runtime.
6. `internal/platform/tracectx/trace.go` — carries one Trace ID through every layer.

## Trace map

```text
main.go
  creates MemoryStore → injects it into Engine → injects Engine into Handler

Transport Handler.Complete(ctx, id)       direct call
  ↓
Engine.CompleteTask(ctx, id)              direct call
  ↓
TaskStore.Find(ctx, id)                   interface dispatch
  ↓
MemoryStore.Find(ctx, id)                 concrete method
  ↓
Task.Complete()                           direct domain call
  ↓
TaskStore.Save(ctx, task)                 interface dispatch
  ↓
MemoryStore.Save(ctx, task)               future database boundary
```

## Important lines to explain aloud

```go
engine := workflow.NewEngine(store)
```

`store` is a `*memory.Store`. Go accepts it because it has the methods required by `workflow.TaskStore`.

```go
value, err := e.store.Find(ctx, id)
```

`ctx` and `id` are inputs. `value` and `err` receive the two outputs. The concrete call goes to `MemoryStore.Find` because `main.go` injected that implementation.

```go
if err := e.store.Save(ctx, value); err != nil
```

`ctx` and `value` are inputs to `Save`. `Save` returns only an error, stored in `err` for the scope of the `if` statement.

## Phase boundary

- Phase 0: establish and trace these backend connections.
- Phase 1: build the visible React workspace.
- Phase 2: deepen the Go domain and CLI skills.
- Phase 3: replace the simulated handler boundary with real Gin HTTP routes.
- Phase 5: replace `MemoryStore` with PostgreSQL/GORM.
