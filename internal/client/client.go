/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package client

import (
	"github.com/simonvetter/modbus"
	"time"
)

func NewModbusClient() (*modbus.ModbusClient, error) {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     "tcp://hostname-or-ip-address:502",
		Timeout: 1 * time.Second,
		//other
	})
	//for an RTU over TCP device/bus (remote serial port or
	//simple TCP-to-serial bridge)
	//client, err = modbus.NewClient(&modbus.ClientConfiguration{
	//	URL:      "rtuovertcp://hostname-or-ip-address:502",
	//	Speed:    19200,                   // serial link speed
	//	Timeout:  1 * time.Second,
	//})
	if err != nil {
		// error out if client creation failed
		return nil, err
	}
	err = client.Open()
	if err != nil {
		// error out if we failed to connect/open the device
		// note: multiple Open() attempts can be made on the same client until
		// the connection succeeds (i.e. err == nil), calling the constructor again
		// is unnecessary.
		// likewise, a client can be opened and closed as many times as needed.
		return nil, err
	}

	//var reg16s []uint16
	//reg16s, err = client.ReadRegisters(100, 4, modbus.INPUT_REGISTER)

	return client, nil
}
