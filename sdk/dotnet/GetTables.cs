// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.PostgreSql
{
    public static class GetTables
    {
        /// <summary>
        /// The ``postgresql.getTables`` data source retrieves a list of table names from a specified PostgreSQL database.
        /// </summary>
        public static Task<GetTablesResult> InvokeAsync(GetTablesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTablesResult>("postgresql:index/getTables:getTables", args ?? new GetTablesArgs(), options.WithDefaults());

        /// <summary>
        /// The ``postgresql.getTables`` data source retrieves a list of table names from a specified PostgreSQL database.
        /// </summary>
        public static Output<GetTablesResult> Invoke(GetTablesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTablesResult>("postgresql:index/getTables:getTables", args ?? new GetTablesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTablesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PostgreSQL database which will be queried for table names.
        /// </summary>
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        [Input("likeAllPatterns")]
        private List<string>? _likeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
        /// </summary>
        public List<string> LikeAllPatterns
        {
            get => _likeAllPatterns ?? (_likeAllPatterns = new List<string>());
            set => _likeAllPatterns = value;
        }

        [Input("likeAnyPatterns")]
        private List<string>? _likeAnyPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
        /// </summary>
        public List<string> LikeAnyPatterns
        {
            get => _likeAnyPatterns ?? (_likeAnyPatterns = new List<string>());
            set => _likeAnyPatterns = value;
        }

        [Input("notLikeAllPatterns")]
        private List<string>? _notLikeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
        /// </summary>
        public List<string> NotLikeAllPatterns
        {
            get => _notLikeAllPatterns ?? (_notLikeAllPatterns = new List<string>());
            set => _notLikeAllPatterns = value;
        }

        /// <summary>
        /// Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
        /// 
        /// Note that all optional arguments can be used in conjunction.
        /// </summary>
        [Input("regexPattern")]
        public string? RegexPattern { get; set; }

        [Input("schemas")]
        private List<string>? _schemas;

        /// <summary>
        /// List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
        /// </summary>
        public List<string> Schemas
        {
            get => _schemas ?? (_schemas = new List<string>());
            set => _schemas = value;
        }

        [Input("tableTypes")]
        private List<string>? _tableTypes;

        /// <summary>
        /// List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use 'BASE TABLE' for normal tables only.
        /// </summary>
        public List<string> TableTypes
        {
            get => _tableTypes ?? (_tableTypes = new List<string>());
            set => _tableTypes = value;
        }

        public GetTablesArgs()
        {
        }
        public static new GetTablesArgs Empty => new GetTablesArgs();
    }

    public sealed class GetTablesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The PostgreSQL database which will be queried for table names.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("likeAllPatterns")]
        private InputList<string>? _likeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ALL`` operators.
        /// </summary>
        public InputList<string> LikeAllPatterns
        {
            get => _likeAllPatterns ?? (_likeAllPatterns = new InputList<string>());
            set => _likeAllPatterns = value;
        }

        [Input("likeAnyPatterns")]
        private InputList<string>? _likeAnyPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``LIKE ANY`` operators.
        /// </summary>
        public InputList<string> LikeAnyPatterns
        {
            get => _likeAnyPatterns ?? (_likeAnyPatterns = new InputList<string>());
            set => _likeAnyPatterns = value;
        }

        [Input("notLikeAllPatterns")]
        private InputList<string>? _notLikeAllPatterns;

        /// <summary>
        /// List of expressions which will be pattern matched against table names in the query using the PostgreSQL ``NOT LIKE ALL`` operators.
        /// </summary>
        public InputList<string> NotLikeAllPatterns
        {
            get => _notLikeAllPatterns ?? (_notLikeAllPatterns = new InputList<string>());
            set => _notLikeAllPatterns = value;
        }

        /// <summary>
        /// Expression which will be pattern matched against table names in the query using the PostgreSQL ``~`` (regular expression match) operator.
        /// 
        /// Note that all optional arguments can be used in conjunction.
        /// </summary>
        [Input("regexPattern")]
        public Input<string>? RegexPattern { get; set; }

        [Input("schemas")]
        private InputList<string>? _schemas;

        /// <summary>
        /// List of PostgreSQL schema(s) which will be queried for table names. Queries all schemas in the database by default.
        /// </summary>
        public InputList<string> Schemas
        {
            get => _schemas ?? (_schemas = new InputList<string>());
            set => _schemas = value;
        }

        [Input("tableTypes")]
        private InputList<string>? _tableTypes;

        /// <summary>
        /// List of PostgreSQL table types which will be queried for table names. Includes all table types by default (including views and temp tables). Use 'BASE TABLE' for normal tables only.
        /// </summary>
        public InputList<string> TableTypes
        {
            get => _tableTypes ?? (_tableTypes = new InputList<string>());
            set => _tableTypes = value;
        }

        public GetTablesInvokeArgs()
        {
        }
        public static new GetTablesInvokeArgs Empty => new GetTablesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTablesResult
    {
        public readonly string Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> LikeAllPatterns;
        public readonly ImmutableArray<string> LikeAnyPatterns;
        public readonly ImmutableArray<string> NotLikeAllPatterns;
        public readonly string? RegexPattern;
        public readonly ImmutableArray<string> Schemas;
        public readonly ImmutableArray<string> TableTypes;
        /// <summary>
        /// A list of PostgreSQL tables retrieved by this data source. Each table consists of the fields documented below.
        /// ___
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTablesTableResult> Tables;

        [OutputConstructor]
        private GetTablesResult(
            string database,

            string id,

            ImmutableArray<string> likeAllPatterns,

            ImmutableArray<string> likeAnyPatterns,

            ImmutableArray<string> notLikeAllPatterns,

            string? regexPattern,

            ImmutableArray<string> schemas,

            ImmutableArray<string> tableTypes,

            ImmutableArray<Outputs.GetTablesTableResult> tables)
        {
            Database = database;
            Id = id;
            LikeAllPatterns = likeAllPatterns;
            LikeAnyPatterns = likeAnyPatterns;
            NotLikeAllPatterns = notLikeAllPatterns;
            RegexPattern = regexPattern;
            Schemas = schemas;
            TableTypes = tableTypes;
            Tables = tables;
        }
    }
}
