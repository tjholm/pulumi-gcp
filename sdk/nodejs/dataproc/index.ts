// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./autoscalingPolicy";
export * from "./autoscalingPolicyIamBinding";
export * from "./autoscalingPolicyIamMember";
export * from "./autoscalingPolicyIamPolicy";
export * from "./cluster";
export * from "./clusterIAMBinding";
export * from "./clusterIAMMember";
export * from "./clusterIAMPolicy";
export * from "./job";
export * from "./jobIAMBinding";
export * from "./jobIAMMember";
export * from "./jobIAMPolicy";
export * from "./metastoreService";
export * from "./metastoreServiceIamBinding";
export * from "./metastoreServiceIamMember";
export * from "./metastoreServiceIamPolicy";
export * from "./workflowTemplate";

// Import resources to register:
import { AutoscalingPolicy } from "./autoscalingPolicy";
import { AutoscalingPolicyIamBinding } from "./autoscalingPolicyIamBinding";
import { AutoscalingPolicyIamMember } from "./autoscalingPolicyIamMember";
import { AutoscalingPolicyIamPolicy } from "./autoscalingPolicyIamPolicy";
import { Cluster } from "./cluster";
import { ClusterIAMBinding } from "./clusterIAMBinding";
import { ClusterIAMMember } from "./clusterIAMMember";
import { ClusterIAMPolicy } from "./clusterIAMPolicy";
import { Job } from "./job";
import { JobIAMBinding } from "./jobIAMBinding";
import { JobIAMMember } from "./jobIAMMember";
import { JobIAMPolicy } from "./jobIAMPolicy";
import { MetastoreService } from "./metastoreService";
import { MetastoreServiceIamBinding } from "./metastoreServiceIamBinding";
import { MetastoreServiceIamMember } from "./metastoreServiceIamMember";
import { MetastoreServiceIamPolicy } from "./metastoreServiceIamPolicy";
import { WorkflowTemplate } from "./workflowTemplate";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gcp:dataproc/autoscalingPolicy:AutoscalingPolicy":
                return new AutoscalingPolicy(name, <any>undefined, { urn })
            case "gcp:dataproc/autoscalingPolicyIamBinding:AutoscalingPolicyIamBinding":
                return new AutoscalingPolicyIamBinding(name, <any>undefined, { urn })
            case "gcp:dataproc/autoscalingPolicyIamMember:AutoscalingPolicyIamMember":
                return new AutoscalingPolicyIamMember(name, <any>undefined, { urn })
            case "gcp:dataproc/autoscalingPolicyIamPolicy:AutoscalingPolicyIamPolicy":
                return new AutoscalingPolicyIamPolicy(name, <any>undefined, { urn })
            case "gcp:dataproc/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "gcp:dataproc/clusterIAMBinding:ClusterIAMBinding":
                return new ClusterIAMBinding(name, <any>undefined, { urn })
            case "gcp:dataproc/clusterIAMMember:ClusterIAMMember":
                return new ClusterIAMMember(name, <any>undefined, { urn })
            case "gcp:dataproc/clusterIAMPolicy:ClusterIAMPolicy":
                return new ClusterIAMPolicy(name, <any>undefined, { urn })
            case "gcp:dataproc/job:Job":
                return new Job(name, <any>undefined, { urn })
            case "gcp:dataproc/jobIAMBinding:JobIAMBinding":
                return new JobIAMBinding(name, <any>undefined, { urn })
            case "gcp:dataproc/jobIAMMember:JobIAMMember":
                return new JobIAMMember(name, <any>undefined, { urn })
            case "gcp:dataproc/jobIAMPolicy:JobIAMPolicy":
                return new JobIAMPolicy(name, <any>undefined, { urn })
            case "gcp:dataproc/metastoreService:MetastoreService":
                return new MetastoreService(name, <any>undefined, { urn })
            case "gcp:dataproc/metastoreServiceIamBinding:MetastoreServiceIamBinding":
                return new MetastoreServiceIamBinding(name, <any>undefined, { urn })
            case "gcp:dataproc/metastoreServiceIamMember:MetastoreServiceIamMember":
                return new MetastoreServiceIamMember(name, <any>undefined, { urn })
            case "gcp:dataproc/metastoreServiceIamPolicy:MetastoreServiceIamPolicy":
                return new MetastoreServiceIamPolicy(name, <any>undefined, { urn })
            case "gcp:dataproc/workflowTemplate:WorkflowTemplate":
                return new WorkflowTemplate(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gcp", "dataproc/autoscalingPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/autoscalingPolicyIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/autoscalingPolicyIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/autoscalingPolicyIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/cluster", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/clusterIAMBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/clusterIAMMember", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/clusterIAMPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/job", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/jobIAMBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/jobIAMMember", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/jobIAMPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/metastoreService", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/metastoreServiceIamBinding", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/metastoreServiceIamMember", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/metastoreServiceIamPolicy", _module)
pulumi.runtime.registerResourceModule("gcp", "dataproc/workflowTemplate", _module)
