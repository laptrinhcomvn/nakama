// Copyright 2019 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/heroiclabs/nakama/console"
)

func (s *ConsoleServer) BanUser(ctx context.Context, in *console.AccountId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *ConsoleServer) UnbanUser(ctx context.Context, in *console.AccountId) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *ConsoleServer) DeleteUsers(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *ConsoleServer) ListUsers(ctx context.Context, in *console.ListUsersRequest) (*console.UserList, error) {
	return &console.UserList{}, nil
}