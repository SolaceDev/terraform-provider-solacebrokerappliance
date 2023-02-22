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
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "msg_vpn_topic_endpoint",
		MarkdownDescription: "A Topic Endpoint attracts messages published to a topic for which the Topic Endpoint has a matching topic subscription. The topic subscription for the Topic Endpoint is specified in the client request to bind a Flow to that Topic Endpoint. Queues are significantly more flexible than Topic Endpoints and are the recommended approach for most applications. The use of Topic Endpoints should be restricted to JMS applications.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\nmsg_vpn_name|x|||\ntopic_endpoint_name|x|||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"vpn/read-only\" is required to perform this operation.\n\nThis has been available since 2.4.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/msgVpns/{msgVpnName}/topicEndpoints/{topicEndpointName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "accessType",
				TerraformName:       "access_type",
				MarkdownDescription: "The access type for delivering messages to consumer flows bound to the Topic Endpoint. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"exclusive\"`. The allowed values and their meaning are:\n\n<pre>\n\"exclusive\" - Exclusive delivery of messages to the first bound consumer flow.\n\"non-exclusive\" - Non-exclusive delivery of messages to all bound consumer flows in a round-robin fashion.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("exclusive", "non-exclusive"),
				},
				Default: "exclusive",
			},
			{
				SempName:            "consumerAckPropagationEnabled",
				TerraformName:       "consumer_ack_propagation_enabled",
				MarkdownDescription: "Enable or disable the propagation of consumer acknowledgements (ACKs) received on the active replication Message VPN to the standby replication Message VPN. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "deadMsgQueue",
				TerraformName:       "dead_msg_queue",
				MarkdownDescription: "The name of the Dead Message Queue (DMQ) used by the Topic Endpoint. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"#DEAD_MSG_QUEUE\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 200),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]+$"), ""),
				},
				Default: "#DEAD_MSG_QUEUE",
			},
			{
				SempName:            "deliveryCountEnabled",
				TerraformName:       "delivery_count_enabled",
				MarkdownDescription: "Enable or disable the ability for client applications to query the message delivery count of messages received from the Topic Endpoint. This is a controlled availability feature. Please contact support to find out if this feature is supported for your use case. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.19.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "deliveryDelay",
				TerraformName:       "delivery_delay",
				MarkdownDescription: "The delay, in seconds, to apply to messages arriving on the Topic Endpoint before the messages are eligible for delivery. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
			},
			{
				SempName:            "egressEnabled",
				TerraformName:       "egress_enabled",
				MarkdownDescription: "Enable or disable the transmission of messages from the Topic Endpoint. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "eventBindCountThreshold",
				TerraformName:       "event_bind_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventRejectLowPriorityMsgLimitThreshold",
				TerraformName:       "event_reject_low_priority_msg_limit_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "eventSpoolUsageThreshold",
				TerraformName:       "event_spool_usage_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 18,
					},
					{
						SempName:            "clearValue",
						TerraformName:       "clear_value",
						MarkdownDescription: "The clear threshold for the absolute value of this counter. Falling below this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"set_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_percent"},
						ConflictsWith:       []string{"clear_value", "set_value"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_value"),
								path.MatchRelative().AtParent().AtName("set_value"),
							),
						},
						Default: 25,
					},
					{
						SempName:            "setValue",
						TerraformName:       "set_value",
						MarkdownDescription: "The set threshold for the absolute value of this counter. Exceeding this value will trigger a corresponding event. This attribute may not be returned in a GET.",
						Requires:            []string{"clear_value"},
						ConflictsWith:       []string{"clear_percent", "set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 4294967295),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_value"),
							),
							schemavalidator.ConflictsWith(
								path.MatchRelative().AtParent().AtName("clear_percent"),
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "ingressEnabled",
				TerraformName:       "ingress_enabled",
				MarkdownDescription: "Enable or disable the reception of messages to the Topic Endpoint. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "maxBindCount",
				TerraformName:       "max_bind_count",
				MarkdownDescription: "The maximum number of consumer flows that can bind to the Topic Endpoint. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 10000),
				},
				Default: 1,
			},
			{
				SempName:            "maxDeliveredUnackedMsgsPerFlow",
				TerraformName:       "max_delivered_unacked_msgs_per_flow",
				MarkdownDescription: "The maximum number of messages delivered but not acknowledged per flow for the Topic Endpoint. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 1000000),
				},
				Default: 10000,
			},
			{
				SempName:            "maxMsgSize",
				TerraformName:       "max_msg_size",
				MarkdownDescription: "The maximum message size allowed in the Topic Endpoint, in bytes (B). Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `10000000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 30000000),
				},
				Default: 1e+07,
			},
			{
				SempName:            "maxRedeliveryCount",
				TerraformName:       "max_redelivery_count",
				MarkdownDescription: "The maximum number of times the Topic Endpoint will attempt redelivery of a message prior to it being discarded or moved to the DMQ. A value of 0 means to retry forever. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 255),
				},
				Default: 0,
			},
			{
				SempName:            "maxSpoolUsage",
				TerraformName:       "max_spool_usage",
				MarkdownDescription: "The maximum message spool usage allowed by the Topic Endpoint, in megabytes (MB). A value of 0 only allows spooling of the last message received and disables quota checking. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `5000`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 6000000),
				},
				Default: 5000,
			},
			{
				SempName:            "maxTtl",
				TerraformName:       "max_ttl",
				MarkdownDescription: "The maximum time in seconds a message can stay in the Topic Endpoint when `respect_ttl_enabled` is `\"true\"`. A message expires when the lesser of the sender assigned time-to-live (TTL) in the message and the `max_ttl` configured for the Topic Endpoint, is exceeded. A value of 0 disables expiry. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
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
				SempName:            "owner",
				TerraformName:       "owner",
				MarkdownDescription: "The Client Username that owns the Topic Endpoint and has permission equivalent to `\"delete\"`. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 189),
				},
				Default: "",
			},
			{
				SempName:            "permission",
				TerraformName:       "permission",
				MarkdownDescription: "The permission level for all consumers of the Topic Endpoint, excluding the owner. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"no-access\"`. The allowed values and their meaning are:\n\n<pre>\n\"no-access\" - Disallows all access.\n\"read-only\" - Read-only access to the messages.\n\"consume\" - Consume (read and remove) messages.\n\"modify-topic\" - Consume messages or modify the topic/selector.\n\"delete\" - Consume messages, modify the topic/selector or delete the Client created endpoint altogether.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("no-access", "read-only", "consume", "modify-topic", "delete"),
				},
				Default: "no-access",
			},
			{
				SempName:            "redeliveryDelayEnabled",
				TerraformName:       "redelivery_delay_enabled",
				MarkdownDescription: "Enable or disable a message redelivery delay. When false, messages are redelivered as-soon-as-possible.  When true, messages are redelivered according to the initial, max and multiplier.  This should only be enabled when redelivery is enabled. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.33.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "redeliveryDelayInitialInterval",
				TerraformName:       "redelivery_delay_initial_interval",
				MarkdownDescription: "The delay to be used between the first 2 redelivery attempts.  This value is in milliseconds. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1000`. Available since 2.33.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 3600000),
				},
				Default: 1000,
			},
			{
				SempName:            "redeliveryDelayMaxInterval",
				TerraformName:       "redelivery_delay_max_interval",
				MarkdownDescription: "The maximum delay to be used between any 2 redelivery attempts.  This value is in milliseconds.  Due to technical limitations, some redelivery attempt delays may slightly exceed this value. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `64000`. Available since 2.33.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 10800000),
				},
				Default: 64000,
			},
			{
				SempName:            "redeliveryDelayMultiplier",
				TerraformName:       "redelivery_delay_multiplier",
				MarkdownDescription: "The amount each delay interval is multiplied by after each failed delivery attempt.  This number is in a fixed-point decimal format in which you must divide by 100 to get the floating point value. For example, a value of 125 would cause the delay to be multiplied by 1.25. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `200`. Available since 2.33.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(100, 500),
				},
				Default: 200,
			},
			{
				SempName:            "redeliveryEnabled",
				TerraformName:       "redelivery_enabled",
				MarkdownDescription: "Enable or disable message redelivery. When enabled, the number of redelivery attempts is controlled by max_redelivery_count. When disabled, the message will never be delivered from the topic-endpoint more than once. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Available since 2.18.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "rejectLowPriorityMsgEnabled",
				TerraformName:       "reject_low_priority_msg_enabled",
				MarkdownDescription: "Enable or disable the checking of low priority messages against the `reject_low_priority_msg_limit`. This may only be enabled if `reject_msg_to_sender_on_discard_behavior` does not have a value of `\"never\"`. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "rejectLowPriorityMsgLimit",
				TerraformName:       "reject_low_priority_msg_limit",
				MarkdownDescription: "The number of messages of any priority in the Topic Endpoint above which low priority messages are not admitted but higher priority messages are allowed. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `0`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 4294967295),
				},
				Default: 0,
			},
			{
				SempName:            "rejectMsgToSenderOnDiscardBehavior",
				TerraformName:       "reject_msg_to_sender_on_discard_behavior",
				MarkdownDescription: "Determines when to return negative acknowledgements (NACKs) to sending clients on message discards. Note that NACKs cause the message to not be delivered to any destination and Transacted Session commits to fail. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as reject_low_priority_msg_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `\"never\"`. The allowed values and their meaning are:\n\n<pre>\n\"always\" - Always return a negative acknowledgment (NACK) to the sending client on message discard.\n\"when-topic-endpoint-enabled\" - Only return a negative acknowledgment (NACK) to the sending client on message discard when the Topic Endpoint is enabled.\n\"never\" - Never return a negative acknowledgment (NACK) to the sending client on message discard.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("always", "when-topic-endpoint-enabled", "never"),
				},
				Default: "never",
			},
			{
				SempName:            "respectMsgPriorityEnabled",
				TerraformName:       "respect_msg_priority_enabled",
				MarkdownDescription: "Enable or disable the respecting of message priority. When enabled, messages contained in the Topic Endpoint are delivered in priority order, from 9 (highest) to 0 (lowest). Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as egress_enabled and ingress_enabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Available since 2.8.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "respectTtlEnabled",
				TerraformName:       "respect_ttl_enabled",
				MarkdownDescription: "Enable or disable the respecting of the time-to-live (TTL) for messages in the Topic Endpoint. When enabled, expired messages are discarded or moved to the DMQ. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "topicEndpointName",
				TerraformName:       "topic_endpoint_name",
				MarkdownDescription: "The name of the Topic Endpoint.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 200),
					stringvalidator.RegexMatches(regexp.MustCompile("^[^*?'<>&;]+$"), ""),
				},
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
