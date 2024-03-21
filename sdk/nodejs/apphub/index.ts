// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ApplicationArgs, ApplicationState } from "./application";
export type Application = import("./application").Application;
export const Application: typeof import("./application").Application = null as any;
utilities.lazyLoad(exports, ["Application"], () => require("./application"));

export { GetDiscoveredServiceArgs, GetDiscoveredServiceResult, GetDiscoveredServiceOutputArgs } from "./getDiscoveredService";
export const getDiscoveredService: typeof import("./getDiscoveredService").getDiscoveredService = null as any;
export const getDiscoveredServiceOutput: typeof import("./getDiscoveredService").getDiscoveredServiceOutput = null as any;
utilities.lazyLoad(exports, ["getDiscoveredService","getDiscoveredServiceOutput"], () => require("./getDiscoveredService"));

export { GetDiscoveredWorkloadArgs, GetDiscoveredWorkloadResult, GetDiscoveredWorkloadOutputArgs } from "./getDiscoveredWorkload";
export const getDiscoveredWorkload: typeof import("./getDiscoveredWorkload").getDiscoveredWorkload = null as any;
export const getDiscoveredWorkloadOutput: typeof import("./getDiscoveredWorkload").getDiscoveredWorkloadOutput = null as any;
utilities.lazyLoad(exports, ["getDiscoveredWorkload","getDiscoveredWorkloadOutput"], () => require("./getDiscoveredWorkload"));

export { ServiceArgs, ServiceState } from "./service";
export type Service = import("./service").Service;
export const Service: typeof import("./service").Service = null as any;
utilities.lazyLoad(exports, ["Service"], () => require("./service"));

export { ServiceProjectAttachmentArgs, ServiceProjectAttachmentState } from "./serviceProjectAttachment";
export type ServiceProjectAttachment = import("./serviceProjectAttachment").ServiceProjectAttachment;
export const ServiceProjectAttachment: typeof import("./serviceProjectAttachment").ServiceProjectAttachment = null as any;
utilities.lazyLoad(exports, ["ServiceProjectAttachment"], () => require("./serviceProjectAttachment"));

export { WorkloadArgs, WorkloadState } from "./workload";
export type Workload = import("./workload").Workload;
export const Workload: typeof import("./workload").Workload = null as any;
utilities.lazyLoad(exports, ["Workload"], () => require("./workload"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:apphub/application:Application":
                return new Application(name, <any>undefined, { urn })
            case "gcp:apphub/service:Service":
                return new Service(name, <any>undefined, { urn })
            case "gcp:apphub/serviceProjectAttachment:ServiceProjectAttachment":
                return new ServiceProjectAttachment(name, <any>undefined, { urn })
            case "gcp:apphub/workload:Workload":
                return new Workload(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "apphub/application", _module)
pulumi.runtime.registerResourceModule("gcp", "apphub/service", _module)
pulumi.runtime.registerResourceModule("gcp", "apphub/serviceProjectAttachment", _module)
pulumi.runtime.registerResourceModule("gcp", "apphub/workload", _module)
