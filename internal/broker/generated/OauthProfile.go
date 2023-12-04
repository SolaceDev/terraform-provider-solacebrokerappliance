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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"regexp"
	"terraform-provider-solacebroker/internal/broker"
)

func init() {
	info := broker.EntityInputs{
		TerraformName:       "oauth_profile",
		MarkdownDescription: "OAuth profiles specify how to securely authenticate to an OAuth provider.\n\n\nAttribute|Identifying|Write-Only|Deprecated|Opaque\n:---|:---:|:---:|:---:|:---:\nclient_secret||x||x\noauth_profile_name|x|||\n\n\n\nA SEMP client authorized with a minimum access scope/level of \"global/read-only\" is required to perform this operation.\n\nThis has been available since SEMP API version 2.24.",
		ObjectType:          broker.StandardObject,
		PathTemplate:        "/oauthProfiles/{oauthProfileName}",
		Version:             0,
		Attributes: []*broker.AttributeInfo{
			{
				BaseType:            broker.String,
				SempName:            "accessLevelGroupsClaimName",
				TerraformName:       "access_level_groups_claim_name",
				MarkdownDescription: "The name of the groups claim. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"groups\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 100),
				},
				Default: "groups",
			},
			{
				BaseType:            broker.String,
				SempName:            "accessLevelGroupsClaimStringFormat",
				TerraformName:       "access_level_groups_claim_string_format",
				MarkdownDescription: "The format of the access level groups claim value when it is a string. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"single\"`. The allowed values and their meaning are:\n\n<pre>\n\"single\" - When the claim is a string, it is interpreted as as single group.\n\"space-delimited\" - When the claim is a string, it is interpreted as a space-delimited list of groups, similar to the \"scope\" claim.\n</pre>\n Available since SEMP API version 2.32.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("single", "space-delimited"),
				},
				Default: "single",
			},
			{
				BaseType:            broker.String,
				SempName:            "clientId",
				TerraformName:       "client_id",
				MarkdownDescription: "The OAuth client id. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "clientRedirectUri",
				TerraformName:       "client_redirect_uri",
				MarkdownDescription: "The OAuth redirect URI. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 300),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "clientRequiredType",
				TerraformName:       "client_required_type",
				MarkdownDescription: "The required value for the TYP field in the ID token header. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"JWT\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
				},
				Default: "JWT",
			},
			{
				BaseType:            broker.String,
				SempName:            "clientScope",
				TerraformName:       "client_scope",
				MarkdownDescription: "The OAuth scope. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"openid email\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "openid email",
			},
			{
				BaseType:            broker.String,
				SempName:            "clientSecret",
				TerraformName:       "client_secret",
				MarkdownDescription: "The OAuth client secret. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4 (refer to the `Notes` section in the SEMP API `Config reference`). Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Sensitive:           true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 512),
				},
				Default: "",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "clientValidateTypeEnabled",
				TerraformName:       "client_validate_type_enabled",
				MarkdownDescription: "Enable or disable verification of the TYP field in the ID token header. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "defaultGlobalAccessLevel",
				TerraformName:       "default_global_access_level",
				MarkdownDescription: "The default global access level for this OAuth profile. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"none\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - User has no access to global data.\n\"read-only\" - User has read-only access to global data.\n\"read-write\" - User has read-write access to most global data.\n\"admin\" - User has read-write access to all global data.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("none", "read-only", "read-write", "admin"),
				},
				Default: "none",
			},
			{
				BaseType:            broker.String,
				SempName:            "defaultMsgVpnAccessLevel",
				TerraformName:       "default_msg_vpn_access_level",
				MarkdownDescription: "The default message VPN access level for the OAuth profile. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"none\"`. The allowed values and their meaning are:\n\n<pre>\n\"none\" - User has no access to a Message VPN.\n\"read-only\" - User has read-only access to a Message VPN.\n\"read-write\" - User has read-write access to most Message VPN settings.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("none", "read-only", "read-write"),
				},
				Default: "none",
			},
			{
				BaseType:            broker.String,
				SempName:            "displayName",
				TerraformName:       "display_name",
				MarkdownDescription: "The user friendly name for the OAuth profile. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 64),
				},
				Default: "",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "enabled",
				TerraformName:       "enabled",
				MarkdownDescription: "Enable or disable the OAuth profile. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `false`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             false,
			},
			{
				BaseType:            broker.String,
				SempName:            "endpointAuthorization",
				TerraformName:       "endpoint_authorization",
				MarkdownDescription: "The OAuth authorization endpoint. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "endpointDiscovery",
				TerraformName:       "endpoint_discovery",
				MarkdownDescription: "The OpenID Connect discovery endpoint or OAuth Authorization Server Metadata endpoint. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "endpointDiscoveryRefreshInterval",
				TerraformName:       "endpoint_discovery_refresh_interval",
				MarkdownDescription: "The number of seconds between discovery endpoint requests. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `86400`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(60, 31536000),
				},
				Default: 86400,
			},
			{
				BaseType:            broker.String,
				SempName:            "endpointIntrospection",
				TerraformName:       "endpoint_introspection",
				MarkdownDescription: "The OAuth introspection endpoint. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "endpointIntrospectionTimeout",
				TerraformName:       "endpoint_introspection_timeout",
				MarkdownDescription: "The maximum time in seconds a token introspection request is allowed to take. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
				Default: 1,
			},
			{
				BaseType:            broker.String,
				SempName:            "endpointJwks",
				TerraformName:       "endpoint_jwks",
				MarkdownDescription: "The OAuth JWKS endpoint. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "endpointJwksRefreshInterval",
				TerraformName:       "endpoint_jwks_refresh_interval",
				MarkdownDescription: "The number of seconds between JWKS endpoint requests. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `86400`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(60, 31536000),
				},
				Default: 86400,
			},
			{
				BaseType:            broker.String,
				SempName:            "endpointToken",
				TerraformName:       "endpoint_token",
				MarkdownDescription: "The OAuth token endpoint. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "endpointTokenTimeout",
				TerraformName:       "endpoint_token_timeout",
				MarkdownDescription: "The maximum time in seconds a token request is allowed to take. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
				Default: 1,
			},
			{
				BaseType:            broker.String,
				SempName:            "endpointUserinfo",
				TerraformName:       "endpoint_userinfo",
				MarkdownDescription: "The OpenID Connect Userinfo endpoint. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.Int64,
				SempName:            "endpointUserinfoTimeout",
				TerraformName:       "endpoint_userinfo_timeout",
				MarkdownDescription: "The maximum time in seconds a userinfo request is allowed to take. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `1`.",
				Type:                types.Int64Type,
				TerraformType:       tftypes.Number,
				Converter:           broker.IntegerConverter{},
				Int64Validators: []validator.Int64{
					int64validator.Between(1, 60),
				},
				Default: 1,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "interactiveEnabled",
				TerraformName:       "interactive_enabled",
				MarkdownDescription: "Enable or disable interactive logins via this OAuth provider. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "interactivePromptForExpiredSession",
				TerraformName:       "interactive_prompt_for_expired_session",
				MarkdownDescription: "The value of the prompt parameter provided to the OAuth authorization server for login requests where the session has expired. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "interactivePromptForNewSession",
				TerraformName:       "interactive_prompt_for_new_session",
				MarkdownDescription: "The value of the prompt parameter provided to the OAuth authorization server for login requests where the session is new or the user has explicitly logged out. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"select_account\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 32),
				},
				Default: "select_account",
			},
			{
				BaseType:            broker.String,
				SempName:            "issuer",
				TerraformName:       "issuer",
				MarkdownDescription: "The Issuer Identifier for the OAuth provider. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "oauthProfileName",
				TerraformName:       "oauth_profile_name",
				MarkdownDescription: "The name of the OAuth profile.",
				Identifying:         true,
				Required:            true,
				RequiresReplace:     true,
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
					stringvalidator.RegexMatches(regexp.MustCompile("^[A-Za-z0-9_]+$"), ""),
				},
			},
			{
				BaseType:            broker.String,
				SempName:            "oauthRole",
				TerraformName:       "oauth_role",
				MarkdownDescription: "The OAuth role of the broker. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"client\"`. The allowed values and their meaning are:\n\n<pre>\n\"client\" - The broker is in the OAuth client role.\n\"resource-server\" - The broker is in the OAuth resource server role.\n</pre>\n",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.OneOf("client", "resource-server"),
				},
				Default: "client",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "resourceServerParseAccessTokenEnabled",
				TerraformName:       "resource_server_parse_access_token_enabled",
				MarkdownDescription: "Enable or disable parsing of the access token as a JWT. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "resourceServerRequiredAudience",
				TerraformName:       "resource_server_required_audience",
				MarkdownDescription: "The required audience value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "resourceServerRequiredIssuer",
				TerraformName:       "resource_server_required_issuer",
				MarkdownDescription: "The required issuer value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 255),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "resourceServerRequiredScope",
				TerraformName:       "resource_server_required_scope",
				MarkdownDescription: "A space-separated list of scopes that must be present in the scope claim. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(0, 200),
				},
				Default: "",
			},
			{
				BaseType:            broker.String,
				SempName:            "resourceServerRequiredType",
				TerraformName:       "resource_server_required_type",
				MarkdownDescription: "The required TYP value. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"at+jwt\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
				},
				Default: "at+jwt",
			},
			{
				BaseType:            broker.Bool,
				SempName:            "resourceServerValidateAudienceEnabled",
				TerraformName:       "resource_server_validate_audience_enabled",
				MarkdownDescription: "Enable or disable verification of the audience claim in the access token or introspection response. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "resourceServerValidateIssuerEnabled",
				TerraformName:       "resource_server_validate_issuer_enabled",
				MarkdownDescription: "Enable or disable verification of the issuer claim in the access token or introspection response. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "resourceServerValidateScopeEnabled",
				TerraformName:       "resource_server_validate_scope_enabled",
				MarkdownDescription: "Enable or disable verification of the scope claim in the access token or introspection response. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "resourceServerValidateTypeEnabled",
				TerraformName:       "resource_server_validate_type_enabled",
				MarkdownDescription: "Enable or disable verification of the TYP field in the access token header. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.Bool,
				SempName:            "sempEnabled",
				TerraformName:       "semp_enabled",
				MarkdownDescription: "Enable or disable authentication of SEMP requests with OAuth tokens. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `true`.",
				Type:                types.BoolType,
				TerraformType:       tftypes.Bool,
				Converter:           broker.SimpleConverter[bool]{TerraformType: tftypes.Bool},
				Default:             true,
			},
			{
				BaseType:            broker.String,
				SempName:            "usernameClaimName",
				TerraformName:       "username_claim_name",
				MarkdownDescription: "The name of the username claim. Changes to this attribute are synchronized to HA mates via config-sync. The default value is `\"sub\"`.",
				Type:                types.StringType,
				TerraformType:       tftypes.String,
				Converter:           broker.SimpleConverter[string]{TerraformType: tftypes.String},
				StringValidators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
				Default: "sub",
			},
		},
	}
	broker.RegisterResource(info)
	broker.RegisterDataSource(info)
}
