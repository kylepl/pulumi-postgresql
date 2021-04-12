# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['GrantRoleArgs', 'GrantRole']

@pulumi.input_type
class GrantRoleArgs:
    def __init__(__self__, *,
                 grant_role: pulumi.Input[str],
                 role: pulumi.Input[str],
                 with_admin_option: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a GrantRole resource.
        :param pulumi.Input[str] grant_role: The name of the role that is added to `role`.
        :param pulumi.Input[str] role: The name of the role that is granted a new membership.
        :param pulumi.Input[bool] with_admin_option: Giving ability to grant membership to others or not for `role`. (Default: false)
        """
        pulumi.set(__self__, "grant_role", grant_role)
        pulumi.set(__self__, "role", role)
        if with_admin_option is not None:
            pulumi.set(__self__, "with_admin_option", with_admin_option)

    @property
    @pulumi.getter(name="grantRole")
    def grant_role(self) -> pulumi.Input[str]:
        """
        The name of the role that is added to `role`.
        """
        return pulumi.get(self, "grant_role")

    @grant_role.setter
    def grant_role(self, value: pulumi.Input[str]):
        pulumi.set(self, "grant_role", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The name of the role that is granted a new membership.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="withAdminOption")
    def with_admin_option(self) -> Optional[pulumi.Input[bool]]:
        """
        Giving ability to grant membership to others or not for `role`. (Default: false)
        """
        return pulumi.get(self, "with_admin_option")

    @with_admin_option.setter
    def with_admin_option(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "with_admin_option", value)


class GrantRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_role: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 with_admin_option: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The ``GrantRole`` resource creates and manages membership in a role to one or more other roles in a non-authoritative way.

        When using ``GrantRole`` resource it is likely because the PostgreSQL role you are modifying was created outside of this provider.

        > **Note:** This resource needs PostgreSQL version 9 or above.

        > **Note:** `GrantRole` **cannot** be used in conjunction with `Role` or they will fight over what your role grants should be.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        grant_root = postgresql.GrantRole("grantRoot",
            grant_role="application",
            role="root",
            with_admin_option=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_role: The name of the role that is added to `role`.
        :param pulumi.Input[str] role: The name of the role that is granted a new membership.
        :param pulumi.Input[bool] with_admin_option: Giving ability to grant membership to others or not for `role`. (Default: false)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrantRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The ``GrantRole`` resource creates and manages membership in a role to one or more other roles in a non-authoritative way.

        When using ``GrantRole`` resource it is likely because the PostgreSQL role you are modifying was created outside of this provider.

        > **Note:** This resource needs PostgreSQL version 9 or above.

        > **Note:** `GrantRole` **cannot** be used in conjunction with `Role` or they will fight over what your role grants should be.

        ## Usage

        ```python
        import pulumi
        import pulumi_postgresql as postgresql

        grant_root = postgresql.GrantRole("grantRoot",
            grant_role="application",
            role="root",
            with_admin_option=True)
        ```

        :param str resource_name: The name of the resource.
        :param GrantRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrantRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 grant_role: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 with_admin_option: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if grant_role is None and not opts.urn:
                raise TypeError("Missing required property 'grant_role'")
            __props__['grant_role'] = grant_role
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['with_admin_option'] = with_admin_option
        super(GrantRole, __self__).__init__(
            'postgresql:index/grantRole:GrantRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            grant_role: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            with_admin_option: Optional[pulumi.Input[bool]] = None) -> 'GrantRole':
        """
        Get an existing GrantRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] grant_role: The name of the role that is added to `role`.
        :param pulumi.Input[str] role: The name of the role that is granted a new membership.
        :param pulumi.Input[bool] with_admin_option: Giving ability to grant membership to others or not for `role`. (Default: false)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["grant_role"] = grant_role
        __props__["role"] = role
        __props__["with_admin_option"] = with_admin_option
        return GrantRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="grantRole")
    def grant_role(self) -> pulumi.Output[str]:
        """
        The name of the role that is added to `role`.
        """
        return pulumi.get(self, "grant_role")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The name of the role that is granted a new membership.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="withAdminOption")
    def with_admin_option(self) -> pulumi.Output[Optional[bool]]:
        """
        Giving ability to grant membership to others or not for `role`. (Default: false)
        """
        return pulumi.get(self, "with_admin_option")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

