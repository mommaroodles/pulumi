// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nestedTypes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"regress-go-15478/nestedTypes/internal"
)

var _ = internal.GetEnvOrDefault

type NestedType struct {
	Data       [][][]string                            `pulumi:"data"`
	NestedMaps map[string]map[string]map[string]string `pulumi:"nestedMaps"`
}

type NestedTypeOutput struct{ *pulumi.OutputState }

func (NestedTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NestedType)(nil)).Elem()
}

func (o NestedTypeOutput) ToNestedTypeOutput() NestedTypeOutput {
	return o
}

func (o NestedTypeOutput) ToNestedTypeOutputWithContext(ctx context.Context) NestedTypeOutput {
	return o
}

func (o NestedTypeOutput) Data() pulumi.StringArrayArrayArrayOutput {
	return o.ApplyT(func(v NestedType) [][][]string { return v.Data }).(pulumi.StringArrayArrayArrayOutput)
}

func (o NestedTypeOutput) NestedMaps() pulumi.StringMapMapMapOutput {
	return o.ApplyT(func(v NestedType) map[string]map[string]map[string]string { return v.NestedMaps }).(pulumi.StringMapMapMapOutput)
}

type NestedTypePtrOutput struct{ *pulumi.OutputState }

func (NestedTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NestedType)(nil)).Elem()
}

func (o NestedTypePtrOutput) ToNestedTypePtrOutput() NestedTypePtrOutput {
	return o
}

func (o NestedTypePtrOutput) ToNestedTypePtrOutputWithContext(ctx context.Context) NestedTypePtrOutput {
	return o
}

func (o NestedTypePtrOutput) Elem() NestedTypeOutput {
	return o.ApplyT(func(v *NestedType) NestedType {
		if v != nil {
			return *v
		}
		var ret NestedType
		return ret
	}).(NestedTypeOutput)
}

func (o NestedTypePtrOutput) Data() pulumi.StringArrayArrayArrayOutput {
	return o.ApplyT(func(v *NestedType) [][][]string {
		if v == nil {
			return nil
		}
		return v.Data
	}).(pulumi.StringArrayArrayArrayOutput)
}

func (o NestedTypePtrOutput) NestedMaps() pulumi.StringMapMapMapOutput {
	return o.ApplyT(func(v *NestedType) map[string]map[string]map[string]string {
		if v == nil {
			return nil
		}
		return v.NestedMaps
	}).(pulumi.StringMapMapMapOutput)
}

func init() {
	pulumi.RegisterOutputType(NestedTypeOutput{})
	pulumi.RegisterOutputType(NestedTypePtrOutput{})
}