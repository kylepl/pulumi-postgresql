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
    /// The ``postgresql.ReplicationSlot`` resource creates and manages a replication slot on a PostgreSQL
    /// server.
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
    ///         var mySlot = new PostgreSql.ReplicationSlot("mySlot", new PostgreSql.ReplicationSlotArgs
    ///         {
    ///             Plugin = "test_decoding",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/replicationSlot:ReplicationSlot")]
    public partial class ReplicationSlot : Pulumi.CustomResource
    {
        /// <summary>
        /// Which database to create the replication slot on. Defaults to provider database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// The name of the replication slot.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sets the output plugin.
        /// </summary>
        [Output("plugin")]
        public Output<string> Plugin { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationSlot resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationSlot(string name, ReplicationSlotArgs args, CustomResourceOptions? options = null)
            : base("postgresql:index/replicationSlot:ReplicationSlot", name, args ?? new ReplicationSlotArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationSlot(string name, Input<string> id, ReplicationSlotState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/replicationSlot:ReplicationSlot", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReplicationSlot resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationSlot Get(string name, Input<string> id, ReplicationSlotState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationSlot(name, id, state, options);
        }
    }

    public sealed class ReplicationSlotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Which database to create the replication slot on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The name of the replication slot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the output plugin.
        /// </summary>
        [Input("plugin", required: true)]
        public Input<string> Plugin { get; set; } = null!;

        public ReplicationSlotArgs()
        {
        }
    }

    public sealed class ReplicationSlotState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Which database to create the replication slot on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// The name of the replication slot.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the output plugin.
        /// </summary>
        [Input("plugin")]
        public Input<string>? Plugin { get; set; }

        public ReplicationSlotState()
        {
        }
    }
}
