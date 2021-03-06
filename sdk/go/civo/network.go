// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Civo Network resource. This can be used to create,
// modify, and delete Networks.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-civo/sdk/go/civo"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := civo.NewNetwork(ctx, "customNet", &civo.NetworkArgs{
// 			Label: pulumi.String("test_network"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Network struct {
	pulumi.CustomResourceState

	// The block ip assigned to the network.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// If is the default network.
	Default pulumi.BoolOutput `pulumi:"default"`
	// The Network label
	Label pulumi.StringOutput `pulumi:"label"`
	// The name of the network.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region where the network was create.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil || args.Label == nil {
		return nil, errors.New("missing required argument 'Label'")
	}
	if args == nil {
		args = &NetworkArgs{}
	}
	var resource Network
	err := ctx.RegisterResource("civo:index/network:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("civo:index/network:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// The block ip assigned to the network.
	Cidr *string `pulumi:"cidr"`
	// If is the default network.
	Default *bool `pulumi:"default"`
	// The Network label
	Label *string `pulumi:"label"`
	// The name of the network.
	Name *string `pulumi:"name"`
	// The region where the network was create.
	Region *string `pulumi:"region"`
}

type NetworkState struct {
	// The block ip assigned to the network.
	Cidr pulumi.StringPtrInput
	// If is the default network.
	Default pulumi.BoolPtrInput
	// The Network label
	Label pulumi.StringPtrInput
	// The name of the network.
	Name pulumi.StringPtrInput
	// The region where the network was create.
	Region pulumi.StringPtrInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// The Network label
	Label string `pulumi:"label"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// The Network label
	Label pulumi.StringInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}
