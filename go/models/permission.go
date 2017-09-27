package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// Permission permission
// swagger:model Permission
type Permission struct {

	//  Enum options - PERMISSION_CONTROLLER, PERMISSION_INTERNAL, PERMISSION_VIRTUALSERVICE, PERMISSION_POOL, PERMISSION_HEALTHMONITOR, PERMISSION_NETWORKPROFILE, PERMISSION_APPLICATIONPROFILE, PERMISSION_HTTPPOLICYSET, PERMISSION_IPADDRGROUP, PERMISSION_STRINGGROUP, PERMISSION_SSLPROFILE, PERMISSION_SSLKEYANDCERTIFICATE, PERMISSION_NETWORKSECURITYPOLICY, PERMISSION_APPLICATIONPERSISTENCEPROFILE, PERMISSION_ANALYTICSPROFILE, PERMISSION_VSDATASCRIPTSET, PERMISSION_TENANT, PERMISSION_PKIPROFILE, PERMISSION_AUTHPROFILE, PERMISSION_CLOUD, PERMISSION_SERVICEENGINE, PERMISSION_SERVICEENGINEGROUP, PERMISSION_NETWORK, PERMISSION_SYSTEMCONFIGURATION, PERMISSION_VRFCONTEXT, PERMISSION_USER, PERMISSION_ROLE, PERMISSION_ALERT, PERMISSION_ALERTCONFIG, PERMISSION_ALERTEMAILCONFIG, PERMISSION_ALERTSYSLOGCONFIG, PERMISSION_ACTIONGROUPCONFIG, PERMISSION_SNMPTRAPPROFILE, PERMISSION_UPGRADE, PERMISSION_REBOOT, PERMISSION_TECHSUPPORT, PERMISSION_EXEMPT, PERMISSION_VIRTUALSERVICE_MAINTENANCE, PERMISSION_POOL_MAINTENANCE, PERMISSION_TRAFFIC_CAPTURE, PERMISSION_MICROSERVICEGROUP, PERMISSION_IPAMDNSPROVIDERPROFILE, PERMISSION_CERTIFICATEMANAGEMENTPROFILE, PERMISSION_POOLGROUP, PERMISSION_PRIORITYLABELS, PERMISSION_POOLGROUPDEPLOYMENTPOLICY, PERMISSION_GSLB, PERMISSION_GSLBSERVICE, PERMISSION_GSLBHEALTHMONITOR, PERMISSION_GSLBGEODBPROFILE, PERMISSION_GSLBAPPLICATIONPERSISTENCEPROFILE, PERMISSION_DNSPOLICY, PERMISSION_TRAFFICCLONEPROFILE.
	Resource string `json:"resource,omitempty"`

	//  Enum options - NO_ACCESS, READ_ACCESS, WRITE_ACCESS.
	Type string `json:"type,omitempty"`
}
