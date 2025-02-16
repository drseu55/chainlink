package chainlink

import (
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap/zapcore"

	v2 "github.com/smartcontractkit/chainlink/core/config/v2"
)

func (g *generalConfig) AppID() uuid.UUID {
	g.appIDOnce.Do(func() {
		if g.c.AppID != (uuid.UUID{}) {
			return // already set (e.g. test override)
		}
		g.c.AppID = uuid.NewV4() // randomize
	})
	return g.c.AppID
}

func (g *generalConfig) DefaultLogLevel() zapcore.Level {
	return g.logLevelDefault
}

func (g *generalConfig) LogLevel() (ll zapcore.Level) {
	g.logMu.RLock()
	ll = zapcore.Level(*g.c.Log.Level)
	g.logMu.RUnlock()
	return
}

func (g *generalConfig) SetLogLevel(lvl zapcore.Level) error {
	g.logMu.Lock()
	g.c.Log.Level = (*v2.LogLevel)(&lvl)
	g.logMu.Unlock()
	return nil
}

func (g *generalConfig) LogSQL() (sql bool) {
	g.logMu.RLock()
	sql = *g.c.Database.LogQueries
	g.logMu.RUnlock()
	return
}

func (g *generalConfig) SetLogSQL(logSQL bool) {
	g.logMu.Lock()
	g.c.Database.LogQueries = &logSQL
	g.logMu.Unlock()
}
