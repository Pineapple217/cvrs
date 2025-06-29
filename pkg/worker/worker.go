package worker

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/Pineapple217/cvrs/pkg/config"
	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/Pineapple217/cvrs/pkg/ent/task"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

type Workforce struct {
	db      *database.Database
	workers []*Worker
	wg      sync.WaitGroup
	tasks   chan *ent.Task
	ctx     context.Context
	cancel  context.CancelFunc
}

type Worker struct {
	db     *database.Database
	id     string
	logger *slog.Logger
	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorkforce(conf config.Workforce, db *database.Database) *Workforce {
	ctx, cancel := context.WithCancel(context.Background())
	ws := []*Worker{}

	for i := range conf.MaxWorkers {
		name := fmt.Sprintf("%d", i)
		logger := slog.With(slog.Group("worker"), slog.String("id", name))
		workerCtx, workerCancel := context.WithCancel(ctx)
		ws = append(ws, &Worker{
			logger: logger,
			ctx:    workerCtx,
			cancel: workerCancel,
			db:     db,
			id:     name,
		})
	}

	w := &Workforce{
		workers: ws,
		tasks:   make(chan *ent.Task, 20),
		db:      db,
		wg:      sync.WaitGroup{},
		ctx:     ctx,
		cancel:  cancel,
	}
	return w
}

func (wf *Workforce) Start() error {
	slog.Info("Starting workforce")
	c, err := wf.db.Client.Task.Update().
		SetStatus(task.StatusPending).
		Where(task.StatusEQ(task.StatusWorking)).
		Save(wf.ctx)
	if err != nil {
		return err
	}
	if c > 0 {
		slog.Info("restored working jobs to pending", "count", c)
	}

	wf.wg.Add(len(wf.workers))
	for _, w := range wf.workers {
		w.Start(&wf.wg, wf.tasks)
	}
	go wf.Fetcher()

	return nil
}

func (wf *Workforce) Stop() {
	slog.Info("Stopping workforce")
	wf.cancel()
	wf.wg.Wait()
}

func (wf *Workforce) Fetcher() {
	wf.wg.Add(1)
	defer wf.wg.Done()
	for {
		select {
		case <-wf.ctx.Done():
			slog.Info("stopped job fetcher")
			return
		default:
			if len(wf.tasks) > 10 {
				time.Sleep(time.Second)
				continue
			}

			tx, err := wf.db.Client.BeginTx(wf.ctx, &sql.TxOptions{})
			if err != nil {
				slog.Warn("failed to fetch tasks", "error", err)
				continue
			}
			tasks, err := tx.Task.Query().
				Where(task.StatusEQ(task.StatusPending)).
				Order(ent.Asc(task.FieldCreatedAt)).
				Limit(10).
				All(wf.ctx)
			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					err = fmt.Errorf("%w: %v", err, rerr)
				}
				slog.Warn("failed to fetch tasks", "error", err)
				continue
			}
			ids := make([]pid.ID, len(tasks))
			for i, t := range tasks {
				ids[i] = t.ID
			}

			// Update their status to 'working'
			_, err = tx.Task.
				Update().
				Where(task.IDIn(ids...)).
				SetStatus(task.StatusWorking).
				Save(wf.ctx)
			if err != nil {
				if rerr := tx.Rollback(); rerr != nil {
					err = fmt.Errorf("%w: %v", err, rerr)
				}
				slog.Warn("failed to fetch tasks", "error", err)
				continue
			}
			err = tx.Commit()
			if err != nil {
				slog.Warn("failed to commit transaction", "error", err)
				continue
			}

			if len(tasks) == 0 {
				slog.Debug("no jobs to fetch, waiting", "sec", 3)
				time.Sleep(time.Second * 3)
			} else {
				slog.Info("adding tracks to task queue", "count", len(tasks))
			}
			for _, task := range tasks {
				wf.tasks <- task
			}
		}
	}
}

func (w *Worker) Start(wg *sync.WaitGroup, tasks chan *ent.Task) {
	go func() {
		defer wg.Done()
		for {
			select {
			case <-w.ctx.Done():
				w.logger.Info("stopping worker")
				return
			default:
				if len(tasks) == 0 {
					time.Sleep(time.Second)
					continue
				}
				t := <-tasks
				w.logger.Info("working", "task", t.ID.String(), "task_id", t.ID.Int(), "type", t.Type)
				err := w.proccesTask(t)
				if err != nil {
					w.logger.Warn("failed to process task", "task", t.ID.String(), "task_id", t.ID.Int(), "error", err)
					err = t.Update().
						SetError(err.Error()).
						SetStatus(task.StatusError).
						Exec(w.ctx)
					if err != nil {
						slog.Error("failed to safe task errro", "error", err)
						w.Stop()
					}
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	w.logger.Info("stopping")
	w.cancel()
}

func (w *Worker) proccesTask(t *ent.Task) error {
	switch t.Type {
	case task.TypeScaleImg:
		err := ScaleImg(t, w.db, w.ctx)
		if err != nil {
			return err
		}
		return w.db.Client.Task.UpdateOne(t).SetStatus(task.StatusDone).Exec(w.ctx)
	default:
		return fmt.Errorf("%s is not a valid task type", t.Type)
	}
}
