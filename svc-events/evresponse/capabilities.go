//(C) Copyright [2022] Hewlett Packard Enterprise Development LP
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

// Package evresponse have error and response struct
// and also have functionality to create error response
package evresponse

// SubscriptionCreationCapibilitiesResp ...
var SubscriptionCreationCapibilitiesResp = map[string]interface{}{
	"@odata.id":                                     "/redfish/v1/EventService/Subscriptions/Capabilities",
	"@odata.type":                                   "#EventDestination.v1_12_0.EventDestination",
	"Name":                                          "Capabilities for the subscription collection",
	"ID":                                            "Capabilities",
	"Name@Redfish.OptionalOnCreate":                 true,
	"Name@Redfish.SetOnlyOnCreate":                  true,
	"Destination@Redfish.RequiredOnCreate":          true,
	"Destination@Redfish.SetOnlyOnCreate":           true,
	"EventTypes@Redfish.OptionalOnCreate":           true,
	"EventTypes@Redfish.SetOnlyOnCreate":            true,
	"EventTypes@Redfish.AllowableValues":            []interface{}{"Alert", "MetricReport", "ResourceAdded", "ResourceRemoved", "ResourceUpdated", "StatusChange", "Other"},
	"MessageIds@Redfish.OptionalOnCreate":           true,
	"MessageIds@Redfish.SetOnlyOnCreate":            true,
	"ResourceTypes@Redfish.OptionalOnCreate":        true,
	"ResourceTypes@Redfish.SetOnlyOnCreate":         true,
	"ResourceTypes@Redfish.AllowableValues":         []interface{}{"Job", "MemoryMetrics", "SerialInterfaces", "Switch", "Bios", "NetworkPort", "ComputerSystem", "MessageRegistry", "Thermal", "Event", "Manager", "Port", "Drive", "NetworkInterface", "PhysicalContext", "Processor", "Protocol", "BootOption", "IPAddresses", "MemoryChunks", "Storage", "Task", "AddressPool", "ManagerAccount", "ProcessorCollection", "VLanNetworkInterface", "Assembly", "EventService", "MemoryDomain", "MessageRegistryFile", "NetworkDeviceFunction", "PCIeFunction", "PrivilegeRegistry", "Resource", "Sensor", "EthernetInterface", "Fabric", "ManagerNetworkProtocol", "Redundancy", "Role", "SecureBoot", "Endpoint", "JobService", "NetworkAdapter", "Session", "LogEntry", "Memory", "Message", "Privileges", "ProcessorMetrics", "Chassis", "LogService", "PCIeDevice", "AccelerationFunction", "EventDestination", "HostInterface", "PCIeSlots", "Power", "Volume", "Zone"},
	"Context@Redfish.OptionalOnCreate":              true,
	"Context@Redfish.SetOnlyOnCreate":               true,
	"Protocol@Redfish.RequiredOnCreate":             true,
	"Protocol@Redfish.SetOnlyOnCreate":              true,
	"Protocol@Redfish.AllowableValues":              []interface{}{"Redfish"},
	"SubscriptionType@Redfish.OptionalOnCreate":     true,
	"SubscriptionType@Redfish.SetOnlyOnCreate":      true,
	"SubscriptionType@Redfish.AllowableValues":      []interface{}{"RedfishEvent"},
	"EventFormatType@Redfish.OptionalOnCreate":      true,
	"EventFormatType@Redfish.SetOnlyOnCreate":       true,
	"EventFormatType@Redfish.AllowableValues":       []interface{}{"Event", "MetricReport"},
	"SubordinateResources@Redfish.OptionalOnCreate": true,
	"SubordinateResources@Redfish.SetOnlyOnCreate":  true,
	"OriginResources@Redfish.OptionalOnCreate":      true,
	"OriginResources@Redfish.SetOnlyOnCreate":       true,
	"DeliveryRetryPolicy@Redfish.OptionalOnCreate":  true,
	"DeliveryRetryPolicy@Redfish.SetOnlyOnCreate":   true,
}
