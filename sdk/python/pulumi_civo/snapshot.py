# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class Snapshot(pulumi.CustomResource):
    completed_at: pulumi.Output[str]
    """
    The date where the snapshot was completed.
    """
    cron_timing: pulumi.Output[str]
    """
    If a valid cron string is passed, the snapshot will be saved as an automated snapshot 
    continuing to automatically update based on the schedule of the cron sequence provided
    The default is nil meaning the snapshot will be saved as a one-off snapshot.
    """
    hostname: pulumi.Output[str]
    """
    The hostname of the instance.
    """
    instance_id: pulumi.Output[str]
    """
    The ID of the Instance from which the snapshot will be taken.
    """
    name: pulumi.Output[str]
    """
    A name for the instance snapshot.
    """
    next_execution: pulumi.Output[str]
    """
    if cron was define this date will be the next execution date.
    """
    region: pulumi.Output[str]
    """
    The region where the snapshot was take.
    """
    requested_at: pulumi.Output[str]
    """
    The date where the snapshot was requested.
    """
    safe: pulumi.Output[bool]
    """
    If `true` the instance will be shut down during the snapshot to ensure all files 
    are in a consistent state (e.g. database tables aren't in the middle of being optimised
    and hence risking corruption). The default is `false` so you experience no interruption
    of service, but a small risk of corruption.
    """
    size_gb: pulumi.Output[float]
    """
    The size of the snapshot in GB.
    """
    state: pulumi.Output[str]
    """
    The status of the snapshot.
    """
    template_id: pulumi.Output[str]
    """
    The template id.
    """
    def __init__(__self__, resource_name, opts=None, cron_timing=None, instance_id=None, name=None, safe=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource which can be used to create a snapshot from an existing Civo Instance.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_civo as civo

        myinstance_backup = civo.Snapshot("myinstance-backup", instance_id=civo_instance["myinstance"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cron_timing: If a valid cron string is passed, the snapshot will be saved as an automated snapshot 
               continuing to automatically update based on the schedule of the cron sequence provided
               The default is nil meaning the snapshot will be saved as a one-off snapshot.
        :param pulumi.Input[str] instance_id: The ID of the Instance from which the snapshot will be taken.
        :param pulumi.Input[str] name: A name for the instance snapshot.
        :param pulumi.Input[bool] safe: If `true` the instance will be shut down during the snapshot to ensure all files 
               are in a consistent state (e.g. database tables aren't in the middle of being optimised
               and hence risking corruption). The default is `false` so you experience no interruption
               of service, but a small risk of corruption.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cron_timing'] = cron_timing
            if instance_id is None:
                raise TypeError("Missing required property 'instance_id'")
            __props__['instance_id'] = instance_id
            __props__['name'] = name
            __props__['safe'] = safe
            __props__['completed_at'] = None
            __props__['hostname'] = None
            __props__['next_execution'] = None
            __props__['region'] = None
            __props__['requested_at'] = None
            __props__['size_gb'] = None
            __props__['state'] = None
            __props__['template_id'] = None
        super(Snapshot, __self__).__init__(
            'civo:index/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, completed_at=None, cron_timing=None, hostname=None, instance_id=None, name=None, next_execution=None, region=None, requested_at=None, safe=None, size_gb=None, state=None, template_id=None):
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] completed_at: The date where the snapshot was completed.
        :param pulumi.Input[str] cron_timing: If a valid cron string is passed, the snapshot will be saved as an automated snapshot 
               continuing to automatically update based on the schedule of the cron sequence provided
               The default is nil meaning the snapshot will be saved as a one-off snapshot.
        :param pulumi.Input[str] hostname: The hostname of the instance.
        :param pulumi.Input[str] instance_id: The ID of the Instance from which the snapshot will be taken.
        :param pulumi.Input[str] name: A name for the instance snapshot.
        :param pulumi.Input[str] next_execution: if cron was define this date will be the next execution date.
        :param pulumi.Input[str] region: The region where the snapshot was take.
        :param pulumi.Input[str] requested_at: The date where the snapshot was requested.
        :param pulumi.Input[bool] safe: If `true` the instance will be shut down during the snapshot to ensure all files 
               are in a consistent state (e.g. database tables aren't in the middle of being optimised
               and hence risking corruption). The default is `false` so you experience no interruption
               of service, but a small risk of corruption.
        :param pulumi.Input[float] size_gb: The size of the snapshot in GB.
        :param pulumi.Input[str] state: The status of the snapshot.
        :param pulumi.Input[str] template_id: The template id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["completed_at"] = completed_at
        __props__["cron_timing"] = cron_timing
        __props__["hostname"] = hostname
        __props__["instance_id"] = instance_id
        __props__["name"] = name
        __props__["next_execution"] = next_execution
        __props__["region"] = region
        __props__["requested_at"] = requested_at
        __props__["safe"] = safe
        __props__["size_gb"] = size_gb
        __props__["state"] = state
        __props__["template_id"] = template_id
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
