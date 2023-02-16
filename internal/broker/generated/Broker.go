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
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "broker",
		MarkdownDescription: "This object contains global configuration for the message broker.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\ntlsServerCertContent||x||x\ntlsServerCertPassword||x||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/none\" is required to perform this operation. Requests which include the following attributes require greater access scope/level:\n\n\nAttribute|Access Scope/Level\n:---|:---:\nconfigSyncAuthenticationClientCertMaxChainDepth|global/read-only\nconfigSyncAuthenticationClientCertValidateDateEnabled|global/read-only\nconfigSyncClientProfileTcpInitialCongestionWindow|global/read-only\nconfigSyncClientProfileTcpKeepaliveCount|global/read-only\nconfigSyncClientProfileTcpKeepaliveIdle|global/read-only\nconfigSyncClientProfileTcpKeepaliveInterval|global/read-only\nconfigSyncClientProfileTcpMaxWindow|global/read-only\nconfigSyncClientProfileTcpMss|global/read-only\nconfigSyncEnabled|global/read-only\nconfigSyncSynchronizeUsernameEnabled|global/read-only\nconfigSyncTlsEnabled|global/read-only\nguaranteedMsgingDefragmentationScheduleDayList|global/read-only\nguaranteedMsgingDefragmentationScheduleEnabled|global/read-only\nguaranteedMsgingDefragmentationScheduleTimeList|global/read-only\nguaranteedMsgingDefragmentationThresholdEnabled|global/read-only\nguaranteedMsgingDefragmentationThresholdFragmentationPercentage|global/read-only\nguaranteedMsgingDefragmentationThresholdMinInterval|global/read-only\nguaranteedMsgingDefragmentationThresholdUsagePercentage|global/read-only\nguaranteedMsgingDiskArrayWwn|global/read-only\nguaranteedMsgingDiskLocation|global/read-only\nguaranteedMsgingEnabled|global/read-only\nguaranteedMsgingEventCacheUsageThreshold.clearPercent|global/read-only\nguaranteedMsgingEventCacheUsageThreshold.clearValue|global/read-only\nguaranteedMsgingEventCacheUsageThreshold.setPercent|global/read-only\nguaranteedMsgingEventCacheUsageThreshold.setValue|global/read-only\nguaranteedMsgingEventDeliveredUnackedThreshold.clearPercent|global/read-only\nguaranteedMsgingEventDeliveredUnackedThreshold.setPercent|global/read-only\nguaranteedMsgingEventDiskUsageThreshold.clearPercent|global/read-only\nguaranteedMsgingEventDiskUsageThreshold.setPercent|global/read-only\nguaranteedMsgingEventEgressFlowCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventEgressFlowCountThreshold.clearValue|global/read-only\nguaranteedMsgingEventEgressFlowCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventEgressFlowCountThreshold.setValue|global/read-only\nguaranteedMsgingEventEndpointCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventEndpointCountThreshold.clearValue|global/read-only\nguaranteedMsgingEventEndpointCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventEndpointCountThreshold.setValue|global/read-only\nguaranteedMsgingEventIngressFlowCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventIngressFlowCountThreshold.clearValue|global/read-only\nguaranteedMsgingEventIngressFlowCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventIngressFlowCountThreshold.setValue|global/read-only\nguaranteedMsgingEventMsgCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventMsgCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventMsgSpoolFileCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventMsgSpoolFileCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventMsgSpoolUsageThreshold.clearPercent|global/read-only\nguaranteedMsgingEventMsgSpoolUsageThreshold.clearValue|global/read-only\nguaranteedMsgingEventMsgSpoolUsageThreshold.setPercent|global/read-only\nguaranteedMsgingEventMsgSpoolUsageThreshold.setValue|global/read-only\nguaranteedMsgingEventTransactedSessionCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventTransactedSessionCountThreshold.clearValue|global/read-only\nguaranteedMsgingEventTransactedSessionCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventTransactedSessionCountThreshold.setValue|global/read-only\nguaranteedMsgingEventTransactedSessionResourceCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventTransactedSessionResourceCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventTransactionCountThreshold.clearPercent|global/read-only\nguaranteedMsgingEventTransactionCountThreshold.clearValue|global/read-only\nguaranteedMsgingEventTransactionCountThreshold.setPercent|global/read-only\nguaranteedMsgingEventTransactionCountThreshold.setValue|global/read-only\nguaranteedMsgingMaxCacheUsage|global/read-only\nguaranteedMsgingMaxMsgSpoolUsage|global/read-only\nguaranteedMsgingTransactionReplicationCompatibilityMode|global/read-only\nguaranteedMsgingVirtualRouterWhenActiveActive|global/read-only\noauthProfileDefault|global/read-only\nserviceAmqpEnabled|global/read-only\nserviceAmqpTlsListenPort|global/read-only\nserviceEventConnectionCountThreshold.clearPercent|global/read-only\nserviceEventConnectionCountThreshold.clearValue|global/read-only\nserviceEventConnectionCountThreshold.setPercent|global/read-only\nserviceEventConnectionCountThreshold.setValue|global/read-only\nserviceHealthCheckEnabled|global/read-only\nserviceHealthCheckListenPort|global/read-only\nserviceMqttEnabled|global/read-only\nserviceMsgBackboneEnabled|global/read-only\nserviceRestEventOutgoingConnectionCountThreshold.clearPercent|global/read-only\nserviceRestEventOutgoingConnectionCountThreshold.clearValue|global/read-only\nserviceRestEventOutgoingConnectionCountThreshold.setPercent|global/read-only\nserviceRestEventOutgoingConnectionCountThreshold.setValue|global/read-only\nserviceRestIncomingEnabled|global/read-only\nserviceRestOutgoingEnabled|global/read-only\nserviceSempCorsAllowAnyHostEnabled|global/read-only\nserviceSempLegacyTimeoutEnabled|global/read-only\nserviceSempPlainTextEnabled|global/read-only\nserviceSempPlainTextListenPort|global/read-only\nserviceSempSessionIdleTimeout|global/read-only\nserviceSempSessionMaxLifetime|global/read-only\nserviceSempTlsEnabled|global/read-only\nserviceSempTlsListenPort|global/read-only\nserviceSmfCompressionListenPort|global/read-only\nserviceSmfEnabled|global/read-only\nserviceSmfEventConnectionCountThreshold.clearPercent|global/read-only\nserviceSmfEventConnectionCountThreshold.clearValue|global/read-only\nserviceSmfEventConnectionCountThreshold.setPercent|global/read-only\nserviceSmfEventConnectionCountThreshold.setValue|global/read-only\nserviceSmfPlainTextListenPort|global/read-only\nserviceSmfRoutingControlListenPort|global/read-only\nserviceSmfTlsListenPort|global/read-only\nserviceTlsEventConnectionCountThreshold.clearPercent|global/read-only\nserviceTlsEventConnectionCountThreshold.clearValue|global/read-only\nserviceTlsEventConnectionCountThreshold.setPercent|global/read-only\nserviceTlsEventConnectionCountThreshold.setValue|global/read-only\nserviceWebTransportEnabled|global/read-only\nserviceWebTransportPlainTextListenPort|global/read-only\nserviceWebTransportTlsListenPort|global/read-only\nserviceWebTransportWebUrlSuffix|global/read-only\ntlsBlockVersion10Enabled|global/read-only\ntlsBlockVersion11Enabled|global/read-only\ntlsCipherSuiteManagementList|global/read-only\ntlsCipherSuiteMsgBackboneList|global/read-only\ntlsCipherSuiteSecureShellList|global/read-only\ntlsCrimeExploitProtectionEnabled|global/read-only\ntlsServerCertContent|global/read-only\ntlsStandardDomainCertificateAuthoritiesEnabled|vpn/read-only\ntlsTicketLifetime|global/read-only\nwebManagerAllowUnencryptedWizardsEnabled|vpn/read-only\nwebManagerCustomization|vpn/read-only\nwebManagerRedirectHttpEnabled|vpn/read-only\nwebManagerRedirectHttpOverrideTlsPort|vpn/read-only\n\n\n\nThis has been available since 2.13.",
		ObjectType:          broker.SingletonObject,
		PathTemplate:        "/",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				SempName:            "authClientCertRevocationCheckMode",
				TerraformName:       "auth_client_cert_revocation_check_mode",
				MarkdownDescription: "The client certificate revocation checking mode used when a client authenticates with a client certificate. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"none\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - Do not perform any certificate revocation checking.\n\"ocsp\" - Use the Open Certificate Status Protcol (OCSP) for certificate revocation checking.\n\"crl\" - Use Certificate Revocation Lists (CRL) for certificate revocation checking.\n\"ocsp-crl\" - Use OCSP first, but if OCSP fails to return an unambiguous result, then check via CRL.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("none", "ocsp", "crl", "ocsp-crl"),
				},
				Default: "none",
			},
			{
				SempName:            "configSyncAuthenticationClientCertMaxChainDepth",
				TerraformName:       "config_sync_authentication_client_cert_max_chain_depth",
				MarkdownDescription: "The maximum depth for a client certificate chain. The depth of a chain is defined as the number of signing CA certificates that are present in the chain back to a trusted self-signed root CA certificate. The default value is `3`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 8),
				},
				Default: 3,
			},
			{
				SempName:            "configSyncAuthenticationClientCertValidateDateEnabled",
				TerraformName:       "config_sync_authentication_client_cert_validate_date_enabled",
				MarkdownDescription: "Enable or disable validation of the \"Not Before\" and \"Not After\" validity dates in the authentication certificate(s). The default value is `true`. Available since 2.22.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "configSyncClientProfileTcpInitialCongestionWindow",
				TerraformName:       "config_sync_client_profile_tcp_initial_congestion_window",
				MarkdownDescription: "The TCP initial congestion window size for Config Sync clients, in multiples of the TCP Maximum Segment Size (MSS). Changing the value from its default of 2 results in non-compliance with RFC 2581. Contact support before changing this value. The default value is `2`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 7826),
				},
				Default: 2,
			},
			{
				SempName:            "configSyncClientProfileTcpKeepaliveCount",
				TerraformName:       "config_sync_client_profile_tcp_keepalive_count",
				MarkdownDescription: "The number of TCP keepalive retransmissions to a client using the Client Profile before declaring that it is not available. The default value is `5`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(2, 5),
				},
				Default: 5,
			},
			{
				SempName:            "configSyncClientProfileTcpKeepaliveIdle",
				TerraformName:       "config_sync_client_profile_tcp_keepalive_idle",
				MarkdownDescription: "The amount of time a client connection using the Client Profile must remain idle before TCP begins sending keepalive probes, in seconds. The default value is `3`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(3, 120),
				},
				Default: 3,
			},
			{
				SempName:            "configSyncClientProfileTcpKeepaliveInterval",
				TerraformName:       "config_sync_client_profile_tcp_keepalive_interval",
				MarkdownDescription: "The amount of time between TCP keepalive retransmissions to a client using the Client Profile when no acknowledgement is received, in seconds. The default value is `1`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 30),
				},
				Default: 1,
			},
			{
				SempName:            "configSyncClientProfileTcpMaxWindow",
				TerraformName:       "config_sync_client_profile_tcp_max_window",
				MarkdownDescription: "The TCP maximum window size for clients using the Client Profile, in kilobytes. Changes are applied to all existing connections. The default value is `256`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(16, 65536),
				},
				Default: 256,
			},
			{
				SempName:            "configSyncClientProfileTcpMss",
				TerraformName:       "config_sync_client_profile_tcp_mss",
				MarkdownDescription: "The TCP maximum segment size for clients using the Client Profile, in bytes. Changes are applied to all existing connections. The default value is `1460`. Available since 2.22.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(256, 1460),
				},
				Default: 1460,
			},
			{
				SempName:            "configSyncEnabled",
				TerraformName:       "config_sync_enabled",
				MarkdownDescription: "Enable or disable configuration synchronization for High Availability or Disaster Recovery. The default value is `false`. Available since 2.22.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "configSyncSynchronizeUsernameEnabled",
				TerraformName:       "config_sync_synchronize_username_enabled",
				MarkdownDescription: "Enable or disable the synchronizing of usernames within High Availability groups. The transition from not synchronizing to synchronizing will cause the High Availability mate to fall out of sync. Recommendation: leave this as enabled. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`. Available since 2.22.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "configSyncTlsEnabled",
				TerraformName:       "config_sync_tls_enabled",
				MarkdownDescription: "Enable or disable the use of TLS encryption of the configuration synchronization communications between brokers in High Availability groups and/or Disaster Recovery sites. The default value is `false`. Available since 2.22.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "guaranteedMsgingDefragmentationScheduleDayList",
				TerraformName:       "guaranteed_msging_defragmentation_schedule_day_list",
				MarkdownDescription: "The days of the week to schedule defragmentation runs, specified as \"daily\" or as a comma-separated list of days. Days must be specified as \"Sun\", \"Mon\", \"Tue\", \"Wed\", \"Thu\", \"Fri, or \"Sat\", with no spaces, and in sorted order from Sunday to Saturday. Please note \"Sun,Mon,Tue,Wed,Thu,Fri,Sat\" is not allowed, use \"daily\" instead. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"daily\"`. Available since 2.25.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 100),
				},
				Default: "daily",
			},
			{
				SempName:            "guaranteedMsgingDefragmentationScheduleEnabled",
				TerraformName:       "guaranteed_msging_defragmentation_schedule_enabled",
				MarkdownDescription: "Enable or disable schedule-based defragmentation of Guaranteed Messaging spool files. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.25.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "guaranteedMsgingDefragmentationScheduleTimeList",
				TerraformName:       "guaranteed_msging_defragmentation_schedule_time_list",
				MarkdownDescription: "The times of the day to schedule defragmentation runs, specified as \"hourly\" or as a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times (up to 4) must be in sorted order from 0:00 to 23:59. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"0:00\"`. Available since 2.25.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(1, 100),
				},
				Default: "0:00",
			},
			{
				SempName:            "guaranteedMsgingDefragmentationThresholdEnabled",
				TerraformName:       "guaranteed_msging_defragmentation_threshold_enabled",
				MarkdownDescription: "Enable or disable threshold-based defragmentation of Guaranteed Messaging spool files. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.25.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "guaranteedMsgingDefragmentationThresholdFragmentationPercentage",
				TerraformName:       "guaranteed_msging_defragmentation_threshold_fragmentation_percentage",
				MarkdownDescription: "Percentage of spool fragmentation needed to trigger defragmentation run. The minimum value allowed is 30%. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `50`. Available since 2.25.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(30, 100),
				},
				Default: 50,
			},
			{
				SempName:            "guaranteedMsgingDefragmentationThresholdMinInterval",
				TerraformName:       "guaranteed_msging_defragmentation_threshold_min_interval",
				MarkdownDescription: "Minimum interval of time (in minutes) between defragmentation runs triggered by thresholds. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `15`. Available since 2.25.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 4294967295),
				},
				Default: 15,
			},
			{
				SempName:            "guaranteedMsgingDefragmentationThresholdUsagePercentage",
				TerraformName:       "guaranteed_msging_defragmentation_threshold_usage_percentage",
				MarkdownDescription: "Percentage of spool usage needed to trigger defragmentation run. The minimum value allowed is 30%. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `50`. Available since 2.25.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(30, 100),
				},
				Default: 50,
			},
			{
				SempName:            "guaranteedMsgingDiskArrayWwn",
				TerraformName:       "guaranteed_msging_disk_array_wwn",
				MarkdownDescription: "The WWN number to use when accessing a LUN on an external disk array. The default value is `\"\"`. Available since 2.18.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 64),
				},
				Default: "",
			},
			{
				SempName:            "guaranteedMsgingDiskLocation",
				TerraformName:       "guaranteed_msging_disk_location",
				MarkdownDescription: "The disk location for the the guaranteed message spool (required for high availability with guaranteed messaging). When external is chosen the guaranteed message spool is stored on an external disk array attached to the router. If internal storage is currently used, changing to external causes message spooling on the router to stop and messages spooled on the internal storage to be deleted. If internal is chosen the guaranteed message spool is stored on an external disk array attached to the router. If internal storage is currently used, changing to external causes message spooling on the router to stop and messages spooled on the internal storage to be deleted. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as guaranteedMsgingEnabled will be temporarily set to false to apply the change. The default value is `\"external\"`. The allowed values and their meaning are:\n\n<pre>\n\"external\" - The guaranteed message spool is stored on an external disk array attached to the appliance.\n\"internal\" - The guaranteed message spool is stored internally on the appliance.\n</pre>\n Available since 2.18.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("external", "internal"),
				},
				Default: "external",
			},
			{
				SempName:            "guaranteedMsgingEnabled",
				TerraformName:       "guaranteed_msging_enabled",
				MarkdownDescription: "Enable or disable Guaranteed Messaging. The default value is `false`. Available since 2.18.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "guaranteedMsgingEventCacheUsageThreshold",
				TerraformName:       "guaranteed_msging_event_cache_usage_threshold",
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
				SempName:            "guaranteedMsgingEventDeliveredUnackedThreshold",
				TerraformName:       "guaranteed_msging_event_delivered_unacked_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.",
						Requires:            []string{"set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.",
						Requires:            []string{"clear_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "guaranteedMsgingEventDiskUsageThreshold",
				TerraformName:       "guaranteed_msging_event_disk_usage_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.",
						Requires:            []string{"set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.",
						Requires:            []string{"clear_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "guaranteedMsgingEventEgressFlowCountThreshold",
				TerraformName:       "guaranteed_msging_event_egress_flow_count_threshold",
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
				SempName:            "guaranteedMsgingEventEndpointCountThreshold",
				TerraformName:       "guaranteed_msging_event_endpoint_count_threshold",
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
				SempName:            "guaranteedMsgingEventIngressFlowCountThreshold",
				TerraformName:       "guaranteed_msging_event_ingress_flow_count_threshold",
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
				SempName:            "guaranteedMsgingEventMsgCountThreshold",
				TerraformName:       "guaranteed_msging_event_msg_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.",
						Requires:            []string{"set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.",
						Requires:            []string{"clear_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "guaranteedMsgingEventMsgSpoolFileCountThreshold",
				TerraformName:       "guaranteed_msging_event_msg_spool_file_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.",
						Requires:            []string{"set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.",
						Requires:            []string{"clear_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "guaranteedMsgingEventMsgSpoolUsageThreshold",
				TerraformName:       "guaranteed_msging_event_msg_spool_usage_threshold",
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
				SempName:            "guaranteedMsgingEventTransactedSessionCountThreshold",
				TerraformName:       "guaranteed_msging_event_transacted_session_count_threshold",
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
				SempName:            "guaranteedMsgingEventTransactedSessionResourceCountThreshold",
				TerraformName:       "guaranteed_msging_event_transacted_session_resource_count_threshold",
				MarkdownDescription: "",
				Attributes: []*broker.AttributeInfo{
					{
						SempName:            "clearPercent",
						TerraformName:       "clear_percent",
						MarkdownDescription: "The clear threshold for the value of this counter as a percentage of its maximum value. Falling below this value will trigger a corresponding event.",
						Requires:            []string{"set_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("set_percent"),
							),
						},
					},
					{
						SempName:            "setPercent",
						TerraformName:       "set_percent",
						MarkdownDescription: "The set threshold for the value of this counter as a percentage of its maximum value. Exceeding this value will trigger a corresponding event.",
						Requires:            []string{"clear_percent"},
						Type:                types.Int64Type,
						TerraformType:       tftypes.Number,
						Converter:           broker.IntegerConverter{},
						Validators: []tfsdk.AttributeValidator{
							int64validator.Between(0, 100),
							schemavalidator.AlsoRequires(
								path.MatchRelative().AtParent().AtName("clear_percent"),
							),
						},
					},
				},
			},
			{
				SempName:            "guaranteedMsgingEventTransactionCountThreshold",
				TerraformName:       "guaranteed_msging_event_transaction_count_threshold",
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
				SempName:            "guaranteedMsgingMaxCacheUsage",
				TerraformName:       "guaranteed_msging_max_cache_usage",
				MarkdownDescription: "Guaranteed messaging cache usage limit. Expressed as a maximum percentage of the NAB's egress queueing. resources that the guaranteed message cache is allowed to use. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `10`. Available since 2.18.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 50),
				},
				Default: 10,
			},
			{
				SempName:            "guaranteedMsgingMaxMsgSpoolUsage",
				TerraformName:       "guaranteed_msging_max_msg_spool_usage",
				MarkdownDescription: "The maximum total message spool usage allowed across all VPNs on this broker, in megabytes. Recommendation: the maximum value should be less than 90% of the disk space allocated for the guaranteed message spool. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `60000`. Available since 2.18.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 6000000),
				},
				Default: 60000,
			},
			{
				SempName:            "guaranteedMsgingTransactionReplicationCompatibilityMode",
				TerraformName:       "guaranteed_msging_transaction_replication_compatibility_mode",
				MarkdownDescription: "The replication compatibility mode for the router. The default value is `\"legacy\"`. The allowed values and their meaning are:\"legacy\" - All transactions originated by clients are replicated to the standby site without using transactions.\"transacted\" - All transactions originated by clients are replicated to the standby site using transactions. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"legacy\"`. The allowed values and their meaning are:\n\n<pre>\n\"legacy\" - All transactions originated by clients are replicated to the standby site without using transactions.\n\"transacted\" - All transactions originated by clients are replicated to the standby site using transactions.\n</pre>\n Available since 2.18.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("legacy", "transacted"),
				},
				Default: "legacy",
			},
			{
				SempName:            "guaranteedMsgingVirtualRouterWhenActiveActive",
				TerraformName:       "guaranteed_msging_virtual_router_when_active_active",
				MarkdownDescription: "The High Availability role for this broker if using the legacy Active/Active configuration for high availability (not recommended). Note: for Active/Standby high availability configuration, this setting is ignored. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as guaranteedMsgingEnabled will be temporarily set to false to apply the change. The default value is `\"primary\"`. The allowed values and their meaning are:\n\n<pre>\n\"primary\" - The primary virtual router.\n\"backup\" - The backup virtual router.\n</pre>\n Available since 2.18.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.OneOf("primary", "backup"),
				},
				Default: "primary",
			},
			{
				SempName:            "oauthProfileDefault",
				TerraformName:       "oauth_profile_default",
				MarkdownDescription: "The default OAuth profile for OAuth authenticated SEMP requests. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`. Available since 2.24.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				SempName:            "serviceAmqpEnabled",
				TerraformName:       "service_amqp_enabled",
				MarkdownDescription: "Enable or disable the AMQP service. When disabled new AMQP Clients may not connect through the global or per-VPN AMQP listen-ports, and all currently connected AMQP Clients are immediately disconnected. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceAmqpTlsListenPort",
				TerraformName:       "service_amqp_tls_listen_port",
				MarkdownDescription: "TCP port number that AMQP clients can use to connect to the broker using raw TCP over TLS. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceAmqpEnabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `0`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 65535),
				},
				Default: 0,
			},
			{
				SempName:            "serviceEventConnectionCountThreshold",
				TerraformName:       "service_event_connection_count_threshold",
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
							int64validator.Between(0, 200000),
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
							int64validator.Between(0, 200000),
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
				SempName:            "serviceHealthCheckEnabled",
				TerraformName:       "service_health_check_enabled",
				MarkdownDescription: "Enable or disable the health-check service. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceHealthCheckListenPort",
				TerraformName:       "service_health_check_listen_port",
				MarkdownDescription: "The port number for the health-check service. The port must be unique across the message backbone. The health-check service must be disabled to change the port. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceHealthCheckEnabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `5550`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 5550,
			},
			{
				SempName:            "serviceMqttEnabled",
				TerraformName:       "service_mqtt_enabled",
				MarkdownDescription: "Enable or disable the MQTT service. When disabled new MQTT Clients may not connect through the per-VPN MQTT listen-ports, and all currently connected MQTT Clients are immediately disconnected. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceMsgBackboneEnabled",
				TerraformName:       "service_msg_backbone_enabled",
				MarkdownDescription: "Enable or disable the msg-backbone service. When disabled new Clients may not connect through global or per-VPN listen-ports, and all currently connected Clients are immediately disconnected. The default value is `true`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "serviceRestEventOutgoingConnectionCountThreshold",
				TerraformName:       "service_rest_event_outgoing_connection_count_threshold",
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
							int64validator.Between(0, 6000),
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
							int64validator.Between(0, 6000),
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
				SempName:            "serviceRestIncomingEnabled",
				TerraformName:       "service_rest_incoming_enabled",
				MarkdownDescription: "Enable or disable the REST service incoming connections on the router. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceRestOutgoingEnabled",
				TerraformName:       "service_rest_outgoing_enabled",
				MarkdownDescription: "Enable or disable the REST service outgoing connections on the router. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceSempCorsAllowAnyHostEnabled",
				TerraformName:       "service_semp_cors_allow_any_host_enabled",
				MarkdownDescription: "Enable or disable cross origin resource requests for the SEMP service. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`. Available since 2.24.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "serviceSempLegacyTimeoutEnabled",
				TerraformName:       "service_semp_legacy_timeout_enabled",
				MarkdownDescription: "Enable or disable extended SEMP timeouts for paged GETs. When a request times out, it returns the current page of content, even if the page is not full.  When enabled, the timeout is 60 seconds. When disabled, the timeout is 5 seconds.  The recommended setting is disabled (no legacy-timeout).  This parameter is intended as a temporary workaround to be used until SEMP clients can handle short pages.  This setting will be removed in a future release. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.18.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceSempPlainTextEnabled",
				TerraformName:       "service_semp_plain_text_enabled",
				MarkdownDescription: "Enable or disable plain-text SEMP service. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "serviceSempPlainTextListenPort",
				TerraformName:       "service_semp_plain_text_listen_port",
				MarkdownDescription: "The TCP port for plain-text SEMP client connections. This attribute cannot be cannot be changed while serviceSempPlainTextEnabled are set to true. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `80`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 80,
			},
			{
				SempName:            "serviceSempSessionIdleTimeout",
				TerraformName:       "service_semp_session_idle_timeout",
				MarkdownDescription: "The session idle timeout, in minutes. Sessions will be invalidated if there is no activity in this period of time. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `15`. Available since 2.21.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 5256000),
				},
				Default: 15,
			},
			{
				SempName:            "serviceSempSessionMaxLifetime",
				TerraformName:       "service_semp_session_max_lifetime",
				MarkdownDescription: "The maximum lifetime of a session, in minutes. Sessions will be invalidated after this period of time, regardless of activity. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `43200`. Available since 2.21.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 5256000),
				},
				Default: 43200,
			},
			{
				SempName:            "serviceSempTlsEnabled",
				TerraformName:       "service_semp_tls_enabled",
				MarkdownDescription: "Enable or disable TLS SEMP service. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "serviceSempTlsListenPort",
				TerraformName:       "service_semp_tls_listen_port",
				MarkdownDescription: "The TCP port for TLS SEMP client connections. This attribute cannot be cannot be changed while serviceSempTlsEnabled are set to true. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1943`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 1943,
			},
			{
				SempName:            "serviceSmfCompressionListenPort",
				TerraformName:       "service_smf_compression_listen_port",
				MarkdownDescription: "TCP port number that SMF clients can use to connect to the broker using raw compression TCP. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceSmfEnabled will be temporarily set to false to apply the change. The default value is `55003`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 55003,
			},
			{
				SempName:            "serviceSmfEnabled",
				TerraformName:       "service_smf_enabled",
				MarkdownDescription: "Enable or disable the SMF service. When disabled new SMF Clients may not connect through the global listen-ports, and all currently connected SMF Clients are immediately disconnected. The default value is `true`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "serviceSmfEventConnectionCountThreshold",
				TerraformName:       "service_smf_event_connection_count_threshold",
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
							int64validator.Between(0, 200000),
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
							int64validator.Between(0, 200000),
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
				SempName:            "serviceSmfPlainTextListenPort",
				TerraformName:       "service_smf_plain_text_listen_port",
				MarkdownDescription: "TCP port number that SMF clients can use to connect to the broker using raw TCP. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceSmfEnabled will be temporarily set to false to apply the change. The default value is `55555`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 55555,
			},
			{
				SempName:            "serviceSmfRoutingControlListenPort",
				TerraformName:       "service_smf_routing_control_listen_port",
				MarkdownDescription: "TCP port number that SMF clients can use to connect to the broker using raw routing control TCP. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceSmfEnabled will be temporarily set to false to apply the change. The default value is `55556`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 55556,
			},
			{
				SempName:            "serviceSmfTlsListenPort",
				TerraformName:       "service_smf_tls_listen_port",
				MarkdownDescription: "TCP port number that SMF clients can use to connect to the broker using raw TCP over TLS. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceSmfEnabled will be temporarily set to false to apply the change. The default value is `55443`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 55443,
			},
			{
				SempName:            "serviceTlsEventConnectionCountThreshold",
				TerraformName:       "service_tls_event_connection_count_threshold",
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
							int64validator.Between(0, 200000),
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
							int64validator.Between(0, 200000),
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
				SempName:            "serviceWebTransportEnabled",
				TerraformName:       "service_web_transport_enabled",
				MarkdownDescription: "Enable or disable the web-transport service. When disabled new web-transport Clients may not connect through the global listen-ports, and all currently connected web-transport Clients are immediately disconnected. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.17.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "serviceWebTransportPlainTextListenPort",
				TerraformName:       "service_web_transport_plain_text_listen_port",
				MarkdownDescription: "The TCP port for plain-text WEB client connections. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceWebTransportEnabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `80`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 80,
			},
			{
				SempName:            "serviceWebTransportTlsListenPort",
				TerraformName:       "service_web_transport_tls_listen_port",
				MarkdownDescription: "The TCP port for TLS WEB client connections. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceWebTransportEnabled will be temporarily set to false to apply the change. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `443`. Available since 2.17.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 65535),
				},
				Default: 443,
			},
			{
				SempName:            "serviceWebTransportWebUrlSuffix",
				TerraformName:       "service_web_transport_web_url_suffix",
				MarkdownDescription: "Used to specify the Web URL suffix that will be used by Web clients when communicating with the broker. Modifying this attribute while the object (or the relevant part of the object) is administratively enabled may be service impacting as serviceWebTransportEnabled will be temporarily set to false to apply the change. The default value is `\"\"`. Available since 2.17.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 127),
				},
				Default: "",
			},
			{
				SempName:            "tlsBlockVersion10Enabled",
				TerraformName:       "tls_block_version10_enabled",
				MarkdownDescription: "Enable or disable the blocking of incoming TLS version 1.0 connections. When blocked, existing TLS 1.0 connections from Clients and SEMP users remain connected while new connections are blocked. Note that support for TLS 1.0 will eventually be discontinued, at which time TLS 1.0 connections will be blocked regardless of this setting. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "tlsBlockVersion11Enabled",
				TerraformName:       "tls_block_version11_enabled",
				MarkdownDescription: "Enable or disable the blocking of TLS version 1.1 connections. When blocked, all existing incoming and outgoing TLS 1.1 connections with Clients, SEMP users, and LDAP servers remain connected while new connections are blocked. Note that support for TLS 1.1 will eventually be discontinued, at which time TLS 1.1 connections will be blocked regardless of this setting. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "tlsCipherSuiteManagementList",
				TerraformName:       "tls_cipher_suite_management_list",
				MarkdownDescription: "The colon-separated list of cipher suites used for TLS management connections (e.g. SEMP, LDAP). The value \"default\" implies all supported suites ordered from most secure to least secure. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 1559),
				},
				Default: "default",
			},
			{
				SempName:            "tlsCipherSuiteMsgBackboneList",
				TerraformName:       "tls_cipher_suite_msg_backbone_list",
				MarkdownDescription: "The colon-separated list of cipher suites used for TLS data connections (e.g. client pub/sub). The value \"default\" implies all supported suites ordered from most secure to least secure. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 1559),
				},
				Default: "default",
			},
			{
				SempName:            "tlsCipherSuiteSecureShellList",
				TerraformName:       "tls_cipher_suite_secure_shell_list",
				MarkdownDescription: "The colon-separated list of cipher suites used for TLS secure shell connections (e.g. SSH, SFTP, SCP). The value \"default\" implies all supported suites ordered from most secure to least secure. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"default\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 1559),
				},
				Default: "default",
			},
			{
				SempName:            "tlsCrimeExploitProtectionEnabled",
				TerraformName:       "tls_crime_exploit_protection_enabled",
				MarkdownDescription: "Enable or disable protection against the CRIME exploit. When enabled, TLS+compressed messaging performance is degraded. This protection should only be disabled if sufficient ACL and authentication features are being employed such that a potential attacker does not have sufficient access to trigger the exploit. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "tlsServerCertContent",
				TerraformName:       "tls_server_cert_content",
				MarkdownDescription: "The PEM formatted content for the server certificate used for TLS connections. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `\"\"`.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 32768),
				},
				Default: "",
			},
			{
				SempName:            "tlsServerCertPassword",
				TerraformName:       "tls_server_cert_password",
				MarkdownDescription: "The password for the server certificate used for TLS connections. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is `\"\"`.",
				Sensitive:           true,
				Requires:            []string{"tls_server_cert_content"},
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					schemavalidator.AlsoRequires(
						path.MatchRelative().AtParent().AtName("tls_server_cert_content"),
					),
					stringvalidator.LengthBetween(0, 32768),
				},
				Default: "",
			},
			{
				SempName:            "tlsStandardDomainCertificateAuthoritiesEnabled",
				TerraformName:       "tls_standard_domain_certificate_authorities_enabled",
				MarkdownDescription: "Enable or disable the standard domain certificate authority list. The default value is `true`. Available since 2.19.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "tlsTicketLifetime",
				TerraformName:       "tls_ticket_lifetime",
				MarkdownDescription: "The TLS ticket lifetime in seconds. When a client connects with TLS, a session with a session ticket is created using the TLS ticket lifetime which determines how long the client has to resume the session. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `86400`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(1, 86400),
				},
				Default: 86400,
			},
			{
				SempName:            "webManagerAllowUnencryptedWizardsEnabled",
				TerraformName:       "web_manager_allow_unencrypted_wizards_enabled",
				MarkdownDescription: "Enable or disable the use of unencrypted wizards in the Web-based Manager UI. This setting should be left at its default on all production systems or other systems that need to be secure.  Enabling this option will permit the broker to forward plain-text data to other brokers, making important information or credentials available for snooping. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`. Available since 2.28.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				SempName:            "webManagerCustomization",
				TerraformName:       "web_manager_customization",
				MarkdownDescription: "Reserved for internal use by Solace. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`. Available since 2.25.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				Validators: []tfsdk.AttributeValidator{
					stringvalidator.LengthBetween(0, 1024),
				},
				Default: "",
			},
			{
				SempName:            "webManagerRedirectHttpEnabled",
				TerraformName:       "web_manager_redirect_http_enabled",
				MarkdownDescription: "Enable or disable redirection of HTTP requests for the broker manager to HTTPS. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`. Available since 2.24.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				SempName:            "webManagerRedirectHttpOverrideTlsPort",
				TerraformName:       "web_manager_redirect_http_override_tls_port",
				MarkdownDescription: "The HTTPS port that HTTP requests will be redirected towards in a HTTP 301 redirect response. Zero is a special value that means use the value specified for the SEMP TLS port value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `0`. Available since 2.24.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Validators: []tfsdk.AttributeValidator{
					int64validator.Between(0, 65535),
				},
				Default: 0,
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
