# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'FunctionArgArgs',
    'ProviderClientcertArgs',
    'SchemaPolicyArgs',
]

@pulumi.input_type
class FunctionArgArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 default: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The type of the argument.
        :param pulumi.Input[str] default: An expression to be used as default value if the parameter is not specified.
        :param pulumi.Input[str] mode: Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
        :param pulumi.Input[str] name: The name of the argument.
        """
        pulumi.set(__self__, "type", type)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the argument.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[str]]:
        """
        An expression to be used as default value if the parameter is not specified.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Can be one of IN, INOUT, OUT, or VARIADIC. Default is IN.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the argument.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ProviderClientcertArgs:
    def __init__(__self__, *,
                 cert: pulumi.Input[str],
                 key: pulumi.Input[str]):
        pulumi.set(__self__, "cert", cert)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def cert(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)


@pulumi.input_type
class SchemaPolicyArgs:
    def __init__(__self__, *,
                 create: Optional[pulumi.Input[bool]] = None,
                 create_with_grant: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 usage: Optional[pulumi.Input[bool]] = None,
                 usage_with_grant: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] create: Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        :param pulumi.Input[bool] create_with_grant: Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        :param pulumi.Input[str] role: The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        :param pulumi.Input[bool] usage: Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        :param pulumi.Input[bool] usage_with_grant: Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
        """
        if create is not None:
            pulumi.set(__self__, "create", create)
        if create_with_grant is not None:
            pulumi.set(__self__, "create_with_grant", create_with_grant)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if usage_with_grant is not None:
            pulumi.set(__self__, "usage_with_grant", usage_with_grant)

    @property
    @pulumi.getter
    def create(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "create")

    @create.setter
    def create(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create", value)

    @property
    @pulumi.getter(name="createWithGrant")
    def create_with_grant(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        """
        return pulumi.get(self, "create_with_grant")

    @create_with_grant.setter
    def create_with_grant(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_with_grant", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def usage(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "usage")

    @usage.setter
    def usage(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "usage", value)

    @property
    @pulumi.getter(name="usageWithGrant")
    def usage_with_grant(self) -> Optional[pulumi.Input[bool]]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
        """
        return pulumi.get(self, "usage_with_grant")

    @usage_with_grant.setter
    def usage_with_grant(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "usage_with_grant", value)


