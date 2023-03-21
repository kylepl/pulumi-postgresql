// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql.UserMapping`` resource creates and manages a user mapping on a PostgreSQL server.
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const extPostgresFdw = new postgresql.Extension("extPostgresFdw", {});
 * const myserverPostgres = new postgresql.Server("myserverPostgres", {
 *     serverName: "myserver_postgres",
 *     fdwName: "postgres_fdw",
 *     options: {
 *         host: "foo",
 *         dbname: "foodb",
 *         port: "5432",
 *     },
 * }, {
 *     dependsOn: [extPostgresFdw],
 * });
 * const remoteRole = new postgresql.Role("remoteRole", {});
 * const remoteUserMapping = new postgresql.UserMapping("remoteUserMapping", {
 *     serverName: myserverPostgres.serverName,
 *     userName: remoteRole.name,
 *     options: {
 *         user: "admin",
 *         password: "pass",
 *     },
 * });
 * ```
 */
export class UserMapping extends pulumi.CustomResource {
    /**
     * Get an existing UserMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserMappingState, opts?: pulumi.CustomResourceOptions): UserMapping {
        return new UserMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/userMapping:UserMapping';

    /**
     * Returns true if the given object is an instance of UserMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserMapping.__pulumiType;
    }

    /**
     * This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
     */
    public readonly options!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of an existing server for which the user mapping is to be created.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a UserMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserMappingArgs | UserMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserMappingState | undefined;
            resourceInputs["options"] = state ? state.options : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as UserMappingArgs | undefined;
            if ((!args || args.serverName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serverName'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["options"] = args ? args.options : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserMapping resources.
 */
export interface UserMappingState {
    /**
     * This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
     */
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of an existing server for which the user mapping is to be created.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     */
    serverName?: pulumi.Input<string>;
    /**
     * The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserMapping resource.
 */
export interface UserMappingArgs {
    /**
     * This clause specifies the options of the user mapping. The options typically define the actual user name and password of the mapping. Option names must be unique. The allowed option names and values are specific to the server's foreign-data wrapper.
     */
    options?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of an existing server for which the user mapping is to be created.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     */
    serverName: pulumi.Input<string>;
    /**
     * The name of an existing user that is mapped to foreign server. CURRENT_ROLE, CURRENT_USER, and USER match the name of the current user. When PUBLIC is specified, a so-called public mapping is created that is used when no user-specific mapping is applicable.
     * Changing this value
     * will force the creation of a new resource as this value can only be set
     * when the user mapping is created.
     */
    userName: pulumi.Input<string>;
}