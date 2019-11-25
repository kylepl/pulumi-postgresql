// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Postgresql
{
    /// <summary>
    /// The ``postgresql..Schema`` resource creates and manages [schema
    /// objects](https://www.postgresql.org/docs/current/static/ddl-schemas.html) within
    /// a PostgreSQL database.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/schema.html.markdown.
    /// </summary>
    public partial class Schema : Pulumi.CustomResource
    {
        /// <summary>
        /// When true, use the existing schema if it exists. (Default: true)
        /// </summary>
        [Output("ifNotExists")]
        public Output<bool?> IfNotExists { get; private set; } = null!;

        /// <summary>
        /// The name of the schema. Must be unique in the PostgreSQL
        /// database instance where it is configured.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ROLE who owns the schema.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Can be specified multiple times for each policy.  Each
        /// policy block supports fields documented below.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<Outputs.SchemaPolicies>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a Schema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schema(string name, SchemaArgs? args = null, CustomResourceOptions? options = null)
            : base("postgresql:index/schema:Schema", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Schema(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
            : base("postgresql:index/schema:Schema", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Schema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schema Get(string name, Input<string> id, SchemaState? state = null, CustomResourceOptions? options = null)
        {
            return new Schema(name, id, state, options);
        }
    }

    public sealed class SchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, use the existing schema if it exists. (Default: true)
        /// </summary>
        [Input("ifNotExists")]
        public Input<bool>? IfNotExists { get; set; }

        /// <summary>
        /// The name of the schema. Must be unique in the PostgreSQL
        /// database instance where it is configured.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ROLE who owns the schema.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("policies")]
        private InputList<Inputs.SchemaPoliciesArgs>? _policies;

        /// <summary>
        /// Can be specified multiple times for each policy.  Each
        /// policy block supports fields documented below.
        /// </summary>
        public InputList<Inputs.SchemaPoliciesArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.SchemaPoliciesArgs>());
            set => _policies = value;
        }

        public SchemaArgs()
        {
        }
    }

    public sealed class SchemaState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, use the existing schema if it exists. (Default: true)
        /// </summary>
        [Input("ifNotExists")]
        public Input<bool>? IfNotExists { get; set; }

        /// <summary>
        /// The name of the schema. Must be unique in the PostgreSQL
        /// database instance where it is configured.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ROLE who owns the schema.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        [Input("policies")]
        private InputList<Inputs.SchemaPoliciesGetArgs>? _policies;

        /// <summary>
        /// Can be specified multiple times for each policy.  Each
        /// policy block supports fields documented below.
        /// </summary>
        public InputList<Inputs.SchemaPoliciesGetArgs> Policies
        {
            get => _policies ?? (_policies = new InputList<Inputs.SchemaPoliciesGetArgs>());
            set => _policies = value;
        }

        public SchemaState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class SchemaPoliciesArgs : Pulumi.ResourceArgs
    {
        [Input("create")]
        public Input<bool>? Create { get; set; }

        [Input("createWithGrant")]
        public Input<bool>? CreateWithGrant { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("usage")]
        public Input<bool>? Usage { get; set; }

        [Input("usageWithGrant")]
        public Input<bool>? UsageWithGrant { get; set; }

        public SchemaPoliciesArgs()
        {
        }
    }

    public sealed class SchemaPoliciesGetArgs : Pulumi.ResourceArgs
    {
        [Input("create")]
        public Input<bool>? Create { get; set; }

        [Input("createWithGrant")]
        public Input<bool>? CreateWithGrant { get; set; }

        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("usage")]
        public Input<bool>? Usage { get; set; }

        [Input("usageWithGrant")]
        public Input<bool>? UsageWithGrant { get; set; }

        public SchemaPoliciesGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class SchemaPolicies
    {
        public readonly bool? Create;
        public readonly bool? CreateWithGrant;
        public readonly string? Role;
        public readonly bool? Usage;
        public readonly bool? UsageWithGrant;

        [OutputConstructor]
        private SchemaPolicies(
            bool? create,
            bool? createWithGrant,
            string? role,
            bool? usage,
            bool? usageWithGrant)
        {
            Create = create;
            CreateWithGrant = createWithGrant;
            Role = role;
            Usage = usage;
            UsageWithGrant = usageWithGrant;
        }
    }
    }
}
