// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AlertPolicyArgs, AlertPolicyState } from "./alertPolicy";
export type AlertPolicy = import("./alertPolicy").AlertPolicy;
export const AlertPolicy: typeof import("./alertPolicy").AlertPolicy = null as any;
utilities.lazyLoad(exports, ["AlertPolicy"], () => require("./alertPolicy"));

export { CustomServiceArgs, CustomServiceState } from "./customService";
export type CustomService = import("./customService").CustomService;
export const CustomService: typeof import("./customService").CustomService = null as any;
utilities.lazyLoad(exports, ["CustomService"], () => require("./customService"));

export { DashboardArgs, DashboardState } from "./dashboard";
export type Dashboard = import("./dashboard").Dashboard;
export const Dashboard: typeof import("./dashboard").Dashboard = null as any;
utilities.lazyLoad(exports, ["Dashboard"], () => require("./dashboard"));

export { GenericServiceArgs, GenericServiceState } from "./genericService";
export type GenericService = import("./genericService").GenericService;
export const GenericService: typeof import("./genericService").GenericService = null as any;
utilities.lazyLoad(exports, ["GenericService"], () => require("./genericService"));

export { GetAppEngineServiceArgs, GetAppEngineServiceResult, GetAppEngineServiceOutputArgs } from "./getAppEngineService";
export const getAppEngineService: typeof import("./getAppEngineService").getAppEngineService = null as any;
export const getAppEngineServiceOutput: typeof import("./getAppEngineService").getAppEngineServiceOutput = null as any;
utilities.lazyLoad(exports, ["getAppEngineService","getAppEngineServiceOutput"], () => require("./getAppEngineService"));

export { GetClusterIstioServiceArgs, GetClusterIstioServiceResult, GetClusterIstioServiceOutputArgs } from "./getClusterIstioService";
export const getClusterIstioService: typeof import("./getClusterIstioService").getClusterIstioService = null as any;
export const getClusterIstioServiceOutput: typeof import("./getClusterIstioService").getClusterIstioServiceOutput = null as any;
utilities.lazyLoad(exports, ["getClusterIstioService","getClusterIstioServiceOutput"], () => require("./getClusterIstioService"));

export { GetIstioCanonicalServiceArgs, GetIstioCanonicalServiceResult, GetIstioCanonicalServiceOutputArgs } from "./getIstioCanonicalService";
export const getIstioCanonicalService: typeof import("./getIstioCanonicalService").getIstioCanonicalService = null as any;
export const getIstioCanonicalServiceOutput: typeof import("./getIstioCanonicalService").getIstioCanonicalServiceOutput = null as any;
utilities.lazyLoad(exports, ["getIstioCanonicalService","getIstioCanonicalServiceOutput"], () => require("./getIstioCanonicalService"));

export { GetMeshIstioServiceArgs, GetMeshIstioServiceResult, GetMeshIstioServiceOutputArgs } from "./getMeshIstioService";
export const getMeshIstioService: typeof import("./getMeshIstioService").getMeshIstioService = null as any;
export const getMeshIstioServiceOutput: typeof import("./getMeshIstioService").getMeshIstioServiceOutput = null as any;
utilities.lazyLoad(exports, ["getMeshIstioService","getMeshIstioServiceOutput"], () => require("./getMeshIstioService"));

export { GetNotificationChannelArgs, GetNotificationChannelResult, GetNotificationChannelOutputArgs } from "./getNotificationChannel";
export const getNotificationChannel: typeof import("./getNotificationChannel").getNotificationChannel = null as any;
export const getNotificationChannelOutput: typeof import("./getNotificationChannel").getNotificationChannelOutput = null as any;
utilities.lazyLoad(exports, ["getNotificationChannel","getNotificationChannelOutput"], () => require("./getNotificationChannel"));

export { GetSecretVersionArgs, GetSecretVersionResult, GetSecretVersionOutputArgs } from "./getSecretVersion";
export const getSecretVersion: typeof import("./getSecretVersion").getSecretVersion = null as any;
export const getSecretVersionOutput: typeof import("./getSecretVersion").getSecretVersionOutput = null as any;
utilities.lazyLoad(exports, ["getSecretVersion","getSecretVersionOutput"], () => require("./getSecretVersion"));

export { GetUptimeCheckIPsResult } from "./getUptimeCheckIPs";
export const getUptimeCheckIPs: typeof import("./getUptimeCheckIPs").getUptimeCheckIPs = null as any;
utilities.lazyLoad(exports, ["getUptimeCheckIPs"], () => require("./getUptimeCheckIPs"));

export { GroupArgs, GroupState } from "./group";
export type Group = import("./group").Group;
export const Group: typeof import("./group").Group = null as any;
utilities.lazyLoad(exports, ["Group"], () => require("./group"));

export { MetricDescriptorArgs, MetricDescriptorState } from "./metricDescriptor";
export type MetricDescriptor = import("./metricDescriptor").MetricDescriptor;
export const MetricDescriptor: typeof import("./metricDescriptor").MetricDescriptor = null as any;
utilities.lazyLoad(exports, ["MetricDescriptor"], () => require("./metricDescriptor"));

export { MonitoredProjectArgs, MonitoredProjectState } from "./monitoredProject";
export type MonitoredProject = import("./monitoredProject").MonitoredProject;
export const MonitoredProject: typeof import("./monitoredProject").MonitoredProject = null as any;
utilities.lazyLoad(exports, ["MonitoredProject"], () => require("./monitoredProject"));

export { NotificationChannelArgs, NotificationChannelState } from "./notificationChannel";
export type NotificationChannel = import("./notificationChannel").NotificationChannel;
export const NotificationChannel: typeof import("./notificationChannel").NotificationChannel = null as any;
utilities.lazyLoad(exports, ["NotificationChannel"], () => require("./notificationChannel"));

export { SloArgs, SloState } from "./slo";
export type Slo = import("./slo").Slo;
export const Slo: typeof import("./slo").Slo = null as any;
utilities.lazyLoad(exports, ["Slo"], () => require("./slo"));

export { UptimeCheckConfigArgs, UptimeCheckConfigState } from "./uptimeCheckConfig";
export type UptimeCheckConfig = import("./uptimeCheckConfig").UptimeCheckConfig;
export const UptimeCheckConfig: typeof import("./uptimeCheckConfig").UptimeCheckConfig = null as any;
utilities.lazyLoad(exports, ["UptimeCheckConfig"], () => require("./uptimeCheckConfig"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:monitoring/alertPolicy:AlertPolicy":
                return new AlertPolicy(name, <any>undefined, { urn })
            case "gcp:monitoring/customService:CustomService":
                return new CustomService(name, <any>undefined, { urn })
            case "gcp:monitoring/dashboard:Dashboard":
                return new Dashboard(name, <any>undefined, { urn })
            case "gcp:monitoring/genericService:GenericService":
                return new GenericService(name, <any>undefined, { urn })
            case "gcp:monitoring/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "gcp:monitoring/metricDescriptor:MetricDescriptor":
                return new MetricDescriptor(name, <any>undefined, { urn })
            case "gcp:monitoring/monitoredProject:MonitoredProject":
                return new MonitoredProject(name, <any>undefined, { urn })
            case "gcp:monitoring/notificationChannel:NotificationChannel":
                return new NotificationChannel(name, <any>undefined, { urn })
            case "gcp:monitoring/slo:Slo":
                return new Slo(name, <any>undefined, { urn })
            case "gcp:monitoring/uptimeCheckConfig:UptimeCheckConfig":
                return new UptimeCheckConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "monitoring/alertPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/customService", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/dashboard", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/genericService", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/group", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/metricDescriptor", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/monitoredProject", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/notificationChannel", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/slo", _module)
pulumi.runtime.registerResourceModule("gcp", "monitoring/uptimeCheckConfig", _module)
