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
    /// The `postgresql.Publication` resource creates and manages a publication on a PostgreSQL
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
    ///         var publication = new PostgreSql.Publication("publication", new PostgreSql.PublicationArgs
    ///         {
    ///             Tables = 
    ///             {
    ///                 "public.test",
    ///                 "another_schema.test",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [PostgreSqlResourceType("postgresql:index/publication:Publication")]
    public partial class Publication : Pulumi.CustomResource
    {
        /// <summary>
        /// Should be ALL TABLES added to the publication. Defaults to 'false'
        /// </summary>
        [Output("allTables")]
        public Output<bool> AllTables { get; private set; } = null!;

        /// <summary>
        /// Which database to create the publication on. Defaults to provider database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// Should all subsequent resources of the publication be dropped. Defaults to 'false'
        /// </summary>
        [Output("dropCascade")]
        public Output<bool?> DropCascade { get; private set; } = null!;

        /// <summary>
        /// The name of the publication.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Who owns the publication. Defaults to provider user.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Which 'publish' options should be turned on. Default to 'insert','update','delete'
        /// </summary>
        [Output("publishParams")]
        public Output<ImmutableArray<string>> PublishParams { get; private set; } = null!;

        /// <summary>
        /// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        /// </summary>
        [Output("publishViaPartitionRootParam")]
        public Output<bool?> PublishViaPartitionRootParam { get; private set; } = null!;

        /// <summary>
        /// Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
        /// </summary>
        [Output("tables")]
        public Output<ImmutableArray<string>> Tables { get; private set; } = null!;


        /// <summary>
        /// Create a Publication resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Publication(string name, PublicationArgs? args = null, CustomResourceOptions? options = null)
            : base("postgresql:index/publication:Publication", name, args ?? new PublicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Publication(string name, Input<string> id, PublicationState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/publication:Publication", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Publication resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Publication Get(string name, Input<string> id, PublicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Publication(name, id, state, options);
        }
    }

    public sealed class PublicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should be ALL TABLES added to the publication. Defaults to 'false'
        /// </summary>
        [Input("allTables")]
        public Input<bool>? AllTables { get; set; }

        /// <summary>
        /// Which database to create the publication on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Should all subsequent resources of the publication be dropped. Defaults to 'false'
        /// </summary>
        [Input("dropCascade")]
        public Input<bool>? DropCascade { get; set; }

        /// <summary>
        /// The name of the publication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Who owns the publication. Defaults to provider user.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("publishParams")]
        private InputList<string>? _publishParams;

        /// <summary>
        /// Which 'publish' options should be turned on. Default to 'insert','update','delete'
        /// </summary>
        public InputList<string> PublishParams
        {
            get => _publishParams ?? (_publishParams = new InputList<string>());
            set => _publishParams = value;
        }

        /// <summary>
        /// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        /// </summary>
        [Input("publishViaPartitionRootParam")]
        public Input<bool>? PublishViaPartitionRootParam { get; set; }

        [Input("tables")]
        private InputList<string>? _tables;

        /// <summary>
        /// Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
        /// </summary>
        public InputList<string> Tables
        {
            get => _tables ?? (_tables = new InputList<string>());
            set => _tables = value;
        }

        public PublicationArgs()
        {
        }
    }

    public sealed class PublicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Should be ALL TABLES added to the publication. Defaults to 'false'
        /// </summary>
        [Input("allTables")]
        public Input<bool>? AllTables { get; set; }

        /// <summary>
        /// Which database to create the publication on. Defaults to provider database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Should all subsequent resources of the publication be dropped. Defaults to 'false'
        /// </summary>
        [Input("dropCascade")]
        public Input<bool>? DropCascade { get; set; }

        /// <summary>
        /// The name of the publication.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Who owns the publication. Defaults to provider user.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("publishParams")]
        private InputList<string>? _publishParams;

        /// <summary>
        /// Which 'publish' options should be turned on. Default to 'insert','update','delete'
        /// </summary>
        public InputList<string> PublishParams
        {
            get => _publishParams ?? (_publishParams = new InputList<string>());
            set => _publishParams = value;
        }

        /// <summary>
        /// Should be option 'publish_via_partition_root' be turned on. Default to 'false'
        /// </summary>
        [Input("publishViaPartitionRootParam")]
        public Input<bool>? PublishViaPartitionRootParam { get; set; }

        [Input("tables")]
        private InputList<string>? _tables;

        /// <summary>
        /// Which tables add to the publication. By defaults no tables added. Format of table is `&lt;schema_name&gt;.&lt;table_name&gt;`. If `&lt;schema_name&gt;` is not specified - default database schema will be used.
        /// </summary>
        public InputList<string> Tables
        {
            get => _tables ?? (_tables = new InputList<string>());
            set => _tables = value;
        }

        public PublicationState()
        {
        }
    }
}
