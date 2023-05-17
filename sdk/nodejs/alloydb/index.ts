// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BackupArgs, BackupState } from "./backup";
export type Backup = import("./backup").Backup;
export const Backup: typeof import("./backup").Backup = null as any;
utilities.lazyLoad(exports, ["Backup"], () => require("./backup"));

export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { GetLocationsArgs, GetLocationsResult, GetLocationsOutputArgs } from "./getLocations";
export const getLocations: typeof import("./getLocations").getLocations = null as any;
export const getLocationsOutput: typeof import("./getLocations").getLocationsOutput = null as any;
utilities.lazyLoad(exports, ["getLocations","getLocationsOutput"], () => require("./getLocations"));

export { GetSupportedDatabaseFlagsArgs, GetSupportedDatabaseFlagsResult, GetSupportedDatabaseFlagsOutputArgs } from "./getSupportedDatabaseFlags";
export const getSupportedDatabaseFlags: typeof import("./getSupportedDatabaseFlags").getSupportedDatabaseFlags = null as any;
export const getSupportedDatabaseFlagsOutput: typeof import("./getSupportedDatabaseFlags").getSupportedDatabaseFlagsOutput = null as any;
utilities.lazyLoad(exports, ["getSupportedDatabaseFlags","getSupportedDatabaseFlagsOutput"], () => require("./getSupportedDatabaseFlags"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:alloydb/backup:Backup":
                return new Backup(name, <any>undefined, { urn })
            case "gcp:alloydb/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "gcp:alloydb/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "alloydb/backup", _module)
pulumi.runtime.registerResourceModule("gcp", "alloydb/cluster", _module)
pulumi.runtime.registerResourceModule("gcp", "alloydb/instance", _module)
