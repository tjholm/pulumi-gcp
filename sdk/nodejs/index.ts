// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./provider";
export * from "./utils";

// Export sub-modules:
import * as accessapproval from "./accessapproval";
import * as accesscontextmanager from "./accesscontextmanager";
import * as activedirectory from "./activedirectory";
import * as apigateway from "./apigateway";
import * as apigee from "./apigee";
import * as appengine from "./appengine";
import * as artifactregistry from "./artifactregistry";
import * as assuredworkloads from "./assuredworkloads";
import * as bigquery from "./bigquery";
import * as bigtable from "./bigtable";
import * as billing from "./billing";
import * as binaryauthorization from "./binaryauthorization";
import * as certificateauthority from "./certificateauthority";
import * as cloudasset from "./cloudasset";
import * as cloudbuild from "./cloudbuild";
import * as cloudfunctions from "./cloudfunctions";
import * as cloudfunctionsv2 from "./cloudfunctionsv2";
import * as cloudidentity from "./cloudidentity";
import * as cloudrun from "./cloudrun";
import * as cloudscheduler from "./cloudscheduler";
import * as cloudtasks from "./cloudtasks";
import * as composer from "./composer";
import * as compute from "./compute";
import * as config from "./config";
import * as container from "./container";
import * as containeranalysis from "./containeranalysis";
import * as datacatalog from "./datacatalog";
import * as dataflow from "./dataflow";
import * as datafusion from "./datafusion";
import * as dataloss from "./dataloss";
import * as dataproc from "./dataproc";
import * as datastore from "./datastore";
import * as deploymentmanager from "./deploymentmanager";
import * as diagflow from "./diagflow";
import * as dns from "./dns";
import * as endpoints from "./endpoints";
import * as essentialcontacts from "./essentialcontacts";
import * as eventarc from "./eventarc";
import * as filestore from "./filestore";
import * as firebase from "./firebase";
import * as firebaserules from "./firebaserules";
import * as firestore from "./firestore";
import * as folder from "./folder";
import * as gameservices from "./gameservices";
import * as gkehub from "./gkehub";
import * as healthcare from "./healthcare";
import * as iam from "./iam";
import * as iap from "./iap";
import * as identityplatform from "./identityplatform";
import * as iot from "./iot";
import * as kms from "./kms";
import * as logging from "./logging";
import * as memcache from "./memcache";
import * as ml from "./ml";
import * as monitoring from "./monitoring";
import * as networkconnectivity from "./networkconnectivity";
import * as networkmanagement from "./networkmanagement";
import * as networkservices from "./networkservices";
import * as notebooks from "./notebooks";
import * as organizations from "./organizations";
import * as orgpolicy from "./orgpolicy";
import * as osconfig from "./osconfig";
import * as oslogin from "./oslogin";
import * as projects from "./projects";
import * as pubsub from "./pubsub";
import * as recaptcha from "./recaptcha";
import * as redis from "./redis";
import * as resourcemanager from "./resourcemanager";
import * as runtimeconfig from "./runtimeconfig";
import * as secretmanager from "./secretmanager";
import * as securitycenter from "./securitycenter";
import * as serviceaccount from "./serviceaccount";
import * as servicedirectory from "./servicedirectory";
import * as servicenetworking from "./servicenetworking";
import * as serviceusage from "./serviceusage";
import * as sourcerepo from "./sourcerepo";
import * as spanner from "./spanner";
import * as sql from "./sql";
import * as storage from "./storage";
import * as tags from "./tags";
import * as tpu from "./tpu";
import * as types from "./types";
import * as vertex from "./vertex";
import * as vpcaccess from "./vpcaccess";
import * as workflows from "./workflows";

export {
    accessapproval,
    accesscontextmanager,
    activedirectory,
    apigateway,
    apigee,
    appengine,
    artifactregistry,
    assuredworkloads,
    bigquery,
    bigtable,
    billing,
    binaryauthorization,
    certificateauthority,
    cloudasset,
    cloudbuild,
    cloudfunctions,
    cloudfunctionsv2,
    cloudidentity,
    cloudrun,
    cloudscheduler,
    cloudtasks,
    composer,
    compute,
    config,
    container,
    containeranalysis,
    datacatalog,
    dataflow,
    datafusion,
    dataloss,
    dataproc,
    datastore,
    deploymentmanager,
    diagflow,
    dns,
    endpoints,
    essentialcontacts,
    eventarc,
    filestore,
    firebase,
    firebaserules,
    firestore,
    folder,
    gameservices,
    gkehub,
    healthcare,
    iam,
    iap,
    identityplatform,
    iot,
    kms,
    logging,
    memcache,
    ml,
    monitoring,
    networkconnectivity,
    networkmanagement,
    networkservices,
    notebooks,
    organizations,
    orgpolicy,
    osconfig,
    oslogin,
    projects,
    pubsub,
    recaptcha,
    redis,
    resourcemanager,
    runtimeconfig,
    secretmanager,
    securitycenter,
    serviceaccount,
    servicedirectory,
    servicenetworking,
    serviceusage,
    sourcerepo,
    spanner,
    sql,
    storage,
    tags,
    tpu,
    types,
    vertex,
    vpcaccess,
    workflows,
};

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("gcp", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:gcp") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
