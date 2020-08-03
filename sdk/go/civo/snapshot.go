// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource which can be used to create a snapshot from an existing Civo Instance.
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
// 		_, err := civo.NewSnapshot(ctx, "myinstance_backup", &civo.SnapshotArgs{
// 			InstanceId: pulumi.Any(civo_instance.Myinstance.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Snapshot struct {
	pulumi.CustomResourceState

	// The date where the snapshot was completed.
	CompletedAt pulumi.StringOutput `pulumi:"completedAt"`
	// If a valid cron string is passed, the snapshot will be saved as an automated snapshot
	// continuing to automatically update based on the schedule of the cron sequence provided
	// The default is nil meaning the snapshot will be saved as a one-off snapshot.
	CronTiming pulumi.StringPtrOutput `pulumi:"cronTiming"`
	// The hostname of the instance.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// The ID of the Instance from which the snapshot will be taken.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// A name for the instance snapshot.
	Name pulumi.StringOutput `pulumi:"name"`
	// if cron was define this date will be the next execution date.
	NextExecution pulumi.StringOutput `pulumi:"nextExecution"`
	// The region where the snapshot was take.
	Region pulumi.StringOutput `pulumi:"region"`
	// The date where the snapshot was requested.
	RequestedAt pulumi.StringOutput `pulumi:"requestedAt"`
	// If `true` the instance will be shut down during the snapshot to ensure all files
	// are in a consistent state (e.g. database tables aren't in the middle of being optimised
	// and hence risking corruption). The default is `false` so you experience no interruption
	// of service, but a small risk of corruption.
	Safe pulumi.BoolPtrOutput `pulumi:"safe"`
	// The size of the snapshot in GB.
	SizeGb pulumi.IntOutput `pulumi:"sizeGb"`
	// The status of the snapshot.
	State pulumi.StringOutput `pulumi:"state"`
	// The template id.
	TemplateId pulumi.StringOutput `pulumi:"templateId"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil || args.InstanceId == nil {
		return nil, errors.New("missing required argument 'InstanceId'")
	}
	if args == nil {
		args = &SnapshotArgs{}
	}
	var resource Snapshot
	err := ctx.RegisterResource("civo:index/snapshot:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("civo:index/snapshot:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
	// The date where the snapshot was completed.
	CompletedAt *string `pulumi:"completedAt"`
	// If a valid cron string is passed, the snapshot will be saved as an automated snapshot
	// continuing to automatically update based on the schedule of the cron sequence provided
	// The default is nil meaning the snapshot will be saved as a one-off snapshot.
	CronTiming *string `pulumi:"cronTiming"`
	// The hostname of the instance.
	Hostname *string `pulumi:"hostname"`
	// The ID of the Instance from which the snapshot will be taken.
	InstanceId *string `pulumi:"instanceId"`
	// A name for the instance snapshot.
	Name *string `pulumi:"name"`
	// if cron was define this date will be the next execution date.
	NextExecution *string `pulumi:"nextExecution"`
	// The region where the snapshot was take.
	Region *string `pulumi:"region"`
	// The date where the snapshot was requested.
	RequestedAt *string `pulumi:"requestedAt"`
	// If `true` the instance will be shut down during the snapshot to ensure all files
	// are in a consistent state (e.g. database tables aren't in the middle of being optimised
	// and hence risking corruption). The default is `false` so you experience no interruption
	// of service, but a small risk of corruption.
	Safe *bool `pulumi:"safe"`
	// The size of the snapshot in GB.
	SizeGb *int `pulumi:"sizeGb"`
	// The status of the snapshot.
	State *string `pulumi:"state"`
	// The template id.
	TemplateId *string `pulumi:"templateId"`
}

type SnapshotState struct {
	// The date where the snapshot was completed.
	CompletedAt pulumi.StringPtrInput
	// If a valid cron string is passed, the snapshot will be saved as an automated snapshot
	// continuing to automatically update based on the schedule of the cron sequence provided
	// The default is nil meaning the snapshot will be saved as a one-off snapshot.
	CronTiming pulumi.StringPtrInput
	// The hostname of the instance.
	Hostname pulumi.StringPtrInput
	// The ID of the Instance from which the snapshot will be taken.
	InstanceId pulumi.StringPtrInput
	// A name for the instance snapshot.
	Name pulumi.StringPtrInput
	// if cron was define this date will be the next execution date.
	NextExecution pulumi.StringPtrInput
	// The region where the snapshot was take.
	Region pulumi.StringPtrInput
	// The date where the snapshot was requested.
	RequestedAt pulumi.StringPtrInput
	// If `true` the instance will be shut down during the snapshot to ensure all files
	// are in a consistent state (e.g. database tables aren't in the middle of being optimised
	// and hence risking corruption). The default is `false` so you experience no interruption
	// of service, but a small risk of corruption.
	Safe pulumi.BoolPtrInput
	// The size of the snapshot in GB.
	SizeGb pulumi.IntPtrInput
	// The status of the snapshot.
	State pulumi.StringPtrInput
	// The template id.
	TemplateId pulumi.StringPtrInput
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// If a valid cron string is passed, the snapshot will be saved as an automated snapshot
	// continuing to automatically update based on the schedule of the cron sequence provided
	// The default is nil meaning the snapshot will be saved as a one-off snapshot.
	CronTiming *string `pulumi:"cronTiming"`
	// The ID of the Instance from which the snapshot will be taken.
	InstanceId string `pulumi:"instanceId"`
	// A name for the instance snapshot.
	Name *string `pulumi:"name"`
	// If `true` the instance will be shut down during the snapshot to ensure all files
	// are in a consistent state (e.g. database tables aren't in the middle of being optimised
	// and hence risking corruption). The default is `false` so you experience no interruption
	// of service, but a small risk of corruption.
	Safe *bool `pulumi:"safe"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// If a valid cron string is passed, the snapshot will be saved as an automated snapshot
	// continuing to automatically update based on the schedule of the cron sequence provided
	// The default is nil meaning the snapshot will be saved as a one-off snapshot.
	CronTiming pulumi.StringPtrInput
	// The ID of the Instance from which the snapshot will be taken.
	InstanceId pulumi.StringInput
	// A name for the instance snapshot.
	Name pulumi.StringPtrInput
	// If `true` the instance will be shut down during the snapshot to ensure all files
	// are in a consistent state (e.g. database tables aren't in the middle of being optimised
	// and hence risking corruption). The default is `false` so you experience no interruption
	// of service, but a small risk of corruption.
	Safe pulumi.BoolPtrInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}