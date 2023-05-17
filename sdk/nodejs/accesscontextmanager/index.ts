// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccessLevelArgs, AccessLevelState } from "./accessLevel";
export type AccessLevel = import("./accessLevel").AccessLevel;
export const AccessLevel: typeof import("./accessLevel").AccessLevel = null as any;
utilities.lazyLoad(exports, ["AccessLevel"], () => require("./accessLevel"));

export { AccessLevelConditionArgs, AccessLevelConditionState } from "./accessLevelCondition";
export type AccessLevelCondition = import("./accessLevelCondition").AccessLevelCondition;
export const AccessLevelCondition: typeof import("./accessLevelCondition").AccessLevelCondition = null as any;
utilities.lazyLoad(exports, ["AccessLevelCondition"], () => require("./accessLevelCondition"));

export { AccessLevelsArgs, AccessLevelsState } from "./accessLevels";
export type AccessLevels = import("./accessLevels").AccessLevels;
export const AccessLevels: typeof import("./accessLevels").AccessLevels = null as any;
utilities.lazyLoad(exports, ["AccessLevels"], () => require("./accessLevels"));

export { AccessPolicyArgs, AccessPolicyState } from "./accessPolicy";
export type AccessPolicy = import("./accessPolicy").AccessPolicy;
export const AccessPolicy: typeof import("./accessPolicy").AccessPolicy = null as any;
utilities.lazyLoad(exports, ["AccessPolicy"], () => require("./accessPolicy"));

export { AccessPolicyIamBindingArgs, AccessPolicyIamBindingState } from "./accessPolicyIamBinding";
export type AccessPolicyIamBinding = import("./accessPolicyIamBinding").AccessPolicyIamBinding;
export const AccessPolicyIamBinding: typeof import("./accessPolicyIamBinding").AccessPolicyIamBinding = null as any;
utilities.lazyLoad(exports, ["AccessPolicyIamBinding"], () => require("./accessPolicyIamBinding"));

export { AccessPolicyIamMemberArgs, AccessPolicyIamMemberState } from "./accessPolicyIamMember";
export type AccessPolicyIamMember = import("./accessPolicyIamMember").AccessPolicyIamMember;
export const AccessPolicyIamMember: typeof import("./accessPolicyIamMember").AccessPolicyIamMember = null as any;
utilities.lazyLoad(exports, ["AccessPolicyIamMember"], () => require("./accessPolicyIamMember"));

export { AccessPolicyIamPolicyArgs, AccessPolicyIamPolicyState } from "./accessPolicyIamPolicy";
export type AccessPolicyIamPolicy = import("./accessPolicyIamPolicy").AccessPolicyIamPolicy;
export const AccessPolicyIamPolicy: typeof import("./accessPolicyIamPolicy").AccessPolicyIamPolicy = null as any;
utilities.lazyLoad(exports, ["AccessPolicyIamPolicy"], () => require("./accessPolicyIamPolicy"));

export { AuthorizedOrgsDescArgs, AuthorizedOrgsDescState } from "./authorizedOrgsDesc";
export type AuthorizedOrgsDesc = import("./authorizedOrgsDesc").AuthorizedOrgsDesc;
export const AuthorizedOrgsDesc: typeof import("./authorizedOrgsDesc").AuthorizedOrgsDesc = null as any;
utilities.lazyLoad(exports, ["AuthorizedOrgsDesc"], () => require("./authorizedOrgsDesc"));

export { EgressPolicyArgs, EgressPolicyState } from "./egressPolicy";
export type EgressPolicy = import("./egressPolicy").EgressPolicy;
export const EgressPolicy: typeof import("./egressPolicy").EgressPolicy = null as any;
utilities.lazyLoad(exports, ["EgressPolicy"], () => require("./egressPolicy"));

export { GcpUserAccessBindingArgs, GcpUserAccessBindingState } from "./gcpUserAccessBinding";
export type GcpUserAccessBinding = import("./gcpUserAccessBinding").GcpUserAccessBinding;
export const GcpUserAccessBinding: typeof import("./gcpUserAccessBinding").GcpUserAccessBinding = null as any;
utilities.lazyLoad(exports, ["GcpUserAccessBinding"], () => require("./gcpUserAccessBinding"));

export { IngressPolicyArgs, IngressPolicyState } from "./ingressPolicy";
export type IngressPolicy = import("./ingressPolicy").IngressPolicy;
export const IngressPolicy: typeof import("./ingressPolicy").IngressPolicy = null as any;
utilities.lazyLoad(exports, ["IngressPolicy"], () => require("./ingressPolicy"));

export { ServicePerimeterArgs, ServicePerimeterState } from "./servicePerimeter";
export type ServicePerimeter = import("./servicePerimeter").ServicePerimeter;
export const ServicePerimeter: typeof import("./servicePerimeter").ServicePerimeter = null as any;
utilities.lazyLoad(exports, ["ServicePerimeter"], () => require("./servicePerimeter"));

export { ServicePerimeterResourceArgs, ServicePerimeterResourceState } from "./servicePerimeterResource";
export type ServicePerimeterResource = import("./servicePerimeterResource").ServicePerimeterResource;
export const ServicePerimeterResource: typeof import("./servicePerimeterResource").ServicePerimeterResource = null as any;
utilities.lazyLoad(exports, ["ServicePerimeterResource"], () => require("./servicePerimeterResource"));

export { ServicePerimetersArgs, ServicePerimetersState } from "./servicePerimeters";
export type ServicePerimeters = import("./servicePerimeters").ServicePerimeters;
export const ServicePerimeters: typeof import("./servicePerimeters").ServicePerimeters = null as any;
utilities.lazyLoad(exports, ["ServicePerimeters"], () => require("./servicePerimeters"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:accesscontextmanager/accessLevel:AccessLevel":
                return new AccessLevel(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/accessLevelCondition:AccessLevelCondition":
                return new AccessLevelCondition(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/accessLevels:AccessLevels":
                return new AccessLevels(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/accessPolicy:AccessPolicy":
                return new AccessPolicy(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/accessPolicyIamBinding:AccessPolicyIamBinding":
                return new AccessPolicyIamBinding(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/accessPolicyIamMember:AccessPolicyIamMember":
                return new AccessPolicyIamMember(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/accessPolicyIamPolicy:AccessPolicyIamPolicy":
                return new AccessPolicyIamPolicy(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/authorizedOrgsDesc:AuthorizedOrgsDesc":
                return new AuthorizedOrgsDesc(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/egressPolicy:EgressPolicy":
                return new EgressPolicy(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/gcpUserAccessBinding:GcpUserAccessBinding":
                return new GcpUserAccessBinding(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/ingressPolicy:IngressPolicy":
                return new IngressPolicy(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/servicePerimeter:ServicePerimeter":
                return new ServicePerimeter(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/servicePerimeterResource:ServicePerimeterResource":
                return new ServicePerimeterResource(name, <any>undefined, { urn })
            case "gcp:accesscontextmanager/servicePerimeters:ServicePerimeters":
                return new ServicePerimeters(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessLevel", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessLevelCondition", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessLevels", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessPolicyIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessPolicyIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/accessPolicyIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/authorizedOrgsDesc", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/egressPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/gcpUserAccessBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/ingressPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/servicePerimeter", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/servicePerimeterResource", _module)
pulumi.runtime.registerResourceModule("gcp", "accesscontextmanager/servicePerimeters", _module)
