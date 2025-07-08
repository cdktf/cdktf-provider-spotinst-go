// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notificationcenter

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationCenterRegisteredUsersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotificationCenterRegisteredUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationCenterRegisteredUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterRegisteredUsersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterRegisteredUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterRegisteredUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterRegisteredUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationCenterRegisteredUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

