// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package oceanrightsizingrule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateInterpolationForAttributeParameters(property *string) error {
	if property == nil {
		return fmt.Errorf("parameter property is required, but nil was provided")
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validatePutMonthlyRepetitionBasisParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasis:
		value := value.(*[]*OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasis)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasis:
		value_ := value.([]*OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasis)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasis; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validatePutWeeklyRepetitionBasisParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktf.IResolvable:
		// ok
	case *[]*OceanRightSizingRuleRecommendationApplicationIntervalsWeeklyRepetitionBasis:
		value := value.(*[]*OceanRightSizingRuleRecommendationApplicationIntervalsWeeklyRepetitionBasis)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*OceanRightSizingRuleRecommendationApplicationIntervalsWeeklyRepetitionBasis:
		value_ := value.([]*OceanRightSizingRuleRecommendationApplicationIntervalsWeeklyRepetitionBasis)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktf.IResolvable, *[]*OceanRightSizingRuleRecommendationApplicationIntervalsWeeklyRepetitionBasis; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (o *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateResolveParameters(_context cdktf.IResolveContext) error {
	if _context == nil {
		return fmt.Errorf("parameter _context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktf.IResolvable:
		// ok
	case *OceanRightSizingRuleRecommendationApplicationIntervals:
		val := val.(*OceanRightSizingRuleRecommendationApplicationIntervals)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case OceanRightSizingRuleRecommendationApplicationIntervals:
		val_ := val.(OceanRightSizingRuleRecommendationApplicationIntervals)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktf.IResolvable, *OceanRightSizingRuleRecommendationApplicationIntervals; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateSetRepetitionBasisParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_OceanRightSizingRuleRecommendationApplicationIntervalsOutputReference) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewOceanRightSizingRuleRecommendationApplicationIntervalsOutputReferenceParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

