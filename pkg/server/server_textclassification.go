// Copyright 2022 NLP Odyssey Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	textclassificationv1 "github.com/nlpodyssey/cybertron/pkg/server/gen/proto/go/textclassification/v1"
	"github.com/nlpodyssey/cybertron/pkg/tasks/textclassification"
	"google.golang.org/grpc"
)

// serverForTextClassification is a server that provides gRPC and HTTP/2 APIs for Text Classification task.
type serverForTextClassification struct {
	textclassificationv1.UnimplementedTextClassificationServiceServer
	classifier textclassification.Interface
}

// registerTextClassificationFunc registers the Text Classification functions.
func registerTextClassificationFunc(classifier textclassification.Interface) (*RegisterFuncs, error) {
	s := &serverForTextClassification{classifier: classifier}
	return &RegisterFuncs{
		RegisterServer: func(r grpc.ServiceRegistrar) error {
			textclassificationv1.RegisterTextClassificationServiceServer(r, s)
			return nil
		},
		RegisterHandlerServer: func(ctx context.Context, mux *runtime.ServeMux) error {
			return textclassificationv1.RegisterTextClassificationServiceHandlerServer(ctx, mux, s)
		},
	}, nil
}

// Classify handles the Classify request.
func (s *serverForTextClassification) Classify(_ context.Context, req *textclassificationv1.ClassifyRequest) (*textclassificationv1.ClassifyResponse, error) {
	result, err := s.classifier.Classify(req.GetInput())
	if err != nil {
		return nil, err
	}

	labels := make([]string, len(result.Labels))
	for i := range result.Labels {
		labels[i] = fmt.Sprintf("%d", result.Labels[i]) // TODO: use label names
	}

	resp := &textclassificationv1.ClassifyResponse{
		Labels: labels,
		Scores: result.Scores,
	}
	return resp, nil
}