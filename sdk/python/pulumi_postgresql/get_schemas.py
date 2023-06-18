# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetSchemasResult',
    'AwaitableGetSchemasResult',
    'get_schemas',
    'get_schemas_output',
]

@pulumi.output_type
class GetSchemasResult:
    """
    A collection of values returned by getSchemas.
    """
    def __init__(__self__, database=None, id=None, include_system_schemas=None, like_all_patterns=None, like_any_patterns=None, not_like_all_patterns=None, regex_pattern=None, schemas=None):
        if database and not isinstance(database, str):
            raise TypeError("Expected argument 'database' to be a str")
        pulumi.set(__self__, "database", database)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_system_schemas and not isinstance(include_system_schemas, bool):
            raise TypeError("Expected argument 'include_system_schemas' to be a bool")
        pulumi.set(__self__, "include_system_schemas", include_system_schemas)
        if like_all_patterns and not isinstance(like_all_patterns, list):
            raise TypeError("Expected argument 'like_all_patterns' to be a list")
        pulumi.set(__self__, "like_all_patterns", like_all_patterns)
        if like_any_patterns and not isinstance(like_any_patterns, list):
            raise TypeError("Expected argument 'like_any_patterns' to be a list")
        pulumi.set(__self__, "like_any_patterns", like_any_patterns)
        if not_like_all_patterns and not isinstance(not_like_all_patterns, list):
            raise TypeError("Expected argument 'not_like_all_patterns' to be a list")
        pulumi.set(__self__, "not_like_all_patterns", not_like_all_patterns)
        if regex_pattern and not isinstance(regex_pattern, str):
            raise TypeError("Expected argument 'regex_pattern' to be a str")
        pulumi.set(__self__, "regex_pattern", regex_pattern)
        if schemas and not isinstance(schemas, list):
            raise TypeError("Expected argument 'schemas' to be a list")
        pulumi.set(__self__, "schemas", schemas)

    @property
    @pulumi.getter
    def database(self) -> str:
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeSystemSchemas")
    def include_system_schemas(self) -> Optional[bool]:
        return pulumi.get(self, "include_system_schemas")

    @property
    @pulumi.getter(name="likeAllPatterns")
    def like_all_patterns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "like_all_patterns")

    @property
    @pulumi.getter(name="likeAnyPatterns")
    def like_any_patterns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "like_any_patterns")

    @property
    @pulumi.getter(name="notLikeAllPatterns")
    def not_like_all_patterns(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "not_like_all_patterns")

    @property
    @pulumi.getter(name="regexPattern")
    def regex_pattern(self) -> Optional[str]:
        return pulumi.get(self, "regex_pattern")

    @property
    @pulumi.getter
    def schemas(self) -> Sequence[str]:
        """
        A list of full names of found schemas.
        """
        return pulumi.get(self, "schemas")


class AwaitableGetSchemasResult(GetSchemasResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSchemasResult(
            database=self.database,
            id=self.id,
            include_system_schemas=self.include_system_schemas,
            like_all_patterns=self.like_all_patterns,
            like_any_patterns=self.like_any_patterns,
            not_like_all_patterns=self.not_like_all_patterns,
            regex_pattern=self.regex_pattern,
            schemas=self.schemas)


def get_schemas(database: Optional[str] = None,
                include_system_schemas: Optional[bool] = None,
                like_all_patterns: Optional[Sequence[str]] = None,
                like_any_patterns: Optional[Sequence[str]] = None,
                not_like_all_patterns: Optional[Sequence[str]] = None,
                regex_pattern: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSchemasResult:
    """
    The `_get_schemas_` data source retrieves a list of schema names from a specified PostgreSQL database.

    ## Usage

    ```python
    import pulumi
    import pulumi_postgresql as postgresql

    my_schemas = postgresql.get_schemas(database="my_database")
    ```


    :param str database: The PostgreSQL database which will be queried for schema names.
    :param bool include_system_schemas: Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
    :param Sequence[str] like_all_patterns: List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
    :param Sequence[str] like_any_patterns: List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
    :param Sequence[str] not_like_all_patterns: List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
    :param str regex_pattern: Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
    """
    __args__ = dict()
    __args__['database'] = database
    __args__['includeSystemSchemas'] = include_system_schemas
    __args__['likeAllPatterns'] = like_all_patterns
    __args__['likeAnyPatterns'] = like_any_patterns
    __args__['notLikeAllPatterns'] = not_like_all_patterns
    __args__['regexPattern'] = regex_pattern
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('postgresql:index/getSchemas:getSchemas', __args__, opts=opts, typ=GetSchemasResult).value

    return AwaitableGetSchemasResult(
        database=__ret__.database,
        id=__ret__.id,
        include_system_schemas=__ret__.include_system_schemas,
        like_all_patterns=__ret__.like_all_patterns,
        like_any_patterns=__ret__.like_any_patterns,
        not_like_all_patterns=__ret__.not_like_all_patterns,
        regex_pattern=__ret__.regex_pattern,
        schemas=__ret__.schemas)


@_utilities.lift_output_func(get_schemas)
def get_schemas_output(database: Optional[pulumi.Input[str]] = None,
                       include_system_schemas: Optional[pulumi.Input[Optional[bool]]] = None,
                       like_all_patterns: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       like_any_patterns: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       not_like_all_patterns: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                       regex_pattern: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSchemasResult]:
    """
    The `_get_schemas_` data source retrieves a list of schema names from a specified PostgreSQL database.

    ## Usage

    ```python
    import pulumi
    import pulumi_postgresql as postgresql

    my_schemas = postgresql.get_schemas(database="my_database")
    ```


    :param str database: The PostgreSQL database which will be queried for schema names.
    :param bool include_system_schemas: Determines whether to include system schemas (pg_ prefix and information_schema). 'public' will always be included. Defaults to ``false``.
    :param Sequence[str] like_all_patterns: List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ALL`` operators.
    :param Sequence[str] like_any_patterns: List of expressions which will be pattern matched in the query using the PostgreSQL ``LIKE ANY`` operators.
    :param Sequence[str] not_like_all_patterns: List of expressions which will be pattern matched in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
    :param str regex_pattern: Expression which will be pattern matched in the query using the PostgreSQL ``~`` (regular expression match) operator.
    """
    ...