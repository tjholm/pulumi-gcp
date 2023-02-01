// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccessBoundaryPolicyArgs, AccessBoundaryPolicyState } from "./accessBoundaryPolicy";
export type AccessBoundaryPolicy = import("./accessBoundaryPolicy").AccessBoundaryPolicy;
export const AccessBoundaryPolicy: typeof import("./accessBoundaryPolicy").AccessBoundaryPolicy = null as any;
utilities.lazyLoad(exports, ["AccessBoundaryPolicy"], () => require("./accessBoundaryPolicy"));

export { DenyPolicyArgs, DenyPolicyState } from "./denyPolicy";
export type DenyPolicy = import("./denyPolicy").DenyPolicy;
export const DenyPolicy: typeof import("./denyPolicy").DenyPolicy = null as any;
utilities.lazyLoad(exports, ["DenyPolicy"], () => require("./denyPolicy"));

export { GetRuleArgs, GetRuleResult, GetRuleOutputArgs } from "./getRule";
export const getRule: typeof import("./getRule").getRule = null as any;
export const getRuleOutput: typeof import("./getRule").getRuleOutput = null as any;
utilities.lazyLoad(exports, ["getRule","getRuleOutput"], () => require("./getRule"));

export { GetTestablePermissionsArgs, GetTestablePermissionsResult, GetTestablePermissionsOutputArgs } from "./getTestablePermissions";
export const getTestablePermissions: typeof import("./getTestablePermissions").getTestablePermissions = null as any;
export const getTestablePermissionsOutput: typeof import("./getTestablePermissions").getTestablePermissionsOutput = null as any;
utilities.lazyLoad(exports, ["getTestablePermissions","getTestablePermissionsOutput"], () => require("./getTestablePermissions"));

export { GetWorkloadIdentityPoolArgs, GetWorkloadIdentityPoolResult, GetWorkloadIdentityPoolOutputArgs } from "./getWorkloadIdentityPool";
export const getWorkloadIdentityPool: typeof import("./getWorkloadIdentityPool").getWorkloadIdentityPool = null as any;
export const getWorkloadIdentityPoolOutput: typeof import("./getWorkloadIdentityPool").getWorkloadIdentityPoolOutput = null as any;
utilities.lazyLoad(exports, ["getWorkloadIdentityPool","getWorkloadIdentityPoolOutput"], () => require("./getWorkloadIdentityPool"));

export { GetWorkloadIdentityPoolProviderArgs, GetWorkloadIdentityPoolProviderResult, GetWorkloadIdentityPoolProviderOutputArgs } from "./getWorkloadIdentityPoolProvider";
export const getWorkloadIdentityPoolProvider: typeof import("./getWorkloadIdentityPoolProvider").getWorkloadIdentityPoolProvider = null as any;
export const getWorkloadIdentityPoolProviderOutput: typeof import("./getWorkloadIdentityPoolProvider").getWorkloadIdentityPoolProviderOutput = null as any;
utilities.lazyLoad(exports, ["getWorkloadIdentityPoolProvider","getWorkloadIdentityPoolProviderOutput"], () => require("./getWorkloadIdentityPoolProvider"));

export { WorkforcePoolArgs, WorkforcePoolState } from "./workforcePool";
export type WorkforcePool = import("./workforcePool").WorkforcePool;
export const WorkforcePool: typeof import("./workforcePool").WorkforcePool = null as any;
utilities.lazyLoad(exports, ["WorkforcePool"], () => require("./workforcePool"));

export { WorkforcePoolProviderArgs, WorkforcePoolProviderState } from "./workforcePoolProvider";
export type WorkforcePoolProvider = import("./workforcePoolProvider").WorkforcePoolProvider;
export const WorkforcePoolProvider: typeof import("./workforcePoolProvider").WorkforcePoolProvider = null as any;
utilities.lazyLoad(exports, ["WorkforcePoolProvider"], () => require("./workforcePoolProvider"));

export { WorkloadIdentityPoolArgs, WorkloadIdentityPoolState } from "./workloadIdentityPool";
export type WorkloadIdentityPool = import("./workloadIdentityPool").WorkloadIdentityPool;
export const WorkloadIdentityPool: typeof import("./workloadIdentityPool").WorkloadIdentityPool = null as any;
utilities.lazyLoad(exports, ["WorkloadIdentityPool"], () => require("./workloadIdentityPool"));

export { WorkloadIdentityPoolProviderArgs, WorkloadIdentityPoolProviderState } from "./workloadIdentityPoolProvider";
export type WorkloadIdentityPoolProvider = import("./workloadIdentityPoolProvider").WorkloadIdentityPoolProvider;
export const WorkloadIdentityPoolProvider: typeof import("./workloadIdentityPoolProvider").WorkloadIdentityPoolProvider = null as any;
utilities.lazyLoad(exports, ["WorkloadIdentityPoolProvider"], () => require("./workloadIdentityPoolProvider"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:iam/accessBoundaryPolicy:AccessBoundaryPolicy":
                return new AccessBoundaryPolicy(name, <any>undefined, { urn })
            case "gcp:iam/denyPolicy:DenyPolicy":
                return new DenyPolicy(name, <any>undefined, { urn })
            case "gcp:iam/workforcePool:WorkforcePool":
                return new WorkforcePool(name, <any>undefined, { urn })
            case "gcp:iam/workforcePoolProvider:WorkforcePoolProvider":
                return new WorkforcePoolProvider(name, <any>undefined, { urn })
            case "gcp:iam/workloadIdentityPool:WorkloadIdentityPool":
                return new WorkloadIdentityPool(name, <any>undefined, { urn })
            case "gcp:iam/workloadIdentityPoolProvider:WorkloadIdentityPoolProvider":
                return new WorkloadIdentityPoolProvider(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "iam/accessBoundaryPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "iam/denyPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "iam/workforcePool", _module)
pulumi.runtime.registerResourceModule("gcp", "iam/workforcePoolProvider", _module)
pulumi.runtime.registerResourceModule("gcp", "iam/workloadIdentityPool", _module)
pulumi.runtime.registerResourceModule("gcp", "iam/workloadIdentityPoolProvider", _module)
