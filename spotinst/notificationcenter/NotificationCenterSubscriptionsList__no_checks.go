// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package notificationcenter

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationCenterSubscriptionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NotificationCenterSubscriptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationCenterSubscriptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterSubscriptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterSubscriptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterSubscriptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationCenterSubscriptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationCenterSubscriptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

