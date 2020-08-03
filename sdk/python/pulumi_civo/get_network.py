# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class GetNetworkResult:
    """
    A collection of values returned by getNetwork.
    """
    def __init__(__self__, cidr=None, default=None, id=None, label=None, name=None, region=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        __self__.cidr = cidr
        """
        The block ip assigned to the network.
        """
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        __self__.default = default
        """
        If is the default network.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        A unique ID that can be used to identify and reference a Network.
        """
        if label and not isinstance(label, str):
            raise TypeError("Expected argument 'label' to be a str")
        __self__.label = label
        """
        The label used in the configuration.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the network.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        The region where the network was create.
        """


class AwaitableGetNetworkResult(GetNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkResult(
            cidr=self.cidr,
            default=self.default,
            id=self.id,
            label=self.label,
            name=self.name,
            region=self.region)


def get_network(id=None, label=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str id: The unique identifier of an existing Network.
    :param str label: The name of an existing Network.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['label'] = label
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('civo:index/getNetwork:getNetwork', __args__, opts=opts).value

    return AwaitableGetNetworkResult(
        cidr=__ret__.get('cidr'),
        default=__ret__.get('default'),
        id=__ret__.get('id'),
        label=__ret__.get('label'),
        name=__ret__.get('name'),
        region=__ret__.get('region'))