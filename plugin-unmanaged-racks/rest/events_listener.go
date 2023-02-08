/*
 * Copyright (c) 2020 Intel Corporation
 * Copyright (c) 2020 Hewlett Packard Enterprise Development LP
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rest

import (
	stdCtx "context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ODIM-Project/ODIM/plugin-unmanaged-racks/db"
	"github.com/ODIM-Project/ODIM/plugin-unmanaged-racks/logging"
	"github.com/ODIM-Project/ODIM/plugin-unmanaged-racks/redfish"

	"github.com/go-redis/redis/v8"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func newEventHandler(dao *db.DAO) context.Handler {
	return (&eventHandler{
		dao: dao,
	}).handleEvent
}

type eventHandler struct {
	dao *db.DAO
}

func (eh *eventHandler) handleEvent(c iris.Context) {
	fmt.Println("Delete is called *************** Event received  ")
	raw := new(json.RawMessage)
	err := c.ReadJSON(raw)
	if err != nil {
		fmt.Println("Step 111   ", raw)
		logging.Errorf("Cannot read message body: %s", err)
		c.StatusCode(http.StatusBadRequest)
		return
	}
	fmt.Println("Step 2222   ", raw)
	message := new(redfish.Event)
	err = json.Unmarshal([]byte(redfish.Translator.RedfishToODIM(string(*raw))), message)
	if err != nil {
		fmt.Println("Step 22222222    ", raw)
		logging.Errorf("Cannot load '%s' message into redfish.Event struct: %s", string(*raw), err)
		c.StatusCode(http.StatusBadRequest)
		return
	}
	fmt.Printf("Step 33333  %+v \n ", message)
	for _, e := range message.Events {
		fmt.Println("Origin of condition ", e.OriginOfCondition.Oid)
		fmt.Printf("Event %+v \n", e)
		ctx := stdCtx.TODO()
		containedInKey := db.CreateContainedInKey("Chassis", e.OriginOfCondition.Oid)
		rackKey, err := eh.dao.Get(ctx, containedInKey.String()).Result()
		fmt.Printf("Step 5555  %s %s %s \n ", containedInKey, rackKey, err)
		if err == redis.Nil {
			fmt.Printf("Step 6666  %+v \n ", err)
			continue
		}
		if err != nil {
			fmt.Printf("Step 6666***  %+v \n ", err)
			continue
		}
		fmt.Printf("Step 7777  %+v \n ", err)
		err = eh.dao.Watch(ctx, func(tx *redis.Tx) error {
			_, err := tx.TxPipelined(ctx, func(pipeliner redis.Pipeliner) error {
				if _, err := pipeliner.Del(
					ctx,
					containedInKey.String(),
				).Result(); err != nil {
					fmt.Printf("Step 8888  %+v \n ", err)
					return fmt.Errorf("del: %s error: %w", containedInKey, err)
				}

				if _, err := pipeliner.SRem(
					ctx,
					db.CreateContainsKey("Chassis", rackKey).String(), e.OriginOfCondition.Oid,
				).Result(); err != nil {
					fmt.Printf("Step 9999  %+v \n ", err)
					return fmt.Errorf("srem: %s error: %w", db.CreateContainsKey("Chassis", rackKey).String(), err)
				}
				return nil
			})
			return err
		}, rackKey)
		fmt.Printf("Step 1010  %+v \n ", err)

		if err != nil {
			fmt.Printf("Step 1011  %+v \n ", err)

			logging.Errorf(
				"cannot consume message(%v): %v",
				message,
				fmt.Errorf("couldn't commit transaction: %w", err),
			)
		}
	}
	fmt.Printf("Step 1022  %+v \n ", err)
}
