package portal

import (
	"context"
	"os"

	"github.com/apex/log"
	"github.com/dpogorzelski/speedrun/proto/portal"
)

func (s *Server) FileRead(ctx context.Context, file *portal.FileReadRequest) (*portal.FileReadResponse, error) {
	fields := log.Fields{
		"context": "file",
		"command": "read",
		"name":    file.GetPath(),
	}
	log := log.WithFields(fields)
	log.Debug("Received file read request")

	content, err := os.ReadFile(file.GetPath())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &portal.FileReadResponse{State: portal.State_UNKNOWN, Content: string(content)}, nil
}
