// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AndroidAppArgs, AndroidAppState } from "./androidApp";
export type AndroidApp = import("./androidApp").AndroidApp;
export const AndroidApp: typeof import("./androidApp").AndroidApp = null as any;
utilities.lazyLoad(exports, ["AndroidApp"], () => require("./androidApp"));

export { GetWebAppArgs, GetWebAppResult, GetWebAppOutputArgs } from "./getWebApp";
export const getWebApp: typeof import("./getWebApp").getWebApp = null as any;
export const getWebAppOutput: typeof import("./getWebApp").getWebAppOutput = null as any;
utilities.lazyLoad(exports, ["getWebApp","getWebAppOutput"], () => require("./getWebApp"));

export { GetWebAppConfigArgs, GetWebAppConfigResult, GetWebAppConfigOutputArgs } from "./getWebAppConfig";
export const getWebAppConfig: typeof import("./getWebAppConfig").getWebAppConfig = null as any;
export const getWebAppConfigOutput: typeof import("./getWebAppConfig").getWebAppConfigOutput = null as any;
utilities.lazyLoad(exports, ["getWebAppConfig","getWebAppConfigOutput"], () => require("./getWebAppConfig"));

export { ProjectArgs, ProjectState } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { ProjectLocationArgs, ProjectLocationState } from "./projectLocation";
export type ProjectLocation = import("./projectLocation").ProjectLocation;
export const ProjectLocation: typeof import("./projectLocation").ProjectLocation = null as any;
utilities.lazyLoad(exports, ["ProjectLocation"], () => require("./projectLocation"));

export { WebAppArgs, WebAppState } from "./webApp";
export type WebApp = import("./webApp").WebApp;
export const WebApp: typeof import("./webApp").WebApp = null as any;
utilities.lazyLoad(exports, ["WebApp"], () => require("./webApp"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:firebase/androidApp:AndroidApp":
                return new AndroidApp(name, <any>undefined, { urn })
            case "gcp:firebase/project:Project":
                return new Project(name, <any>undefined, { urn })
            case "gcp:firebase/projectLocation:ProjectLocation":
                return new ProjectLocation(name, <any>undefined, { urn })
            case "gcp:firebase/webApp:WebApp":
                return new WebApp(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "firebase/androidApp", _module)
pulumi.runtime.registerResourceModule("gcp", "firebase/project", _module)
pulumi.runtime.registerResourceModule("gcp", "firebase/projectLocation", _module)
pulumi.runtime.registerResourceModule("gcp", "firebase/webApp", _module)
