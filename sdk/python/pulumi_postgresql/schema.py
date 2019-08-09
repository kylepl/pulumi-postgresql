# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from . import utilities, tables

class Schema(pulumi.CustomResource):
    if_not_exists: pulumi.Output[bool]
    """
    When true, use the existing schema if it exists. (Default: true)
    """
    name: pulumi.Output[str]
    """
    The name of the schema. Must be unique in the PostgreSQL
    database instance where it is configured.
    """
    owner: pulumi.Output[str]
    """
    The ROLE who owns the schema.
    """
    policies: pulumi.Output[list]
    """
    Can be specified multiple times for each policy.  Each
    policy block supports fields documented below.
    """
    def __init__(__self__, resource_name, opts=None, if_not_exists=None, name=None, owner=None, policies=None, __props__=None, __name__=None, __opts__=None):
        """
        The ``.Schema`` resource creates and manages [schema
        objects](https://www.postgresql.org/docs/current/static/ddl-schemas.html) within
        a PostgreSQL database.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] if_not_exists: When true, use the existing schema if it exists. (Default: true)
        :param pulumi.Input[str] name: The name of the schema. Must be unique in the PostgreSQL
               database instance where it is configured.
        :param pulumi.Input[str] owner: The ROLE who owns the schema.
        :param pulumi.Input[list] policies: Can be specified multiple times for each policy.  Each
               policy block supports fields documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/schema.html.markdown.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

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
    def get(resource_name, id, opts=None, if_not_exists=None, name=None, owner=None, policies=None):
        """
        Get an existing Schema resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] if_not_exists: When true, use the existing schema if it exists. (Default: true)
        :param pulumi.Input[str] name: The name of the schema. Must be unique in the PostgreSQL
               database instance where it is configured.
        :param pulumi.Input[str] owner: The ROLE who owns the schema.
        :param pulumi.Input[list] policies: Can be specified multiple times for each policy.  Each
               policy block supports fields documented below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/schema.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["if_not_exists"] = if_not_exists
        __props__["name"] = name
        __props__["owner"] = owner
        __props__["policies"] = policies
        return Schema(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

