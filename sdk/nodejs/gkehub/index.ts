// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { FeatureArgs, FeatureState } from "./feature";
export type Feature = import("./feature").Feature;
export const Feature: typeof import("./feature").Feature = null as any;
utilities.lazyLoad(exports, ["Feature"], () => require("./feature"));

export { FeatureIamBindingArgs, FeatureIamBindingState } from "./featureIamBinding";
export type FeatureIamBinding = import("./featureIamBinding").FeatureIamBinding;
export const FeatureIamBinding: typeof import("./featureIamBinding").FeatureIamBinding = null as any;
utilities.lazyLoad(exports, ["FeatureIamBinding"], () => require("./featureIamBinding"));

export { FeatureIamMemberArgs, FeatureIamMemberState } from "./featureIamMember";
export type FeatureIamMember = import("./featureIamMember").FeatureIamMember;
export const FeatureIamMember: typeof import("./featureIamMember").FeatureIamMember = null as any;
utilities.lazyLoad(exports, ["FeatureIamMember"], () => require("./featureIamMember"));

export { FeatureIamPolicyArgs, FeatureIamPolicyState } from "./featureIamPolicy";
export type FeatureIamPolicy = import("./featureIamPolicy").FeatureIamPolicy;
export const FeatureIamPolicy: typeof import("./featureIamPolicy").FeatureIamPolicy = null as any;
utilities.lazyLoad(exports, ["FeatureIamPolicy"], () => require("./featureIamPolicy"));

export { FeatureMembershipArgs, FeatureMembershipState } from "./featureMembership";
export type FeatureMembership = import("./featureMembership").FeatureMembership;
export const FeatureMembership: typeof import("./featureMembership").FeatureMembership = null as any;
utilities.lazyLoad(exports, ["FeatureMembership"], () => require("./featureMembership"));

export { GetFeatureIamPolicyArgs, GetFeatureIamPolicyResult, GetFeatureIamPolicyOutputArgs } from "./getFeatureIamPolicy";
export const getFeatureIamPolicy: typeof import("./getFeatureIamPolicy").getFeatureIamPolicy = null as any;
export const getFeatureIamPolicyOutput: typeof import("./getFeatureIamPolicy").getFeatureIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getFeatureIamPolicy","getFeatureIamPolicyOutput"], () => require("./getFeatureIamPolicy"));

export { GetMembershipIamPolicyArgs, GetMembershipIamPolicyResult, GetMembershipIamPolicyOutputArgs } from "./getMembershipIamPolicy";
export const getMembershipIamPolicy: typeof import("./getMembershipIamPolicy").getMembershipIamPolicy = null as any;
export const getMembershipIamPolicyOutput: typeof import("./getMembershipIamPolicy").getMembershipIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getMembershipIamPolicy","getMembershipIamPolicyOutput"], () => require("./getMembershipIamPolicy"));

export { MembershipArgs, MembershipState } from "./membership";
export type Membership = import("./membership").Membership;
export const Membership: typeof import("./membership").Membership = null as any;
utilities.lazyLoad(exports, ["Membership"], () => require("./membership"));

export { MembershipIamBindingArgs, MembershipIamBindingState } from "./membershipIamBinding";
export type MembershipIamBinding = import("./membershipIamBinding").MembershipIamBinding;
export const MembershipIamBinding: typeof import("./membershipIamBinding").MembershipIamBinding = null as any;
utilities.lazyLoad(exports, ["MembershipIamBinding"], () => require("./membershipIamBinding"));

export { MembershipIamMemberArgs, MembershipIamMemberState } from "./membershipIamMember";
export type MembershipIamMember = import("./membershipIamMember").MembershipIamMember;
export const MembershipIamMember: typeof import("./membershipIamMember").MembershipIamMember = null as any;
utilities.lazyLoad(exports, ["MembershipIamMember"], () => require("./membershipIamMember"));

export { MembershipIamPolicyArgs, MembershipIamPolicyState } from "./membershipIamPolicy";
export type MembershipIamPolicy = import("./membershipIamPolicy").MembershipIamPolicy;
export const MembershipIamPolicy: typeof import("./membershipIamPolicy").MembershipIamPolicy = null as any;
utilities.lazyLoad(exports, ["MembershipIamPolicy"], () => require("./membershipIamPolicy"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:gkehub/feature:Feature":
                return new Feature(name, <any>undefined, { urn })
            case "gcp:gkehub/featureIamBinding:FeatureIamBinding":
                return new FeatureIamBinding(name, <any>undefined, { urn })
            case "gcp:gkehub/featureIamMember:FeatureIamMember":
                return new FeatureIamMember(name, <any>undefined, { urn })
            case "gcp:gkehub/featureIamPolicy:FeatureIamPolicy":
                return new FeatureIamPolicy(name, <any>undefined, { urn })
            case "gcp:gkehub/featureMembership:FeatureMembership":
                return new FeatureMembership(name, <any>undefined, { urn })
            case "gcp:gkehub/membership:Membership":
                return new Membership(name, <any>undefined, { urn })
            case "gcp:gkehub/membershipIamBinding:MembershipIamBinding":
                return new MembershipIamBinding(name, <any>undefined, { urn })
            case "gcp:gkehub/membershipIamMember:MembershipIamMember":
                return new MembershipIamMember(name, <any>undefined, { urn })
            case "gcp:gkehub/membershipIamPolicy:MembershipIamPolicy":
                return new MembershipIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "gkehub/feature", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/featureIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/featureIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/featureIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/featureMembership", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/membership", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/membershipIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/membershipIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "gkehub/membershipIamPolicy", _module)
