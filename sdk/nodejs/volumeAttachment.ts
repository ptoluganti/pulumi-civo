// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages attaching a Volume to a Instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as civo from "@pulumi/civo";
 *
 * const db = new civo.Volume("db", {
 *     sizeGb: 60,
 *     bootable: false,
 * });
 * const foobar = new civo.VolumeAttachment("foobar", {
 *     instanceId: civo_instance["my-test-instance"].id,
 *     volumeId: db.id,
 * });
 * ```
 */
export class VolumeAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VolumeAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VolumeAttachmentState, opts?: pulumi.CustomResourceOptions): VolumeAttachment {
        return new VolumeAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'civo:index/volumeAttachment:VolumeAttachment';

    /**
     * Returns true if the given object is an instance of VolumeAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VolumeAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VolumeAttachment.__pulumiType;
    }

    /**
     * ID of the instance to attach the volume to.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * ID of the Volume to be attached to the instance.
     */
    public readonly volumeId!: pulumi.Output<string>;

    /**
     * Create a VolumeAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VolumeAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VolumeAttachmentArgs | VolumeAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VolumeAttachmentState | undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["volumeId"] = state ? state.volumeId : undefined;
        } else {
            const args = argsOrState as VolumeAttachmentArgs | undefined;
            if (!args || args.instanceId === undefined) {
                throw new Error("Missing required property 'instanceId'");
            }
            if (!args || args.volumeId === undefined) {
                throw new Error("Missing required property 'volumeId'");
            }
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["volumeId"] = args ? args.volumeId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VolumeAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VolumeAttachment resources.
 */
export interface VolumeAttachmentState {
    /**
     * ID of the instance to attach the volume to.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * ID of the Volume to be attached to the instance.
     */
    readonly volumeId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VolumeAttachment resource.
 */
export interface VolumeAttachmentArgs {
    /**
     * ID of the instance to attach the volume to.
     */
    readonly instanceId: pulumi.Input<string>;
    /**
     * ID of the Volume to be attached to the instance.
     */
    readonly volumeId: pulumi.Input<string>;
}
