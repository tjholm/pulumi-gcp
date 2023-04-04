// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AddonsConfigArgs, AddonsConfigState } from "./addonsConfig";
export type AddonsConfig = import("./addonsConfig").AddonsConfig;
export const AddonsConfig: typeof import("./addonsConfig").AddonsConfig = null as any;
utilities.lazyLoad(exports, ["AddonsConfig"], () => require("./addonsConfig"));

export { EndpointAttachmentArgs, EndpointAttachmentState } from "./endpointAttachment";
export type EndpointAttachment = import("./endpointAttachment").EndpointAttachment;
export const EndpointAttachment: typeof import("./endpointAttachment").EndpointAttachment = null as any;
utilities.lazyLoad(exports, ["EndpointAttachment"], () => require("./endpointAttachment"));

export { EnvGroupArgs, EnvGroupState } from "./envGroup";
export type EnvGroup = import("./envGroup").EnvGroup;
export const EnvGroup: typeof import("./envGroup").EnvGroup = null as any;
utilities.lazyLoad(exports, ["EnvGroup"], () => require("./envGroup"));

export { EnvGroupAttachmentArgs, EnvGroupAttachmentState } from "./envGroupAttachment";
export type EnvGroupAttachment = import("./envGroupAttachment").EnvGroupAttachment;
export const EnvGroupAttachment: typeof import("./envGroupAttachment").EnvGroupAttachment = null as any;
utilities.lazyLoad(exports, ["EnvGroupAttachment"], () => require("./envGroupAttachment"));

export { EnvKeystoreArgs, EnvKeystoreState } from "./envKeystore";
export type EnvKeystore = import("./envKeystore").EnvKeystore;
export const EnvKeystore: typeof import("./envKeystore").EnvKeystore = null as any;
utilities.lazyLoad(exports, ["EnvKeystore"], () => require("./envKeystore"));

export { EnvReferencesArgs, EnvReferencesState } from "./envReferences";
export type EnvReferences = import("./envReferences").EnvReferences;
export const EnvReferences: typeof import("./envReferences").EnvReferences = null as any;
utilities.lazyLoad(exports, ["EnvReferences"], () => require("./envReferences"));

export { EnvironmentArgs, EnvironmentState } from "./environment";
export type Environment = import("./environment").Environment;
export const Environment: typeof import("./environment").Environment = null as any;
utilities.lazyLoad(exports, ["Environment"], () => require("./environment"));

export { EnvironmentIamBindingArgs, EnvironmentIamBindingState } from "./environmentIamBinding";
export type EnvironmentIamBinding = import("./environmentIamBinding").EnvironmentIamBinding;
export const EnvironmentIamBinding: typeof import("./environmentIamBinding").EnvironmentIamBinding = null as any;
utilities.lazyLoad(exports, ["EnvironmentIamBinding"], () => require("./environmentIamBinding"));

export { EnvironmentIamMemberArgs, EnvironmentIamMemberState } from "./environmentIamMember";
export type EnvironmentIamMember = import("./environmentIamMember").EnvironmentIamMember;
export const EnvironmentIamMember: typeof import("./environmentIamMember").EnvironmentIamMember = null as any;
utilities.lazyLoad(exports, ["EnvironmentIamMember"], () => require("./environmentIamMember"));

export { EnvironmentIamPolicyArgs, EnvironmentIamPolicyState } from "./environmentIamPolicy";
export type EnvironmentIamPolicy = import("./environmentIamPolicy").EnvironmentIamPolicy;
export const EnvironmentIamPolicy: typeof import("./environmentIamPolicy").EnvironmentIamPolicy = null as any;
utilities.lazyLoad(exports, ["EnvironmentIamPolicy"], () => require("./environmentIamPolicy"));

export { FlowhookArgs, FlowhookState } from "./flowhook";
export type Flowhook = import("./flowhook").Flowhook;
export const Flowhook: typeof import("./flowhook").Flowhook = null as any;
utilities.lazyLoad(exports, ["Flowhook"], () => require("./flowhook"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstanceAttachmentArgs, InstanceAttachmentState } from "./instanceAttachment";
export type InstanceAttachment = import("./instanceAttachment").InstanceAttachment;
export const InstanceAttachment: typeof import("./instanceAttachment").InstanceAttachment = null as any;
utilities.lazyLoad(exports, ["InstanceAttachment"], () => require("./instanceAttachment"));

export { KeystoresAliasesKeyCertFileArgs, KeystoresAliasesKeyCertFileState } from "./keystoresAliasesKeyCertFile";
export type KeystoresAliasesKeyCertFile = import("./keystoresAliasesKeyCertFile").KeystoresAliasesKeyCertFile;
export const KeystoresAliasesKeyCertFile: typeof import("./keystoresAliasesKeyCertFile").KeystoresAliasesKeyCertFile = null as any;
utilities.lazyLoad(exports, ["KeystoresAliasesKeyCertFile"], () => require("./keystoresAliasesKeyCertFile"));

export { NatAddressArgs, NatAddressState } from "./natAddress";
export type NatAddress = import("./natAddress").NatAddress;
export const NatAddress: typeof import("./natAddress").NatAddress = null as any;
utilities.lazyLoad(exports, ["NatAddress"], () => require("./natAddress"));

export { OrganizationArgs, OrganizationState } from "./organization";
export type Organization = import("./organization").Organization;
export const Organization: typeof import("./organization").Organization = null as any;
utilities.lazyLoad(exports, ["Organization"], () => require("./organization"));

export { SharedflowArgs, SharedflowState } from "./sharedflow";
export type Sharedflow = import("./sharedflow").Sharedflow;
export const Sharedflow: typeof import("./sharedflow").Sharedflow = null as any;
utilities.lazyLoad(exports, ["Sharedflow"], () => require("./sharedflow"));

export { SharedflowDeploymentArgs, SharedflowDeploymentState } from "./sharedflowDeployment";
export type SharedflowDeployment = import("./sharedflowDeployment").SharedflowDeployment;
export const SharedflowDeployment: typeof import("./sharedflowDeployment").SharedflowDeployment = null as any;
utilities.lazyLoad(exports, ["SharedflowDeployment"], () => require("./sharedflowDeployment"));

export { SyncAuthorizationArgs, SyncAuthorizationState } from "./syncAuthorization";
export type SyncAuthorization = import("./syncAuthorization").SyncAuthorization;
export const SyncAuthorization: typeof import("./syncAuthorization").SyncAuthorization = null as any;
utilities.lazyLoad(exports, ["SyncAuthorization"], () => require("./syncAuthorization"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:apigee/addonsConfig:AddonsConfig":
                return new AddonsConfig(name, <any>undefined, { urn })
            case "gcp:apigee/endpointAttachment:EndpointAttachment":
                return new EndpointAttachment(name, <any>undefined, { urn })
            case "gcp:apigee/envGroup:EnvGroup":
                return new EnvGroup(name, <any>undefined, { urn })
            case "gcp:apigee/envGroupAttachment:EnvGroupAttachment":
                return new EnvGroupAttachment(name, <any>undefined, { urn })
            case "gcp:apigee/envKeystore:EnvKeystore":
                return new EnvKeystore(name, <any>undefined, { urn })
            case "gcp:apigee/envReferences:EnvReferences":
                return new EnvReferences(name, <any>undefined, { urn })
            case "gcp:apigee/environment:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "gcp:apigee/environmentIamBinding:EnvironmentIamBinding":
                return new EnvironmentIamBinding(name, <any>undefined, { urn })
            case "gcp:apigee/environmentIamMember:EnvironmentIamMember":
                return new EnvironmentIamMember(name, <any>undefined, { urn })
            case "gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy":
                return new EnvironmentIamPolicy(name, <any>undefined, { urn })
            case "gcp:apigee/flowhook:Flowhook":
                return new Flowhook(name, <any>undefined, { urn })
            case "gcp:apigee/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "gcp:apigee/instanceAttachment:InstanceAttachment":
                return new InstanceAttachment(name, <any>undefined, { urn })
            case "gcp:apigee/keystoresAliasesKeyCertFile:KeystoresAliasesKeyCertFile":
                return new KeystoresAliasesKeyCertFile(name, <any>undefined, { urn })
            case "gcp:apigee/natAddress:NatAddress":
                return new NatAddress(name, <any>undefined, { urn })
            case "gcp:apigee/organization:Organization":
                return new Organization(name, <any>undefined, { urn })
            case "gcp:apigee/sharedflow:Sharedflow":
                return new Sharedflow(name, <any>undefined, { urn })
            case "gcp:apigee/sharedflowDeployment:SharedflowDeployment":
                return new SharedflowDeployment(name, <any>undefined, { urn })
            case "gcp:apigee/syncAuthorization:SyncAuthorization":
                return new SyncAuthorization(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "apigee/addonsConfig", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/endpointAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/envGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/envGroupAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/envKeystore", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/envReferences", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environmentIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environmentIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environmentIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/flowhook", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/instance", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/instanceAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/keystoresAliasesKeyCertFile", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/natAddress", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/organization", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/sharedflow", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/sharedflowDeployment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/syncAuthorization", _module)
