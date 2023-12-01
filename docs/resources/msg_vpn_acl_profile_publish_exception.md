---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "solacebroker_msg_vpn_acl_profile_publish_exception Resource - solacebroker"
subcategory: ""
description: |-
  A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.
  Attribute|Identifying|Write-Only|Deprecated|Opaque
  :---|:---:|:---:|:---:|:---:
  aclprofilename|x||x|
  msgvpnname|x||x|
  publishexceptiontopic|x||x|
  topic_syntax|x||x|
  A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.
  This has been deprecated since 2.14. Replaced by publishtopicexceptions.
---

# solacebroker_msg_vpn_acl_profile_publish_exception (Resource)

A Publish Topic Exception is an exception to the default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. Exceptions must be expressed as a topic.


Attribute|Identifying|Write-Only|Deprecated|Opaque
:---|:---:|:---:|:---:|:---:
acl_profile_name|x||x|
msg_vpn_name|x||x|
publish_exception_topic|x||x|
topic_syntax|x||x|



A SEMP client authorized with a minimum access scope/level of "vpn/read-only" is required to perform this operation.

This has been deprecated since 2.14. Replaced by publish_topic_exceptions.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `acl_profile_name` (String, Deprecated) The name of the ACL Profile. Deprecated since 2.14. Replaced by publish_topic_exceptions.
- `msg_vpn_name` (String, Deprecated) The name of the Message VPN. Deprecated since 2.14. Replaced by publish_topic_exceptions.
- `publish_exception_topic` (String, Deprecated) The topic for the exception to the default action taken. May include wildcard characters. Deprecated since 2.14. Replaced by publish_topic_exceptions.
- `topic_syntax` (String, Deprecated) The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:

<pre>
"smf" - Topic uses SMF syntax.
"mqtt" - Topic uses MQTT syntax.
</pre>
 Deprecated since 2.14. Replaced by publish_topic_exceptions.

