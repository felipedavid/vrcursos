package repository

import (
	"context"

	"github.com/felipedavid/vrcursos/src/core/model"
)

type ISudentRepository interface {
	Save(ctx context.Context, student model.Student)
}
