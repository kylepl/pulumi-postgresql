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
    /// The ``postgresql..DefaultPrivileges`` resource creates and manages default privileges given to a user for a database schema.
    /// 
    /// &gt; **Note:** This resource needs Postgresql version 9 or above.
    /// </summary>
    public partial class DefaultPrivileges : Pulumi.CustomResource
    {
        /// <summary>
        /// The database to grant default privileges for this role.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to set the default privileges on (one of: table, sequence).
        /// </summary>
        [Output("objectType")]
        public Output<string> ObjectType { get; private set; } = null!;

        /// <summary>
        /// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// The list of privileges to apply as default privileges.
        /// </summary>
        [Output("privileges")]
        public Output<ImmutableArray<string>> Privileges { get; private set; } = null!;

        /// <summary>
        /// The name of the role to which grant default privileges on.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The database schema to set default privileges for this role.
        /// </summary>
        [Output("schema")]
        public Output<string> Schema { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultPrivileges resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultPrivileges(string name, DefaultPrivilegesArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/defaultPrivileges:DefaultPrivileges", name, args ?? new DefaultPrivilegesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultPrivileges(string name, Input<string> id, DefaultPrivilegesState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/defaultPrivileges:DefaultPrivileges", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,                Aliases = { new Alias { Type = "postgresql:index/defaultPrivileg:DefaultPrivileg"} },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DefaultPrivileges resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultPrivileges Get(string name, Input<string> id, DefaultPrivilegesState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultPrivileges(name, id, state, options);
        }
    }

    public sealed class DefaultPrivilegesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database to grant default privileges for this role.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The PostgreSQL object type to set the default privileges on (one of: table, sequence).
        /// </summary>
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        /// <summary>
        /// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
        /// </summary>
        [Input("owner", required: true)]
        public Input<string> Owner { get; set; } = null!;

        [Input("privileges", required: true)]
        private InputList<string>? _privileges;

        /// <summary>
        /// The list of privileges to apply as default privileges.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The name of the role to which grant default privileges on.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The database schema to set default privileges for this role.
        /// </summary>
        [Input("schema", required: true)]
        public Input<string> Schema { get; set; } = null!;

        public DefaultPrivilegesArgs()
        {
        }
    }

    public sealed class DefaultPrivilegesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The database to grant default privileges for this role.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The PostgreSQL object type to set the default privileges on (one of: table, sequence).
        /// </summary>
        [Input("objectType")]
        public Input<string>? ObjectType { get; set; }

        /// <summary>
        /// Role for which apply default privileges (You can change default privileges only for objects that will be created by yourself or by roles that you are a member of).
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("privileges")]
        private InputList<string>? _privileges;

        /// <summary>
        /// The list of privileges to apply as default privileges.
        /// </summary>
        public InputList<string> Privileges
        {
            get => _privileges ?? (_privileges = new InputList<string>());
            set => _privileges = value;
        }

        /// <summary>
        /// The name of the role to which grant default privileges on.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The database schema to set default privileges for this role.
        /// </summary>
        [Input("schema")]
        public Input<string>? Schema { get; set; }

        public DefaultPrivilegesState()
        {
        }
    }
}
