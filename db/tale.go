package db

import (
	"fmt"

	"github.com/go-gorp/gorp"
)

type (
	// Tale describes a tale db entry
	Tale struct {
		ID        string `db:"tale_id"`
		Parent_ID string `db:"parent_id"`
		Content   string `db:"content"`
	}

	TaleFinder interface {
		// Find returns a Tale by ID
		Find(ID string) (*Tale, error)
	}

	TaleUpdater interface {
		// Update updates a Tale entry
		Update(b *Tale) error
	}

	TaleInserter interface {
		// Insert inserts a Tale into db
		Insert(b *Tale) error
	}

	taleManager struct {
		dbMap gorp.SqlExecutor
	}
)

// NewTaleFinder inits and returns an instance of TaleFinder
func NewTaleFinder(dbMap gorp.SqlExecutor) TaleFinder {
	return newTaleManager(dbMap)
}

// NewTaleUpdater inits and returns an instance of TaleUpdater
func NewTaleUpdater(dbMap gorp.SqlExecutor) TaleUpdater {
	return newTaleManager(dbMap)
}

// NewTaleInserter inits and returns an instance of TaleInserter
func NewTaleInserter(dbMap gorp.SqlExecutor) TaleInserter {
	return newTaleManager(dbMap)
}

func newTaleManager(dbMap gorp.SqlExecutor) *taleManager {
	return &taleManager{dbMap}
}

func (m *taleManager) Find(ID string) (*Tale, error) {
	var b Tale

	if err := m.dbMap.SelectOne(&b, "SELECT id, name FROM tale WHERE id = ?", ID); err != nil {
		return nil, fmt.Errorf("taleManager.Find: %s", err)
	}
	return &b, nil
}

func (m *taleManager) Update(b *Tale) error {
	_, err := m.dbMap.Update(b)
	if err != nil {
		return fmt.Errorf("taleManager.Update: %s", err)
	}
	return nil
}

func (m *taleManager) Insert(b *Tale) error {
	if err := m.dbMap.Insert(b); err != nil {
		return fmt.Errorf("taleManager.Insert: %s", err)
	}
	return nil
}
