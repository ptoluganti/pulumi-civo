// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Civo
{
    /// <summary>
    /// Provides a Civo Template resource.
    /// This can be used to create, modify, and delete Templates.
    /// </summary>
    public partial class Template : Pulumi.CustomResource
    {
        /// <summary>
        /// Commonly referred to as 'user-data', this is a customisation script that is run after
        /// the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic
        /// way of configuring cloud servers. If you put `$INITIAL_USER` in your script, this will automatically
        /// be replaced by the initial user chosen when creating the instance, `$INITIAL_PASSWORD` will be
        /// replaced with the random password generated by the system, `$HOSTNAME` is the fully qualified
        /// domain name of the instance and `$SSH_KEY` will be the content of the SSH public key.
        /// (this is technically optional, but you won't really be able to use instances without it -
        /// see our learn guide on templates for more information).
        /// </summary>
        [Output("cloudConfig")]
        public Output<string?> CloudConfig { get; private set; } = null!;

        /// <summary>
        /// This is a unqiue, alphanumerical, short, human readable code for the template.
        /// </summary>
        [Output("code")]
        public Output<string> Code { get; private set; } = null!;

        /// <summary>
        /// The default username to suggest that the user creates
        /// </summary>
        [Output("defaultUsername")]
        public Output<string?> DefaultUsername { get; private set; } = null!;

        /// <summary>
        /// A multi-line description of the template, in Markdown format
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// This is the Image ID of any default template or the ID of another template
        /// either owned by you or global (optional; but must be specified if no volume_id is specified).
        /// </summary>
        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        /// <summary>
        /// This is a short human readable name for the template
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A one line description of the template
        /// </summary>
        [Output("shortDescription")]
        public Output<string?> ShortDescription { get; private set; } = null!;

        /// <summary>
        /// This is the ID of a bootable volume, either owned by you or global
        /// (optional; but must be specified if no image_id is specified)
        /// </summary>
        [Output("volumeId")]
        public Output<string?> VolumeId { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs args, CustomResourceOptions? options = null)
            : base("civo:index/template:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
            : base("civo:index/template:Template", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, TemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new Template(name, id, state, options);
        }
    }

    public sealed class TemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Commonly referred to as 'user-data', this is a customisation script that is run after
        /// the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic
        /// way of configuring cloud servers. If you put `$INITIAL_USER` in your script, this will automatically
        /// be replaced by the initial user chosen when creating the instance, `$INITIAL_PASSWORD` will be
        /// replaced with the random password generated by the system, `$HOSTNAME` is the fully qualified
        /// domain name of the instance and `$SSH_KEY` will be the content of the SSH public key.
        /// (this is technically optional, but you won't really be able to use instances without it -
        /// see our learn guide on templates for more information).
        /// </summary>
        [Input("cloudConfig")]
        public Input<string>? CloudConfig { get; set; }

        /// <summary>
        /// This is a unqiue, alphanumerical, short, human readable code for the template.
        /// </summary>
        [Input("code", required: true)]
        public Input<string> Code { get; set; } = null!;

        /// <summary>
        /// The default username to suggest that the user creates
        /// </summary>
        [Input("defaultUsername")]
        public Input<string>? DefaultUsername { get; set; }

        /// <summary>
        /// A multi-line description of the template, in Markdown format
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This is the Image ID of any default template or the ID of another template
        /// either owned by you or global (optional; but must be specified if no volume_id is specified).
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// This is a short human readable name for the template
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A one line description of the template
        /// </summary>
        [Input("shortDescription")]
        public Input<string>? ShortDescription { get; set; }

        /// <summary>
        /// This is the ID of a bootable volume, either owned by you or global
        /// (optional; but must be specified if no image_id is specified)
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public TemplateArgs()
        {
        }
    }

    public sealed class TemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Commonly referred to as 'user-data', this is a customisation script that is run after
        /// the instance is first booted. We recommend using cloud-config as it's a great distribution-agnostic
        /// way of configuring cloud servers. If you put `$INITIAL_USER` in your script, this will automatically
        /// be replaced by the initial user chosen when creating the instance, `$INITIAL_PASSWORD` will be
        /// replaced with the random password generated by the system, `$HOSTNAME` is the fully qualified
        /// domain name of the instance and `$SSH_KEY` will be the content of the SSH public key.
        /// (this is technically optional, but you won't really be able to use instances without it -
        /// see our learn guide on templates for more information).
        /// </summary>
        [Input("cloudConfig")]
        public Input<string>? CloudConfig { get; set; }

        /// <summary>
        /// This is a unqiue, alphanumerical, short, human readable code for the template.
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// The default username to suggest that the user creates
        /// </summary>
        [Input("defaultUsername")]
        public Input<string>? DefaultUsername { get; set; }

        /// <summary>
        /// A multi-line description of the template, in Markdown format
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// This is the Image ID of any default template or the ID of another template
        /// either owned by you or global (optional; but must be specified if no volume_id is specified).
        /// </summary>
        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        /// <summary>
        /// This is a short human readable name for the template
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A one line description of the template
        /// </summary>
        [Input("shortDescription")]
        public Input<string>? ShortDescription { get; set; }

        /// <summary>
        /// This is the ID of a bootable volume, either owned by you or global
        /// (optional; but must be specified if no image_id is specified)
        /// </summary>
        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        public TemplateState()
        {
        }
    }
}