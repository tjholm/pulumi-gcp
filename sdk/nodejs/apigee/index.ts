// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./endpointAttachment";
export * from "./envGroup";
export * from "./envGroupAttachment";
export * from "./environment";
export * from "./environmentIamBinding";
export * from "./environmentIamMember";
export * from "./environmentIamPolicy";
export * from "./instance";
export * from "./instanceAttachment";
export * from "./organization";

// Import resources to register:
import { EndpointAttachment } from "./endpointAttachment";
import { EnvGroup } from "./envGroup";
import { EnvGroupAttachment } from "./envGroupAttachment";
import { Environment } from "./environment";
import { EnvironmentIamBinding } from "./environmentIamBinding";
import { EnvironmentIamMember } from "./environmentIamMember";
import { EnvironmentIamPolicy } from "./environmentIamPolicy";
import { Instance } from "./instance";
import { InstanceAttachment } from "./instanceAttachment";
import { Organization } from "./organization";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:apigee/endpointAttachment:EndpointAttachment":
                return new EndpointAttachment(name, <any>undefined, { urn })
            case "gcp:apigee/envGroup:EnvGroup":
                return new EnvGroup(name, <any>undefined, { urn })
            case "gcp:apigee/envGroupAttachment:EnvGroupAttachment":
                return new EnvGroupAttachment(name, <any>undefined, { urn })
            case "gcp:apigee/environment:Environment":
                return new Environment(name, <any>undefined, { urn })
            case "gcp:apigee/environmentIamBinding:EnvironmentIamBinding":
                return new EnvironmentIamBinding(name, <any>undefined, { urn })
            case "gcp:apigee/environmentIamMember:EnvironmentIamMember":
                return new EnvironmentIamMember(name, <any>undefined, { urn })
            case "gcp:apigee/environmentIamPolicy:EnvironmentIamPolicy":
                return new EnvironmentIamPolicy(name, <any>undefined, { urn })
            case "gcp:apigee/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "gcp:apigee/instanceAttachment:InstanceAttachment":
                return new InstanceAttachment(name, <any>undefined, { urn })
            case "gcp:apigee/organization:Organization":
                return new Organization(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "apigee/endpointAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/envGroup", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/envGroupAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environmentIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environmentIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/environmentIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/instance", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/instanceAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apigee/organization", _module)
