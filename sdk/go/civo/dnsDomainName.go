// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Civo dns domain name resource.
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
// 		_, err := civo.NewDnsDomainName(ctx, "main", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DnsDomainName struct {
	pulumi.CustomResourceState

	// The id account of the domain
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The name of the domain
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewDnsDomainName registers a new resource with the given unique name, arguments, and options.
func NewDnsDomainName(ctx *pulumi.Context,
	name string, args *DnsDomainNameArgs, opts ...pulumi.ResourceOption) (*DnsDomainName, error) {
	if args == nil {
		args = &DnsDomainNameArgs{}
	}
	var resource DnsDomainName
	err := ctx.RegisterResource("civo:index/dnsDomainName:DnsDomainName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsDomainName gets an existing DnsDomainName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsDomainName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsDomainNameState, opts ...pulumi.ResourceOption) (*DnsDomainName, error) {
	var resource DnsDomainName
	err := ctx.ReadResource("civo:index/dnsDomainName:DnsDomainName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsDomainName resources.
type dnsDomainNameState struct {
	// The id account of the domain
	AccountId *string `pulumi:"accountId"`
	// The name of the domain
	Name *string `pulumi:"name"`
}

type DnsDomainNameState struct {
	// The id account of the domain
	AccountId pulumi.StringPtrInput
	// The name of the domain
	Name pulumi.StringPtrInput
}

func (DnsDomainNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainNameState)(nil)).Elem()
}

type dnsDomainNameArgs struct {
	// The name of the domain
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a DnsDomainName resource.
type DnsDomainNameArgs struct {
	// The name of the domain
	Name pulumi.StringPtrInput
}

func (DnsDomainNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsDomainNameArgs)(nil)).Elem()
}