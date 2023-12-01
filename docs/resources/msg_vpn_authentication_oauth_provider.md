---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solacebroker_msg_vpn_authentication_oauth_provider Resource - solacebroker"
subcategory: ""
description: |-
  OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.
  Attribute|Identifying|Write-Only|Deprecated|Opaque
  :---|:---:|:---:|:---:|:---:
  audienceclaimname|||x|
  audienceclaimsource|||x|
  audienceclaimvalue|||x|
  audiencevalidationenabled|||x|
  authorizationgroupclaimname|||x|
  authorizationgroupclaimsource|||x|
  authorizationgroupenabled|||x|
  disconnectontokenexpirationenabled|||x|
  enabled|||x|
  jwksrefreshinterval|||x|
  jwksuri|||x|
  msgvpnname|x||x|
  oauthprovidername|x||x|
  tokenignoretimelimitsenabled|||x|
  tokenintrospectionparametername|||x|
  tokenintrospectionpassword||x|x|x
  tokenintrospectiontimeout|||x|
  tokenintrospectionuri|||x|
  tokenintrospectionusername|||x|
  usernameclaimname|||x|
  usernameclaimsource|||x|
  usernamevalidateenabled|||x|
  A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.
  This has been deprecated since 2.25. Replaced by authenticationoauth_profiles.
---

# solacebroker_msg_vpn_authentication_oauth_provider (Resource)

OAuth Providers contain information about the issuer of an OAuth token that is needed to validate the token and derive a client username from it.


Attribute|Identifying|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:
audience_claim_name|||x|
audience_claim_source|||x|
audience_claim_value|||x|
audience_validation_enabled|||x|
authorization_group_claim_name|||x|
authorization_group_claim_source|||x|
authorization_group_enabled|||x|
disconnect_on_token_expiration_enabled|||x|
enabled|||x|
jwks_refresh_interval|||x|
jwks_uri|||x|
msg_vpn_name|x||x|
oauth_provider_name|x||x|
token_ignore_time_limits_enabled|||x|
token_introspection_parameter_name|||x|
token_introspection_password||x|x|x
token_introspection_timeout|||x|
token_introspection_uri|||x|
token_introspection_username|||x|
username_claim_name|||x|
username_claim_source|||x|
username_validate_enabled|||x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.25. Replaced by authenticationoauth_profiles.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `msg_vpn_name` (String, Deprecated) The name of the Message VPN. Deprecated since 2.25. Replaced by authenticationoauth_profiles.
- `oauth_provider_name` (String, Deprecated) The name of the OAuth Provider. Deprecated since 2.25. Replaced by authenticationoauth_profiles.

### Optional

- `audience_claim_name` (String, Deprecated) The audience claim name, indicating which part of the object to use for determining the audience. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"aud"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `audience_claim_source` (String, Deprecated) The audience claim source, indicating where to search for the audience value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"id-token"`. The allowed values and their meaning are:

<pre>
"access-token" - The OAuth v2 access_token.
"id-token" - The OpenID Connect id_token.
"introspection" - The result of introspecting the OAuth v2 access_token.
</pre>
 Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `audience_claim_value` (String, Deprecated) The required audience value for a token to be considered valid. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `""`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `audience_validation_enabled` (Boolean, Deprecated) Enable or disable audience validation. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `authorization_group_claim_name` (String, Deprecated) The authorization group claim name, indicating which part of the object to use for determining the authorization group. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"scope"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `authorization_group_claim_source` (String, Deprecated) The authorization group claim source, indicating where to search for the authorization group name. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"id-token"`. The allowed values and their meaning are:

<pre>
"access-token" - The OAuth v2 access_token.
"id-token" - The OpenID Connect id_token.
"introspection" - The result of introspecting the OAuth v2 access_token.
</pre>
 Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `authorization_group_enabled` (Boolean, Deprecated) Enable or disable OAuth based authorization. When enabled, the configured authorization type for OAuth clients is overridden. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `disconnect_on_token_expiration_enabled` (Boolean, Deprecated) Enable or disable the disconnection of clients when their tokens expire. Changing this value does not affect existing clients, only new client connections. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `true`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `enabled` (Boolean, Deprecated) Enable or disable OAuth Provider client authentication. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `jwks_refresh_interval` (Number, Deprecated) The number of seconds between forced JWKS public key refreshing. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `86400`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `jwks_uri` (String, Deprecated) The URI where the OAuth provider publishes its JWKS public keys. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `""`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `token_ignore_time_limits_enabled` (Boolean, Deprecated) Enable or disable whether to ignore time limits and accept tokens that are not yet valid or are no longer valid. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `token_introspection_parameter_name` (String, Deprecated) The parameter name used to identify the token during access token introspection. A standards compliant OAuth introspection server expects "token". Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"token"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `token_introspection_password` (String, Sensitive, Deprecated) The password to use when logging into the token introspection URI. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `""`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `token_introspection_timeout` (Number, Deprecated) The maximum time in seconds a token introspection is allowed to take. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `1`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `token_introspection_uri` (String, Deprecated) The token introspection URI of the OAuth authentication server. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `""`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `token_introspection_username` (String, Deprecated) The username to use when logging into the token introspection URI. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `""`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `username_claim_name` (String, Deprecated) The username claim name, indicating which part of the object to use for determining the username. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"sub"`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `username_claim_source` (String, Deprecated) The username claim source, indicating where to search for the username value. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `"id-token"`. The allowed values and their meaning are:

<pre>
"access-token" - The OAuth v2 access_token.
"id-token" - The OpenID Connect id_token.
"introspection" - The result of introspecting the OAuth v2 access_token.
</pre>
 Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.
- `username_validate_enabled` (Boolean, Deprecated) Enable or disable whether the API provided username will be validated against the username calculated from the token(s); the connection attempt is rejected if they differ. Changes to this attribute are synchronized to HA mates and replication sites via config-sync. The default value is `false`. Deprecated since 2.25. authenticationOauthProviders replaced by authenticationoauth_profiles.

