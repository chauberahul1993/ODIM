//(C) Copyright [2020] Hewlett Packard Enterprise Development LP
//
//Licensed under the Apache License, Version 2.0 (the "License"); you may
//not use this file except in compliance with the License. You may obtain
//a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//License for the specific language governing permissions and limitations
// under the License.

// Package common ...
package common

import (
	"fmt"

	"github.com/ODIM-Project/ODIM/lib-persistence-manager/persistencemgr"
	"github.com/ODIM-Project/ODIM/lib-utilities/errors"
)

// TruncateDB will clear DB. It will be useful for test cases
// Takes DbFlag of type DbType/int32 to choose Inmemory or OnDisk db to truncate
//dbFlag:
//    InMemory: Truncates InMemory DB
//    OnDisk: Truncates OnDisk DB
func TruncateDB(dbFlag persistencemgr.DbType) *errors.Error {
	conn, err := persistencemgr.GetDBConnection(dbFlag)
	if err != nil {
		return errors.PackError(err.ErrNo(), "unable to connect DB: ", err.Error())
	}
	err = conn.CleanUpDB()
	if err != nil {
		return errors.PackError(err.ErrNo(), "unable to flush out DB: ", err.Error())
	}
	return nil
}

// CheckDBConnection will check both inMemory and onDisk DB connections
// This function is expected to be called at each service startup
func CheckDBConnection() error {
	inMemConn, err := persistencemgr.GetDBConnection(persistencemgr.InMemory)
	if err != nil {
		return fmt.Errorf("unable to create InMemory DB connection: %v", err)
	}
	onDiskConn, err := persistencemgr.GetDBConnection(persistencemgr.OnDisk)
	if err != nil {
		return fmt.Errorf("unable to create OnDisk DB connection: %v", err)
	}

	if err := inMemConn.Ping(); err != nil {
		return fmt.Errorf("unable to ping InMemory DB: %v", err)
	}
	if err := onDiskConn.Ping(); err != nil {
		return fmt.Errorf("unable to ping OnDisk DB: %v", err)
	}

	return nil
}
