# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Schema']


class Schema(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 drop_cascade: Optional[pulumi.Input[bool]] = None,
                 if_not_exists: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SchemaPolicyArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Schema resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
        :param pulumi.Input[bool] drop_cascade: When true, will also drop all the objects that are contained in the schema. (Default: false)
        :param pulumi.Input[bool] if_not_exists: When true, use the existing schema if it exists. (Default: true)
        :param pulumi.Input[str] name: The name of the schema. Must be unique in the PostgreSQL
               database instance where it is configured.
        :param pulumi.Input[str] owner: The ROLE who owns the schema.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SchemaPolicyArgs']]]] policies: Can be specified multiple times for each policy.  Each
               policy block supports fields documented below.
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

            __props__['database'] = database
            __props__['drop_cascade'] = drop_cascade
            __props__['if_not_exists'] = if_not_exists
            __props__['name'] = name
            __props__['owner'] = owner
            __props__['policies'] = policies
        super(Schema, __self__).__init__(
            'postgresql:index/schema:Schema',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            database: Optional[pulumi.Input[str]] = None,
            drop_cascade: Optional[pulumi.Input[bool]] = None,
            if_not_exists: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SchemaPolicyArgs']]]]] = None) -> 'Schema':
        """
        Get an existing Schema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
        :param pulumi.Input[bool] drop_cascade: When true, will also drop all the objects that are contained in the schema. (Default: false)
        :param pulumi.Input[bool] if_not_exists: When true, use the existing schema if it exists. (Default: true)
        :param pulumi.Input[str] name: The name of the schema. Must be unique in the PostgreSQL
               database instance where it is configured.
        :param pulumi.Input[str] owner: The ROLE who owns the schema.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SchemaPolicyArgs']]]] policies: Can be specified multiple times for each policy.  Each
               policy block supports fields documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["database"] = database
        __props__["drop_cascade"] = drop_cascade
        __props__["if_not_exists"] = if_not_exists
        __props__["name"] = name
        __props__["owner"] = owner
        __props__["policies"] = policies
        return Schema(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        The DATABASE in which where this schema will be created. (Default: The database used by your `provider` configuration)
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter(name="dropCascade")
    def drop_cascade(self) -> pulumi.Output[Optional[bool]]:
        """
        When true, will also drop all the objects that are contained in the schema. (Default: false)
        """
        return pulumi.get(self, "drop_cascade")

    @property
    @pulumi.getter(name="ifNotExists")
    def if_not_exists(self) -> pulumi.Output[Optional[bool]]:
        """
        When true, use the existing schema if it exists. (Default: true)
        """
        return pulumi.get(self, "if_not_exists")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the schema. Must be unique in the PostgreSQL
        database instance where it is configured.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        The ROLE who owns the schema.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence['outputs.SchemaPolicy']]:
        """
        Can be specified multiple times for each policy.  Each
        policy block supports fields documented below.
        """
        return pulumi.get(self, "policies")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

