// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./database";
export * from "./defaultPrivileg";
export * from "./defaultPrivileges";
export * from "./extension";
export * from "./grant";
export * from "./grantRole";
export * from "./physicalReplicationSlot";
export * from "./provider";
export * from "./replicationSlot";
export * from "./role";
export * from "./schema";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Database } from "./database";
import { DefaultPrivileg } from "./defaultPrivileg";
import { DefaultPrivileges } from "./defaultPrivileges";
import { Extension } from "./extension";
import { Grant } from "./grant";
import { GrantRole } from "./grantRole";
import { PhysicalReplicationSlot } from "./physicalReplicationSlot";
import { ReplicationSlot } from "./replicationSlot";
import { Role } from "./role";
import { Schema } from "./schema";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "postgresql:index/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "postgresql:index/defaultPrivileg:DefaultPrivileg":
                return new DefaultPrivileg(name, <any>undefined, { urn })
            case "postgresql:index/defaultPrivileges:DefaultPrivileges":
                return new DefaultPrivileges(name, <any>undefined, { urn })
            case "postgresql:index/extension:Extension":
                return new Extension(name, <any>undefined, { urn })
            case "postgresql:index/grant:Grant":
                return new Grant(name, <any>undefined, { urn })
            case "postgresql:index/grantRole:GrantRole":
                return new GrantRole(name, <any>undefined, { urn })
            case "postgresql:index/physicalReplicationSlot:PhysicalReplicationSlot":
                return new PhysicalReplicationSlot(name, <any>undefined, { urn })
            case "postgresql:index/replicationSlot:ReplicationSlot":
                return new ReplicationSlot(name, <any>undefined, { urn })
            case "postgresql:index/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "postgresql:index/schema:Schema":
                return new Schema(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("postgresql", "index/database", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/defaultPrivileg", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/defaultPrivileges", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/extension", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/grant", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/grantRole", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/physicalReplicationSlot", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/replicationSlot", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/role", _module)
pulumi.runtime.registerResourceModule("postgresql", "index/schema", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("postgresql", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:postgresql") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
