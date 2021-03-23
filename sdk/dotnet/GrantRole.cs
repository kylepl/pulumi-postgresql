// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql
{
    /// <summary>
    /// The ``postgresql.GrantRole`` resource creates and manages membership in a role to one or more other roles in a non-authoritative way.
    /// 
    /// When using ``postgresql.GrantRole`` resource it is likely because the PostgreSQL role you are modifying was created outside of this provider.
    /// 
    /// &gt; **Note:** This resource needs PostgreSQL version 9 or above.
    /// 
    /// &gt; **Note:** `postgresql.GrantRole` **cannot** be used in conjunction with `postgresql.Role` or they will fight over what your role grants should be.
    /// 
    /// ## Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using PostgreSql = Pulumi.PostgreSql;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var grantRoot = new PostgreSql.GrantRole("grantRoot", new PostgreSql.GrantRoleArgs
    ///         {
    ///             GrantRole = "application",
    ///             Role = "root",
    ///             WithAdminOption = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/grantRole:GrantRole")]
    public partial class GrantRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the role that is added to `role`.
        /// </summary>
        [Output("grantRole")]
        public Output<string> GrantRoleName { get; private set; } = null!;

        /// <summary>
        /// The name of the role that is granted a new membership.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Giving ability to grant membership to others or not for `role`. (Default: false)
        /// </summary>
        [Output("withAdminOption")]
        public Output<bool?> WithAdminOption { get; private set; } = null!;


        /// <summary>
        /// Create a GrantRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GrantRole(string name, GrantRoleArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/grantRole:GrantRole", name, args ?? new GrantRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GrantRole(string name, Input<string> id, GrantRoleState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/grantRole:GrantRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GrantRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GrantRole Get(string name, Input<string> id, GrantRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new GrantRole(name, id, state, options);
        }
    }

    public sealed class GrantRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the role that is added to `role`.
        /// </summary>
        [Input("grantRole", required: true)]
        public Input<string> GrantRoleName { get; set; } = null!;

        /// <summary>
        /// The name of the role that is granted a new membership.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// Giving ability to grant membership to others or not for `role`. (Default: false)
        /// </summary>
        [Input("withAdminOption")]
        public Input<bool>? WithAdminOption { get; set; }

        public GrantRoleArgs()
        {
        }
    }

    public sealed class GrantRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the role that is added to `role`.
        /// </summary>
        [Input("grantRole")]
        public Input<string>? GrantRoleName { get; set; }

        /// <summary>
        /// The name of the role that is granted a new membership.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Giving ability to grant membership to others or not for `role`. (Default: false)
        /// </summary>
        [Input("withAdminOption")]
        public Input<bool>? WithAdminOption { get; set; }

        public GrantRoleState()
        {
        }
    }
}