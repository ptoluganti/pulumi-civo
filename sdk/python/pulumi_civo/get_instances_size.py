# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import _utilities, _tables


class GetInstancesSizeResult:
    """
    A collection of values returned by getInstancesSize.
    """
    def __init__(__self__, filters=None, id=None, sizes=None, sorts=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if sizes and not isinstance(sizes, list):
            raise TypeError("Expected argument 'sizes' to be a list")
        __self__.sizes = sizes
        if sorts and not isinstance(sorts, list):
            raise TypeError("Expected argument 'sorts' to be a list")
        __self__.sorts = sorts


class AwaitableGetInstancesSizeResult(GetInstancesSizeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstancesSizeResult(
            filters=self.filters,
            id=self.id,
            sizes=self.sizes,
            sorts=self.sorts)


def get_instances_size(filters=None, sorts=None, opts=None):
    """
    Use this data source to access information about an existing resource.


    The **filters** object supports the following:

      * `key` (`str`)
      * `values` (`list`)

    The **sorts** object supports the following:

      * `direction` (`str`)
      * `key` (`str`)
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['sorts'] = sorts
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('civo:index/getInstancesSize:getInstancesSize', __args__, opts=opts).value

    return AwaitableGetInstancesSizeResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        sizes=__ret__.get('sizes'),
        sorts=__ret__.get('sorts'))
