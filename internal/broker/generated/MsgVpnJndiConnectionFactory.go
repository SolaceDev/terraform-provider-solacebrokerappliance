// terraform-provider-solacebroker
//
// Copyright 2023 Solace Corporation. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generated

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_jndi_connection_factory",
		MarkdownDescription: "The message broker provides an internal JNDI store for provisioned Connection Factory objects that clients can access through JNDI lookups.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\nconnectionFactoryName|x|||\nmsgVpnName|x|||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since 2.4.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/jndiConnectionFactories/{connectionFactoryName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "allowDuplicateClientIdEnabled",
				TerraformName:       "allow_duplicate_client_id_enabled",
				MarkdownDescription: "Enable or disable whether new JMS connections can use the same Client identifier (ID) as an existing connection. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "clientDescription",
				TerraformName:       "client_description",
				MarkdownDescription: "The description of the Client. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 256),
				},
				Default: "",
			},
			{
				SempName:            "clientId",
				TerraformName:       "client_id",
				MarkdownDescription: "The Client identifier (ID). If not specified, a unique value for it will be generated. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 256),
				},
				Default: "",
			},
			{
				SempName:            "connectionFactoryName",
				TerraformName:       "connection_factory_name",
				MarkdownDescription: "The name of the JMS Connection Factory.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 256),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^?* ]+$"), ""),
				},
			},
			{
				SempName:            "dtoReceiveOverrideEnabled",
				TerraformName:       "dto_receive_override_enabled",
				MarkdownDescription: "Enable or disable overriding by the Subscriber (Consumer) of the deliver-to-one (DTO) property on messages. When enabled, the Subscriber can receive all DTO tagged messages. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "dtoReceiveSubscriberLocalPriority",
				TerraformName:       "dto_receive_subscriber_local_priority",
				MarkdownDescription: "The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on the local broker that the Subscriber is directly connected to. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 4),
				},
				Default: 1,
			},
			{
				SempName:            "dtoReceiveSubscriberNetworkPriority",
				TerraformName:       "dto_receive_subscriber_network_priority",
				MarkdownDescription: "The priority for receiving deliver-to-one (DTO) messages by the Subscriber (Consumer) if the messages are published on a remote broker. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 4),
				},
				Default: 1,
			},
			{
				SempName:            "dtoSendEnabled",
				TerraformName:       "dto_send_enabled",
				MarkdownDescription: "Enable or disable the deliver-to-one (DTO) property on messages sent by the Publisher (Producer). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "dynamicEndpointCreateDurableEnabled",
				TerraformName:       "dynamic_endpoint_create_durable_enabled",
				MarkdownDescription: "Enable or disable whether a durable endpoint will be dynamically created on the broker when the client calls \"Session.createDurableSubscriber()\" or \"Session.createQueue()\". The created endpoint respects the message time-to-live (TTL) according to the \"dynamicEndpointRespectTtlEnabled\" property. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "dynamicEndpointRespectTtlEnabled",
				TerraformName:       "dynamic_endpoint_respect_ttl_enabled",
				MarkdownDescription: "Enable or disable whether dynamically created durable and non-durable endpoints respect the message time-to-live (TTL) property. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "guaranteedReceiveAckTimeout",
				TerraformName:       "guaranteed_receive_ack_timeout",
				MarkdownDescription: "The timeout for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer), in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(20, 1500),
				},
				Default: 1000,
			},
			{
				SempName:            "guaranteedReceiveReconnectRetryCount",
				TerraformName:       "guaranteed_receive_reconnect_retry_count",
				MarkdownDescription: "The maximum number of attempts to reconnect to the host or list of hosts after the guaranteed  messaging connection has been lost. The value \"-1\" means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `-1`. Available since 2.14.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-1, 2147483647),
				},
				Default: -1,
			},
			{
				SempName:            "guaranteedReceiveReconnectRetryWait",
				TerraformName:       "guaranteed_receive_reconnect_retry_wait",
				MarkdownDescription: "The amount of time to wait before making another attempt to connect or reconnect to the host after the guaranteed messaging connection has been lost, in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3000`. Available since 2.14.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(50, 2147483647),
				},
				Default: 3000,
			},
			{
				SempName:            "guaranteedReceiveWindowSize",
				TerraformName:       "guaranteed_receive_window_size",
				MarkdownDescription: "The size of the window for guaranteed messages received by the Subscriber (Consumer), in messages. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `18`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 255),
				},
				Default: 18,
			},
			{
				SempName:            "guaranteedReceiveWindowSizeAckThreshold",
				TerraformName:       "guaranteed_receive_window_size_ack_threshold",
				MarkdownDescription: "The threshold for sending the acknowledgement (ACK) for guaranteed messages received by the Subscriber (Consumer) as a percentage of `guaranteedReceiveWindowSize`. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `60`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 75),
				},
				Default: 60,
			},
			{
				SempName:            "guaranteedSendAckTimeout",
				TerraformName:       "guaranteed_send_ack_timeout",
				MarkdownDescription: "The timeout for receiving the acknowledgement (ACK) for guaranteed messages sent by the Publisher (Producer), in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `2000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(20, 60000),
				},
				Default: 2000,
			},
			{
				SempName:            "guaranteedSendWindowSize",
				TerraformName:       "guaranteed_send_window_size",
				MarkdownDescription: "The size of the window for non-persistent guaranteed messages sent by the Publisher (Producer), in messages. For persistent messages the window size is fixed at 1. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `255`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 255),
				},
				Default: 255,
			},
			{
				SempName:            "messagingDefaultDeliveryMode",
				TerraformName:       "messaging_default_delivery_mode",
				MarkdownDescription: "The default delivery mode for messages sent by the Publisher (Producer). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"persistent\"`. The allowed values and their meaning are:\n\n<pre>\n\"persistent\" - The broker spools messages (persists in the Message Spool) as part of the send operation.\n\"non-persistent\" - The broker does not spool messages (does not persist in the Message Spool) as part of the send operation.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("persistent", "non-persistent"),
				},
				Default: "persistent",
			},
			{
				SempName:            "messagingDefaultDmqEligibleEnabled",
				TerraformName:       "messaging_default_dmq_eligible_enabled",
				MarkdownDescription: "Enable or disable whether messages sent by the Publisher (Producer) are Dead Message Queue (DMQ) eligible by default. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "messagingDefaultElidingEligibleEnabled",
				TerraformName:       "messaging_default_eliding_eligible_enabled",
				MarkdownDescription: "Enable or disable whether messages sent by the Publisher (Producer) are Eliding eligible by default. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "messagingJmsxUserIdEnabled",
				TerraformName:       "messaging_jmsx_user_id_enabled",
				MarkdownDescription: "Enable or disable inclusion (adding or replacing) of the JMSXUserID property in messages sent by the Publisher (Producer). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "messagingTextInXmlPayloadEnabled",
				TerraformName:       "messaging_text_in_xml_payload_enabled",
				MarkdownDescription: "Enable or disable encoding of JMS text messages in Publisher (Producer) messages as XML payload. When disabled, JMS text messages are encoded as a binary attachment. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "msgVpnName",
				TerraformName:       "msg_vpn_name",
				MarkdownDescription: "The name of the Message VPN.",
				Identifying:         true,
				Required:            true,
				ReadOnly:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?]+$"), ""),
				},
			},
			{
				SempName:            "transportCompressionLevel",
				TerraformName:       "transport_compression_level",
				MarkdownDescription: "The ZLIB compression level for the connection to the broker. The value \"0\" means no compression, and the value \"-1\" means the compression level is specified in the JNDI Properties file. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `-1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-1, 9),
				},
				Default: -1,
			},
			{
				SempName:            "transportConnectRetryCount",
				TerraformName:       "transport_connect_retry_count",
				MarkdownDescription: "The maximum number of retry attempts to establish an initial connection to the host or list of hosts. The value \"0\" means a single attempt (no retries), and the value \"-1\" means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-1, 2147483647),
				},
				Default: 0,
			},
			{
				SempName:            "transportConnectRetryPerHostCount",
				TerraformName:       "transport_connect_retry_per_host_count",
				MarkdownDescription: "The maximum number of retry attempts to establish an initial connection to each host on the list of hosts. The value \"0\" means a single attempt (no retries), and the value \"-1\" means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-1, 2147483647),
				},
				Default: 0,
			},
			{
				SempName:            "transportConnectTimeout",
				TerraformName:       "transport_connect_timeout",
				MarkdownDescription: "The timeout for establishing an initial connection to the broker, in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `30000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 2147483647),
				},
				Default: 30000,
			},
			{
				SempName:            "transportDirectTransportEnabled",
				TerraformName:       "transport_direct_transport_enabled",
				MarkdownDescription: "Enable or disable usage of Direct Transport mode. When enabled, NON-PERSISTENT messages are sent as direct messages and non-durable topic consumers and temporary queue consumers consume using direct subscriptions rather than from guaranteed endpoints. If disabled all messaging uses guaranteed transport. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "transportKeepaliveCount",
				TerraformName:       "transport_keepalive_count",
				MarkdownDescription: "The maximum number of consecutive application-level keepalive messages sent without the broker response before the connection to the broker is closed. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(3, 2147483647),
				},
				Default: 3,
			},
			{
				SempName:            "transportKeepaliveEnabled",
				TerraformName:       "transport_keepalive_enabled",
				MarkdownDescription: "Enable or disable usage of application-level keepalive messages to maintain a connection with the broker. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "transportKeepaliveInterval",
				TerraformName:       "transport_keepalive_interval",
				MarkdownDescription: "The interval between application-level keepalive messages, in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(50, 2147483647),
				},
				Default: 3000,
			},
			{
				SempName:            "transportMsgCallbackOnIoThreadEnabled",
				TerraformName:       "transport_msg_callback_on_io_thread_enabled",
				MarkdownDescription: "Enable or disable delivery of asynchronous messages directly from the I/O thread. Contact support before enabling this property. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "transportOptimizeDirectEnabled",
				TerraformName:       "transport_optimize_direct_enabled",
				MarkdownDescription: "Enable or disable optimization for the Direct Transport delivery mode. If enabled, the client application is limited to one Publisher (Producer) and one non-durable Subscriber (Consumer). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "transportPort",
				TerraformName:       "transport_port",
				MarkdownDescription: "The connection port number on the broker for SMF clients. The value \"-1\" means the port is specified in the JNDI Properties file. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `-1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-1, 65535),
				},
				Default: -1,
			},
			{
				SempName:            "transportReadTimeout",
				TerraformName:       "transport_read_timeout",
				MarkdownDescription: "The timeout for reading a reply from the broker, in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 2147483647),
				},
				Default: 10000,
			},
			{
				SempName:            "transportReceiveBufferSize",
				TerraformName:       "transport_receive_buffer_size",
				MarkdownDescription: "The size of the receive socket buffer, in bytes. It corresponds to the SO_RCVBUF socket option. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `65536`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 2147483647),
				},
				Default: 65536,
			},
			{
				SempName:            "transportReconnectRetryCount",
				TerraformName:       "transport_reconnect_retry_count",
				MarkdownDescription: "The maximum number of attempts to reconnect to the host or list of hosts after the connection has been lost. The value \"-1\" means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(-1, 2147483647),
				},
				Default: 3,
			},
			{
				SempName:            "transportReconnectRetryWait",
				TerraformName:       "transport_reconnect_retry_wait",
				MarkdownDescription: "The amount of time before making another attempt to connect or reconnect to the host after the connection has been lost, in milliseconds. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `3000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 60000),
				},
				Default: 3000,
			},
			{
				SempName:            "transportSendBufferSize",
				TerraformName:       "transport_send_buffer_size",
				MarkdownDescription: "The size of the send socket buffer, in bytes. It corresponds to the SO_SNDBUF socket option. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `65536`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 2147483647),
				},
				Default: 65536,
			},
			{
				SempName:            "transportTcpNoDelayEnabled",
				TerraformName:       "transport_tcp_no_delay_enabled",
				MarkdownDescription: "Enable or disable the TCP_NODELAY option. When enabled, Nagle's algorithm for TCP/IP congestion control (RFC 896) is disabled. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "xaEnabled",
				TerraformName:       "xa_enabled",
				MarkdownDescription: "Enable or disable this as an XA Connection Factory. When enabled, the Connection Factory can be cast to \"XAConnectionFactory\", \"XAQueueConnectionFactory\" or \"XATopicConnectionFactory\". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
