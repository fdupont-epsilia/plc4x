/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package bacnetip

// TODO: implement
type Capability struct {
}

func NewCapability() *Capability {
	return &Capability{}
}

// TODO: implement
type Collector struct {
}

func NewCollector() *Collector {
	return &Collector{}
}

func (c *Collector) searchCapability() {
	// TODO: implement
	return
}

func (c *Collector) CapabilityFunctions(fn string) []func(args Args, kwargs KWArgs) error {
	// TODO: implement
	return nil
}

func (c *Collector) AddCapability(cls any) {
	// TODO: implement
	return
}